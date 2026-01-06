package updater

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/bytedance/sonic"
)

func updateFromRemote(apiPath string) []*ApiGroup {
	url := apiPath
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get api list: %v", err)
		return nil
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read response body: %v", err)
		return nil
	}

	var Response struct {
		ApiList []*ApiGroup `json:"resources"`
	}

	if err = sonic.Unmarshal(buf, &Response); err != nil {

		fmt.Printf("failed to unmarshal api list: %v\n", err)
		//fmt.Printf("%s\n", string(buf))
		return nil
	}

	return Response.ApiList
}

func ParseTemplate(pkgName string, folder string, apiPath string) {
	t, err := template.New("tmpl").Parse(dataTmpl)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// also prepare model template
	m, err := template.New("model").Parse(modelTmpl)
	if err != nil {
		fmt.Printf("Error parsing model template: %v\n", err)
		return
	}

	apis := updateFromRemote(apiPath)

	os.MkdirAll(folder, os.ModePerm)
	for _, apiGroup := range apis {
		apiGroup.fixed()
		subFolder := folder + "/" + apiGroup.ApiGroupName
		os.Mkdir(subFolder, os.ModePerm)

		file, err := os.Create(subFolder + "/" + apiGroup.ApiGroupName + ".go")
		if err = t.Execute(file, map[string]any{
			"PkgName":      pkgName,
			"ApiGroupName": apiGroup.ApiGroupName,
			"Apis":         apiGroup.Apis,
		}); err != nil {
			fmt.Printf("Error executing template: %v\n", err)
			return
		}
		file.Close()
		fmt.Printf("generate api %s\n", apiGroup.ApiGroupName)

		// For each API in the group, generate a separate model file named {{.Name}}.go
		for _, api := range apiGroup.Apis {
			modelFilePath := subFolder + "/" + api.Name + ".model.go"
			mf, err := os.Create(modelFilePath)
			if err != nil {
				fmt.Printf("Error creating model file %s: %v\n", modelFilePath, err)
				continue
			}
			if err = m.Execute(mf, map[string]any{
				"PkgName":      pkgName,
				"ApiGroupName": apiGroup.ApiGroupName,
				"Apis":         apiGroup.Apis,
				"Name":         api.Name,
			}); err != nil {
				fmt.Printf("Error executing model template for %s: %v\n", api.Name, err)
				mf.Close()
				continue
			}
			mf.Close()
			fmt.Printf("generate model %s\n", api.Name)
		}
	}
	return
	// Generate the index file
	indexFile, err := os.Create(folder + "entry.go")
	if err != nil {
		fmt.Printf("Error creating index file: %v\n", err)
		return
	}
	indexTmpl, err := template.New("index").Parse(apiListTmpl)
	if err != nil {
		fmt.Printf("Error parsing index template: %v\n", err)
		return
	}
	if err = indexTmpl.Execute(indexFile, map[string]any{
		"PkgName": pkgName,
		"Apis":    apis,
	}); err != nil {
		fmt.Printf("Error executing index template: %v\n", err)
		return
	}
}

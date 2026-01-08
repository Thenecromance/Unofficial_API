package updater

import (
	automodel "Unofficial_API/tools/updater/autoModel"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/bytedance/sonic"
	log "github.com/sirupsen/logrus"
)

func AnyToString(v any) string {
	switch val := v.(type) {
	case nil:
		return ""
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		return fmt.Sprint(v)
	}
}

func RequestApiList(pkgName string, apiPath string, category string) []*ApiGroup {
	url := apiPath
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("failed to get api list: %v", err)
		return nil
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("failed to read response body: %v", err)
		return nil
	}

	var Response struct {
		ApiList []*ApiGroup `json:"resources"`
	}

	if err = sonic.Unmarshal(buf, &Response); err != nil {
		log.Errorf("failed to unmarshal api list: %v", err)
		return nil
	}
	for i := range Response.ApiList {
		Response.ApiList[i].Game = pkgName
		Response.ApiList[i].Category = category
	}

	return Response.ApiList
}

type GenerateHandler func(pkgName string, folder string, apiList []*ApiGroup)

// writeFile creates the file (and any necessary directories) and executes the template into it.
func writeFile(filePath string, t *template.Template, data any) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := t.Execute(file, data); err != nil {
		return err
	}
	return nil
}

func read(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func parse(path string) (*template.Template, error) {
	file, err := read(path)
	if err != nil {
		return nil, err
	}
	log.Debugf("Parsing template from %s", path)
	return template.New(path).Parse(file)
}

func GenerateApi(pkgName string, folder string, apiList []*ApiGroup) {
	t, err := parse("./templates/api.tmpl")
	if err != nil {
		log.Errorf("Error parsing template: %v", err)
		return
	}

	for _, apiGroup := range apiList {
		// e.g. folder/GroupName/GroupName.go
		filePath := filepath.Join(folder, apiGroup.ApiGroupName, apiGroup.ApiGroupName+".go")

		data := map[string]any{
			"PkgName":            pkgName,
			"ApiGroupName":       apiGroup.ApiGroupName,
			"Apis":               apiGroup.Apis,
			"NeedStrconv":        apiGroup.NeedStrconv(), // Change: Call the method explicitly
			"HasURIBinding":      apiGroup.HasURIBinding(),
			"ProcessChineseData": apiGroup.ProcessChineseData(),
		}

		if err := writeFile(filePath, t, data); err != nil {
			log.Error("Error generating API file %s: %v", filePath, err)
			continue
		}
		log.Infof("generate api %s", apiGroup.ApiGroupName)
	}
}

func GenerateModels(pkgName string, folder string, apiList []*ApiGroup) {
	t, err := parse("./templates/model.tmpl")
	if err != nil {
		log.Error("Error parsing template: %v", err)
		return
	}

	for _, apiGroup := range apiList {
		for _, api := range apiGroup.Apis {
			// e.g. folder/GroupName/ApiName.model.go
			modelFilePath := filepath.Join(folder, apiGroup.ApiGroupName, api.Name+".model.go")

			data := map[string]any{
				"PkgName":      pkgName,
				"ApiGroupName": apiGroup.ApiGroupName,
				"Apis":         apiGroup.Apis,
				"Name":         api.Name,
			}

			args := make(map[string]any)
			for _, param := range api.Params {
				if param.Type == "string" { // todo: this part of code need to be improved
					args[param.SourceName] = AnyToString(param.DefaultValue)
					continue
				} else {
					args[param.ParamName] = AnyToString(param.DefaultValue)

				}
			}

			str, err := automodel.Request(&automodel.Fields{
				Method:      api.Method,
				Path:        api.Path,
				StructName:  api.Name + "Model",
				PackageName: pkgName + "_" + apiGroup.ApiGroupName,
				Params:      args})
			if err == nil {
				data["AutoModel"] = str
			}

			if err := writeFile(modelFilePath, t, data); err != nil {
				log.Errorf("Error generating model file %s: %v", modelFilePath, err)
				continue
			}
			log.Infof("generate model %s", api.Name)
		}
	}
}

func GenerateRouters(pkgName string, folder string, apiList []*ApiGroup) {
	log.Debug("Generating routers...")
	t, err := parse("./templates/routers.tmpl")
	if err != nil {
		log.Errorf("Error parsing template: %v", err)
		return
	}

	log.Debugf("Target folder: %s", folder)

	for _, apiGroup := range apiList {
		// e.g. folder/GroupName.go
		filePath := filepath.Join(folder, apiGroup.ApiGroupName+".go")

		data := map[string]any{
			"Game":         apiGroup.Game,
			"Category":     apiGroup.Category,
			"ApiGroupName": apiGroup.ApiGroupName,
			"Apis":         apiGroup.Apis,
			"NeedStrconv":  apiGroup.NeedStrconv(), // Change: Call the method explicitly
		}

		if err := writeFile(filePath, t, data); err != nil {
			log.Errorf("Error generating router file %s: %v", filePath, err)
			continue
		}
		log.Infof("generate router %s", apiGroup.ApiGroupName)
	}
}

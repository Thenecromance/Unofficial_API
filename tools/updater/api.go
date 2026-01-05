package updater

import (
	"fmt"
	"strings"
)

type ApiGroup struct {
	ApiGroupName string `json:"name"`    // just like "Achievement API"
	Apis         []*Api `json:"methods"` // just like "Achievement API"
}

func (ap *ApiGroup) fixed() {
	if strings.Contains(ap.ApiGroupName, " API") {
		fixedName := ap.ApiGroupName[:len(ap.ApiGroupName)-4]
		fixedName = strings.ReplaceAll(fixedName, " ", "")
		ap.ApiGroupName = fixedName
	}
	for i := 0; i < len(ap.Apis); i++ {
		ap.Apis[i].fixed()
	}

}

type Api struct {
	Name        string        `json:"name"`        // just like "Achievement API"
	Description string        `json:"description"` // what this api do
	Path        string        `json:"path"`        // just like "/wow/achievement"
	Method      string        `json:"method"`      // just like "GET"
	Params      []*Parameters `json:"parameters"`  // just like "id"
	NameSpace   string        `json:"-"`           // the real place which will write to api file
}

func (a *Api) fixed() {
	a.Name = strings.ReplaceAll(a.Name, " ", "") // Remove spaces
	a.Description = a.Name + " " + a.Description
	a.fixParams()
	//a.fixPath()

}

func (a *Api) fixParams() {
	var filtered []*Parameters
	for _, p := range a.Params {
		if strings.Contains(p.Name, "locale") {
			continue
		} else if strings.Contains(p.Name, "namespace") {
			a.NameSpace = strings.Split(p.DefaultValue.(string), "-")[0] // just need this api is static or dynamic
		} else if strings.Contains(p.Name, ".") {
			p.Name = strings.ReplaceAll(p.Name, ".", "Dot")
			fmt.Printf("wrong params need to be fixed %s\n", p.Name)
			filtered = append(filtered, p)
		} else {
			filtered = append(filtered, p)
		}
	}
	a.Params = filtered

	for i := 0; i < len(a.Params); i++ {
		a.Params[i].fixed()
	}
}

func (a *Api) formatPath(template string, params map[string]string) string {
	var args = make([]string, 0)
	for key, value := range params {
		placeholder := "{" + key + "}"
		switch value {
		case "string":

			template = strings.Replace(template, placeholder, "%s", -1)
			args = append(args, key)
		case "int":
			template = strings.Replace(template, placeholder, "%d", -1)
			args = append(args, key)
		case "float":
			template = strings.Replace(template, placeholder, "%f", -1)
			args = append(args, key)
		default:
			panic(fmt.Errorf("unsupported type: %s", value))
		}
	}

	fmt.Sprintf("fmt.Sprinf(template, args...) ")

	result := "fmt.Sprintf(\"" + template + "\","
	for _, arg := range args {
		result += arg + ","
	}
	result = strings.TrimRight(result, ",")
	result += ")"

	return result
}

func (a *Api) fixPath() {
	// remove unnecessary slashes
	if strings.Contains(a.Path, "/data/wow/") {
		a.Path = strings.ReplaceAll(a.Path, "/data/wow/", "") // Remove slashes
	}

	// path with params
	if len(a.Params) == 0 {
		a.Path = "\"" + a.Path + "\""
		return
	}
	// Create a map to hold the parameters
	params := make(map[string]string)

	var notRequired bool
	// Iterate over the parameters and add them to the map
	for _, param := range a.Params {
		params[param.Name] = param.Type
		if !param.Required {
			notRequired = true
		}
	}

	// Format the path with the parameters
	// the final params will be like ===>fmt.Sprintf("apis/%s" , ...args )
	// but if the params contains the not required params,  the final path should be like====> fmt.Sprintf("api/?paramsName=paramsValue")
	// not done yet lol
	if !notRequired {
		a.Path = a.formatPath(a.Path, params)
	} else {
		a.Path = fmt.Sprintf("\"%s\"", a.Path)
	}
}

type Parameters struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
	DefaultValue any    `json:"defaultValue"`
}

func (p *Parameters) fixed() {
	p.Name = strings.ReplaceAll(p.Name, "{", "") // remove
	p.Name = strings.ReplaceAll(p.Name, "}", "")

	if strings.HasPrefix(p.Name, ":") {
		p.Name = strings.ReplaceAll(p.Name, ":", "")
	}

	// todo: parse Type to support golang types
	switch p.Type {
	case "integer":
		p.Type = "int"
	}
	// todo: fix the DefaultValue ( which just like static-us or dynamic-us)
}

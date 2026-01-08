package updater

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type ApiGroup struct {
	Game         string `json:"package"` // just like "Achievement"
	Category     string `json:"category"`
	ApiGroupName string `json:"name"`    // just like "Achievement API"
	Apis         []*Api `json:"methods"` // just like "Achievement API"
}

func (ap *ApiGroup) PackageName() string {
	return ap.Game + "_" + ap.Category
}

func (ap *ApiGroup) Fixed() {
	if strings.HasSuffix(ap.ApiGroupName, " API") {
		ap.ApiGroupName = strings.TrimSuffix(ap.ApiGroupName, " API")
	}
	ap.ApiGroupName = strings.ReplaceAll(ap.ApiGroupName, " ", "")

	for _, api := range ap.Apis {
		api.fixed()
	}
}

func (ap *ApiGroup) NeedStrconv() bool {
	for _, p := range ap.Apis {
		if p.NeedStrconv() {
			return true
		}
	}
	return false
}

func (ap *ApiGroup) HasURIBinding() bool {
	for _, p := range ap.Apis {
		if p.HasURIBinding() {
			return true
		}
	}
	return false
}

func (ap *ApiGroup) ProcessChineseData() bool {
	for _, p := range ap.Apis {
		if p.ProcessChineseData() {
			return true
		}
	}
	return false
}

type Api struct {
	Name        string        `json:"name"`        // just like "Achievement API"
	Description string        `json:"description"` // what this api do
	Path        string        `json:"path"`        // just like "/wow/achievement"
	GinPath     string        `json:"gin_path"`    // the real gin path with :id
	Method      string        `json:"httpMethod"`  // just like "GET"
	CnRegion    bool          `json:"cnRegion"`    // just like true
	Params      []*Parameters `json:"parameters"`  // just like "id"
	NameSpace   string        `json:"name_space"`  // the real place which will write to api file
}

func (a *Api) HasURIBinding() bool {
	for _, p := range a.Params {
		if p != nil && p.IsBindingUri {
			return true
		}
	}
	return false
}

func (a *Api) NeedStrconv() bool {
	for _, p := range a.Params {

		if !p.IsBindingUri && p.Type == "int" {
			return true
		}
	}
	return false
}

func (a *Api) ProcessChineseData() bool {
	var containName, containRealm bool
	containName = strings.Contains(a.Path, "characterName")
	containGuild := strings.Contains(a.Path, "nameSlug")
	containRealm = strings.Contains(a.Path, "realmSlug")
	return (containName || containGuild) && containRealm
}

func (a *Api) IsGuild() bool {
	return strings.Contains(a.Path, "nameSlug")
}

func (a *Api) fixed() {
	if strings.Contains(a.Name, " Card") {
		log.Info(a.Name)
	}
	a.Name = strings.ReplaceAll(a.Name, " ", "") // Remove spaces
	a.Name = strings.ReplaceAll(a.Name, "(US,EU,KR,TW)", "")
	a.Name = strings.ReplaceAll(a.Name, "(CN)", "CN")

	a.Name = strings.Title(a.Name)

	a.Description = a.Name + " " + a.Description
	a.fixParams()
	a.fixPath()
}

func (a *Api) fixParams() {
	// 使用 pre-allocate 优化性能
	filtered := make([]*Parameters, 0, len(a.Params))

	for _, p := range a.Params {
		// 1. 优先检查并删除已弃用的参数
		if strings.Contains(p.SourceName, "(deprecated)") {
			log.Infof("found deprecated param: %s at API:%s", p.SourceName, a.Name)
			continue // 跳过 append，即为删除
		}

		// 2. 修复无效的字符
		if strings.Contains(p.SourceName, ".") {
			oldName := p.SourceName
			p.SourceName = strings.ReplaceAll(p.SourceName, ".", "")
			log.Infof("fixed param name from %s to %s", oldName, p.SourceName)
		}

		filtered = append(filtered, p)
	}
	a.Params = filtered

	for _, p := range a.Params {
		p.fixed()
	}
}

func (a *Api) fixPath() {
	lists := strings.Split(a.Path, "/")
	for i, segment := range lists {
		if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
			paramName := strings.ReplaceAll(strings.ReplaceAll(segment, "{", ""), "}", "")

			lists[i] = ":" + paramName // just mark it
		}

		if strings.HasPrefix(segment, ":") {
			paramName := strings.ReplaceAll(strings.ReplaceAll(segment, ":", ""), "}", "")

			// bypass regionID and regionId
			if paramName == "regionID" {
				paramName = "regionId"
			}
			if paramName == "realmID" {
				paramName = "realmId"
			}

			lists[i] = ":" + paramName // just mark it
		}
	}
	a.GinPath = strings.Join(lists, "/")
}

type Parameters struct {
	SourceName   string `json:"name"`
	Name         string `json:"title"`      // original name
	ParamName    string `json:"param_name"` // fixed name for golang
	IsBindingUri bool   `json:"is_binding_uri"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
	DefaultValue any    `json:"defaultValue"`
}

func (p *Parameters) fixed() {
	if p == nil {
		return
	}

	if strings.Contains(p.SourceName, "{") && strings.Contains(p.SourceName, "}") {
		// fmt.Printf("found binding uri param: %s\n", p.SourceName)
		p.SourceName = strings.ReplaceAll(p.SourceName, "{", "") // remove
		p.SourceName = strings.ReplaceAll(p.SourceName, "}", "")
		p.IsBindingUri = true
	}

	if strings.HasPrefix(p.SourceName, ":") {
		p.SourceName = strings.ReplaceAll(p.SourceName, ":", "")
		p.IsBindingUri = true
	}
	p.ParamName = p.SourceName
	p.Name = strings.Title(p.SourceName) // capitalize the first letter

	switch p.Type {
	case "integer":
		p.Type = "int"
	case "number":
		p.Type = "int"
	case "numbers":
		p.Type = "[]int"
	}
}

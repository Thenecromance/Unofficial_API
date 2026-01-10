package postmanconverter

type Content struct {
	Info     Info       `json:"info"`
	Items    []Item     `json:"item"`
	Auth     Auth       `json:"auth"`
	Events   []Event    `json:"event"`
	Variable []Variable `json:"variable"`
}

type Info struct {
	PostmanId  string `json:"_postman_id"`
	Name       string `json:"name"`
	Schema     string `json:"schema"`
	ExporterId string `json:"_exporter_id"`
}

type Item struct {
	Name     string    `json:"name"`
	SubItems []SubItem `json:"item"`
}
type SubItem struct {
	Name string        `json:"name"`
	Item []SubItemInfo `json:"item"`
}
type SubItemInfo struct {
	Name    string `json:"name"`
	Request struct {
		Method string        `json:"method"`
		Header []interface{} `json:"header"`
		Url    struct {
			Raw   string   `json:"raw"`
			Host  []string `json:"host"`
			Path  []string `json:"path"`
			Query []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"query"`
		} `json:"url"`
	} `json:"request"`
	Response []interface{} `json:"response"`
}
type Auth struct {
	Type   string `json:"type"`
	Oauth2 struct {
		TokenName      string `json:"tokenName"`
		AccessTokenUrl string `json:"accessTokenUrl"`
		GrantType      string `json:"grant_type"`
		AddTokenTo     string `json:"addTokenTo"`
		AuthUrl        string `json:"authUrl"`
		ClientSecret   string `json:"clientSecret"`
		ClientId       string `json:"clientId"`
	} `json:"oauth2"`
}

type Event struct {
	Listen string `json:"listen"`
	Script struct {
		Type     string `json:"type"`
		Packages struct {
		} `json:"packages"`
		Requests struct {
		} `json:"requests"`
		Exec []string `json:"exec"`
	} `json:"script"`
}
type Variable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

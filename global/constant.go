package global

import "strings"

const (
	NspcStatic  = "static"
	NspcDynamic = "dynamic"
	NspcProfile = "profile"
)

const (
	RegionCN = "cn"
	RegionUS = "us"
	RegionEU = "eu"
	RegionKR = "kr"
)

const (
	LocaleCN = "zh_CN"
	LocaleUS = "en_US"
	LocaleSP = "es_MX"
	LocalePT = "pt_BR"
	LocaleDE = "de_DE"
	LocaleGB = "en_GB"
	LocaleES = "es_ES"
	LocaleFR = "fr_FR"
	LocaleIT = "it_IT"
	LocaleRU = "ru_RU"
	LocaleKR = "ko_KR"
	LocaleTW = "zh_TW"
)

func Locale() string {
	return LocaleCN
}

func Region() string {
	return "us"
}

func Profile() string {
	var builder strings.Builder
	builder.WriteString(NspcProfile)
	builder.WriteString("-")
	builder.WriteString(Region())
	return builder.String()
}

func Static() string {
	var builder strings.Builder
	builder.WriteString(NspcStatic)
	builder.WriteString("-")
	builder.WriteString(Region())
	return builder.String()
}
func Dynamic() string {
	var builder strings.Builder
	builder.WriteString(NspcDynamic)
	builder.WriteString("-")
	builder.WriteString(Region())
	return builder.String()
}

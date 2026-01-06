package utils

type RequestMethod struct {
	Methods  string
	Path     string
	CnRegion bool
}

type CNRequestMethod struct {
	Name      string
	RealmSlug string
}

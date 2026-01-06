package client

import (
	"Unofficial_API/Interface"
	"Unofficial_API/bridge/basicImpl"
	"io"
)

var _cli Interface.Request

func init() {
	_cli = basicImpl.NewRequest()
}

func GET(url string, args ...interface{}) (response string, err error) {
	return _cli.GET(url, args...)
}
func GETWithHeaders(url string, headers map[string]string, args ...interface{}) (response string, err error) {
	return _cli.GETWithHeaders(url, headers, args...)
}
func GETWithCookies(url string, cookies map[string]string, args ...interface{}) (response string, err error) {
	return _cli.GETWithCookies(url, cookies, args...)
}
func POST(url string, args ...interface{}) (response string, err error) {
	return _cli.POST(url, args...)
}
func POSTWithBody(url string, body io.Reader) (response string, err error) {
	return _cli.POSTWithBody(url, body)
}

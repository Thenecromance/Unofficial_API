package Interface

import "io"

type Request interface {
	GET(url string, args ...interface{}) (response string, err error)
	GETWithHeaders(url string, headers map[string]string, args ...interface{}) (response string, err error)
	GETWithCookies(url string, cookies map[string]string, args ...interface{}) (response string, err error)
	POST(url string, args ...interface{}) (response string, err error)
	POSTWithBody(url string, body io.Reader) (response string, err error)
}

package utils

import (
	"Unofficial_API/Interface"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type simpleRequest struct {
	_cli http.Client
}

func (sr *simpleRequest) constructURL(url string, args ...interface{}) string {
	argLen := len(args)

	// if there are no arguments or args is not equal key,value pair, just return the original URL
	if argLen == 0 || argLen%2 != 0 {
		return url
	}
	var builder strings.Builder
	builder.WriteString(url)
	builder.WriteString("?")
	for i := 0; i < argLen; i += 2 {
		if i != 0 {
			builder.WriteString("&")
		}
		builder.WriteString(args[i].(string))
		builder.WriteString("=")
		builder.WriteString(args[i+1].(string))
	}
	return builder.String()
}

func (sr *simpleRequest) GET(url string, args ...interface{}) (response string, err error) {
	url = sr.constructURL(url, args...)
	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Add("Accept", "*/*")

	resp, err := sr._cli. /*Get(url)*/ Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBytes), nil
}

func (sr *simpleRequest) GETWithHeaders(url string, headers map[string]string, args ...interface{}) (response string, err error) {
	req, err := http.NewRequest("GET", sr.constructURL(url, args...), nil)
	//req.Header.Add("Accept", "*/*")
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := sr._cli. /*Get(url)*/ Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBytes := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(responseBytes)
	if err != nil {
		return "", err
	}

	return string(responseBytes), nil
}

func (sr *simpleRequest) GETWithCookies(url string, cookies map[string]string, args ...interface{}) (response string, err error) {
	url = sr.constructURL(url, args...)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0")

	//req.Header.Add("Accept", "*/*")
	for key, value := range cookies {
		req.AddCookie(&http.Cookie{Name: key, Value: value})
	}

	resp, err := sr._cli. /*Get(url)*/ Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(responseBytes), nil
}

func (sr *simpleRequest) POST(url string, args ...interface{}) (response string, err error) {
	req, err := http.NewRequest("POST", sr.constructURL(url, args...), nil)

	//req.Header.Add("Content-Type", "application/json; charset=utf-8")
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	//req.Header.Add("Accept-Charset", "gzip, deflate, br, zstd")
	//req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	//req.Header.Add("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0")

	resp, err := sr._cli. /*Get(url)*/ Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(responseBytes), nil
}

func (sr *simpleRequest) POSTWithBody(url string, body io.Reader) (response string, err error) {
	req, err := http.NewRequest("POST", url, body)

	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	//req.Header.Add("Accept-Charset", "gzip, deflate, br, zstd")
	//req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	//req.Header.Add("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36 Edg/144.0.0.0")
	req.Host = "worldofwarcraft.blizzard.com"
	resp, err := sr._cli. /*Get(url)*/ Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(responseBytes), nil
}

var _instance Interface.Request = nil

func NewRequest() Interface.Request {
	if _instance == nil {
		_instance = &simpleRequest{
			_cli: http.Client{},
		}
	}
	return _instance
}

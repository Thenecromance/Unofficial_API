package Authentication

import (
	"Unofficial_API/bridge/log"
	"Unofficial_API/global"
	"bytes"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/bytedance/sonic"
)

var tokenRefresherMtx sync.RWMutex // TokenRefresherMtx protects access to the token
var token string                   // token holds the current access token
var expiration int64               // expiration holds the token expiration time in seconds

func init() {
	authenticate()
}

func GetToken() string {
	tokenRefresherMtx.RLock()
	defer tokenRefresherMtx.RUnlock()
	return token
}

func authenticate() error {
	tokenRefresherMtx.Lock()
	defer tokenRefresherMtx.Unlock()

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	body := bytes.NewBufferString(data.Encode())
	req, _ := http.NewRequest("POST", "https://us.battle.net/oauth/token", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	log.Debug(global.GetClientID(), global.GetClientSecret())
	req.SetBasicAuth(global.GetClientID(), global.GetClientSecret())

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Debug("Blizzard API auth response:", string(buf))
	var response struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"` // no need to care about this so far
		ExpiresIn   int    `json:"expires_in"`
	}
	err = sonic.Unmarshal(buf, &response)
	if err != nil {
		log.Warn(err)
		return err
	}
	token = response.AccessToken
	expiration = int64(response.ExpiresIn) - 300 // refresh 5 minutes before expiration
	return nil
}

func tokenRefresher() {
	for { // todo: add global shutdown signal to exit this loop
		err := authenticate()
		if err != nil {
			log.Warn("failed to authenticate: %v\n", err)
		} else {
			log.Debug("Successfully authenticated with Blizzard API")
		}
		// Refresh token every 90 minutes
		select {
		case <-time.After(time.Duration(expiration) * time.Second):
		}
	}
}

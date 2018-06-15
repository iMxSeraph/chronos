package pushover

import (
	"encoding/json"
	"net/http"
	"bytes"
)

const (
	pushApi = "https://api.pushover.net/1/messages.json"
	Muxix   = "muxin"
	Yating	= "yating"
)

type auth struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

type message struct {
	auth
	Message  string `json:"message"`
	Title    string `json:"title,omitempty"`
	Url      string `json:"url,omitempty"`
	UrlTitle string `json:"url_title,omitempty"`
}

var authList = map[string]auth{
	"muxin":  {Token: "a96ttsq2cdjwtse33ksowumhwgg3wx", User: "uvnzwbxxbovejpzso2mfz58geztd7y"},
	"yating": {Token: "a1f6fbivqopn2k9qshvceet4urkaxi", User: "upd74n3j9up7eky83gn564x1kchx9n"}}

func Send(target string, title, msg string) int {
	reqObj := message{auth: authList[target], Title: title, Message: msg}
	reqBody, _ := json.Marshal(reqObj)
	resp, _ := http.Post(pushApi, "application/json", bytes.NewReader(reqBody))
	defer resp.Body.Close()
	return resp.StatusCode
}

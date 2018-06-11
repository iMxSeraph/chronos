package models

type Error struct {
	ErrCode int `json:"errCode"`
	ErrMsg string `json:"errMsg"`
}


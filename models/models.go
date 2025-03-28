package models

type Req struct {
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

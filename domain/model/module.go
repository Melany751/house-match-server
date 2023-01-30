package model

type Module struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Order       int    `json:"order"`
}

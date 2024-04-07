package model

type AppHelth struct {
	Name    string `json:"name,omitempty"`
	Version int64 `json:"version,omitempty"`
	Status  string `json:"status,omitempty"`
}

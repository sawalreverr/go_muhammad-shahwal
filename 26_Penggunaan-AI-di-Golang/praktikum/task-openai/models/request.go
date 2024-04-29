package models

type Request struct {
	Budget  int    `json:"budget"`
	Purpose string `json:"purpose"`
	Merk    string `json:"merk"`
	RAM     string `json:"ram"`
	GPU     string `json:"gpu"`
}

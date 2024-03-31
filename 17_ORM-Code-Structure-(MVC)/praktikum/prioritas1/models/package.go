package models

type Package struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Sender         string  `json:"sender"`
	Receiver       string  `json:"receiver"`
	SenderLocation string  `json:"sender_location"`
	SenderReceiver string  `json:"sender_receiver"`
	Fee            int     `json:"fee"`
	Weight         float64 `json:"weight"`
}

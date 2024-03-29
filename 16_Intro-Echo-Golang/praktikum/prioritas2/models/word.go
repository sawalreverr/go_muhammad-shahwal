package models

type Word struct {
	Word        string `json:"word"`
	Length      int    `json:"length"`
	NumOfVocals int    `json:"num_of_vocals"`
	Palindrome  bool   `json:"palindrome"`
}

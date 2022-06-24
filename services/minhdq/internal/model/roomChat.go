package model

type Room struct {
	History []string `json:"history"`
	ID      string   `json:"ID"`
	User    []string `json:"user"`
}

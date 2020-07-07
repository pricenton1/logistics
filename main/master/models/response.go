package models

type Responses struct {
	Messages string `json:"messages"`
	Status   int    `json:"status"`
	Data     interface{}
}

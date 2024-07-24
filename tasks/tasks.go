package models

type Task struct {
	ID       int    `json:"id"`
	DESC     string `json:"desc"`
	COMPLETE bool   `json:"complete"`
}

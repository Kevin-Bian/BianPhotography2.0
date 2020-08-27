package models

type User struct {
	ID       int64  `json:"userid"`
	Location string `json:"location"`
	UserName string `json:"username"`
}

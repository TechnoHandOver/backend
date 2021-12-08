package models

type User struct {
	Id     uint32 `json:"id"`
	VkId   uint32 `json:"vkId"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Users []*User

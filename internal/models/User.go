package models

type User struct {
	Id     uint32 `json:"id"`
	VkId   uint32 `json:"vkId"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type UserUpdate struct {
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

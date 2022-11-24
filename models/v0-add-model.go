package models

import "mime/multipart"

type AddRequest struct {
	Id           string                //`form:"id" binding:"required"`
	FileReceived *multipart.FileHeader `form:"file" binding:"required"`
}

type AddResponse struct {
	Bytes int64  `json:"Bytes"`
	Hash  string `json:"Hash"`
	Name  string `json:"Name"`
	Size  string `json:"Size"`
	Link  string `json:"Link"`
	Owner string `json:"Owner"`
}

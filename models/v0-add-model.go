package models

import "mime/multipart"

type AddRequest struct {
	Id           string                //`form:"id" binding:"required"`
	FileReceived *multipart.FileHeader `form:"file" binding:"required"`
}

type AddResponse struct {
	Hash  string `json:"Hash"`
	Name  string `json:"Name"`
	Size  string `json:"Size"`
}

package models

import "mime/multipart"

// AddRequest example
type AddRequest struct {
	Id           string                //`form:"id" binding:"required"`
	Owner        string                `json:"owner" binding:"required"`
	FileReceived *multipart.FileHeader `form:"file" binding:"required"`
}

// AddResponse example
type AddResponse struct {
	Hash  string `json:"Hash" example:"HASH-FOR-FILE"`
	Name  string `json:"Name" example:"NAME-OF-FILE"`
	Size  string `json:"Size" example:"SIZE-OF-FILE"`
	Link  string `json:"Link" example:"LINK-TO-FILE"`
	Owner string `json:"Owner" example:"OWNER-OF-FILE"`
}

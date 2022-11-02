package models

//import "mime/multipart"

type GetRequest struct {
	FileType          string              
	Content 		  string
}

type GetResponse struct {
	FileType string `json: "type"`
	Content  string `json:"Content"`
}

package models

type UserFileResponse struct {
	Userfiles []AddResponse `json:"userfiles"  binding:"required"`
}

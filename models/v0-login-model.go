package models

type LoginRequest struct {
	Username     string `json:"Username"  binding:"required"`
	UserPassword string `json:"UserPassword"  binding:"required"`
}

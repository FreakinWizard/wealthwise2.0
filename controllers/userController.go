package controllers

import "github.com/gin-gonic/gin"

func Signup(r *gin.Context){

	var body struct {
		Email string
		Password string
	}
}
package controllers

import (
	"net/http"
	"secure/challenge-1/domain"

	"github.com/gin-gonic/gin"
)

// Post element data
func PostElement(ctx *gin.Context){
	var postReq domain.ElementRequest
	if err := ctx.BindJSON(&postReq); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hitung status untuk water
	var waterStatus string
	if postReq.Water < 5 {
		waterStatus = "aman"
	} else if postReq.Water >= 5 && postReq.Water <= 8 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	// Hitung status untuk wind
	var windStatus string
	if postReq.Wind < 6 {
		windStatus = "aman"
	} else if postReq.Wind >= 6 && postReq.Wind <= 15 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	// Kirim response dalam format JSON
	postRes := domain.ElementResponse{
		Water:  postReq.Water,
		Wind:   postReq.Wind,
		WaterStatus: waterStatus,
		WindStatus: windStatus,
	}
	ctx.JSON(http.StatusOK, postRes)
}
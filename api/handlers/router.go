package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
)

type T struct {
	Image string `json:"image"`
}

func FindMatch(ctx *gin.Context) {

	ctx.JSON(200, T{Image: "img"})
}

func SubmitQuiz(ctx *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	println(string(body))
}

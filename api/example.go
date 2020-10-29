package api

import (
	"github.com/gin-gonic/gin"
)

// Example godoc
// @Summary summary
// @Description description
// @ID get-example
// @Accept  json
// @Produce  json
// @Param id path int true "Example ID"
// @Success 200 {object} Example
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /example [get]
func Example(ctx *gin.Context) {
	ctx.JSON(200, "")
}

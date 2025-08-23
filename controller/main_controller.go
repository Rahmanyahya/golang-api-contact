package controller

import (
	"golang-api-contact/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func NewMainController() *MainController {
	return &MainController{}
}

func (h *MainController) MainController(c *gin.Context) {
	c.JSON(http.StatusOK, response.APIResponse{
		Code: "SUCCESS",
		Message: "API Contact Management is running",
	})
}
package controllers

import (
	"bus-booking/models"
	"bus-booking/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllCodes(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	codes := make([]models.Codes, 0)
	err = models.AllCodes(&codes, &session)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"codes": codes,
	})
}
func FAcodes(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	codes := models.Codes{CodeID: c.Param("codeID")}
	err = models.FAcode(&codes, &session)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

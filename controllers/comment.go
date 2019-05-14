package controllers

import (
	"bus-booking/models"
	"bus-booking/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Commentlist(c *gin.Context) {
	comments := make([]models.Comment, 0)
	busID := c.Param("busID")
	err := models.Commentlist(&comments, &busID)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"comment": comments,
	})
}

func InsertComment(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	busID := c.PostForm("busID")
	content := c.PostForm("content")
	starss, err := strconv.Atoi(c.PostForm("stars"))
	var stars = uint8(starss)
	bus := models.Bus{BusID: busID}
	comment := models.Comment{Content: content, Stars: stars}
	err = models.InsertComment(&comment, &bus, &session)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}
func Deletecomment(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	commentID := c.Query("commentID")
	log.Print(session)
	log.Print(commentID)
	err = models.Deletecomment(&session, &commentID)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}
func ReplyComment(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	commentID := c.Param("commentID")
	contentReplied := c.PostForm("contentReplied")
	comment := models.Comment{CommentID: commentID, ContentReplied: contentReplied}
	err = models.ReplyComment(&comment, &session)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

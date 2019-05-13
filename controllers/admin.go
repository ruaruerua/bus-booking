package controllers
import(
	"bus-booking/models"
	"bus-booking/util"
	"github.com/gin-gonic/gin"
	"net/http"
)
func Alogin(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	user := models.User{Account: account, Password: password}
	session, err := models.Alogin(&user)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.SetCookie("session", session, 3600, "/", util.Domain, false, false)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"session": session,
	})
}
package httplayer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type user struct {
	Name   string `json:"name"`
	Handle string `json:"handle"`
}

func (h *httpApi) createUser(c *gin.Context) {
	newUser := &user{}
	err := c.BindJSON(newUser)
	if err != nil {
		logrus.Println(err.Error())
		return
	}

	err = h.app.CreateUser(c, newUser.Name, newUser.Handle)
	if err != nil {
		logrus.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusAccepted, "success")
}

func (h *httpApi) getAllUsers(c *gin.Context) {
	users, err := h.app.GetAllUsers(c)
	if err != nil {
		logrus.Println(err.Error())
		return
	}
	var userList []gin.H
	for _, user := range users {
		userList = append(userList, gin.H{
			"handle": user.Handle,
			"name":   user.Name,
		})
	}
	c.JSONP(http.StatusOK, userList)
}

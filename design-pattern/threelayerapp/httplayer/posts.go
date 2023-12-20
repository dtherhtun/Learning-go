package httplayer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type post struct {
	Content string `json:"content"`
}

func (h *httpApi) createPost(c *gin.Context) {
	newPost := &post{}
	err := c.BindJSON(newPost)
	if err != nil {
		logrus.Println("failed to bind in create post: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = h.app.CreatePost(c, newPost.Content, "someone")
	if err != nil {
		logrus.Println("failed to create post: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "created successfully",
	})
}

func (h *httpApi) getAllPosts(c *gin.Context) {
	posts, err := h.app.GetAllPosts(c)
	if err != nil {
		logrus.Println("failed to get posts: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	var postList []gin.H
	for _, post := range posts {
		postList = append(postList, gin.H{
			"content": post.Content,
			"owner":   post.Owner,
		})
	}
	c.JSON(http.StatusOK, postList)
}

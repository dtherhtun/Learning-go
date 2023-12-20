package httplayer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dtherhtun/Learning-go/design-pattern/threelayerapp/applayer"
)

type httpApi struct {
	engine *gin.Engine
	app    applayer.App
}

func New(app applayer.App) *httpApi {
	a := &httpApi{
		engine: gin.New(),
		app:    app,
	}
	a.SetupRouters()
	return a
}

func (h *httpApi) SetupRouters() {
	h.engine.Use(gin.Recovery())
	api := h.engine.Group("/api")
	{
		api.GET("/ping", pong)
		users := api.Group("/users")
		{
			users.GET("", h.getAllUsers)
			users.POST("", h.createUser)
		}
		posts := api.Group("/posts")
		{
			posts.GET("", h.getAllPosts)
			posts.POST("", h.createPost)
		}
	}
}

func (h *httpApi) Engage() {
	h.engine.Run()
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

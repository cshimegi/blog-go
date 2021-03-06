package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"posts/services"
)

var (
	postService services.PostService
)

// PostAPI is the interface for post API
type PostAPI interface {
	ListAll(c *gin.Context)
}

// PostController defines type of post controller
type PostController struct{}

func init() {
	postService = services.NewPostService()
}

// NewPostAPI defines api paths to post service
func NewPostAPI() *PostController {
	return &PostController{}
}

// ListAll lists all posts
func (p PostController) ListAll(c *gin.Context) {
	posts, err := postService.ListAll(c)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

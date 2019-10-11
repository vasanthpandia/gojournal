package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/vasanthpandia/gojournal/internal/controllers"
	"github.com/vasanthpandia/gojournal/internal/models"
	"github.com/vasanthpandia/gojournal/internal/jsonerrors"
)


func CreatePost(c *gin.Context) {
	controller := c.MustGet("PostsController").(*controllers.PostsController)
	logger := c.MustGet("Logger").(*zap.Logger)
	currentUser := c.MustGet("CurrentUser").(*models.User)

	request := &controllers.PostCreatePayload{}
	err := c.BindJSON(&request)
	if err != nil {
		logger.Error("Json Bind Error", zap.Error(err))
		c.JSON(http.StatusBadRequest, jsonerrors.New("Bad Request Body"))
		c.Abort()
	}

	request.UserID = currentUser.ID

	logger.Info("Post Create Payload", zap.String("text : ", request.Text))

	post, err := controller.Create(request)

	if err != nil {
		logger.Error("Error Creating Post", zap.Error(err))
		c.JSON(http.StatusInternalServerError, jsonerrors.New(err.Error()))
		return
	}

	c.JSON(http.StatusOK, post)
}

func ReadPost(c *gin.Context) {
	controller := c.MustGet("PostsController").(*controllers.PostsController)
	logger := c.MustGet("Logger").(*zap.Logger)
	currentUser := c.MustGet("CurrentUser").(*models.User)

	request := &controllers.PostReadPayload{}
	request.UserID = currentUser.ID
	request.ID = c.Param("postId")

	logger.Info("Params : ", zap.String("UserId", request.UserID), zap.String("PostId", request.ID))

	post, err := controller.Read(request)
	if err != nil {
		logger.Error("Error Fetching Post", zap.Error(err))
		c.JSON(http.StatusNotFound, jsonerrors.ResourceNotFound)
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	controller := c.MustGet("PostsController").(*controllers.PostsController)
	logger := c.MustGet("Logger").(*zap.Logger)
	currentUser := c.MustGet("CurrentUser").(*models.User)

	request := &controllers.PostDeletePayload{}
	request.UserID = currentUser.ID
	request.ID = c.Param("postId")

	logger.Info("Params : ", zap.String("UserId", request.UserID), zap.String("PostId", request.ID))

	err := controller.Delete(request)
	if err != nil {
		logger.Error("Error Deleting Post", zap.Error(err))
		c.JSON(http.StatusNotFound, jsonerrors.ResourceNotFound)
		return
	}

	c.JSON(http.StatusOK, nil)
}

package route

import (
	"mapping_func/pkg/handler"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, h *handler.UserHandler) {
	r.POST("/users", h.UserPost)
}

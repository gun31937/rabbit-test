package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"net/http"
	"rabbit-test/app/handlers/http/shorturl"
	"rabbit-test/app/middlewares"
	"rabbit-test/app/usecases"
	"rabbit-test/app/utils"
)

// Demo auth middleware.
// Use to prove for unauthorized user can not access
// endpoints that use `authMiddleWare`.
var authMiddleWare = middlewares.AuthMiddleware()

func NewRouterShortURL(engine *gin.Engine, remittanceUseCase usecases.ShortURL) {
	handler := shorturl.Handler{ShortURLUseCase: remittanceUseCase}

	engine.GET("/:shortCode", handler.Get)
	engine.POST("", handler.Create)

	authorized := engine.Group("")
	authorized.Use(authMiddleWare.MiddlewareFunc())
	{
		authorized.GET("", handler.List)
		authorized.DELETE("/:shortCode", handler.Delete)
	}

}

// NewRouterAdmin Using demo auth middleware.
func NewRouterAdmin(engine *gin.Engine) {
	engine.POST("/login", authMiddleWare.LoginHandler)
}

// NewRouterHealth Just make sure services run correctly.
func NewRouterHealth(engine *gin.Engine, dbConnect *gorm.DB, rdbConnect *redis.Client) {

	engine.GET("/health", func(c *gin.Context) {
		if err := dbConnect.DB().Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
			return
		}

		if _, err := rdbConnect.Ping(c).Result(); err != nil {
			c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
			return
		}
		c.JSON(http.StatusOK, utils.NewSuccessResponse(nil))
	})
}

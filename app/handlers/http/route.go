package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"net/http"
	"rabbit-test/app/handlers/http/shorturl"
	"rabbit-test/app/usecases"
	"rabbit-test/app/utils"
)

func NewRouterShortURL(engine *gin.Engine, remittanceUseCase usecases.ShortURL) {
	handler := shorturl.Handler{ShortURLUseCase: remittanceUseCase}
	endpoint := engine.Group("/short-url")
	{
		endpoint.POST("/create", handler.Create)
	}

	engine.GET("/:shortCode", handler.Get)

}

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

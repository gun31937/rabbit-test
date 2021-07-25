package shorturl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rabbit-test/app/utils"
)

func (handler *Handler) Get(ctx *gin.Context) {

	shortCode := ctx.Param("shortCode")

	useCaseResponse, err := handler.ShortURLUseCase.Get(ctx, shortCode)
	if err != nil {
		status, message := utils.GetHTTPStatusCodeWithMessage(err)
		ctx.JSON(status, utils.NewErrorResponse(message))
		return
	}

	ctx.Redirect(http.StatusFound, *useCaseResponse)
}

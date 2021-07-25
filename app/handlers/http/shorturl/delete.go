package shorturl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rabbit-test/app/utils"
)

func (handler *Handler) Delete(ctx *gin.Context) {

	shortCode := ctx.Param("shortCode")

	err := handler.ShortURLUseCase.Delete(ctx, shortCode)
	if err != nil {
		status, message := utils.GetHTTPStatusCodeWithMessage(err)
		ctx.JSON(status, utils.NewErrorResponse(message))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(nil))
}

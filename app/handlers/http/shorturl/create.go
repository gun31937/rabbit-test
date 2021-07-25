package shorturl

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"rabbit-test/app/usecases/shorturl"
	"rabbit-test/app/utils"
)

type createShortURLRequest struct {
	FullURL   string `json:"fullURL" binding:"required"`
	ExpiresIn *int   `json:"expiresIn"` // expires in x minutes.
}

type createShortURLResponse struct {
	ShortURL string  `json:"shortURL"`
	Expiry   *string `json:"expiry"`
}

func (handler *Handler) Create(ctx *gin.Context) {
	request, err := new(createShortURLRequest).parseRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	useCaseResponse, err := handler.ShortURLUseCase.Create(ctx, request.FullURL, request.ExpiresIn)
	if err != nil {
		status, message := utils.GetHTTPStatusCodeWithMessage(err)
		ctx.JSON(status, utils.NewErrorResponse(message))
		return
	}

	response := new(createShortURLResponse).parseResponse(useCaseResponse)
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(response))
}

func (m *createShortURLRequest) parseRequest(c *gin.Context) (*createShortURLRequest, error) {
	if err := c.ShouldBindJSON(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *createShortURLResponse) parseResponse(response *shorturl.CreateShortURLResponse) *createShortURLResponse {
	_ = copier.Copy(m, response)
	return m
}

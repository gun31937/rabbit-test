package shorturl

import (
	"encoding/json"
	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"rabbit-test/app/usecases/shorturl"
	"rabbit-test/app/utils"
	"time"
)

type listURLRequest struct {
	ShortCode *string `form:"shortCode"`
	Keyword   *string `form:"keyword"`
}

type listURL struct {
	ID        int        `json:"id"`
	ShortCode string     `json:"shortCode"`
	FullURL   string     `json:"fullURL"`
	Expiry    *time.Time `json:"expiry"`
	Hits      int        `json:"hits"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type listURLResponse []listURL

func (handler *Handler) List(ctx *gin.Context) {
	model, err := new(listURLRequest).parseRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewErrorResponse(err.Error()))
		return
	}

	request := shorturl.ListURLRequest{}
	_ = copier.Copy(&request, &model)

	useCaseResponse, err := handler.ShortURLUseCase.List(request)
	if err != nil {
		status, message := utils.GetHTTPStatusCodeWithMessage(err)
		ctx.JSON(status, utils.NewErrorResponse(message))
		return
	}

	response := new(listURLResponse).parseResponse(useCaseResponse)
	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(response))
}

func (m *listURLRequest) parseRequest(c *gin.Context) (*listURLRequest, error) {
	if err := c.ShouldBindQuery(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (listURLResponse *listURLResponse) parseResponse(data []shorturl.URL) *listURLResponse {
	_ = copier.Copy(listURLResponse, data)
	return listURLResponse
}

func (list *listURL) MarshalJSON() ([]byte, error) {
	type puppetModel listURL

	location, _ := time.LoadLocation("Asia/Bangkok")
	var timeFormat = "2006-01-02 15:04:05"

	var expiry *string
	var deletedAt *string

	if list.Expiry != nil {
		expiry = pointer.ToString(list.Expiry.In(location).Format(timeFormat))
	}

	if list.DeletedAt != nil {
		deletedAt = pointer.ToString(list.DeletedAt.In(location).Format(timeFormat))
	}

	return json.Marshal(&struct {
		*puppetModel
		Expiry    *string `json:"expiry"`
		CreatedAt string  `json:"createdAt"`
		UpdatedAt string  `json:"updatedAt"`
		DeletedAt *string `json:"deletedAt"`
	}{
		Expiry:      expiry,
		CreatedAt:   list.CreatedAt.In(location).Format(timeFormat),
		UpdatedAt:   list.UpdatedAt.In(location).Format(timeFormat),
		DeletedAt:   deletedAt,
		puppetModel: (*puppetModel)(list),
	})
}

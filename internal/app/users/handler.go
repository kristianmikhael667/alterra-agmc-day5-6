package users

import (
	"main/internal/factory"
	"main/internal/pkg/util"
	pkgdto "main/pkg/dto"
	"main/pkg/util/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	_, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(pkgdto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	result, err := h.service.Find(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.CustomSuccessBuilder(http.StatusOK, result.Data, "Get Users Success", &result.PaginationInfo).Send(c)
}

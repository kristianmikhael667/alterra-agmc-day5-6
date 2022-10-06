package auth

import (
	"main/internal/dto"
	"main/internal/factory"
	"main/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *Handler {
	return &Handler{
		service: NewService(f),
	}
}

func (h *Handler) LoginByEmailAndPassword(c echo.Context) error {
	payload := new(dto.ByEmailAndPasswordRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	users, err := h.service.LoginByEmailAndPassword(c.Request().Context(), payload)

	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}
	return response.SuccessResponse(users).Send(c)
}

func (h *Handler) RegisterByEmailAndPassword(c echo.Context) error {
	payload := new(dto.RegisterUserRequestBody)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}
	payload.FillDefaults()

	users, err := h.service.RegisterByEmailAndPassword(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(users).Send(c)
}

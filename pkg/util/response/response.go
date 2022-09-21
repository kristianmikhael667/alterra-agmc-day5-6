package response

import "main/pkg/dto"

type Meta struct {
	Success bool
	Message string              `json:"success" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}

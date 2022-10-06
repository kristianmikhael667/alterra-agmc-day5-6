package users

import (
	"context"
	"main/internal/dto"
	"main/internal/factory"
	"main/internal/repository"
	pkgdto "main/pkg/dto"
	"main/pkg/util/response"
)

type service struct {
	UsersRepository repository.Users
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.UserResponse], error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UsersRepository: f.UsersRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.UserResponse], error) {
	users, info, err := s.UsersRepository.FindAll(ctx, payload, &payload.Pagination)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	var data []dto.UserResponse

	for _, user := range users {
		data = append(data, dto.UserResponse{
			ID:       user.ID,
			Fullname: user.Fullname,
			Email:    user.Email,
		})
	}

	result := new(pkgdto.SearchGetResponse[dto.UserResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}

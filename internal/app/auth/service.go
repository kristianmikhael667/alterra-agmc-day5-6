package auth

import (
	"context"
	"errors"
	"main/internal/dto"
	"main/internal/factory"
	"main/internal/pkg/util"
	"main/pkg/constant"
	pkgutil "main/pkg/util"
	res "main/pkg/util/response"
	"main/repository"
)

type service struct {
	UsersRepository repository.Users
}

type Service interface {
}

func NewService(f *factory.Factory) Service {
	return &service{
		UsersRepository: f.UsersRepository,
	}
}

func (s *service) LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.UserWithJWTResponse, error) {
	var result *dto.UserWithJWTResponse

	data, err := s.UsersRepository.FindByEmail(ctx, &payload.Email)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if !(pkgutil.CompareHashPassword(payload.Password, data.Password)) {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			errors.New(res.ErrorConstant.EmailOrPasswordIncorrect.Response.Meta.Message),
		)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID, data.RoleID, data.DivisionID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError, errors.New("Error when generating token"),
		)
	}

	result = &dto.UserWithJWTResponse{
		UserResponse: dto.UserResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		JWT: token,
	}

	return result, nil
}

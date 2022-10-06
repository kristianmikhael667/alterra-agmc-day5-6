package auth

import (
	"context"
	"errors"
	"main/internal/dto"
	"main/internal/factory"
	"main/internal/pkg/util"
	"main/internal/repository"
	"main/pkg/constant"
	pkgutil "main/pkg/util"
	"main/pkg/util/response"
	res "main/pkg/util/response"
)

type service struct {
	UsersRepository repository.Users
}

type Service interface {
	LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.UserWithJWTResponse, error)
	RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterUserRequestBody) (*dto.UserWithJWTResponse, error)
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

func (s *service) RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterUserRequestBody) (*dto.UserWithJWTResponse, error) {
	var result *dto.UserWithJWTResponse
	isExist, err := s.UsersRepository.ExistByEmail(ctx, &payload.Email)
	if err != nil {
		return result, response.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("Users already exists"))
	}

	hashedPassword, err := pkgutil.HashPassword(payload.Password)
	if err != nil {
		return result, response.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	payload.Password = hashedPassword

	data, err := s.UsersRepository.Save(ctx, payload)
	if err != nil {
		return result, response.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	claims := util.CreateJWTClaims(data.Email, data.ID, data.RoleID, data.DivisionID)

	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, errors.New("Error when genering token"))
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

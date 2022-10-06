package repository

import (
	"context"
	"main/internal/dto"
	"main/internal/models"
	pkgdto "main/pkg/dto"
	"strings"

	"gorm.io/gorm"
)

type Users interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]models.User, *pkgdto.PaginationInfo, error)
	FindByEmail(ctx context.Context, email *string) (*models.User, error)
	ExistByEmail(ctx context.Context, email *string) (bool, error)
	Save(ctx context.Context, users *dto.RegisterUserRequestBody) (models.User, error)
}

type users struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *users {
	return &users{
		db,
	}
}

func (r *users) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]models.User, *pkgdto.PaginationInfo, error) {
	var users []models.User
	var count int64

	query := r.Db.WithContext(ctx).Model(&models.User{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? of lower(email) Like ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *users) FindByEmail(ctx context.Context, email *string) (*models.User, error) {
	var data models.User
	err := r.Db.WithContext(ctx).Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *users) ExistByEmail(ctx context.Context, email *string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *users) Save(ctx context.Context, employee *dto.RegisterUserRequestBody) (models.User, error) {
	newUsers := models.User{
		Fullname:   employee.Fullname,
		Email:      employee.Email,
		Password:   employee.Password,
		RoleID:     *employee.RoleID,
		DivisionID: *employee.DivisionID,
	}
	if err := r.Db.WithContext(ctx).Save(&newUsers).Error; err != nil {
		return newUsers, err
	}
	return newUsers, nil
}

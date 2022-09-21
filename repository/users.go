package repository

import (
	"context"
	"main/internal/models"
	pkgdto "main/pkg/dto"
	"strings"

	"gorm.io/gorm"
)

type Users interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]models.User, *pkgdto.PaginationInfo, error)
	FindByEmail(ctx context.Context, email *string) (*models.User, error)
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

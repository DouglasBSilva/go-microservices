package database

import (
	"context"
	"errors"

	"github.com/DouglasBSilva/go-microservices/internal/dberrors"
	"github.com/DouglasBSilva/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service

	result := c.DB.WithContext(ctx).Find(&services)

	return services, result.Error
}

func (c Client) AddService(ctx context.Context, service *models.Service) (*models.Service, error) {
	service.ServiceID = uuid.NewString()

	result := c.DB.WithContext(ctx).Create(&service)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}

		return nil, result.Error
	}

	return service, nil
}

func (c Client) GetService(ctx context.Context, serviceId string) (*models.Service, error) {
	var service models.Service
	result := c.DB.WithContext(ctx).
		Where(models.Service{ServiceID: serviceId}).
		First(&service)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{}
		}

		return nil, result.Error
	}

	return &service, result.Error
}

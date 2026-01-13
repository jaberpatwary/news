package service

import (
	models "app/src/model"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ComService interface {
	CreateCom(c *fiber.Ctx) (*models.Com, error)
	GetAll(c *fiber.Ctx, params *validation.QueryCom) ([]models.Com, error)
	//GetByCommentId(c *fiber.Ctx, id string) (*models.Comment, error)
	//GetByPhoneNumber(c *fiber.Ctx)
	//UpdateComment(c *fiber.Ctx, id string, payload *validation.UpdateComment) (*models.Comment, error)
	//DeleteComment(c *fiber.Ctx, id string) error
}

// Define methods for expert service

type comService struct {
	DB *gorm.DB
}

// DB servie init
func NewComService(db *gorm.DB) ComService {
	return &comService{DB: db}
}

// Create
func (s *comService) CreateCom(c *fiber.Ctx) (*models.Com, error) {
	var com models.Com

	if err := c.BodyParser(&com); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&com).Error; err != nil {
		return nil, err
	}

	// Reload the created record (optional but clean)
	if err := s.DB.First(&com, com.ID).Error; err != nil {
		return nil, err
	}

	return &com, nil
}

// Get All
func (s *comService) GetAll(c *fiber.Ctx, params *validation.QueryCom) ([]models.Com, error) {
	var coms []models.Com

	page := params.Page
	if page <= 0 {
		page = 1
	}

	limit := params.Limit
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := s.DB.
		WithContext(c.Context()).
		Model(&models.Com{}).
		Order("Createdtime asc")

	if search := params.Search; search != "" {
		query = query.Where(
			"commenttext LIKE ? OR useridentity LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)

	}

	result := query.
		Limit(limit).
		Offset(offset).
		Find(&coms)

	if result.Error != nil {
		return nil, result.Error
	}

	return coms, nil
}

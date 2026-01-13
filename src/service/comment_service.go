package service

import (
	models "app/src/model"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CommentService interface {
	CreateComment(c *fiber.Ctx) (*models.Comment, error)
	GetAll(c *fiber.Ctx, params *validation.QueryComment) ([]models.Comment, error)
	GetByCommentId(c *fiber.Ctx, id string) (*models.Comment, error)
	//GetByPhoneNumber(c *fiber.Ctx)
	UpdateComment(c *fiber.Ctx, id string, payload *validation.UpdateComment) (*models.Comment, error)
	DeleteComment(c *fiber.Ctx, id string) error
}

// Define methods for expert service

type commentService struct {
	DB *gorm.DB
}

// DB servie init
func NewCommentService(db *gorm.DB) CommentService {
	return &commentService{DB: db}
}

// Create
func (s *commentService) CreateComment(c *fiber.Ctx) (*models.Comment, error) {
	var comment models.Comment
	if err := c.BodyParser(&comment); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&comment).Error; err != nil { //new
		return nil, err

	}
	if err := s.DB.Preload("User").First(&comment, comment.ID).Error; err != nil { //new
		return nil, err
	}

	return &comment, nil
}

// Get All
func (s *commentService) GetAll(c *fiber.Ctx, params *validation.QueryComment) ([]models.Comment, error) {
	var comments []models.Comment

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
		Model(&models.Comment{}).
		Preload("User").
		Order("created_at asc")

	if search := params.Search; search != "" {
		query = query.Where("name LIKE ? OR phone LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	result := query.
		Limit(limit).
		Offset(offset).
		Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

// GetbyId
func (s *commentService) GetByCommentId(c *fiber.Ctx, id string) (*models.Comment, error) {
	var comment models.Comment

	if err := s.DB.
		WithContext(c.Context()).
		Preload("User").
		First(&comment, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

// Update
func (s *commentService) UpdateComment(c *fiber.Ctx, id string, payload *validation.UpdateComment) (*models.Comment, error) {
	var comment models.Comment

	if err := s.DB.
		WithContext(c.Context()).
		First(&comment, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if payload.Content != nil {
		comment.Content = *payload.Content
	}

	if payload.IsAnonymous != nil {
		comment.IsAnonymous = *payload.IsAnonymous
	}

	if payload.IsDeleted != nil {
		comment.IsDeleted = *payload.IsDeleted
	}

	if err := s.DB.
		WithContext(c.Context()).
		Omit("User").
		Save(&comment).Error; err != nil {
		return nil, err
	}

	if err := s.DB.
		WithContext(c.Context()).
		Preload("User").
		First(&comment, comment.ID).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

// Delete
func (s *commentService) DeleteComment(c *fiber.Ctx, id string) error {
	result := s.DB.
		WithContext(c.Context()).
		Delete(&models.Comment{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "comment not found")
	}

	return nil
}

package service

import (
	"fmt"

	models "app/src/model"
	"app/src/validation"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type AdminService interface {
	CreateAdmin(c *fiber.Ctx) (*models.Admin, error)
	GetAll(c *fiber.Ctx, params *validation.QueryAdmin) ([]models.Admin, error)
	//GetByCommentId(c *fiber.Ctx, id string) (*models.Comment, error)
}

type adminService struct {
	DB *gorm.DB
}

func NewAdminService(db *gorm.DB) AdminService {
	return &adminService{DB: db}
}

func (s *adminService) CreateAdmin(c *fiber.Ctx) (*models.Admin, error) {
	var req AdminCreateRequest

	fmt.Println("RAW BODY:", string(c.Body()))

	if err := c.BodyParser(&req); err != nil {
		return nil, err
	}

	fmt.Println("PARSED PASSWORD:", req.Password)

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("HASH:", string(hash))

	admin := models.Admin{
		Username:     req.Username,
		PasswordHash: string(hash),
		FullName:     req.FullName,
		Email:        req.Email,
		Phone:        req.Phone,
		IsActive:     true,
	}

	fmt.Println("MODEL HASH:", admin.PasswordHash)

	if err := s.DB.Debug().Create(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

func (s *adminService) GetAll(c *fiber.Ctx, params *validation.QueryAdmin) ([]models.Admin, error) {

	var admins []models.Admin

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
		Model(&models.Admin{}).
		Order("createdtime asc")

	if search := params.Search; search != "" {
		query = query.Where(
			"fullname ILIKE ? OR phone ILIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	result := query.
		Limit(limit).
		Offset(offset).
		Find(&admins)

	if result.Error != nil {
		return nil, result.Error
	}

	return admins, nil
}

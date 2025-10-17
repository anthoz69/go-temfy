package handlers

import (
	"errors"
	"strconv"

	"github.com/anthoz69/go-temfy/internal/domain/entities"
	"github.com/anthoz69/go-temfy/internal/services"
	"github.com/anthoz69/go-temfy/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService *services.UserService
	validator   *validator.Validate
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6"`
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator.New(),
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "User creation request"
// @Success 201 {object} entities.User "User created successfully"
// @Failure 400 {object} utils.BaseErrorResponse "Invalid request body or validation failed"
// @Failure 409 {object} utils.BaseErrorResponse "User already exists"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", "0001", nil)
	}

	if err := h.validator.Struct(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed: "+err.Error(), "0001", nil)
	}

	user := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.userService.CreateUser(user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusConflict, err.Error(), "0002", nil)
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, user)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a specific user by their ID
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entities.User "User retrieved successfully"
// @Failure 400 {object} utils.BaseErrorResponse "Invalid user ID"
// @Failure 404 {object} utils.BaseErrorResponse "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", "0001", nil)
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found", "0002", nil)
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get user", "0003", nil)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update an existing user by ID
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param request body UpdateUserRequest true "User update request"
// @Success 200 {object} entities.User "User updated successfully"
// @Failure 400 {object} utils.BaseErrorResponse "Invalid user ID or request body"
// @Failure 404 {object} utils.BaseErrorResponse "User not found"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", "0001", nil)
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", "0001", nil)
	}

	if err := h.validator.Struct(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed: "+err.Error(), "0001", nil)
	}

	user := &entities.User{
		ID:    uint(id),
		Name:  req.Name,
		Email: req.Email,
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	if err := h.userService.UpdateUser(user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found", "0002", nil)
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), "0003", nil)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete an existing user by ID
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User deleted successfully"
// @Failure 400 {object} utils.BaseErrorResponse "Invalid user ID"
// @Failure 404 {object} utils.BaseErrorResponse "User not found"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", "0001", nil)
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found", "0002", nil)
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), "0003", nil)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a paginated list of all users
// @Tags users
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param limit query int false "Number of users per page" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} entities.User "Users retrieved successfully"
// @Failure 500 {object} utils.BaseErrorResponse "Failed to get users"
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	users, err := h.userService.GetAllUsers(limit, offset)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get users", "0001", nil)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, users)
}

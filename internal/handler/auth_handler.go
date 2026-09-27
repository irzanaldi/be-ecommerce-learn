package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/irzan-aldi/be-ecommerce/internal/service"
	"github.com/irzan-aldi/be-ecommerce/pkg/response"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	user, err := h.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		if appErr, ok := err.(*response.AppError); ok {
			response.Error(c, appErr.Code, appErr.Message, nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "internal error", nil)
		return
	}

	response.Success(c, http.StatusCreated, "register success", user.ToResponse())
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	tokens, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		if appErr, ok := err.(*response.AppError); ok {
			response.Error(c, appErr.Code, appErr.Message, nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "internal error", nil)
		return
	}

	response.Success(c, http.StatusOK, "login success", tokens)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID := c.GetUint("userID") // di-set oleh JWT middleware

	user, err := h.authService.GetByID(userID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "user not found", nil)
		return
	}

	response.Success(c, http.StatusOK, "success", user.ToResponse())
}

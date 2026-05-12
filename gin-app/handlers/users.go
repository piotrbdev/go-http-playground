package handlers

import (
	"gin-app/models"
	"gin-app/storage"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type UserHandler struct {
	store storage.UsersStore
}

func NewUserHandler(store storage.UsersStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(models.NewBadRequest("invalid request body"))
		return
	}

	if !h.store.IsEmailAvailable(req.Email) {
		c.Error(models.NewConflict("email already exists in DB"))
		return
	}

	user, err := h.store.AddUser(req.Name, req.Email, req.Password)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {

				c.Error(models.NewBadRequest("email already exists"))
				return
			}
		}
		c.Error(models.NewBadRequest(err.Error()))
		return
	}
	c.JSON(
		http.StatusCreated,
		models.UserResponse{Email: user.Email, Name: user.Name},
	)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(models.NewBadRequest("invalid request body"))
		return
	}

	user, err := h.store.MatchUserPassword(req.Email, req.Password)
	if err != nil {
		c.Error(models.NewUnauthorized("bad email or password"))
		return
	}

	jwtClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.Error(models.NewInternal("could not generate token"))
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{
		Token: tokenString,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(models.NewBadRequest("invalid id"))
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(models.NewBadRequest("invalid request body"))
		return
	}

	user, err := h.store.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		c.Error(models.NewNotFound("user not found"))
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(models.NewBadRequest("invalid id"))
		return
	}
	err2 := h.store.DeleteUser(id)
	if err2 != nil {
		c.Error(models.NewNotFound("user not found"))
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "user deleted"})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	filters := make(map[string]string)
	if email := c.Query("email"); email != "" {
		filters["email"] = email
	}
	if name := c.Query("name"); name != "" {
		filters["name"] = name
	}

	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	offset := (page - 1) * limit

	users, err := h.store.GetUsers(limit, offset, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: "database error",
		})
		return
	}
	total, err := h.store.CountUsers(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: "database error",
		})
		return
	}
	response := models.PaginatedUsersResponse{
		Data:  users,
		Page:  page,
		Limit: limit,
		Total: int(total),
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserByName(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, models.Response{Message: "name is required"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "User " + name})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(models.NewBadRequest("invalid id"))
		return
	}

	user, err := h.store.GetUser(id)
	if err != nil {
		c.Error(models.NewNotFound("user not found"))
		return
	}
	c.JSON(http.StatusOK, user)
}

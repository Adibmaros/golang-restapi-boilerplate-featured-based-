package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{
		service: service,
	}
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Get a list of all registered users
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users/ [get]
func (h *controller) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "gagal mengambil data user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil mengambil data user",
		"data":    users,
	})
}

// GetUserByID godoc
// @Summary      Get current user profile
// @Description  Get details of the currently authenticated user
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /users/profile [get]
func (h *controller) GetUserByID(c *gin.Context) {
	userId, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id user tidak ditemukan",
		})
		return
	}

	uid, ok := userId.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id user tidak valid",
		})
		return
	}

	user, err := h.service.GetUserByID(uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "gagal mengambil data user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil mengambil data user",
		"data":    user,
	})
}

// RegisterUser godoc
// @Summary      Register a new user
// @Description  Register a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      RegisterUserDTO         true  "User Registration Payload"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /users/register [post]
func (h *controller) RegisterUser(c *gin.Context) {
	var req RegisterUserDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "tipe data tidak valid",
		})
		return
	}

	user, err := h.service.RegisterUser(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "gagal register data user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil menambahkan data",
		"data":    user,
	})
}

// LoginUser godoc
// @Summary      User Login
// @Description  Authenticate user with email and password to receive JWT token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      LoginUserDTO            true  "User Login Payload"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Failure      401      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /users/login [post]
func (h *controller) LoginUser(c *gin.Context) {
	var req LoginUserDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "tipe data tidak valid",
		})
		return
	}

	token, user, err := h.service.LoginUser(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login berhasil",
		"data": LoginResponseDTO{
			Token: token,
			User:  user,
		},
	})
}

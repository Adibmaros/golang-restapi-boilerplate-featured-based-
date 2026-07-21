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

package product

import (
	"net/http"
	"strconv"

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

// GetAllProducts godoc
// @Summary      Get all products
// @Description  Get a list of all products with creator details
// @Tags         products
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /products/ [get]
func (h *controller) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "gagal mengambil data product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil mengambil data product",
		"data":    products,
	})
}

// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Get details of a specific product by its ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /products/{id} [get]
func (h *controller) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id product tidak valid",
		})
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "product tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil mengambil data product",
		"data":    product,
	})
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product attached to the currently authenticated user
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateProductDTO        true  "Create Product Payload"
// @Success      201      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Failure      401      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /products/ [post]
func (h *controller) CreateProduct(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user tidak terotentikasi",
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

	var req CreateProductDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "tipe data tidak valid",
		})
		return
	}

	product, err := h.service.CreateProduct(uid, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "gagal membuat product baru",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "berhasil menambahkan product baru",
		"data":    product,
	})
}

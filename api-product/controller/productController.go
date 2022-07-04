package controller

import (
	"net/http"
	"strconv"
	"test-azuralabs/api-product/models"
	"test-azuralabs/api-product/service"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{service: service}
}

func (ctr *controller) GetAllProduct(c *gin.Context) {
	products, err := ctr.service.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	var productResponses []models.ProductResponse
	for _, i := range products {
		productResponse := models.ProductResponse{
			ID:          int(i.ID),
			Name:        i.Name,
			Description: i.Description,
		}
		productResponses = append(productResponses, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponses,
	})
}

func (ctr *controller) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	product, err := ctr.service.GetProductByID(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if product.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "product not available",
		})
		return
	}

	productResponse := models.ProductResponse{
		ID:          int(product.ID),
		Name:        product.Name,
		Description: product.Description,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (ctr *controller) PostProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	newProduct, err := ctr.service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	productResponse := models.ProductResponse{
		ID:          int(newProduct.ID),
		Name:        newProduct.Name,
		Description: newProduct.Description,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (ctr *controller) PutProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	productChanged, err := ctr.service.UpdateProduct(idInt, newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	productResponse := models.ProductResponse{
		ID:          idInt,
		Name:        productChanged.Name,
		Description: productChanged.Description,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (ctr *controller) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	if err := ctr.service.DeleteProduct(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meesage": "Success deleted product",
	})
}

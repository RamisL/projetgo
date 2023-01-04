package handler

import (
	"github.com/RamisL/server/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type productResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (th *productHandler) CreateProduct(c *gin.Context) {
	// Get json body
	var input product.InputProduct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := &productResponse{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProduct, err := th.productService.CreateProduct(input)
	if err != nil {
		response := &productResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &productResponse{
		Success: true,
		Message: "New product created",
		Data:    newProduct,
	}
	c.JSON(http.StatusCreated, response)
}
func (th *productHandler) GetAllProduct(c *gin.Context) {
	products, err := th.productService.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &productResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &productResponse{
		Success: true,
		Data:    products,
	})
}

func (th *productHandler) GetByIdProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &productResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	product, err := th.productService.GetByIdProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &productResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &productResponse{
		Success: true,
		Data:    product,
	})
}
func (th *productHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &productResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	// Get json body
	var input product.InputProduct
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := &productResponse{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	uProduct, err := th.productService.UpdateProduct(id, input)
	if err != nil {
		response := &productResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &productResponse{
		Success: true,
		Message: "Update product created",
		Data:    uProduct,
	}
	c.JSON(http.StatusCreated, response)
}
func (th *productHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &productResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	err = th.productService.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &productResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &productResponse{
		Success: true,
		Message: "Product successfully deleted",
	})
}

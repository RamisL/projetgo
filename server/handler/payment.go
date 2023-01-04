package handler

import (
	"net/http"
	"strconv"

	"github.com/RamisL/server/adapter"
	"github.com/RamisL/server/broadcast"
	"github.com/RamisL/server/payment"

	"github.com/gin-gonic/gin"
)

type paymentResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type paymentHandler struct {
	paymentService payment.Service
	broadcaster    broadcast.Broadcaster
}

func NewPaymentHandler(paymentService payment.Service, broadcaster broadcast.Broadcaster) *paymentHandler {
	return &paymentHandler{paymentService, broadcaster}
}

func (th *paymentHandler) CreatePayment(c *gin.Context) {
	// Get json body

	var input payment.InputPayment

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := &paymentResponse{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//var price = db.Select("price").Find(&product).Where("id = ?", input.ProductId)
	newPayment, err := th.paymentService.CreatePayment(input)
	if err != nil {
		response := &paymentResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	th.broadcaster.Submit(adapter.MessageInput{
		Text:  "New payment received",
		Name:  newPayment.Product.Name,
		Price: newPayment.Product.Price,
	})

	response := &paymentResponse{
		Success: true,
		Message: "New payment created",
		Data:    newPayment,
	}
	c.JSON(http.StatusCreated, response)
}
func (th *paymentHandler) GetAllPayment(c *gin.Context) {
	payments, err := th.paymentService.GetAllPayment()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &paymentResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &paymentResponse{
		Success: true,
		Data:    payments,
	})
}
func (th *paymentHandler) GetByIdPayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &paymentResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	payment, err := th.paymentService.GetByIdPayment(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &paymentResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &paymentResponse{
		Success: true,
		Data:    payment,
	})
}
func (th *paymentHandler) UpdatePayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &paymentResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	// Get json body
	var input payment.InputPayment
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := &paymentResponse{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	uPayment, err := th.paymentService.UpdatePayment(id, input)
	if err != nil {
		response := &paymentResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	th.broadcaster.Submit(adapter.MessageInput{
		Text:  "Payment Updated",
		Name:  uPayment.Product.Name,
		Price: uPayment.Product.Price,
	})

	response := &paymentResponse{
		Success: true,
		Message: "Update Payment created",
		Data:    uPayment,
	}
	c.JSON(http.StatusCreated, response)
}
func (th *paymentHandler) DeletePayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &paymentResponse{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	err = th.paymentService.DeletePayment(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &paymentResponse{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &paymentResponse{
		Success: true,
		Message: "Payment successfully deleted",
	})
}

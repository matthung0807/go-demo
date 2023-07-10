package handler

import (
	"context"
	"net/http"

	order "abc.com/demo/internal/usecase/order"
	"github.com/gin-gonic/gin"
)

type CreateOrderHandler struct {
	usecase *order.CreateUseCase
}

func NewCreateOrderHandler(usecase *order.CreateUseCase) *CreateOrderHandler {
	return &CreateOrderHandler{
		usecase: usecase,
	}
}

// @Tags Order
// @Success 203 {string} string
// @Router /order [post]
// @Param X-Correlation-Id header string true "uuid"
func (h *CreateOrderHandler) Exec(c *gin.Context) {
	corId := c.Request.Header.Get("X-Correlation-Id")
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "corId", corId)
	err := h.usecase.Exec(ctx)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusAccepted, nil)
}

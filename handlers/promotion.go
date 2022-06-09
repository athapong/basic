package handlers

import (
	"basic/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promoService services.PromotionService
}

func NewPromotionHandler(promoService services.PromotionService) PromotionHandler {
	return &promotionHandler{promoService: promoService}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	// http://localhost:8000/calculate?amount=100
	amount := c.Query("amount")
	amoun, err := strconv.Atoi(amount)
	if err != nil {
		return c.Status(400).SendString("invalid amount")
	}

	discount, err := h.promoService.CalculateDiscount(amoun)
	if err != nil {
		return c.Status(500).SendString("internal server error")
	}
	return c.JSON(map[string]int{"discount": discount})
}

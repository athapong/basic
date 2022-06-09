//go:build integration

package handlers_test

import (
	"basic/handlers"
	"basic/repositories"
	"basic/services"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := `{"discount":80}`
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     amount,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)
		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		resp, _ := app.Test(req)

		if assert.Equal(t, fiber.StatusOK, resp.StatusCode) {
			body, _ := io.ReadAll(resp.Body)
			assert.JSONEq(t, expected, string(body))
		}
	})

}

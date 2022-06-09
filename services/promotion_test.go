package services_test

import (
	"basic/repositories"
	"basic/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	type testCases struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCases{
		{name: "invalid amount", purchaseMin: 100, discountPercent: 10, amount: 0, expected: 0},
		{name: "success discount", purchaseMin: 100, discountPercent: 10, amount: 100, expected: 90},
		{name: "success discount", purchaseMin: 200, discountPercent: 12, amount: 200, expected: 176},
		{name: "success discount", purchaseMin: 300, discountPercent: 15, amount: 300, expected: 255},
		{name: "success discount", purchaseMin: 400, discountPercent: 20, amount: 400, expected: 320},
		{name: "success discount", purchaseMin: 500, discountPercent: 25, amount: 500, expected: 375},
		{name: "success discount", purchaseMin: 500, discountPercent: 25, amount: 90, expected: 90},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Arrange
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(
				repositories.Promotion{ID: 1, PurchaseMin: c.purchaseMin, DiscountPercent: c.discountPercent},
				nil)
			promoService := services.NewPromotionService(promoRepo)

			// Act
			expected, err := promoService.CalculateDiscount(c.amount)
			if err != nil {
				assert.ErrorIs(t, err, services.ErrInvalidAmount)
			} else {
				// Assert
				assert.Equal(t, c.expected, expected)
			}

		})
	}

	t.Run("error repository", func(t *testing.T) {
		// Arrange
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(
			repositories.Promotion{ID: 1, PurchaseMin: 100, DiscountPercent: 10},
			errors.New("error"))

		promoService := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoService.CalculateDiscount(100)
		assert.ErrorIs(t, err, services.ErrRepository)
	})

}

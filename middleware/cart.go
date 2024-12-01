package middleware

import (
	"coba/model"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Cart() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("user").(model.User)
		random := rand.New(rand.NewSource(time.Now().Unix()))
		var cart model.Cart

		if err := model.DB.Where("user_id =? ", userId.ID).First(&cart).Error; err != nil {

			randomID := random.Uint32()

			cart = model.Cart{
				ID: uint(randomID),
				UserID: uint(userId.ID),
				CreatedAt: time.Now(),
			}

			model.DB.Create(&cart)

			c.Locals("cartID", randomID)
			return c.Next()

		}

		c.Locals("cartID",cart.ID)
		return c.Next()
	}
}
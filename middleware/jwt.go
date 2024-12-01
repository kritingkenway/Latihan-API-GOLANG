package middleware

import (
	"coba/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTProtectedRoute() fiber.Handler {
	return func (c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status" : false,
				"message": "There is No Required Token",
			})
		}

		tokenString := authHeader[len("Bearer "):]

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": false,
				"message" : "Empty Token",
			})
		}

		var userToken model.UserToken

		if err := model.DB.Where("token = ?", tokenString).First(&userToken).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status" : false,
				"message": "Token is Invalid",
			})
		}

		token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface {}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil , jwt.NewValidationError("unexpected Signing Method", jwt.ValidationErrorSignatureInvalid)
			}

			return []byte("SuperRahasiaJWT"), nil

		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status" : false,
				"message" : "Invalid Token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := uint(claims["user_id"].(float64))
			var user model.User
			if err := model.DB.Where("id = ?", userId).First(&user).Error; err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status" : false,
					"message" : "Invalid Token",
				})
			}

			c.Locals("user",user)

			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status" : false,
			"message" : "Invalid Token",
		})
	}
}
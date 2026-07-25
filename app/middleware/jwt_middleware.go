package middleware

import (
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/iskandar221201/goigniter/system"
)

func JWTProtected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if len(auth) < 7 || auth[:7] != "Bearer " {
			return system.Error(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		token := auth[7:]
		claims, err := system.ParseToken(token, secret)
		if err != nil {
			return system.Error(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		c.Locals("user", claims)
		return c.Next()
	}
}

func RoleGuard(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*system.JWTClaims)
		if !ok {
			return system.Error(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		if slices.Contains(roles, claims.Role) {
			return c.Next()
		}

		return system.Error(c, fiber.StatusForbidden, "Forbidden")
	}
}

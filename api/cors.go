package api

import (
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func sameDomainCORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		origin := c.Get(fiber.HeaderOrigin)
		if origin == "" {
			return c.Next()
		}

		originURL, err := url.Parse(origin)
		if err != nil {
			return c.Next()
		}

		originHost := originURL.Hostname()
		requestHost := c.Hostname()

		if strings.EqualFold(originHost, requestHost) || isSameBaseDomain(originHost, requestHost) {
			c.Set(fiber.HeaderAccessControlAllowOrigin, origin)
			c.Set(fiber.HeaderAccessControlAllowCredentials, "true")
			c.Set(fiber.HeaderAccessControlAllowMethods, "GET,POST,PUT,PATCH,DELETE,OPTIONS,HEAD")
			c.Set(fiber.HeaderAccessControlAllowHeaders, "Origin,Content-Type,Accept,Authorization,Cache-Control")

			if c.Method() == fiber.MethodOptions {
				c.Set(fiber.HeaderAccessControlMaxAge, "86400")
				return c.SendStatus(fiber.StatusNoContent)
			}
		}

		return c.Next()
	}
}

func isSameBaseDomain(host1, host2 string) bool {
	parts1 := strings.Split(host1, ".")
	parts2 := strings.Split(host2, ".")

	if len(parts1) < 2 || len(parts2) < 2 {
		return false
	}

	base1 := strings.Join(parts1[len(parts1)-2:], ".")
	base2 := strings.Join(parts2[len(parts2)-2:], ".")

	return strings.EqualFold(base1, base2)
}

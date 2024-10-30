package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(ctx *fiber.Ctx, name, token string) {
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Secure = true
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour * 7)
	cookie.MaxAge = 60 * 60 * 24 * 7
	cookie.Path = "/"
	ctx.Cookie(cookie)
}

func GetCookie(ctx *fiber.Ctx) (*fiber.Cookie, error) {
	c := new(fiber.Cookie)
	if err := ctx.CookieParser(c); err != nil {
		return nil, err
	}
	return c, nil
}

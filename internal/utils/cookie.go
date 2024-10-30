package utils

import (
	"github.com/gofiber/fiber/v2"
)

func SetCookie(ctx *fiber.Ctx, name, token string) {
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.MaxAge = 60 * 60 * 24 * 7
	ctx.Cookie(cookie)
}

func GetCookie(ctx *fiber.Ctx) (*fiber.Cookie, error) {
	// value := ctx.Cookies(name)
	// if value == "" {
	// 	return "", fiber.ErrUnauthorized
	// }
	// return value, nil

	c := new(fiber.Cookie)
	if err := ctx.CookieParser(c); err != nil {
		return nil, err
	}
	return c, nil
}

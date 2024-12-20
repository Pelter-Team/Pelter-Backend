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

func ClearCookie(ctx *fiber.Ctx, name string) {
	cookie := new(fiber.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.Secure = true
	cookie.HTTPOnly = true
	cookie.Expires = time.Now()
	cookie.MaxAge = -1
	cookie.Path = "/"
	ctx.Cookie(cookie)
}

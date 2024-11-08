package api

import (
	"github.com/gofiber/fiber/v2"
)

// ProductHandler  an interface with operations to be implemented by a specific handler, ie http, gRCP
type ProductHandler interface {
	Login(ctx *fiber.Ctx)
	Get(ctx *fiber.Ctx)
	Post(ctx *fiber.Ctx)
	Put(ctx *fiber.Ctx)
	Delete(ctx *fiber.Ctx)
	GetAll(ctx *fiber.Ctx)
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nikitamirzani323/btangkas-client/controller"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(compress.New())
	// Custom config
	app.Static("/", "frontend/dist", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("api/healthz", controller.HealthCheck)
	app.Post("api/checktoken", controller.CheckToken)
	app.Post("api/listinvoice", controller.Listinvoice)
	app.Post("api/listinvoicedetail", controller.Listinvoicedetail)
	app.Post("api/savetransaksi", controller.SaveTransaksi)
	app.Post("api/savetransaksidetail", controller.SaveTransaksiDetail)
	return app
}

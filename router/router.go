package router

import (
	"bufio"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nikitamirzani323/btangkas-client/controller"
	"github.com/valyala/fasthttp"
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

	app.Get("/sse", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")
			var i int
			id_invoice := 1
			time_data := 30
			time_sleep := 20
			status := "LOCK"
			invoice := "20240301"
			invoice = invoice + strconv.Itoa(id_invoice)
			for {
				i++
				msg := ""
				if time_data < 1 {
					status = "LOCK"
					time_sleep = time_sleep - 1
					if time_sleep < 1 {
						time_data = 30
						time_sleep = 20
						id_invoice = id_invoice + 1
						invoice = invoice + strconv.Itoa(id_invoice)
					}
					msg = fmt.Sprintf("%s|%s|%s", strconv.Itoa(time_data), invoice, status)
				} else {
					status = "OPEN"
					msg = fmt.Sprintf("%s|%s|%s", strconv.Itoa(time_data), invoice, status)
					time_data = time_data - 1

				}
				fmt.Fprintf(w, "data: Message: %s\n\n", msg)
				fmt.Println(msg)
				err := w.Flush()
				if err != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)

					break
				}
				time.Sleep(2 * time.Second)
			}
		}))

		return nil
	})
	return app
}

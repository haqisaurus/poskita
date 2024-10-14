package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggingRoute(app *fiber.App) {

	app.Use(logger.New(logger.Config{
		TimeFormat: time.RFC1123Z,
		TimeZone: "Asia/Jakarta",
		Format:   `>>>>>>>>>>> incoming >>>>>>>>>>>
${time} | ${method} | ${status} | ${path} | ${queryParams}
Headers: ${reqHeaders}
======== REQUEST ${method} ========
${body}
======== RESPONSE =======
${resBody}
########   END   ########
`,
	}))

}


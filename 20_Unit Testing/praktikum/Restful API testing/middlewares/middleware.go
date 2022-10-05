package middlewares

import (
	"echo-book/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human} \n",
	}))

}

func CheckTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := auth.ExtractToken(c)

		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{

				"message": "invalid token",
			})

		}
		return next(c)
	}
}

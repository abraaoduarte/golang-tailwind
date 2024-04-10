package routes

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func RegisterRouters(db *pgx.Conn, router *echo.Echo) {
	ur := UserRouter{db}

	ur.SetRoutes(router)
}

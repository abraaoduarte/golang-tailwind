package routes

import (
	"example.com/resume/internal/handler"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type UserRouter struct {
	db *pgx.Conn
}

func NewUserRouter(db *pgx.Conn) *UserRouter {
	return &UserRouter{db}
}

func (ur UserRouter) SetRoutes(router *echo.Echo) {
	user := router.Group("/users")
	uh := handler.NewUserHandler(ur.db)

    user.GET("/", uh.GetUsers)
}

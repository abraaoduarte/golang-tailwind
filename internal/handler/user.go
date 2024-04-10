package handler

import (
	userTempl "example.com/resume/internal/templates/user"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	db *pgx.Conn
}

func NewUserHandler(db *pgx.Conn) *UserHandler {
	return &UserHandler{db}
}

func (uh *UserHandler) GetUsers(context echo.Context) error {
	return renderTempl(context, userTempl.Show())
}

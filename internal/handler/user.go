package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	errorTempl "example.com/resume/internal/templates/errors"
	userTempl "example.com/resume/internal/templates/user"
	"example.com/resume/internal/types"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	db *pgx.Conn
}

func NewUserHandler(db *pgx.Conn) *UserHandler {
	return &UserHandler{db}
}

func (uh *UserHandler) Index(ctx echo.Context) error {
	row, err := uh.db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return renderTempl(ctx, errorTempl.UnknownError())
	}
	defer row.Close()

	users := []types.User{}
	for row.Next() {
		u := types.User{}
		err = row.Scan(&u.Id, &u.Name, &u.Email, &u.Birthdate, &u.IsAdmin, &u.CreateAt)
		if err != nil {
			return renderTempl(ctx, errorTempl.UnknownError())
		}
		users = append(users, u)
	}

	return renderTempl(ctx, userTempl.ListUsers(users))
}

func (uh *UserHandler) Create(ctx echo.Context) error {
	return renderTempl(ctx, userTempl.Create())
}

func (uh *UserHandler) Store(ctx echo.Context) error {
	tx, err := uh.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var objRequest types.User

	date, _ := time.Parse("2006-01-02", ctx.FormValue("birthdate"))

	if err := ctx.Bind(&objRequest); err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, "")
	}

	objRequest.Birthdate = date
	objRequest.CreateAt = time.Now()

	query := "INSERT INTO users (name, email, birthdate, is_admin, created_at) VALUES (@name, @email, @birthdate, @is_admin, @created_at)"
	args := pgx.NamedArgs{
		"name":       objRequest.Name,
		"email":      objRequest.Email,
		"birthdate":  objRequest.Birthdate,
		"is_admin":   objRequest.IsAdmin,
		"created_at": objRequest.CreateAt,
	}
	_, err = tx.Exec(context.Background(), query, args)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, "Something happened")
	}

	tx.Commit(context.Background())

	return ctx.Redirect(http.StatusFound, "/users/")
}

func (uh *UserHandler) Edit(ctx echo.Context) error {
	query := "SELECT * FROM users WHERE id=$1"

	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return renderTempl(ctx, errorTempl.UnknownError())
	}

	user := types.User{}
	err = uh.db.QueryRow(context.Background(), query, userID).Scan(&user.Id, &user.Name, &user.Email, &user.Birthdate, &user.IsAdmin, &user.CreateAt)
	if err != nil {
		return renderTempl(ctx, errorTempl.UnknownError())
	}

	return renderTempl(ctx, userTempl.Edit(user))
}

func (uh *UserHandler) Update(ctx echo.Context) error {
	tx, err := uh.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var objRequest types.User

	date, _ := time.Parse("2006-01-02", ctx.FormValue("birthdate"))

	if err := ctx.Bind(&objRequest); err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, "")
	}

	objRequest.Birthdate = date

	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return renderTempl(ctx, errorTempl.UnknownError())
	}

	query := "UPDATE users SET name=@name, email=@email, birthdate=@birthdate, is_admin=@is_admin WHERE id=@id"
	args := pgx.NamedArgs{
		"name":      objRequest.Name,
		"email":     objRequest.Email,
		"birthdate": objRequest.Birthdate,
		"is_admin":  objRequest.IsAdmin,
		"id":        userID,
	}
	_, err = tx.Exec(context.Background(), query, args)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, "Something happened")
	}

	tx.Commit(context.Background())

	return ctx.Redirect(http.StatusFound, "/users/")
}

func (uh *UserHandler) Delete(ctx echo.Context) error {
	tx, err := uh.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	userID, err := strconv.Atoi(ctx.Param("id"))

	if err != err {
		return renderTempl(ctx, errorTempl.UnknownError())
	}

	query := "DELETE FROM users WHERE id=$1"
	_, err = uh.db.Exec(context.Background(), query, userID)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, "Something wrong happened")
	}
	tx.Commit(context.Background())

	return ctx.Redirect(http.StatusFound, "/users/")
}

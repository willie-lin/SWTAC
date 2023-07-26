package handler

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/account"
	"SWTAC/utils"
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterForm struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
}

func Register(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var form RegisterForm
		if err := c.Bind(&form); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		// 开启一个事务
		tx, err := client.Tx(context.Background())
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		//
		pwd, err := utils.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		form.Password = string(pwd)
		// 在 Account 表中创建一个新账户，并将其与新用户关联起来
		a, err := tx.Account.
			Create().
			SetUsername(form.Username).
			SetEmail(form.Email).
			SetPhone(form.Phone).
			SetPassword(form.Password).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if ent.IsConstraintError(err) {
			tx.Rollback()
			return c.JSON(http.StatusConflict, err.Error())
		}
		if err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		//在 User 表中创建一个新用户
		u, err := tx.User.Create().
			SetNickname(form.Nickname).
			AddAccounts(a).
			AddGroups().
			AddRoles().
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if ent.IsConstraintError(err) {
			tx.Rollback()
			return c.JSON(http.StatusConflict, err.Error())
		}
		if err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		// 提交事务
		if err := tx.Commit(); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, u)
	}
}

type LoginRequest struct {
	UsernameOrEmailOrPhone string `json:"usernameOrEmailOrPhone" form:"usernameOrEmailOrPhone" query:"usernameOrEmailOrPhone"`
	Password               string `json:"password" form:"password" query:"password"`
}

func Login(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var req LoginRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		//在这里，您可以使用Ent框架提供的API来验证用户名、电子邮件或电话号码和密码
		//例如：
		account, err := client.Account.
			Query().
			Where(
				account.Or(
					account.Username(req.UsernameOrEmailOrPhone),
					account.Email(req.UsernameOrEmailOrPhone),
					account.Phone(req.UsernameOrEmailOrPhone),
				),
			).
			Where(account.Password(req.Password)).
			Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if err != nil {
			return echo.ErrUnauthorized
		}

		//return c.JSON(http.StatusOK, map[string]string{"message": "Logged in successfully"})
		return c.JSON(http.StatusOK, account)
	}
}

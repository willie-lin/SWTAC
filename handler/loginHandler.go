package handler

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/account"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

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

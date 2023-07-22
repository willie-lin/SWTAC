package handler

import (
	"SWTAC/datasource/ent"
	_ "SWTAC/datasource/ent"
	"SWTAC/utils"
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

//type RegisterForm struct {
//	Username string `form:"username"`
//	Password string `form:"password"`
//	Email    string `form:"email"`
//	Name     string `form:"name"`
//	Address  string `form:"address"`
//	Phone    string `form:"phone"`
//}
//
//func handleRegister(username, password, email, name, address, phone string) error {
//	// 创建新账户
//	account, err := client.Account.
//		Create().
//		SetUsername(username).
//		SetPassword(password).
//		SetEmail(email).
//		Save(context.Background())
//	if err != nil {
//		return err
//	}
//
//	// 创建新用户
//	_, err = client.User.
//		Create().
//		SetName(name).
//		SetAddress(address).
//		SetPhone(phone).
//		AddAccounts(account).
//		Save(context.Background())
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//e.POST("/register", func(c echo.Context) error {
//	var form RegisterForm
//	if err := c.Bind(&form); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{
//			"message": err.Error(),
//		})
//	}
//
//	err := handleRegister(form.Username, form.Password, form.Email, form.Name, form.Address, form.Phone)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{
//			"message": err.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]string{
//		"message": "Registered successfully",
//	})
//})
//
//e.POST("/register", func(c echo.Context) error {
//	var form RegisterForm
//	if err := c.Bind(&form); err != nil {
//		return c.String(http.StatusBadRequest, err.Error())
//	}
//
//	err := handleRegister(form.Username, form.Password, form.Email, form.Name, form.Address, form.Phone)
//	if err != nil {
//		return c.String(http.StatusInternalServerError, err.Error())
//	}
//
//	return c.String(http.StatusOK, "Registered successfully")
//})
//
//
//func handleUpdate(name, address, phone, editor string) error {
//	// 更新用户
//	user, err := client.User.
//		UpdateOneID(1). // 假设更新 ID 为 1 的用户
//		SetName(name).
//		SetAddress(address).
//		SetPhone(phone).
//		SetEditor(editor). // 设置编辑者
//		Save(context.Background())
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//e.POST("/register", func(c echo.Context) error {
//	var form RegisterForm
//	if err := c.Bind(&form); err != nil {
//		return c.String(http.StatusBadRequest, err.Error())
//	}
//
//	err := handleRegister(form.Username, form.Password, form.Email, form.Name, form.Address, form.Phone)
//	if err != nil {
//		return c.String(http.StatusInternalServerError, err.Error())
//	}
//
//	return c.String(http.StatusOK, "Registered successfully")
//})
//
//
//e.POST("/update", func(c echo.Context) error {
//	var form UpdateForm
//	if err := c.Bind(&form); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{
//			"message": err.Error(),
//		})
//	}
//
//	// 获取当前登录账户的用户名
//	editor := getCurrentUsername(c)
//
//	err := handleUpdate(form.Name, form.Address, form.Phone, editor)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{
//			"message": err.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, map[string]string{
//		"message": "Updated successfully",
//	})
//})

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
		u, err := tx.User.
			Create().SetNickname(form.Nickname).
			AddAccount(a).
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

package handler

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/user"
	"SWTAC/utils"
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// GetAllUsers 获取所有用户
func GetAllUsers(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := client.User.Query().All(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, users)
	}
}

// GetUserByUsername 根据用户名查找
func GetUserByUsername(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)
		//直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.UsernameEQ(u.Username)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// GetUserById  根据ID查找
func GetUserById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.IDEQ(u.ID)).Only(context.Background())

		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// GetUserByEmail 根据email 查找用户
func GetUserByEmail(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {

		u := new(ent.User)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		user, err := client.User.Query().Where(user.EmailEQ(u.Email)).Only(context.Background())

		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// GetUserByPhone  根据手机号码查找用户
func GetUserByPhone(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.PhoneEQ(u.Phone)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// CreateUser 创建用户
func CreateUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		u := new(ent.User)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		pwd, err := utils.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		u.Password = string(pwd)

		user, err := client.User.Create().
			SetUsername(u.Username).
			SetPassword(u.Password).
			SetEmail(u.Email).
			SetNickname(u.Nickname).
			SetCity(u.City).
			SetAvatar(u.Avatar).
			SetPhone(u.Phone).
			SetAge(u.Age).
			SetIntroduction(u.Introduction).
			SetState(u.State).
			SetCreator(u.Creator).
			SetEditor(u.Editor).
			SetDeleted(u.Deleted).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if ent.IsConstraintError(err) {
			return c.JSON(http.StatusConflict, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, user)
	}
}

// UpdateUserById  更新用户
func UpdateUserById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.UpdateOneID(u.ID).
			SetEmail(u.Email).
			SetNickname(u.Nickname).
			SetCity(u.City).
			SetAvatar(u.Avatar).
			SetPhone(u.Phone).
			SetAge(u.Age).
			SetIntroduction(u.Introduction).
			SetState(u.State).
			SetCreator(u.Creator).
			SetEditor(u.Editor).
			SetDeleted(u.Deleted).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if ent.IsConstraintError(err) {
			return c.JSON(http.StatusConflict, err.Error())
		}
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// UpdateUser UpdateOneUser 更新用户
func UpdateUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}

		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.IDEQ(u.ID)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		user, err = user.Update().
			SetEmail(u.Email).
			SetNickname(u.Nickname).
			SetCity(u.City).
			SetAvatar(u.Avatar).
			SetPhone(u.Phone).
			SetAge(u.Age).
			SetIntroduction(u.Introduction).
			SetState(u.State).
			SetCreator(u.Creator).
			SetEditor(u.Editor).
			SetDeleted(u.Deleted).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// DeleteUser 删除用户
func DeleteUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)

		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.UsernameEQ(u.Username)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		err = client.User.DeleteOne(user).Exec(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}

// DeleteUserById 删除用户
func DeleteUserById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)

		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err := client.User.DeleteOneID(u.ID).Exec(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}

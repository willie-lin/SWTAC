package handler

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/usergroup"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// GetAllUserGroups  获取所有用户组
func GetAllUserGroups(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		supergroups, err := client.UserGroup.Query().All(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, supergroups)
	}
}

// GetUserGroupByName  根据用户组名查找
func GetUserGroupByName(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		ug := new(ent.UserGroup)
		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&ug); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user_group, err := client.UserGroup.Query().Where(usergroup.NameEQ(ug.Name)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, user_group)
	}
}

// GetUserGroupById   根据ID查找
func GetUserGroupById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		ug := new(ent.UserGroup)
		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&ug); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		usergroup, err := client.UserGroup.Query().Where(usergroup.IDEQ(ug.ID)).Only(context.Background())

		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, usergroup)
	}
}

// CreateUserGroup  创建用户组
func CreateUserGroup(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		ug := new(ent.UserGroup)

		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&ug); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		//
		//pwd, err := utils.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		//if err != nil {
		//	//fmt.Println("加密密码失败", err)
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		//fmt.Println(pwd)
		//u.Password = string(pwd)
		//fmt.Println(pwd)

		usergroup, err := client.UserGroup.Create().
			SetName(ug.Name).
			SetCreator(ug.Creator).
			SetEditor(ug.Editor).
			SetDeleted(ug.Deleted).
			SetCode(ug.Code).
			SetIntro(ug.Intro).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())

		if ent.IsConstraintError(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, usergroup)
	}
}

// UpdateUserGroupById 更新用户组
func UpdateUserGroupById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		ug := new(ent.UserGroup)
		// 解析json 并绑定到u
		if err := json.NewDecoder(c.Request().Body).Decode(&ug); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		usergroup, err := client.UserGroup.UpdateOneID(ug.ID).
			SetName(ug.Name).
			SetCreator(ug.Creator).
			SetEditor(ug.Editor).
			SetDeleted(ug.Deleted).
			SetCode(ug.Code).
			SetIntro(ug.Intro).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if ent.IsConstraintError(err) {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, usergroup)
	}
}

// UpdateUserGroup   更新用户组
func UpdateUserGroup(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		ug := new(ent.UserGroup)
		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&ug); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		usergroup, err := client.UserGroup.Query().Where(usergroup.IDEQ(ug.ID)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		usergroup, err = usergroup.Update().
			SetName(ug.Name).
			SetCreator(ug.Creator).
			SetEditor(ug.Editor).
			SetDeleted(ug.Deleted).
			SetCode(ug.Code).
			SetIntro(ug.Intro).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		return c.JSON(http.StatusOK, usergroup)
	}
}

// DeleteUserGroup  删除用户组
func DeleteUserGroup(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		ug := new(ent.UserGroup)

		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&ug); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		usergroup, err := client.UserGroup.Query().Where(usergroup.NameEQ(ug.Name)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err = client.UserGroup.DeleteOne(usergroup).Exec(context.Background())
		return c.NoContent(http.StatusOK)
	}
}

// DeleteUserGroupById 删除用户组
func DeleteUserGroupById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		ug := new(ent.UserGroup)

		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&ug); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err := client.UserGroup.DeleteOneID(ug.ID).Exec(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

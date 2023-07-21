package handler

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/group"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// GetAllGroups   获取所有组
func GetAllGroups(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		groups, err := client.Group.Query().All(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, groups)
	}
}

// GetGroupByName    根据组名查找
func GetGroupByName(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		g := new(ent.Group)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&g); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}

		if err := c.Bind(g); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		group, err := client.Group.Query().Where(group.NameEQ(g.Name)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, group)
	}
}

// GetGroupById   根据ID查找
func GetGroupById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		g := new(ent.Group)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&g); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}

		if err := c.Bind(g); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		group, err := client.Group.Query().Where(group.IDEQ(g.ID)).Only(context.Background())

		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, group)
	}
}

// CreateGroup  创建组
func CreateGroup(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		g := new(ent.Group)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&g); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(g); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		//pwd, err := utils.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		//if err != nil {
		//	//fmt.Println("加密密码失败", err)
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		//fmt.Println(pwd)
		//u.Password = string(pwd)
		//fmt.Println(pwd)

		group, err := client.Group.Create().
			SetName(g.Name).
			SetCreator(g.Creator).
			SetEditor(g.Editor).
			SetDeleted(g.Deleted).
			SetCode(g.Code).
			SetIntro(g.Intro).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if ent.IsConstraintError(err) {
			return c.JSON(http.StatusConflict, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, group)
	}
}

// UpdateGroupById  更新用户组
func UpdateGroupById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		g := new(ent.Group)
		// 解析json 并绑定到u
		//if err := json.NewDecoder(c.Request().Body).Decode(&g); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(g); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		group, err := client.Group.UpdateOneID(g.ID).
			SetName(g.Name).
			SetCreator(g.Creator).
			SetEditor(g.Editor).
			SetDeleted(g.Deleted).
			SetCode(g.Code).
			SetIntro(g.Intro).
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
		return c.JSON(http.StatusOK, group)
	}
}

// UpdateGroup    更新用户组
func UpdateGroup(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		g := new(ent.Group)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&g); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}

		if err := c.Bind(g); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		group, err := client.Group.Query().Where(group.IDEQ(g.ID)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		group, err = group.Update().
			SetName(g.Name).
			SetCreator(g.Creator).
			SetEditor(g.Editor).
			SetDeleted(g.Deleted).
			SetCode(g.Code).
			SetIntro(g.Intro).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, group)
	}
}

// DeleteGroup DeleteUserGroup  删除用户组
func DeleteGroup(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		g := new(ent.Group)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&g); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}

		if err := c.Bind(g); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		group, err := client.Group.Query().Where(group.NameEQ(g.Name)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		err = client.Group.DeleteOne(group).Exec(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

// DeleteGroupById  删除用户组
func DeleteGroupById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		g := new(ent.Group)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&g); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(g); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err := client.Group.DeleteOneID(g.ID).Exec(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

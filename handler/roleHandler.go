package handler

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/role"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// GetAllRoles 获取所有角色
func GetAllRoles(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		roles, err := client.Role.Query().All(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, roles)
	}
}

// GetRoleByRoleName 根据角色名查找
func GetRoleByRoleName(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(ent.Role)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		role, err := client.Role.Query().Where(role.NameEQ(r.Name)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, role)
	}
}

// GetRoleById   根据ID查找
func GetRoleById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(ent.Role)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		role, err := client.Role.Query().Where(role.IDEQ(r.ID)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, role)
	}
}

// CreateRole  创建角色
func CreateRole(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		r := new(ent.Role)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}

		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		role, err := client.Role.Create().
			SetName(r.Name).
			SetIntro(r.Intro).
			SetCode(r.Code).
			SetCreator(r.Creator).
			SetEditor(r.Editor).
			SetDeleted(r.Deleted).
			SetParentID(r.ParentID).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if ent.IsConstraintError(err) {
			return c.JSON(http.StatusConflict, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, role)
	}
}

// UpdateRoleById   更新角色
func UpdateRoleById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(ent.Role)
		// 解析json 并绑定到
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		role, err := client.Role.UpdateOneID(r.ID).
			SetName(r.Name).
			SetIntro(r.Intro).
			SetCode(r.Code).
			SetCreator(r.Creator).
			SetEditor(r.Editor).
			SetDeleted(r.Deleted).
			SetParentID(r.ParentID).
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
		return c.JSON(http.StatusOK, role)
	}
}

// UpdateRole  更新角色
func UpdateRole(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(ent.Role)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}

		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		role, err := client.Role.Query().Where(role.IDEQ(r.ID)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		role, err = role.Update().
			SetName(r.Name).
			SetIntro(r.Intro).
			SetCode(r.Code).
			SetCreator(r.Creator).
			SetEditor(r.Editor).
			SetDeleted(r.Deleted).
			SetParentID(r.ParentID).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, role)
	}

}

// DeleteRole  删除角色
func DeleteRole(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(ent.Role)

		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}

		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		role, err := client.Role.Query().Where(role.NameEQ(r.Name)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		err = client.Role.DeleteOne(role).Exec(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

// DeleteRoleById  删除角色
func DeleteRoleById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(ent.Role)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err := client.Role.DeleteOneID(r.ID).Exec(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

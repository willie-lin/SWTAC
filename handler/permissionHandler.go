package handler

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/permission"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// GetAllPermissions 获取所有角色
func GetAllPermissions(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		permissions, err := client.Permission.Query().All(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, permissions)
	}
}

// GetPermissionByName 根据角色名查找
func GetPermissionByName(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(ent.Permission)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		permission, err := client.Permission.Query().Where(permission.NameEQ(p.Name)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, permission)
	}
}

// GetPermissionById GetRoleById   根据ID查找
func GetPermissionById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(ent.Permission)
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		permission, err := client.Permission.Query().Where(permission.IDEQ(p.ID)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, permission)
	}
}

// CreatePermission   创建角色
func CreatePermission(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		p := new(ent.Permission)

		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		permission, err := client.Permission.Create().
			SetName(p.Name).
			SetDescription(p.Description).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if ent.IsConstraintError(err) {
			return c.JSON(http.StatusConflict, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, permission)
	}
}

// UpdatePermissionById  更新角色
func UpdatePermissionById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(ent.Permission)

		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		permission, err := client.Role.UpdateOneID(p.ID).
			SetName(p.Name).
			SetDescription(p.Description).
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
		return c.JSON(http.StatusOK, permission)
	}
}

// UpdatePermission  更新角色
func UpdatePermission(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(ent.Permission)

		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		permission, err := client.Permission.Query().Where(permission.IDEQ(p.ID)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		permission, err = permission.Update().
			SetName(p.Name).
			SetDescription(p.Description).
			SetUpdatedAt(time.Now()).
			Save(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, permission)
	}

}

// DeletePermission   删除角色
func DeletePermission(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(ent.Permission)

		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		permission, err := client.Permission.Query().Where(permission.NameEQ(p.Name)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		err = client.Permission.DeleteOne(permission).Exec(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

// DeletePermissionById   删除角色
func DeletePermissionById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(ent.Permission)
		// 直接解析raw数据为json
		//if err := json.NewDecoder(c.Request().Body).Decode(&r); err != nil {
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err := client.Permission.DeleteOneID(p.ID).Exec(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusOK)
	}
}

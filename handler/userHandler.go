package handler

import (
	"SWTAC/datasource/ent"
	"SWTAC/datasource/ent/user"
	"SWTAC/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// GetAllUsers 获取所有用户
func GetAllUsers(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := client.User.Query().All(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, users)
	}
}

// GetUserByUsername 根据用户名查找
func GetUserByUsername(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)
		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.UsernameEQ(u.Username)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// GetUserById  根据ID查找
func GetUserById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)
		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.IDEQ(u.ID)).Only(context.Background())

		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())

		}
		return c.JSON(http.StatusOK, user)
	}
}

// GetUserByEmail 根据email 查找用户
func GetUserByEmail(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {

		u := new(ent.User)

		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.EmailEQ(u.Email)).Only(context.Background())

		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())

		}
		return c.JSON(http.StatusOK, user)
	}
}

// CreateUser 创建用户
func CreateUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		u := new(ent.User)

		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		pwd, err := utils.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			//fmt.Println("加密密码失败", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		//fmt.Println(pwd)
		u.Password = string(pwd)
		//fmt.Println(pwd)

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
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusCreated, user)
	}
}

// UpdateUserById  更新用户
func UpdateUserById(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)
		// 解析json 并绑定到u
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
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
			c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}

// UpdateUser UpdateOneUser 更新用户
func UpdateUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(ent.User)

		// 直接解析raw数据为json
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := client.User.Query().Where(user.IDEQ(u.ID)).Only(context.Background())
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusBadRequest, err.Error())
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
		return c.JSON(http.StatusOK, user)
	}

}

func DeleteUser(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		//client, err := database.Client()
		//client, err := config.NewClient()
		//if err != nil {
		//	panic(err)
		//	return err
		//}
		u := new(ent.User)

		//// 接收raw数据
		//result, err := ioutil.ReadAll(c.Request().Body)
		//if err != nil {
		//	fmt.Println("ioutil.ReadAll err:", err)
		//	return err
		//}
		//// 解析raw为json
		//err = json.Unmarshal(result, &u)
		//if err != nil {
		//	fmt.Println("json.Unmarshal err:", err)
		//	return err
		//}

		// 直接解析raw数据为json
		log, _ := zap.NewProduction()
		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return err
		}
		//// or for DisallowUnknownFields() wrapped in your custom func
		//decoder := json.NewDecoder(c.Request().Body)
		//decoder.DisallowUnknownFields()
		//if err := decoder.Decode(&payload); err != nil {
		//	return err
		//}

		fmt.Println(1111)
		fmt.Println(u.Username)
		fmt.Println(22222)
		us, err := client.User.Query().Where(user.UsernameEQ(u.Username)).Only(context.Background())
		if err != nil {
			//panic(err)
			//log.Fatal("Query user error:", zap.Error(err))
			return fmt.Errorf("failed querying user: %v", err)
		}
		fmt.Println(us.ID)

		//err = client.User.DeleteOneID(u.ID).Exec(context.Background())
		err = client.User.DeleteOne(us).Exec(context.Background())
		if err != nil {
			//panic(err)
			//fmt.Println("Delete user err: ", err)
			//log.Fatal("Delete user error:", zap.Error(err))
			return err
		}
		return c.NoContent(http.StatusOK)
	}
}

//
//// @Title DeleteUser
//// @Description 删除用户信息
//// @Accept  json
//// @Param username formData string true "用户名称"
//// @Success 200 "删除信息成功"
//// @Failure 400 "删除信息失败"
//// @Router /handler.DeleteUser [DELETE]
//// 删除用户
//// 根据ID删除用户
//func DeleteUserById(client *ent.Client) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		//client, err := database.Client()
//		//client, err := config.NewClient()
//		//if err != nil {
//		//	panic(err)
//		//	return err
//		//}
//		u := new(ent.User)
//
//		//// 接收raw数据
//		//result, err := ioutil.ReadAll(c.Request().Body)
//		//if err != nil {
//		//	fmt.Println("ioutil.ReadAll err:", err)
//		//	return err
//		//}
//		//// 解析raw为json
//		//err = json.Unmarshal(result, &u)
//		//if err != nil {
//		//	fmt.Println("json.Unmarshal err:", err)
//		//	return err
//		//}
//
//		// 直接解析raw数据为json
//		log, _ := zap.NewProduction()
//		if err := json.NewDecoder(c.Request().Body).Decode(&u); err != nil {
//			log.Fatal("json decode error", zap.Error(err))
//			return err
//		}
//		//// or for DisallowUnknownFields() wrapped in your custom func
//		//decoder := json.NewDecoder(c.Request().Body)
//		//decoder.DisallowUnknownFields()
//		//if err := decoder.Decode(&payload); err != nil {
//		//	return err
//		//}
//
//		fmt.Println(1111)
//		fmt.Println(u.Username)
//		fmt.Println(22222)
//		us, err := client.User.Query().Where(user.UsernameEQ(u.Username)).Only(context.Background())
//		if err != nil {
//			//panic(err)
//			//log.Fatal("Query user error:", zap.Error(err))
//			return fmt.Errorf("failed querying user: %v", err)
//		}
//		fmt.Println(us.ID)
//
//		//err = client.User.DeleteOneID(u.ID).Exec(context.Background())
//		err = client.User.DeleteOneID(us.ID).Exec(context.Background())
//		if err != nil {
//			//panic(err)
//			//log.Fatal("Delete user error:", zap.Error(err))
//			//fmt.Println("Delete user err: ", err)
//			return err
//		}
//		//return c.NoContent(http.StatusNoContent)
//		return c.NoContent(http.StatusOK)
//	}
//}
//
////u, err := client.User.Create().SetID()
////var us ent.User
////if err := c.Bind(&us); err != nil {
////	return err
////}
////
////pass, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
////if err != nil {
////	fmt.Println(err)
////}
////us.Password = string(pass)
////us.ID = utils.UUID()
//
////u1:=User{}
////u1.Password=encodePWD //模拟从数据库中读取到的 经过bcrypt.GenerateFromPassword处理的密码值
////loginPwd:="pwd" //用户登录时输入的密码
////// 密码验证
////err = bcrypt.CompareHashAndPassword([]byte(u1.Password), []byte(loginPwd)) //验证（对比）
////if err != nil {
////	fmt.Println("pwd wrong")
////} else {
////	fmt.Println("pwd ok")
////}

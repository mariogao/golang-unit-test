package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mariogao/golang-unit-test/model"
)

// UserControl
type UserControl struct {
	UserModel model.Iuser
}

// NewUserControl 实例化
func NewUserControl() *UserControl {
	return &UserControl{
		UserModel: &model.UserModel{},
	}
}

type paramGetUser struct {
	Name string `josn:"name"`
	Age  int    `json:"age"`
}

// GetUser handler
func (u *UserControl) GetUser(c echo.Context) error {
	var user = &model.User{}
	err := c.Bind(user)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("get接收：", user)
	resu, _ := u.UserModel.Get(user)
	fmt.Println("get输出u:", resu)
	return c.JSON(200, map[string]interface{}{"code": 0, "res": resu})
}

//UpdateUser handler
func (u *UserControl) UpdateUser(c echo.Context) error {
	var user = &model.User{}
	err := c.Bind(user)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("post接收：", user)
	err = u.UserModel.Update(user)
	fmt.Println("post err：", err)
	return c.JSON(200, map[string]interface{}{"code": 0, "res": map[string]interface{}{}})

}

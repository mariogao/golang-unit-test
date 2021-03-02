package model

import "log"

//User 数据库对应
type User struct {
	Name string
	Age  int
}

//Iuser 接口 mock要用
type Iuser interface {
	Get(u *User) (User, error)
	Update(u *User) error
}

// UserModel  接口实现
type UserModel struct{}

// Get 方法实现
func (uc *UserModel) Get(u *User) (User, error) {
	// 操作数据库的逻辑

	log.Println("执行了Get方法")
	var resu = User{Name: "数据库读到的name", Age: 15}
	return resu, nil
}

// Update 实现
func (uc *UserModel) Update(u *User) error {
	// 操作数据库的逻辑

	log.Println("执行了Update方法")
	return nil
}

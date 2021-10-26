package controller

import (
	"net/http"
	"test/controller/mocks"
	"test/model"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type UserTest struct {
	suite.Suite
	ex           *httpexpect.Expect
	mCtrl        *gomock.Controller
	muserControl *UserControl
}

// SetupSuite 所有测试用例执行前进行的操作
func (ut *UserTest) SetupSuite() {
	//连接数据库

	//加载路由
	e := echo.New()
	ut.muserControl = NewUserControl()
	e.GET("/get", ut.muserControl.GetUser)
	e.POST("/update", ut.muserControl.UpdateUser)
	ut.ex = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(e),
		},
		Reporter: httpexpect.NewAssertReporter(ut.T()),
	})
}

// TearDownSuite 所有测试用例结束后进行的操作
func (ut *UserTest) TearDownSuite() {

}

//SetupTest 每个测试用例开始前进行的操作
func (ut *UserTest) SetupTest() {
	ut.mCtrl = gomock.NewController(ut.T())
}

//TearDownTest 每个测试用例结束后进行的操作
func (ut *UserTest) TearDownTest() {
	ut.mCtrl.Finish()
}

func TestUserController(t *testing.T) {
	suite.Run(t, new(UserTest))
}

//正常走流程的测试用例
func (ut *UserTest) TestGetUser() {
	ut.ex.GET("/get").Expect().Status(200).JSON().Object().ValueEqual("code", 0) //返回 {"code":0,"res":{"Name":"数据库读到的name","Age":15}}
}

//跳过model层Get方法，用自定义的Get方法 来测试接口的用例
func (ut *UserTest) TestGetUserSkipDB() {
	//首先 生成mock文件 mockgen -source=model/user.go -destination=controller/mocks/user.go -package=mocks Iuser

	//mock生成的新的结构体，该结构体同样实现了 Iuser的所有方法
	getUserSkip := mocks.NewMockIuser(ut.mCtrl)
	// //把mock生成的结构体赋值给 ut.muserControl  实现最终调用Iuser接口的方法时都用mock实现的方法
	ut.muserControl.UserModel = getUserSkip
	// 自定义 Get(u *User) (User, error) 方法返回值
	resu := model.User{Name: "自定义的name", Age: 20}
	getUserSkip.EXPECT().Get(gomock.Any()).Return(resu, nil)
	//测试接口现在的返回值
	ut.ex.GET("/get").Expect().Status(200).JSON().Object().ValueEqual("code", 0) //返回 {"code":0,"res":{"Name":"自定义的name","Age":20}}

	// 测试完毕后，将 ut.muserControl.UserModel 的值初始化，避免其它用例本应调用ut.muserControl.UserModel的方法时，由于此处赋值的原因而最张调用了getUserSkip（mocks） 的方法
	ut.muserControl.UserModel = &model.UserModel{}
}

func (ut *UserTest) TestUpdateUser() {
	ut.ex.POST("/update").WithHeader("Content-Type", "application/json").WithJSON(&model.User{Name: "test name", Age: 19}).Expect().
		Status(200).JSON().Object().ValueEqual("code", 0)
}

func (ut *UserTest) TestUpdateUserSkipDB() {
	UpdateUserSkip := mocks.NewMockIuser(ut.mCtrl)
	ut.muserControl.UserModel = UpdateUserSkip
	UpdateUserSkip.EXPECT().Update(gomock.Any()).Return(nil)
	ut.ex.POST("/update").WithHeader("Content-Type", "application/json").WithJSON(&model.User{Name: "test name", Age: 18}).Expect().
		Status(200).JSON().Object().ValueEqual("code", 0)
}

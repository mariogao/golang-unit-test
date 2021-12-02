// package controller ...
package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/mariogao/golang-unit-test/model"

	. "github.com/agiledragon/gomonkey"
	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetUserSkip(t *testing.T) {
	Convey("test GetUser", t, func(c C) {
		var TG = NewUserControl()
		patches := ApplyMethod(reflect.TypeOf(TG.UserModel), "Get", func(_ *model.UserModel, _ *model.User) (model.User, error) {
			return model.User{Name: "convey test"}, nil
		})
		defer patches.Reset()
		reqData := &model.User{Name: "test"}
		reqBody, _ := json.Marshal(reqData)
		req := httptest.NewRequest(http.MethodGet, "/get", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		echoTest := e.NewContext(req, rec)
		getUser := TG.GetUser(echoTest)
		fmt.Printf("=== %+v", getUser)
		c.So(getUser, ShouldBeNil)
	})
}

func TestGetUser(t *testing.T) {
	Convey("test GetUser", t, func(c C) {
		var TG = NewUserControl()
		reqData := &model.User{Name: "test"}
		reqBody, _ := json.Marshal(reqData)
		req := httptest.NewRequest(http.MethodGet, "/get", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		echoTest := e.NewContext(req, rec)
		getUser := TG.GetUser(echoTest)
		fmt.Printf("=== %+v", getUser)
		c.So(getUser, ShouldBeNil)
	})
}

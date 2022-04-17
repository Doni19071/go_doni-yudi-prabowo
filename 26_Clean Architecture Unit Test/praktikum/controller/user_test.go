package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUserControllerGet(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	ps := service.NewMockUserService()
	ps.Add(model.User{Email: "Budi@gmail.com", Password: "Bandung123"})
	ps.Add(model.User{Email: "Chandra@gmail.com", Password: "Cilegon123"})

	pc := NewUserController(ps)
	pc.Get(c)

	var users []model.User
	if err := json.Unmarshal(rec.Body.Bytes(), &users); err != nil {
		t.Errorf("unmarshalling returned user failed")
	}
	if len(users) != 2 {
		t.Errorf("Expecting len(users) to be 2, get %d", len(users))
	}
	if users[0].Email != "Budi@gmail.com" {
		t.Errorf("Expecting users[0].Name to be Budi, get %s", users[0].Email)
	}
	if users[1].Email != "Chandra@gmail.com" {
		t.Errorf("Expecting users[1].Name to be Chandra, get %s", users[1].Email)
	}
}

func TestUserControllerAdd(t *testing.T) {
	e := echo.New()

	newUserJson, _ := json.Marshal(map[string]string{
		"email":    "Anton@gmail.com",
		"password": "Aceh123",
	})
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newUserJson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	ps := service.NewMockUserService()
	pc := NewUserController(ps)
	pc.Add(c)

	users, err := ps.Get()
	if err != nil {
		t.Error(err)
	}
	if len(users) != 1 {
		t.Errorf("Expecting len(users) to be 1, get %d", len(users))
	}
	if users[0].Email != "Anton@gmail.com" {
		t.Errorf("Expecting users[0].Name to be Anton, get %s", users[0].Email)
	}
	if users[0].Password != "Aceh123" {
		t.Errorf("Expecting users[0].Address to be Aceh, get %s", users[0].Password)
	}

	// TODO: also check response
}

package integrationtest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/litonshil/crud_go_echo/pkg/connection"
	"github.com/litonshil/crud_go_echo/pkg/controllers"
	repoImpl "github.com/litonshil/crud_go_echo/pkg/repository"
	svcImpl "github.com/litonshil/crud_go_echo/pkg/svc"
	"github.com/stretchr/testify/assert"
)

var db *gorm.DB

func TestRegistration(t *testing.T) {
	connection.Connect()
	db = connection.GetDB()
	userRepo := repoImpl.NewUsersRepository(db)
	authSvc := svcImpl.NewAuthService(userRepo)
	authCr := controllers.NewAuthController(authSvc)

	url := "/registration"
	// method := "POST"

	payload := strings.NewReader(`{
    	"name": "bbbbbccsccdddcc",
    	"address": "bbbbsbcdcddcccc",
    	"email": "bbbbbcccsdcddcc@gmail.com",
    	"type": "normal-user",
    	"password": "bbbbdsdcbccdccc"
	}`)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, url, payload)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code, "testtt")
	if assert.NoError(t, authCr.Registration(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
	// err := authCr.Registration(c)
	// fmt.Println(err)
	// assert.NoError(t, err)
}

// func TestRegistration(t *testing.T) {

// 	url := "http://localhost:8080/user/registration"
// 	method := "POST"

// 	payload := strings.NewReader(`{
//     	"name": "bbbbbccccc",
//     	"address": "bbbbbccccc",
//     	"email": "bbbbbccccc@gmail.com",
//     	"type": "normal-user",
//     	"password": "bbbbbccccc"
// 	}`)

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)

// 	if err != nil {
// 		assert.Fail(t, err.Error())
// 	}
// 	req.Header.Add("Content-Type", "application/json")

// 	res, err := client.Do(req)
// 	if err != nil {
// 		assert.Fail(t, err.Error())
// 	}

// 	assert.Equal(t, http.StatusCreated, res.StatusCode)
// 	// assert.Condition(t, func() bool {
// 	// 	return res.StatusCode == 200
// 	// })

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		assert.Fail(t, err.Error())
// 	}
// 	resp := string(body)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)
// 	fmt.Println(resp)
// }

package integrationtest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {

	url := "http://localhost:8080/user/registration"
	method := "POST"

	payload := strings.NewReader(`{
    	"name": "bbbbcccc",
    	"address": "bbbbcccc",
    	"email": "bbbbcccc@gmail.com",
    	"type": "normal-user",
    	"password": "bbbbcccc"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		assert.Fail(t, err.Error())
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, http.StatusCreated, res.StatusCode)
	// assert.Condition(t, func() bool {
	// 	return res.StatusCode == 200
	// })

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	resp := string(body)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp)
}

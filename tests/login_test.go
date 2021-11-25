package Tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// func TestSignUp(t *testing.T) {

// 	client := resty.New()

// 	// POST JSON string
// 	// No need to set content type, if you have client level setting
// 	resp, err := client.R().
// 		SetHeader("Content-Type", "application/json").
// 		SetBody(`{
// 		"id" : 0,
// 		"username" : "sharyo1",
// 		"email": "ahmed1@sharyo.com",
// 		"password": "passw1",
// 		"authorityLevel" : 1

// 	}`).Post("http://localhost:8080/signup")

// 	fmt.Println("  Body       :\n", string(resp.Body()))
// 	//fmt.Println("  Body       :\n", err)
// 	assert.Equal(t, nil, err, err)

// 	assert.Equal(t, 200, resp.StatusCode(), resp.StatusCode())
// }

func TestLogin(t *testing.T) {

	client := resty.New()

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{
		"email": "ahmed@sharyo.com",
		"password": "passw"
	}`).Post("http://localhost:8080/login")

	//fmt.Println("  Body       :\n", resp.StatusCode())
	//fmt.Println("  Body       :\n", err)
	assert.Equal(t, nil, err, err)

	assert.Equal(t, 200, resp.StatusCode(), resp.StatusCode())
}

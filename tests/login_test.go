package Tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {

	client := resty.New()

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{
		"email": "bikash.dulal@wesionary.team",
		"password": "testing"
	}`).Post("http://localhost:8080/login")

	//fmt.Println("  Body       :\n", resp.StatusCode())
	//fmt.Println("  Body       :\n", err)
	assert.Equal(t, nil, err, err)

	assert.Equal(t, 200, resp.StatusCode(), resp.StatusCode())
}

package test_auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
	"github.com/ioannuwu/git-diff-as-a-service/internal/test"
)

func register(ctx context.Context, log *logger.Logger, url string) error {

	requestBody := `
	{
    	"login": "login",
    	"password": "1234567890",

    	"full_name": "Alex",
    	"email": "alex@gmail.com"
	}
	`

	integrationTest := test.IntegrationTest{
		Method:             http.MethodPost,
		URL:                url,
		RequestBody:        requestBody,
		ExpectedStatusCode: http.StatusCreated,
		Check:              nil,
	}

	return integrationTest.Do(ctx, log)
}

func registerAgain(ctx context.Context, log *logger.Logger, baseURL string) error {

	requestBody := `
	{
    	"login": "login",
    	"password": "12345678333390",

    	"full_name": "Alexander",
    	"email": "alexander@gmail.com"
	}
	`

	integrationTest := test.IntegrationTest{
		Method:             http.MethodPost,
		URL:                baseURL,
		RequestBody:        requestBody,
		ExpectedStatusCode: http.StatusConflict,

		Check: func(r *http.Response) error {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return err
			}

			if !strings.Contains(string(body),"already exists") {
				return fmt.Errorf("registered user should already exist")
			}
			return nil
		},
	}

	return integrationTest.Do(ctx, log)
}

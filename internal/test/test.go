package test

import (
	"fmt"
	"bytes"
	"context"
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

type IntegrationTest struct {
	Method             string
	URL                string
	RequestBody        string
	ExpectedStatusCode int

	Check func(*http.Response) error
}

func (i *IntegrationTest) Do(ctx context.Context, log *logger.Logger) error {

	request, err := http.NewRequestWithContext(
		ctx, 
		http.MethodPost, 
		i.URL, 
		bytes.NewBufferString(i.RequestBody),
	)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != i.ExpectedStatusCode {
		log.Error(
			"Wrong status code", 
			logger.String("got_status", fmt.Sprint(response.StatusCode)),
			logger.String("expected_status", fmt.Sprint(i.ExpectedStatusCode)),
		)
		return nil
	}

	if i.Check != nil {
		if err := i.Check(response); err != nil {
			log.Error(err.Error())
			return err
		}
	}

	return nil
}

package test_auth

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

func RunAuthTests(ctx context.Context, log *logger.Logger) {

	{
		baseURL := "http://localhost:5050/api/v1/auth/register"

		err := register(ctx, log, baseURL)
		if err != nil {
			panic(err)
		}

		err = registerAgain(ctx, log, baseURL)
		if err != nil {
			panic(err)
		}
	}
}

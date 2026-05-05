package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
	test_auth "github.com/ioannuwu/git-diff-as-a-service/internal/test/auth"
)

func main() {
	
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log := logger.MustNewLogger(logger.MustNewConfig())

	log.Debug("Starting integration tests for git-diff-app")

	test_auth.RunAuthTests(ctx, log)
	
	log.Debug("All tests in git-diff-app were sucessful")
}
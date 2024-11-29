package services

import (
	"context"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"gotest.tools/v3/assert"
)

func TestCreateNextUID(t *testing.T) {

	t.Run("", func(t *testing.T) {
		ctx := context.Background()
		logger := logger.NewLogrusLogger("info", ctx)
		store := storager.New(nil, nil, nil)
		service := NewService(logger, store)
		first, err := service.CreateNextUID()
		if err != nil {
			t.Errorf(err.Error())
		}
		assert.Equal(t, service.GetUsersCount(), first)
		assert.Equal(t, 1, first)
	})
}

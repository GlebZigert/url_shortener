package services

import (
	"context"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/db"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func TestShort(t *testing.T) {

	type request struct {
		origin string
		user   int
	}

	type want struct {
		err error
	}

	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "",
			request: request{
				"example",
				0,
			},
			want: want{
				nil,
			},
		},
		{
			name: "",
			request: request{
				"example",
				0,
			},
			want: want{
				&packerr.ErrConflict409{S: "попытка сократить уже имеющийся в базе URL"},
			},
		},
	}
	cfg := config.NewConfig("prog", []string{})
	ctx := context.Background()

	db.Init(cfg.DatabaseDSN)
	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := NewService(logger, store)

	//srv, _ := NewServer(cfg, mdl, logger, service)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//t.Log("req: ", test.request.method, " ", test.request.url)

			var err error

			_, errr := service.Short(test.request.origin, test.request.user)

			if test.want.err != nil {
				assert.ErrorType(t, errr, test.want.err)
			} else {
				assert.NilError(t, errr)
			}

			//errors.As(err, &conflict)

			require.NoError(t, err)

		})
	}
}

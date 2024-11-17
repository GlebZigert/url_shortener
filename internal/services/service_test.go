package services

import (
	"context"
	"testing"

	"github.com/GlebZigert/url_shortener.git/internal/config"
	"github.com/GlebZigert/url_shortener.git/internal/file_reader"
	"github.com/GlebZigert/url_shortener.git/internal/logger"
	"github.com/GlebZigert/url_shortener.git/internal/packerr"
	"github.com/GlebZigert/url_shortener.git/internal/storager"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func TestService(t *testing.T) {
	//Сначала тест Short
	type request struct {
		value string
		user  int
	}

	type want struct {
		err    error
		origin string
	}

	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "Short",
			request: request{
				"example1",
				0,
			},
			want: want{
				nil,
				"",
			},
		},
		{
			name: "Short same value",
			request: request{
				"example1",
				0,
			},
			want: want{
				&packerr.Conflict,
				"",
			},
		},
	}
	cfg, err := config.NewConfig("prog", []string{}, file_reader.New())
	if err != nil {
		t.Errorf("error parse config")
	}
	ctx := context.Background()

	store := storager.New(cfg)

	logger := logger.NewLogrusLogger(cfg.FlagLogLevel, ctx)

	service := NewService(logger, store)

	//srv, _ := NewServer(cfg, mdl, logger, service)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//t.Log("req: ", test.request.method, " ", test.request.url)

			_, errr := service.Short(test.request.value, test.request.user)

			if test.want.err != nil {
				assert.Equal(t, errr, test.want.err)
			} else {
				assert.NilError(t, errr)
			}

			//errors.As(err, &conflict)

			require.NoError(t, err)

		})
	}

	short := "example"
	origin, err := service.Short("example", 0)
	if err != nil {
		t.Error(err)
	}
	//Origin
	tests = []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "Origin",
			request: request{
				origin,
				0,
			},
			want: want{
				nil,
				short,
			},
		},
	}

	//srv, _ := NewServer(cfg, mdl, logger, service)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//t.Log("req: ", test.request.method, " ", test.request.url)

			var err error

			result, errr := service.Origin(test.request.value)

			if test.want.err != nil {

				assert.Equal(t, err, test.want.err)
			} else {
				assert.NilError(t, errr)
			}

			assert.Equal(t, result, test.want.origin)

			//errors.As(err, &conflict)

			require.NoError(t, err)

		})
	}
	//Delete

	type requestToDelete struct {
		value string
		user  int
	}

	type wantToDelete struct {
		err error
	}

	testsToDelete := []struct {
		name    string
		request requestToDelete
		want    wantToDelete
	}{
		{
			name: "TesDelete",
			request: requestToDelete{
				value: "fffffff",
				user:  0,
			},
			want: wantToDelete{
				err: nil,
			},
		},
	}

	for _, test := range testsToDelete {
		t.Run(test.name, func(t *testing.T) {
			//t.Log("req: ", test.request.method, " ", test.request.url)
			short, err := service.Short(test.request.value, test.request.user)

			if err != nil {
				assert.Error(t, err, "")
			}

			origin, err := service.Origin(short)

			if err != nil {
				assert.Error(t, err, "")
			}

			assert.Equal(t, origin, test.request.value)

			err = service.Delete([]string{short}, test.request.user)

			if err != nil {
				t.Error(err)
			}

			_, err = service.Origin(short)
			str := "шорт " + short + " удален"

			customErr := err.(*packerr.ErrDeleted)

			assert.Equal(t, str, customErr.S)

		})
	}

}

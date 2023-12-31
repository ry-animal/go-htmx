package server

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ry-animal/go-htmx/ent"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
	os.Exit(m.Run())
}

func setupDatabase(t require.TestingT) *ent.Client {
	const (
		driver = "sqlite3"
		dbURL  = "file:indie?mode=memory&cache=shared&_fk=1"
	)
	client, err := ent.Open(driver, dbURL)
	require.NoError(t, err)

	err = client.Schema.Create(context.Background())
	require.NoError(t, err)

	return client
}

func TestExample(t *testing.T) {
	t.Parallel()

	client := setupDatabase(t)

	app := NewApp(client)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
}

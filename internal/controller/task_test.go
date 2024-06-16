package controller_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/symonk/toodoo/cmd/server"
	"github.com/symonk/toodoo/internal/db"
	"github.com/symonk/toodoo/internal/logging"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Manage a postgres testcontainer for integration testing against
// a real database.
func TestMain(m *testing.M) {
	ctx := context.Background()
	pgres, err := postgres.RunContainer(
		ctx,
		postgres.WithInitScripts(filepath.Join("../../", "testdata", "testbootstrap.sql")),
		postgres.WithDatabase("toodoo"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic("unable to start postgres test container")
	}
	defer func() {
		if err := pgres.Terminate(ctx); err != nil {
			panic("could not stop the pgres container after tests!")
		}
	}()
	db.Init("postgresql://postgres:postgres@localhost:5432/toodoo?sslmode=disable")
	logging.Init()
	os.Exit(m.Run())
}

func TestFetchAllTasksSuccess(t *testing.T) {
	t.Parallel()
	router := server.NewRouter()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/task", nil)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
}

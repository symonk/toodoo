package controller_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/symonk/toodoo/cmd/server"
	"github.com/symonk/toodoo/internal/config"
	"github.com/symonk/toodoo/internal/db"
	"github.com/symonk/toodoo/internal/logging"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TasksTestSuite struct {
	suite.Suite
	container *postgres.PostgresContainer
	cfg       *config.Config
}

func TestTasksTestSuite(t *testing.T) {
	suite.Run(t, new(TasksTestSuite))
}

func (t *TasksTestSuite) SetupSuite() {
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
				WithStartupTimeout(20*time.Second)),
	)
	if err != nil {
		panic(fmt.Errorf("unable to start postgres test container %s", err.Error()))
	}
	t.container = pgres
	connStr, err := t.container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(fmt.Errorf("err with connection string %s", err.Error()))
	}
	db.Init(connStr)
	logging.Init()
}

func (t *TasksTestSuite) TeardownSuite() {
	ctx := context.Background()
	defer func() {
		if err := t.container.Terminate(ctx); err != nil {
			panic("could not stop the pgres container after tests!")
		}
	}()
}

func (suite *TasksTestSuite) TestFetchAllTasksSuccess() {
	router := server.NewRouter()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/task", nil)
	router.ServeHTTP(recorder, request)
	suite.Equal(200, recorder.Code)
}

func (suite *TasksTestSuite) TestFetchASingleTask() {
	router := server.NewRouter()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/task/1", nil)
	router.ServeHTTP(recorder, request)
	suite.Equal(200, recorder.Code)
}

func (suite *TasksTestSuite) TestFetchWithInvalidTaskID() {
	router := server.NewRouter()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/task/1337", nil)
	router.ServeHTTP(recorder, request)
	suite.Equal(500, recorder.Code)
}

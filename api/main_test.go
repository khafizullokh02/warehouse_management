package api

import (
	"os"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	db "github.com/khafizullokh02/warehouse_management/db/sqlc"
	"github.com/khafizullokh02/warehouse_management/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   fake.RandomStringWithLength(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

var (
	fake faker.Faker
)

func TestMain(m *testing.M) {
	fake = faker.New()

	os.Exit(m.Run())
}

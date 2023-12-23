package db

import (
	"context"
	"os"
	"testing"

	"github.com/jaswdr/faker"
)

var (
	fake        faker.Faker
	testStore   Store
	postgresURI = "postgresql://root:secret@localhost:5432/warehouse_management?sslmode=disable"
)

func TestMain(m *testing.M) {
	fake = faker.New()
	testStore = NewStore(context.Background(), postgresURI)

	os.Exit(m.Run())
}

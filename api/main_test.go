package api

import (
	"os"
	"testing"

	"github.com/jaswdr/faker"
)

var (
	fake faker.Faker
)

func TestMain(m *testing.M) {
	fake = faker.New()

	os.Exit(m.Run())
}

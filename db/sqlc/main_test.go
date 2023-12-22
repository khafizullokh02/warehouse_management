package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/jaswdr/faker"
)

var (
	fake      faker.Faker
	testStore Store
)

func TestMain(m *testing.M) {
	fake = faker.New()

	for i := 0; i < 10; i++ {
		fmt.Println(fake.Address().Address())
	}

	os.Exit(m.Run())
}

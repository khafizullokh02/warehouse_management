package db

import (
	"testing"
)

func createRandomUser(t *testing.T) User {
	hahsedPassword, err := util.HashedPassword(fake.RandomStringWithLength(6))
}

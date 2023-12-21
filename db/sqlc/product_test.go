package db

import (
	"context"
	"testing"

	"github.com/jaswdr/faker"
)

func TestCreateProduct(t *testing.T) {

	fake = faker.New()

	arg := CreateProductParams{
		ID:    fake.Int32(),
		Name:  fake.Person().Name(),
		Image: fake.Lorem().Text(50),
		Brand: fake.Lorem().Text(10),
	}

	product, err := createProduct(context.Background(), arg)
}

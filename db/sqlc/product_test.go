package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Name:  fake.Food().Vegetable(),
		Image: fake.Food().Vegetable(),
		Brand: fake.RandomStringElement([]string{"apple", "samsung", "xiaomi"}),
		
	}

	product, err := testStore.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Image, product.Image)
	require.Equal(t, arg.Brand, product.Brand)

	require.NotZero(t, product.Name)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testStore.GetProduct(
		context.Background(),
		product1.Name,
	)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Image, product2.Image)
	require.Equal(t, product1.Brand, product2.Brand)
	// require.WithinDuration(t, product1.CreatedAt, user2.CreatedAt, time.Second)
}

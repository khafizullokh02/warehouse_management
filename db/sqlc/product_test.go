package db

import (
	"context"
	"testing"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Name:           fake.Food().Vegetable(),
		Image:          fake.Food().Vegetable(),
		Brand:          fake.RandomStringElement([]string{"apple", "samsung", "xiaomi"}),
		SupCode:        fake.RandomStringWithLength(10),
		BarCode:        cast.ToString(fake.RandomNumber(10000)),
		WholesalePrice: fake.Float64(2, 100, 1000),
		RetailPrice:    fake.Float64(2, 100, 1000),
	}

	product, err := testStore.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Image, product.Image)
	require.Equal(t, arg.Brand, product.Brand)
	require.Equal(t, arg.SupCode, product.SupCode)
	require.Equal(t, arg.BarCode, product.BarCode)
	require.Equal(t, arg.WholesalePrice, product.WholesalePrice)
	require.Equal(t, arg.RetailPrice, product.RetailPrice)

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
}

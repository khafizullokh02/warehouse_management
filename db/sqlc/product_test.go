package db

import (
	"context"
	"testing"
	"time"

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

func TestGetProduct(t *testing.T) { // query has been changed from: name to id
	product1 := createRandomProduct(t)
	product2, err := testStore.GetProduct(
		context.Background(),
		product1.ID,
	)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1, product2)
	require.WithinDuration(t, product1.CreatedAt.Time, product2.CreatedAt.Time, time.Second)
}

func TestUpdateProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	arg := UpdateProductParams{
		ID:             product1.ID,             // query has been updated from: sqlc.narg() to: sqlc.arg()
		Name:           product1.Name,           //
		Image:          product1.Image,          //
		Brand:          product1.Brand,          //
		SupCode:        product1.SupCode,        //
		BarCode:        product1.BarCode,        //
		WholesalePrice: product1.WholesalePrice, //
		RetailPrice:    product1.RetailPrice,    //
		Discount:       product1.Discount,       //
	}

	product2, err := testStore.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1, product2)
	require.WithinDuration(t, product1.CreatedAt.Time, product2.CreatedAt.Time, time.Second) // there was an error when I wanted to do: product.CreatedAt
}

func TestDeleteProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	err := testStore.DeleteProduct(context.Background(), product1.ID)
	require.NoError(t, err)

	product2, err := testStore.GetProduct(context.Background(), product1.ID)
	require.Error(t, err)
	// require.EqualError(t, err, ErrRecordNotFound.Error()) // this line didn't work because error recorder wasnot declared
	require.Empty(t, product2)
}

func TestListProducts(t *testing.T) {
	var lastProduct Product
	for i := 0; i < 10; i++ {
		lastProduct = createRandomProduct(t)
	}

	arg := ListProductsParams{
		Name:   lastProduct.Name,
		Limit:  5,
		Offset: 0,
	}

	products, err := testStore.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, products)

	for _, product := range products {
		require.NotEmpty(t, product)
		require.Equal(t, lastProduct.Name, product.Name)
	}
}

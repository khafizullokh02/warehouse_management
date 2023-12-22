package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/pgtype"
	"github.com/jaswdr/faker"
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
	product2, err := testStore.GetProduct(context.Background(), product1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Image, product2.Image)
	require.Equal(t, product1.Brand, product2.Brand)
	// require.WithinDuration(t, product1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateProductOnlyName(t *testing.T) {
	oldProduct := createRandomProduct(t)

	newName := fake.Food().Vegetable()
	updatedProduct, err := testStore.UpdateProduct(context.Background(), UpdateProductParams{
		ID: oldProduct.ID,
		Name: pgtype.Text{
			String: newName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldProduct.Name, updatedProduct.Name)
	require.Equal(t, newName, updatedProduct.Name)
	require.Equal(t, oldProduct.Image, updatedProduct.Image)
	require.Equal(t, oldProduct.Brand, updatedProduct.Brand)
}

func TestUpdateProductOnlyImage(t *testing.T) {
	oldProduct := createRandomProduct(t)

	newImage := fake.Food().Vegetable()
	updatedProduct, err := testStore.UpdateProduct(context.Background(), UpdateProductParams{
		ID: oldProduct.ID,
		Image: pgtype.Text{
			String: newImage,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldProduct.Image, updatedProduct.Image)
	require.Equal(t, newImage, updatedProduct.Image)
	require.Equal(t, oldProduct.Image, updatedProduct.Image)
	require.Equal(t, oldProduct.Image, updatedProduct.Image)
}

func TestUpdateProductOnlyBrand(t *testing.T) {
	oldProduct := createRandomProduct(t)

	newBrand := fake.Food().Vegetable()
	updatedProduct, err := testStore.UpdateProduct(context.Background(), UpdateProductParams{
		ID: oldProduct.ID,
		Brand: pgtype.Text{
			String: newBrand,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldProduct.Brand, updatedProduct.Brand)
	require.Equal(t, newBrand, updatedProduct.Brand)
	require.Equal(t, oldProduct.Brand, updatedProduct.Brand)
	require.Equal(t, oldProduct.Brand, updatedProduct.Brand)
}

func TestUpdateProductAllFields(t *testing.T) {
	oldProduct := createRandomProduct(t)

	newName := faker.New().Food().Vegetable()
	newImage := faker.New().Food().Vegetable()
	newBrand := fake.RandomStringElement([]string{"apple", "samsung", "xiaomi"})

	updatedProduct, err := testStore.UpdateProduct(context.Background(), UpdateProductParams{
		Name: oldProduct.Name,
		Image: pgtype.Text{
			String: newImage,
			Valid:  true,
		},
		Brand: pgtype.Text{
			String: newBrand,
			Valid: true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldProduct.Name, updatedProduct.Name)
	require.Equal(t, newName, updatedProduct.Name)
	require.NotEqual(t, oldProduct.Image, updatedProduct.Image)
	require.Equal(t, newImage, updatedProduct.Image)
	require.NotEqual(t, oldProduct.Brand, updatedProduct.Brand)
	require.Equal(t, newBrand, updatedProduct.Brand)
}
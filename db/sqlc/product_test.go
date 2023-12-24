package db

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
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

	require.NotZero(t, product.Name)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	product := createRandomProduct(t)
	assert.NotEmpty(t, product, "returned product should not be empty")
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testStore.GetProduct(
		context.Background(),
		product1.Name,
	)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1, product2)
}

func TestQueries_UpdateProduct(t *testing.T) {
	product := createRandomProduct(t)

	updatingProperties := map[string]any{
		"name": fake.Food().Vegetable(),
	}

	tests := []struct {
		name string
		args UpdateProductParams
	}{
		{
			name: "update name",
			args: UpdateProductParams{
				Name: updatingProperties["name"].(string),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testStore.UpdateProduct(context.Background(), tt.args)
			require.NoError(t, err, "UpdateProduct() error = %v", err)

			want := makeUpdateProductResp(t, &product, tt.args)
			require.Equal(t, want, got, "UpdateProduct() got = %v, want %v", got, want)
		})
	}
}

func makeUpdateProductResp(t *testing.T, product *Product, args UpdateProductParams) *Product {
	bytes, err := json.Marshal(product)
	require.NoError(t, err)

	err = json.Unmarshal(bytes, &product)
	require.NoError(t, err)

	return product
}

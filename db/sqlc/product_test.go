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

func TestUpdateProduct(t *testing.T) {
	product := createRandomProduct(t)

	updatingProperties := map[string]any{
		"name":     fake.Food().Vegetable(),
		"discount": fake.Float64(2, 1, 10),
	}

	tests := []struct {
		name string
		args map[string]any
	}{
		{
			name: "update name",
			args: map[string]any{
				"name": updatingProperties["name"],
				"id":   product.ID,
			},
		},
		{
			name: "discount",
			args: map[string]any{
				"discount": updatingProperties["discount"],
				"id":       product.ID,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbPayload := covnertMapToUpdateProductParams(t, tt.args)
			got, err := testStore.UpdateProduct(context.Background(), dbPayload)
			require.NoError(t, err, "UpdateProduct() error = %v", err)

			want := makeUpdateProductResp(t, &product, tt.args)
			require.Equal(t, product, got, "UpdateProduct() got = %v, want %v", got, want)
		})
	}
}

func covnertMapToUpdateProductParams(t *testing.T, args map[string]any) UpdateProductParams {
	bytes, err := json.Marshal(args)
	require.NoError(t, err)

	var updateProductParams UpdateProductParams
	err = json.Unmarshal(bytes, &updateProductParams)
	require.NoError(t, err)

	return updateProductParams
}

func makeUpdateProductResp(t *testing.T, product *Product, args map[string]any) *Product {
	bytes, err := json.Marshal(args)
	require.NoError(t, err)

	err = json.Unmarshal(bytes, &product)
	require.NoError(t, err)

	return product
}

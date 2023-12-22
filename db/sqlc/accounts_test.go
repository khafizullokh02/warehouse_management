package db

import "testing"

func createRandomAccount(t *testing.T) {
	product := createRandomProduct(t)

	arg := CreateAccountParams{
		ID: product.Owner,//I think I need owner to be here
		UserID: fake.Int32(),
		Name: fake.Food().Vegetable(),
	}
}

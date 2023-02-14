package db

import (
	"context"
	"testing"

	"github.com/gitsuki/simplebank/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	return user
}

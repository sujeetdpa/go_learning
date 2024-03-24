package controllers

import (
	"fmt"

	"github.com/sujeetdpa/go_learning/pgx_learn/repo"
)

func GetAllAppUser() {
	users, err := repo.GetAllAppUser()
	if err != nil {
		fmt.Println("Unable to get All users")
		return
	}

	for idx, user := range users {
		fmt.Println("User Index: ", idx, user)
	}
}

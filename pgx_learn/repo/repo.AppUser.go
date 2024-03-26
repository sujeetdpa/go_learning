package repo

import (
	"context"
	"fmt"

	"github.com/sujeetdpa/go_learning/pgx_learn/db"
	"github.com/sujeetdpa/go_learning/pgx_learn/models"
)

func CreateAppUserTable() error {
	//Creating Table

	_, err := db.PostresDB.Exec(context.Background(), "create table if not exists app_user(id serial primary key, name text, username text, password text)")
	if err != nil {
		fmt.Println("Unable to create table user", err.Error())
		return err
	} else {
		fmt.Println("Successfully created table user")
	}
	return nil
}

func InsertIntoAppUser(appUser *models.AppUser) error {
	//Inseritng data
	_, err := db.PostresDB.Exec(context.Background(), "insert into app_user(name,username,password) values($1,$2,$3)", appUser.Name, appUser.Username, appUser.Password)
	if err != nil {
		fmt.Println("Error while inserting data in to app_user ", err.Error())
		return err
	} else {
		fmt.Println("Successfully inserted data into user")
	}
	tx, err := db.PostresDB.Begin(context.TODO())
	_, err = tx.Exec(context.TODO(), "abcdjhvahcj")
	return nil
	// Uisng name parametes
	// query := "insert into app_user(name,username,password) values(@name,@username,@password)"

	// args := pgx.NamedArgs{
	// 	"name":     "Ramesh",
	// 	"username": "ramesh",
	// 	"password": "ramesh",
	// }
	// _, err = conn.Exec(context.Background(), query, args)
	// if err != nil {
	// 	fmt.Println("Error while inserting data in to app_user ", err.Error())
	// 	return
	// } else {
	// 	fmt.Println("Successfully inserted data into user")
	// }
}

func GetAllAppUser() ([]models.AppUser, error) {
	//Querying Data
	rows, err := db.PostresDB.Query(context.Background(), "select * from app_user")
	if err != nil {
		fmt.Println("Error while fetching all app user: ", err.Error())
		return nil, err
	}
	var appUsers []models.AppUser
	for rows.Next() {
		appUser := models.AppUser{}
		rows.Scan(&appUser.ID, &appUser.Name, &appUser.Username, &appUser.Password)
		appUsers = append(appUsers, appUser)
	}
	return appUsers, nil
}

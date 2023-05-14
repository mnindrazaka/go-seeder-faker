package main

import (
	"database/sql"
	"fmt"

	"github.com/go-faker/faker/v4"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Email string `faker:"email"`
	Name  string `faker:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:((root))@tcp(127.0.0.1:3306)/faker_exercise")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= 100000000; i++ {
		user := User{}
		err = faker.FakeData(&user)

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = db.Exec("INSERT INTO users_with_index SET email=?, name=?", user.Email, user.Name)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Success")
}

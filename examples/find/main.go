package main

import (
	"fmt"
	"github.com/aliforever/go-mongorm"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	_, err := mongorm.New(mongorm.NewConfig().SetDBName("mydb"))

	user, err := mongorm.FindOneWithFilter[User](bson.M{
		"username": "admin",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	users, err := mongorm.FindWithFilter[User](bson.M{})
	if err != nil {
		fmt.Println("HERE:", err)
		return
	}

	fmt.Printf("%#v\n", user)

	for _, m := range users {
		fmt.Println(m.Username)
	}
}

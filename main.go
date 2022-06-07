package main

import (
	"log"

	"github.com/emiliomarin/go-repository/repository"
)

func main() {
	opts := &repository.DBOptions{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "arexdb_dev",
		Password: "arexdb_dev",
		Database: "test_db",
		SSLMode:  "disable",
	}
	db, err := repository.NewDB(opts)
	if err != nil {
		log.Fatalln("error creating DB: ", err)
	}

	fooRepo := repository.NewFooRepo(db)
	foo, err := fooRepo.CreateFoo("some-value")
	if err != nil {
		log.Fatalln("error creating foo: ", err)
	}

	getFoo, err := fooRepo.GetFoo(foo.ID)
	if err != nil {
		log.Fatalln("error getting foo: ", err)
	}

	log.Println("Foo value:", getFoo.Value)
}

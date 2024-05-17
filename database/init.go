package database

import (
	"context"
	"log"
)

func InitDB() {
	if _, err := DB.NewCreateTable().Model((*User)(nil)).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	if _, err := DB.NewCreateTable().Model((*Post)(nil)).Exec(context.TODO()); err != nil {
		log.Fatalln(err)
	}
}
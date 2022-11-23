package main

import (
	"github.com/gogoclouds/gogo-services/admin-service/internal/comm/g"
	"github.com/gogoclouds/gogo-services/common-lib/db"
	"log"
)

func main() {
	var err error
	if g.DB, err = db.Open(); err != nil {
		log.Panicln(err)
	}
}

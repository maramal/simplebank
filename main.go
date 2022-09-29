package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/maramal/simplebank/api"
	db "github.com/maramal/simplebank/db/sqlc"
	"github.com/maramal/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("no se puede leer la configuración:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("No se pudo iniciar el servidor: ", err)
	}

}

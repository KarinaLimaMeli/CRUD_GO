package db

import(
	"database/sql"
	"fmt"
	"github.com/aprendagoland/api-pgsql/configs"
    _ "github.com/lib/pq"
)

func openConnection()(*sql.DB, error){
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	conf.Host, conf.Port, conf.User, conf.Pass,conf.Database)

	conn, err := sql.Open("postgres",sc)
	if err != nil {
		panic(err)
	}
	err =  conn.Ping()

	return conn, err

}



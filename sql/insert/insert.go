package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:otavio@tcp(172.17.0.2:3306)/cursogo")

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec("Pedro")
	if err != nil {
		fmt.Println("ERRO =>", err)
	}

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, err := res.RowsAffected()
	fmt.Println(linhas)
}

package main

// go get -u github.com/go-sql-driver/mysql
// cd /go/src/github.com/go-sql-driver/

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func execQuery(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	// abre conexão, busca em "go-sql-driver/<string>"
	// DSN format = <username>:<pw>@tcp(<HOST>:<port>)/<dbname>
	db, err := sql.Open("mysql", "root:otavio@tcp(172.17.0.2:3306)/otavio")

	if err != nil {
		panic(err)
	}
	defer db.Close() // não vai fechar a conexão aqui nessa linha, mas ao final de main() será fechada a conexão com o banco de dados

	execQuery(db, "Create database if not exists cursogo")
	execQuery(db, "use cursogo")
	execQuery(db, "drop table if exists usuarios")
	// DDL
	execQuery(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
	// execQuery(db, "insert into usuarios values(1, 'otavio')")
}

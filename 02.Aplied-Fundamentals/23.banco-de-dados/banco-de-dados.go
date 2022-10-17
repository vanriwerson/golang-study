package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // import implícito (outros arquivos usarão o pacote)
)

// para abrir a conexão com o db, instalar a seguinte dependência:
// go get github.com/go-sql-driver/mysql
func main() {
	// string de conexão: usuário, senha, configs
	strConexao := "root:password@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", strConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close() // fecha a conexão com o banco

	if erro = db.Ping(); erro != nil { // testa a conexão
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("SELECT * FROM devbook.usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}

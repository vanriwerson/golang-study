package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o mySQL
)

func Conectar() (*sql.DB, error) {
	strConexao := "root:password@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", strConexao)
	if erro != nil {
		return nil, erro //retornando o valor zero para db e um erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil // retornando o db e o valor zero para o erro
}

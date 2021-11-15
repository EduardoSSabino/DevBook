package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //faço essa importação driver manualmente
)

//no terminal eu faço o download do meu driver, github.com/go-sql-driver/mysql

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringDeConexaoBanco) //vamos conectar com o banco usando a string de conexão criada dentro do pacote config
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil { //Ping verifica a conxão com o banco de dados
		db.Close() //fechando a conexão com o banco de dados
		return nil, erro
	}

	return db, nil

	//QUEM VAI CHAMAR ESSA FUNÇÃO SERÁ OS CONTROLLERS!

}

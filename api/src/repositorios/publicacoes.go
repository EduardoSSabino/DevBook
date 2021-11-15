package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(Publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values(?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(Publicacao.Titulo, Publicacao.Conteudo, Publicacao.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarPorId traz uma única publicacao do banco de dados
func (repositorio Publicacoes) BuscarPorId(publicacaoId uint64) (modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select p,*, u.nick from ? 
	publicacoes p inner join usuarios u
	on u.id = p.autor_id where p.id = ?
	`, publicacaoId)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer linhas.Close()

	var publicacao modelos.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan( // a ordem tem que coincidir com o arquivo SQL
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

// Buscar traz as publicações dos usuários seguidos e também a do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select distinct p.*, u.nick from publicacoes p 
		inner join usuarios u o u.id = p.autor_id 
		inner join seguidores s on p.autor_id = s.usuario_id 
		where u.id = ? or s.seguidor_id = ?
		`, usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan( // a ordem tem que coincidir com o arquivo SQL
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// Atualizar altera os dadso de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoId uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare(`
	update publicacoes set titulos = ?, conteudo = ? where id = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.Id); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação no banco de dados
func (repositorio Publicacoes) Deletar(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare(`
	delete from publicacoes where id = ?
	`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

// BscarPorUsuari traz todas as publicações de um usuário específico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select p.*, u.nick from publicacoes p
	join usuarios u on u.id = p.autor_id
	where p.autor_id = ?
	`, usuarioId)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan( // a ordem tem que coincidir com o arquivo SQL
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publicacoes) Curtir(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare(`
	update publicacoes set curtidas = curtidas + 1
	where id = ?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

// Descurtir subtrai uma curtida da publicação
func (repositorio Publicacoes) Descurtir(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare(`
	update publicacoes set curtidas =
	CASE 	when	 curtidas > 0 THEN curtidas -1
	ELSE curtidas END where id = ?
	`)
	if erro != nil {
		return erro
	}
	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

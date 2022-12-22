package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicacao no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	stmt, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer stmt.Close()

	resultado, erro := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}

// BuscarPorID traz uma única publicacao do banco de dados
func (repositorio Publicacoes) BuscarPorID(puclicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query("select p.*, u.nick from publicacoes p inner join usuarios u on p.autor_id = u.id where p.id = ?", puclicacaoID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

// Buscar traz as publicacoes dos usuarios seguidos e também do prprio usuario que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query("select distinct p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id inner join seguidores s on p.autor_id = s.usuario_id where u.id = ? or s.seguidor_id = ? order by 1 desc", usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar Altera os dados de uma publicacao no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	stmt, erro := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Deletar Remove uma publicação do Usuario salva no banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	stmt, erro := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// BuscarPorUsuario traz todas as publicacoes de um determinado usuario
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query("select p.*, u.nick from publicacoes p join usuarios u on u.id = p.autor_id where p.autor_id = ?", usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(&publicacao.ID, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorID, &publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}

// Curtir adiciona uma curtida na publicacao
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	stmt, erro := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where publicacoes.id = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir adiciona uma curtida na publicacao
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	stmt, erro := repositorio.db.Prepare("update publicacoes set curtidas = IF(curtidas > 0, curtidas - 1, curtidas) where publicacoes.id = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

package modelos

import (
	"api/src/seguranca"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O Nome é Obrigatório e não pode estar em Branco")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick é Obrigatório e não pode estar em Branco")
	}
	if usuario.Email == "" {
		return errors.New("O Email é Obrigatório e não pode estar em Branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é invalido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O Senha é Obrigatório e não pode estar em Branco")
	}

	return nil
}

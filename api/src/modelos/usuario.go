package modelos

import (
	"errors"
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

func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O Nome é Obrigatório e não pode estar em Branco")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick é Obrigatório e não pode estar em Branco")
	}
	if usuario.Email == "" {
		return errors.New("O Email é Obrigatório e não pode estar em Branco")
	}
	if usuario.Senha == "" {
		return errors.New("O Senha é Obrigatório e não pode estar em Branco")
	}

	return nil
}

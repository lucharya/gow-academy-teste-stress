package programador

import (
	"errors"
	"time"
)

func ValidarProgramador(p Programador) error {
	if p.Apelido == "" || len(p.Apelido) > 32 {
		return errors.New("apelido inválido")
	}
	if p.Nome == "" || len(p.Nome) > 100 {
		return errors.New("nome inválido")
	}
	if !DataEhValida(p.Nascimento) {
		return errors.New("data de nascimento inválida")
	}
	for _, s := range p.Stack {
		if len(s) > 32 {
			return errors.New("item da stack inválido")
		}
	}
	return nil
}

func DataEhValida(data string) bool {
	_, err := time.Parse("2006-01-02", data)
	return err == nil
}

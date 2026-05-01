package service

import (
	"pulsecore-api/internal/model"
	"pulsecore-api/internal/repository"
)

type CreatePessoaRequest struct {
	Nome string `json:"nome" binding:"required,min=3"`
}

func ListarPessoas() ([]model.Pessoa, error) {
	return repository.Listar()
}

func CriarPessoa(req CreatePessoaRequest) (model.Pessoa, error) {
	p := model.Pessoa{
		Nome: req.Nome,
	}

	return repository.Criar(p)
}

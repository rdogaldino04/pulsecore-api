package service

import "pulsecore-api/internal/repository"

type CreatePessoaRequest struct {
	Nome string `json:"nome" binding:"required,min=3"`
}

func ListarPessoas() ([]repository.Pessoa, error) {
	return repository.Listar()
}

func CriarPessoa(req CreatePessoaRequest) (repository.Pessoa, error) {
	p := repository.Pessoa{
		Nome: req.Nome,
	}

	return repository.Criar(p)
}

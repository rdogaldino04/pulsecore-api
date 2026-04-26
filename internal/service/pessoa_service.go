package service

type Pessoa struct {
	ID   int    `json:"id"`
	Nome string `json:"name"`
}
type CreatePessoaRequest struct {
	Nome string `json:"nome" binding:"required,min=3"`
}

var pessoas = []Pessoa{
	{ID: 1, Nome: "João"},
	{ID: 2, Nome: "Maria"},
}

var nextID = 3

func ListarPessoas() []Pessoa {
	return pessoas
}

func CriarPessoa(req CreatePessoaRequest) Pessoa {
	pessoa := Pessoa{
		ID:   nextID,
		Nome: req.Nome,
	}

	pessoas = append(pessoas, pessoa)
	nextID++

	return pessoa
}

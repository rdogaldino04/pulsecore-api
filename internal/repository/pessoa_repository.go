package repository

type Pessoa struct {
	ID   int    `json:"id"`
	Nome string `json:"name"`
}

var pessoas = []Pessoa{
	{ID: 1, Nome: "João"},
	{ID: 2, Nome: "Maria"},
}

var nextID = 3

func Listar() []Pessoa {
	return pessoas
}

func Criar(p Pessoa) Pessoa {
	p.ID = nextID
	nextID++

	pessoas = append(pessoas, p)
	return p
}

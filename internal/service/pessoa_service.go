package service

type Pessoa struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListarPessoas() []Pessoa {
	return []Pessoa{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
}

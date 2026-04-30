package repository

type Pessoa struct {
	ID   int    `json:"id"`
	Nome string `json:"name"`
}

func Listar() ([]Pessoa, error) {
	rows, err := DB.Query("SELECT id, nome FROM pessoas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pessoas []Pessoa

	for rows.Next() {
		var p Pessoa
		err := rows.Scan(&p.ID, &p.Nome)
		if err != nil {
			return nil, err
		}
		pessoas = append(pessoas, p)
	}

	return pessoas, nil
}

func Criar(p Pessoa) (Pessoa, error) {
	result, err := DB.Exec("INSERT INTO pessoas (nome) VALUES (?)", p.Nome)
	if err != nil {
		return p, err
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)

	return p, nil
}

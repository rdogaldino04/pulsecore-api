package repository

import "pulsecore-api/internal/model"

func Listar() ([]model.Pessoa, error) {
	rows, err := DB.Query("SELECT id, nome FROM pessoas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pessoas []model.Pessoa

	for rows.Next() {
		var p model.Pessoa
		err := rows.Scan(&p.ID, &p.Nome)
		if err != nil {
			return nil, err
		}
		pessoas = append(pessoas, p)
	}

	return pessoas, nil
}

func Criar(p model.Pessoa) (model.Pessoa, error) {
	result, err := DB.Exec("INSERT INTO pessoas (nome) VALUES (?)", p.Nome)
	if err != nil {
		return p, err
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)

	return p, nil
}

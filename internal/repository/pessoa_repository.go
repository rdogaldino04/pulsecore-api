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
	err := DB.QueryRow(
		"INSERT INTO pessoas (nome) VALUES ($1) RETURNING id",
		p.Nome,
	).Scan(&p.ID)

	return p, err
}

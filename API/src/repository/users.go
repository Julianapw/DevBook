package repository

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB // Alterado de bd para db para consistência
}

// NewUserRepository cria um repositório de usuários
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insere um usuário no banco de dados
func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

func (repository users) Search (nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"select, id, name, nick, email, createdAt from users WHERE name LIKE ? or nick LIKE ?",
		nameOrNick,nameOrNick
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	
	for lines.Next(){
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			user.Email,
			user.CreatedAt
		)
	}
	users = append(users, user)
}
return users, nil

func (repository Users) SearchById (ID uint64) (models.User, error) {
	lines, err = repository.db.Query (
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if err != nil {
		return []models.User{}, err

	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt
		); err != nil{
			return models.User{}, err
		}
	}
	return user, nil
}
package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"gin-app/models"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=admin password=secret dbname=appdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("✅ connected to postgres")

	return db, nil
}

func CreateUserTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	);
	`
	_, err := db.Exec(query)
	return err
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

func (p *PostgresStorage) AddUser(name, email string) (models.User, error) {
	var u models.User
	query := `
	INSERT INTO users (name, email)
	VALUES ($1, $2)
	RETURNING id, name, email;
	`
	err := p.db.QueryRow(query, name, email).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

// TODO: these errors are too broad
func (p *PostgresStorage) GetUser(id int) (models.User, error) {
	// var name, email string
	var u models.User
	u.ID = id
	query := `
	SELECT name, email
	FROM users
	WHERE id = $1;
	`
	err := p.db.QueryRow(query, id).Scan(&u.Name, &u.Email)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (p *PostgresStorage) GetUsers() ([]models.User, error) {
	var users []models.User
	query := `
	SELECT id, name, email FROM users;
	`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (p *PostgresStorage) UpdateUser(id int, name, email string) (models.User, error) {
	var u models.User
	query := `
	UPDATE users
	SET name = $1, email = $2
	WHERE id = $3
	RETURNING id, name, email;
	`
	err := p.db.QueryRow(query, name, email, id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("User with ID:%d do not exists", id)
		}
		return models.User{}, err
	}

	return u, nil
}

func (p *PostgresStorage) DeleteUser(id int) error {
	query := `
	DELETE from users
	WHERE id = $1;
	`
	result, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("User with ID:%d does not exist", id)
	}

	return nil
}

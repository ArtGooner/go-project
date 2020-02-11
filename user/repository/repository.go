package repository

import (
	"crypto/sha256"
	"database/sql"
	pb "github.com/ArtGooner/go-project/user/proto"
)

type Repository interface {
	Get(user *pb.Account) (*pb.User, error)
}

type repository struct {
	db *sql.DB
}

func (r repository) Get(acc *pb.Account) (*pb.User, error) {
	h := sha256.New()
	h.Write([]byte(acc.GetPassword()))
	pwd := h.Sum(nil)
	//_, err := r.db.Exec("insert into Users (Email, PasswordHash, Name, Surname, Age) values ($1,$2,$3,$4,$5);", "email", pwd, "Name", "Surname", 100)
	rows, err := r.db.Query("select * from Users where Email=$1 AND PasswordHash = $2;", acc.GetEmail(), pwd)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		usr := &pb.User{}

		err := rows.Scan(&usr.Id, &usr.Email, &usr.PasswordHash, &usr.Name, &usr.Surname, &usr.Age)

		if err != nil {
			return nil, err
		}

		return usr, nil
	}

	return nil, nil
}

func NewRepository() (Repository, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=213612458 dbname=therion sslmode=disable")

	if err != nil {
		return nil, err
	}

	sqlStmt := `create table if not exists users (Id serial primary key, Email text, PasswordHash bytea, Name text, Surname text, Age integer)`
	_, err = db.Exec(sqlStmt)

	if err != nil {
		return nil, err
	}

	return &repository{db: db}, nil
}

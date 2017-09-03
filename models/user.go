package models

import (
	"github.com/michaelbui/bingo-backend/entities"
	"github.com/michaelbui/bingo-backend/helpers"
	"time"
	"errors"
)

type (
	UserModel struct {}
)

var (
	user *UserModel
)

func User() *UserModel {
	if user == nil {
		user = &UserModel{}
	}
	return user
}

func (u *UserModel) Add(e entities.UserEntity) (entities.UserEntity, error) {
	db := helpers.Database().Connect()
	defer db.Close()

	password, err := helpers.Database().HashPassword(e.GetPassword())
	if err != nil {
		return e, err
	}

	numbers, err := e.ParseNumbersForDB()
	if err != nil {
		return e, err
	}

	timestamp := time.Now().Unix()
	result, err := db.Exec(`
		INSERT INTO users (display, email, password, score, numbers, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, e.Display, e.Email, password, 0, string(numbers), timestamp)

	if err != nil {
		return e, err
	}

	rowId, err := result.LastInsertId()
	e.Id = int(rowId)
	e.Timestamp = int(timestamp)

	return e, err
}

func (u *UserModel) Get(email string, password string) (entities.UserEntity, error) {
	db := helpers.Database().Connect()
	defer db.Close()

	row := db.QueryRow(`
		SELECT
			id,
			display,
			email,
			password,
			score,
			numbers,
			timestamp
		FROM users
		WHERE email = $1;
	`, email)

	var (
		numbers string
		dbPassword string
		e entities.UserEntity
	)

	err := row.Scan(&e.Id, &e.Display, &e.Email, &dbPassword, &e.Score, &numbers, &e.Timestamp)
	if err != nil {
		return entities.UserEntity{}, err
	}

	if !helpers.Database().VerifyPassword(password, dbPassword) {
		return entities.UserEntity{}, errors.New("Unauthenticated")
	}

	n, err := e.ParseNumbersFromDB(numbers)
	if err != nil {
		return entities.UserEntity{}, err
	}

	e.Numbers = n
	return e, nil
}

func (u *UserModel) UpdateNumbers(e *entities.UserEntity, numbers []int) error {
	err := e.CheckChosenNumbers(numbers)
	if err != nil {
		return err
	}

	err = Number().CheckCalledNumbers(numbers)
	if err != nil {
		return err
	}

	e.CalculateScore()
	numberJson, err := e.ParseNumbersForDB()
	if err != nil {
		return err
	}

	db := helpers.Database().Connect()
	defer db.Close()
	_, err = db.Exec(
		"UPDATE users SET score=$1, numbers=$2 WHERE id=$3",
		e.Score,
		numberJson,
		e.Id,
	)
	return err
}

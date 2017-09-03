package entities

import (
	"encoding/json"
	"errors"
)

type (
	UserEntity struct {
		Id        int `json:"id"`
		Display   string `json:"display"`
		Email     string `json:"email"`
		Score     int `json:"score"`
		Numbers   UserNumbers `json:"numbers"`
		Timestamp int `json:"timestamp"`
		password  string
	}
	UserNumbers map[int]bool
)

func (u *UserEntity) VerifyPassword(password string, passwordConfirm string) bool {
	if password != passwordConfirm {
		return false
	}

	return len(password) >= 4
}

func (u *UserEntity) SetPassword(password string) *UserEntity {
	u.password = password
	return u
}

func (u *UserEntity) GetPassword() string {
	return u.password
}

func (u *UserEntity) ParseNumbersFromDB(raw string) (UserNumbers, error) {
	var numbers UserNumbers
	err := json.Unmarshal([]byte(raw), &numbers)
	return numbers, err
}

func (u *UserEntity) ParseNumbersForDB() (string, error) {
	numbers, err := json.Marshal(u.Numbers)
	if err != nil {
		return "", err
	}
	return string(numbers), nil
}

func (u *UserEntity) CalculateScore() int {
	score := 0
	for _, checked := range u.Numbers {
		if checked {
			score++
		}
	}
	u.Score = score
	return score
}

func (u *UserEntity) CheckChosenNumbers(numbers []int) error {
	for _, number := range numbers {
		if _, ok := u.Numbers[number]; !ok {
			return errors.New("Non-chosen number(s)")
		}
		u.Numbers[number] = true
	}
	return nil
}


package models

import (
	"github.com/michaelbui/bingo-backend/helpers"
	"github.com/labstack/gommon/log"
	"math/rand"
	"time"
	"fmt"
	"strings"
	"errors"
)

type (
	NumberModel struct{}
	NumberOrder []int
	NumberTimestamp map[int]int
)

type NumberRecord struct {
	Value     int `json:"value"`
	Timestamp int `json:"timestamp"`
}

var (
	number *NumberModel
)

func Number() *NumberModel {
	if number == nil {
		number = &NumberModel{}
	}
	return number
}

func (n *NumberModel) List() (NumberTimestamp, NumberOrder) {
	db := helpers.Database().Connect()
	defer db.Close()

	rows, err := db.Query(`
		SELECT
			value,
			timestamp
		FROM numbers
		ORDER BY timestamp ASC
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		numbers     NumberTimestamp = NumberTimestamp{}
		numberOrder NumberOrder     = NumberOrder{}
	)
	for rows.Next() {
		var value int
		var timestamp int
		err = rows.Scan(&value, &timestamp)
		if err != nil {
			log.Fatal(err)
		}
		numberOrder = append(numberOrder, value)
		numbers[value] = timestamp
	}
	return numbers, numberOrder
}

func (n *NumberModel) Next() int {
	numbers, _ := Number().List()
	randomNumber := Number().generateNewNumber(numbers)
	if randomNumber > 0 {
		Number().insertNewNumberIntoDB(randomNumber)
	}
	return randomNumber
}

func (n *NumberModel) generateNewNumber(numbersTimestamp NumberTimestamp) int {
	remainingNumbers := []int{}
	for i := 1; i <= 99; i++ {
		if _, ok := numbersTimestamp[i]; !ok {
			remainingNumbers = append(remainingNumbers, i)
		}
	}

	remainingCount := len(remainingNumbers)
	if remainingCount == 0 {
		return 0
	}

	randomIndex := rand.Intn(remainingCount)
	return remainingNumbers[randomIndex]
}

func (n *NumberModel) insertNewNumberIntoDB(number int) error {
	db := helpers.Database().Connect()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO numbers (value, timestamp) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(number, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
		return err
	}

	return tx.Commit()
}

func (n *NumberModel) CheckCalledNumbers(numbers []int) error {
	db := helpers.Database().Connect()
	defer db.Close()

	query := fmt.Sprintf(
		"SELECT value FROM numbers WHERE value IN (%s)",
		strings.Join(strings.Split(strings.Repeat("?", len(numbers)), ""), ", "),
	)

	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}

	numbersForDb := []interface{}{}
	for _,v := range numbers {
		numbersForDb = append(numbersForDb, v)
	}
	rows, err := statement.Query(numbersForDb...)
	if err != nil {
		return err
	}
	defer rows.Close()

	calledNumbers := []int{}
	for rows.Next() {
		var number int
		err = rows.Scan(&number)
		if err != nil {
			return err
		}
		calledNumbers = append(calledNumbers, number)
	}
	if len(calledNumbers) != len(numbers) {
		return errors.New("Uncalled Number(s)")
	}

	return nil
}

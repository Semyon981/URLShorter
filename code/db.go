package main

import (
	"database/sql"
	"fmt"
	"math/rand"
)

func GenerateWord() string {

	var runes []rune
	for i := 0; i < 5; i++ {

		asd := rand.Intn(52) + 65
		if asd > 90 {
			asd += 6
		}
		runes = append(runes, rune(asd))
	}

	return string(runes)
}

type DB struct {
	database *sql.DB
}

func (db *DB) GetAll() []string {
	row, err := db.database.Query("SELECT shorturl FROM db.urlshorter;")
	if err != nil {
		panic(err.Error())
	}
	var res []string
	for row.Next() {
		var str string
		row.Scan(&str)
		res = append(res, str)
	}

	return res
}

func (db *DB) GetLurl(surl string) string {
	row, err := db.database.Query("SELECT longurl FROM db.urlshorter WHERE shorturl = '" + surl + "';")
	if err != nil {
		panic(err.Error())
	}
	if row.Next() {
		var lurl string
		row.Scan(&lurl)
		return lurl
	} else {
		panic("Error")
	}
}

func (db *DB) Insert(url string) (string, bool) {

	row, _ := db.database.Query("SELECT shorturl FROM db.urlshorter WHERE longurl = '" + url + "';")
	if row.Next() {
		var surl string
		row.Scan(&surl)
		return surl, false
	} else {
		var surl string = GenerateWord()
		row, _ = db.database.Query("SELECT * FROM db.urlshorter WHERE shorturl = '" + surl + "';")
		for row.Next() {
			surl = GenerateWord()
			row, _ = db.database.Query("SELECT * FROM db.urlshorter WHERE shorturl = '" + surl + "';")
		}

		_, err := db.database.Query("INSERT INTO `db`.`urlshorter` (longurl, shorturl) VALUES ('" + url + "', '" + surl + "');")
		if err != nil {
			panic(err.Error())
		}

		return surl, true
	}

}

func (db *DB) openDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "usr", "qwerty123", "localhost", "db")
	db.database, _ = sql.Open("mysql", dsn)
}
func (db *DB) closeDB() {
	db.database.Close()
}

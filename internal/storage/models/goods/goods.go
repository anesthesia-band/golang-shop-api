package goods

import (
	"github.com/anesthesia-band/golang-shop-api/internal/storage"
)

type GoodType string
type GoodData string

const (
	Laser     GoodType = "laser"
	Stand     GoodType = "stand"
	PowerUnit GoodType = "power_unit"
)

type Good struct {
	Id       int      `db:"id"`
	Name     string   `db:"name"`
	GoodType GoodType `db:"type"`
	Data     GoodData `db:"name"`
	Price    string   `db:"price"`
}

type GoodInsert struct {
	Name     string
	GoodType GoodType
	Data     GoodData
	Price    string
}

// TODO: make fields optional
type GoodUpdate struct {
	Name     string
	GoodType GoodType
	Data     GoodData
	Price    string
}

func GetById(storage *storage.Storage, goodId int) (*Good, error) {
	stmt, err := storage.DB.Prepare("SELECT * FROM goods WHERE id = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(goodId)
	if err != nil {
		return nil, err
	}

	goods := []Good{}
	for result.Next() {
		var good Good
		err := result.Scan(&good)
		if err != nil {
			return nil, err
		}
	}
	return &goods[0], nil
}

func GetAll(storage *storage.Storage, active bool) (*[]Good, error) {
	stmt, err := storage.DB.Prepare("SELECT * FROM goods where active = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(active)
	if err != nil {
		return nil, err
	}

	goods := []Good{}
	for result.Next() {
		var good Good
		err := result.Scan(&good)
		if err != nil {
			continue
		}
	}
	return &goods, nil
}

func InsertGood(storage *storage.Storage, data GoodInsert) error {
	stmt, err := storage.DB.Prepare("INSERT INTO goods(name, type, data, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(data.Name, string(data.GoodType), data.Data, data.Price)
	return err
}

func UpdateGoodById(storage *storage.Storage, goodId int, data GoodUpdate) error {
	stmt, err := storage.DB.Prepare("UPDATE goods SET name = ?, type = ?, data = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(data.Name, string(data.GoodType), data.Data, data.Price, goodId)
	return err
}

func DeleteGoodById(storage *storage.Storage, goodId int) error {
	stmt, err := storage.DB.Prepare("DELETE FROM goods WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(goodId)
	return err
}

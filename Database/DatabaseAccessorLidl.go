package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type LidlPrice struct {
	id         int
	Ing        string
	Price      float32
	Discount   float32
	RepalceFor string
}

type DatabaseAccessorLidl struct {
	databaseLidl *sql.DB
}

func (accessor *DatabaseAccessorLidl) OpenDB(databasePath string) {
	database, err := sql.Open("mysql", databasePath)
	if err != nil {
		fmt.Println("Error occured while creating database")
		return
	}
	accessor.databaseLidl = database
}

func (accessor *DatabaseAccessorLidl) CreateDatabaseLidl() bool {
	statement, err := accessor.databaseLidl.Prepare(`CREATE TABLE IF NOT EXISTS LidlTable
		(id INTEGER AUTO_INCREMENT, ing TEXT, price REAL, discount REAL, replaceFor TEXT, PRIMARY KEY (id))`)
	if err != nil {
		fmt.Println("Error occured while creating database ", err)
		return false
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Println("Error occured while exec creation of database ", err)
		return false
	}

	return true
}

func (accessor *DatabaseAccessorLidl) AddRowToIngs(data LidlPrice) bool {
	statement, err := accessor.databaseLidl.Prepare(`INSERT INTO LidlTable
        (id, ing, price, discount, replaceFor) VALUES (NULL, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Println("Error AddRowToIngs prep" + err.Error())
		return false
	}
	_, err = statement.Exec(data.Ing, data.Price, data.Discount, data.RepalceFor)
	if err != nil {
		fmt.Println("Error AddRowToIngs exec" + err.Error())
		return false
	}
	return true
}

func (accessor *DatabaseAccessorLidl) GetPrice(ing string) *LidlPrice {
	statement, err := accessor.databaseLidl.Prepare(`SELECT * FROM LidlTable
        WHERE replaceFor = ?`)
	if err != nil {
		fmt.Println("Error GetRowFromIngs prep" + err.Error())
		return nil
	}
	defer statement.Close()

	data := new(LidlPrice)
	err = statement.QueryRow(ing).Scan(&data.id, &data.Ing, &data.Price, &data.Discount, &data.RepalceFor)
	if err != nil {
		fmt.Println("Error GetRowFromIngs exec" + err.Error())
		return nil
	}
	return data
}

func (accessor *DatabaseAccessorLidl) Shutdown() bool {
	accessor.databaseLidl.Close()
	return true
}

package Database

import (
	"database/sql"
	"fmt"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

type IngData struct {
	id  int
	Ing string
	GU  uint32
	EC  uint32
	WU  uint32
}

type DatabaseAccessor struct {
	databasePolution *sql.DB
}

func (accessor *DatabaseAccessor) OpenDB(databasePath string) {
	database, err := sql.Open("mysql", databasePath)
	if err != nil {
		fmt.Println("Error occured while creating database")
		return
	}
	accessor.databasePolution = database
}

func (accessor *DatabaseAccessor) CreateDatabasePolution() bool {
	statement, err := accessor.databasePolution.Prepare(`CREATE TABLE IF NOT EXISTS FoodPolutionTable
		(id INTEGER AUTO_INCREMENT, ing TEXT, gw REAL, ec REAL, wu REAL, PRIMARY KEY (id))`)
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

func (accessor *DatabaseAccessor) AddRowToIngs(data IngData) bool {
	statement, err := accessor.databasePolution.Prepare(`INSERT INTO FoodPolutionTable
        (id, ing, gw, ec, we) VALUES (NULL, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Println("Error AddRowToIngs prep" + err.Error())
		return false
	}
	_, err = statement.Exec(data.Ing, data.GU, data.EC, data.WU)
	if err != nil {
		fmt.Println("Error AddRowToIngs exec" + err.Error())
		return false
	}
	return true
}

func (accessor *DatabaseAccessor) GetRowFromIngs(ing string) *IngData {
	statement, err := accessor.databasePolution.Prepare(`SELECT * FROM FoodPolutionTable
        WHERE ing = ?`)
	if err != nil {
		fmt.Println("Error GetRowFromIngs prep" + err.Error())
		return nil
	}
	defer statement.Close()

	data := new(IngData)
	err = statement.QueryRow(ing).Scan(&data)
	if err != nil {
		fmt.Println("Error GetRowFromIngs exec" + err.Error())
		return nil
	}
	return data
}

func (accessor *DatabaseAccessor) Shutdown() bool {
	accessor.databasePolution.Close()
	return true
}

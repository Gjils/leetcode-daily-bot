package groupsService

import (
	"database/sql"
	"fmt"

	dbLayer "leetcodebot/data/db"
)

type GroupsService struct {
	db *sql.DB
}

func (service GroupsService) GetAll() ([]int, error) {
	SQL := "SELECT identifier FROM groups"
	rows, err := service.db.Query(SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var idList []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		idList = append(idList, id)
	}
	return idList, nil
}

func (service GroupsService) Add(id int) error {
	SQL := "INSERT INTO groups (identifier) VALUES (?)"
	statement, err := service.db.Prepare(SQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("Added", id)
	return nil
}


func (service GroupsService) Check(id int) (bool, error) {
	SQL := "SELECT identifier FROM groups WHERE identifier = ?"
	err := service.db.QueryRow(SQL, id).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
				return false, err
		}
		return false, nil
	}
	return true, nil
}



func (service GroupsService) Delete(id int) error {
	SQL := "DELETE FROM groups WHERE identifier = ?"
	statement, err := service.db.Prepare(SQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func GetApi() GroupsService {
	return GroupsService{db: dbLayer.DB}
}

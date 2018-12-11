package db

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"os"
	"strings"
)

var (
	DBCon *gorm.DB
)

const (
	TABLENAME = "restaurants"
)

func InitDB(file string) error {
	sqlStatement, err := parseCsv(file)
	if err != nil {
		return err
	}
	err = DBCon.Exec(fmt.Sprint("TRUNCATE ", TABLENAME)).Error
	if err != nil {
		return err
	}

	err = DBCon.Exec(sqlStatement).Error
	if err != nil {
		return err
	}
	return nil
}

func parseCsv(file string) (string, error) {
	sqlStatement := bytes.Buffer{}
	var sqlStatementString string

	f, err := os.Open(file)
	if err != nil {
		return sqlStatementString, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	sqlStatement.WriteString("Insert into restaurants (")
	columns,_ := csvr.Read()
	sqlStatement.WriteString(strings.Join(columns, ","))
	sqlStatement.WriteString(") VALUES")
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return sqlStatementString, err
		}
		sqlStatement.WriteString("('")
		sqlStatement.WriteString(strings.Join(row, "','"))
		sqlStatement.WriteString("'),")
	}
	sqlStatementString = string(bytes.Trim(sqlStatement.Bytes(), ","))
	return sqlStatementString, nil
}


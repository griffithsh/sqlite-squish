package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func fmtFloat(f float64) string {
	tmp := strings.TrimRight(fmt.Sprintf("%f", f), "0")
	if strings.HasSuffix(tmp, ".") {
		tmp = tmp + "0"
	}
	return tmp
}

func getInserts(db *sql.DB, table string) (string, error) {
	rows, err := db.Query("SELECT * FROM " + table + ";")
	if err != nil {
		return "", fmt.Errorf("Query: %s", err)
	}
	defer rows.Close()

	inserts := []string{}
	for rows.Next() {
		ct, err := rows.ColumnTypes()
		if err != nil {
			return "", fmt.Errorf("ColumnTypes: %s", err)
		}
		tmp := make([]interface{}, len(ct))
		for i, ty := range ct {
			if ty.DatabaseTypeName() == "DOUBLE" {
				tmp[i] = &sql.NullFloat64{}
				continue
			}
			if ty.DatabaseTypeName() == "BOOLEAN" {
				tmp[i] = &sql.NullBool{}
				continue
			}
			if ty.ScanType() == nil {
				tmp[i] = &sql.NullString{}
				continue
			}
			switch ty.ScanType().Kind() {
			case reflect.Int64:
				fallthrough
			case reflect.Int:
				var x int
				tmp[i] = &x
			case reflect.String:
				var x string
				tmp[i] = &x
			case reflect.Float64:
				var x float64
				tmp[i] = &x
			case reflect.Bool:
				panic("bool type detected")
			default:
				return "", fmt.Errorf("unknown column type %s", ty.ScanType())
			}
		}
		err = rows.Scan(tmp...)
		if err != nil {
			fmt.Printf("About to bail on a getInserts %s\n%v\n", table, tmp)
			return "", fmt.Errorf("Scan: %s", err)
		}
		values := []string{}
		for _, it := range tmp {
			switch v := it.(type) {
			case *string:
				values = append(values, fmt.Sprintf("'%s'", strings.Replace(*v, "'", "''", -1)))
			case *int:
				values = append(values, fmt.Sprintf("%d", *v))
			case *float64:
				values = append(values, fmtFloat(*v))
			case *sql.NullFloat64:
				if !v.Valid {
					values = append(values, "NULL")
					continue
				}
				values = append(values, fmtFloat(v.Float64))
			case *sql.NullBool:
				if !v.Valid {
					values = append(values, "NULL")
				} else if v.Bool {
					values = append(values, "1")
				} else {
					values = append(values, "0")
				}
			default:
				values = append(values, "NULL")
			}
		}
		inserts = append(inserts, fmt.Sprintf("INSERT INTO \"%s\" VALUES(%s);", table, strings.Join(values, ",")))
	}

	return strings.Join(inserts, "\n"), nil
}

func inputDBFile(file string) (string, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return "", fmt.Errorf("Open: %s", err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT tbl_name,sql FROM sqlite_master WHERE type = 'table';`)
	if err != nil {
		return "", fmt.Errorf("Query: %s", err)
	}
	defer rows.Close()
	var output string
	for rows.Next() {
		var ntable, nsql sql.NullString
		err := rows.Scan(&ntable, &nsql)
		if err != nil {
			return "", fmt.Errorf("Scan: %s", err)
		}
		table := ntable.String
		sql := nsql.String

		// Don't persist sqlite_sequence
		if table == "sqlite_sequence" {
			continue
		}

		// "CREATE TABLE ..." captured.
		output = output + sql + ";\n"

		// Capture all "INSERT INTO ..."s.
		inserts, err := getInserts(db, table)
		if err != nil {
			return "", fmt.Errorf("getInserts for %s: %s", table, err)
		}
		output = output + inserts
	}

	return output, nil
}

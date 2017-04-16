package database

import (
	"strings"
	"testing"
)

func TestFromString(t *testing.T) {
	d, err := FromString("")
	if len(d.Tables) > 0 {
		t.Errorf("Unexpected length %d", len(d.Tables))
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestSplitStatementsWithSemicolons(t *testing.T) {
	create := `CREATE TABLE [Line] (Id INTEGER PRIMARY KEY AUTOINCREMENT,Scene_Id INTEGER REFERENCES Scene(Id), Content TEXT);`
	insert := `INSERT INTO [Line] (Id,Scene_Id,Content) VALUES (1234,2,'O Romeo, Romeo! wherefore art thou Romeo? Deny thy father and refuse thy name; Or, if thou wilt not, be but sworn my love, and I'll no longer be a Capulet.);`

	stmts := splitStatements(create + insert)
	if len(stmts) != 2 {
		t.Errorf("Unexpected number of statements %d,%v", len(stmts), strings.Join(stmts, ";"))
	}
	if stmts[0] != create {
		t.Errorf("Corruption of statment %s, should be %s", stmts[0], create)
	}
	if stmts[1] != insert {
		t.Errorf("Corruption of statment %s, should be %s", stmts[1], insert)
	}
}

func TestSortedTables(t *testing.T) {
	a := "CREATE TABLE Movie (Id INTEGER AUTOINCREMENT PRIMARY KEY, Director_Id INTEGER REFERENCES Director(Id));"
	b := "CREATE TABLE Director (Id INTEGER AUTOINCREMENT PRIMARY KEY, Person_Id REFERENCES Person(Id));"
	c := "CREATE TABLE Actor (Id INTEGER AUTOINCREMENT PRIMARY KEY, Person_Id REFERENCES Person(Id));"
	d := "CREATE TABLE CastMember (Id INTEGER AUTOINCREMENT PRIMARY KEY, Movie_Id REFERENCES Movie(Id), Actor_Id REFERENCES Actor(Id));"
	e := "CREATE TABLE Person (Id INTEGER AUTOINCREMENT PRIMARY KEY, Name TEXT);"

	input := []string{b, e, d, a, c}

	db, err := FromString(strings.Join(input, ""))
	if err != nil {
		t.Error(err)
		return
	}
	tables, err := db.SortedTables()
	if err != nil {
		t.Error(err)
		return
	}
	if tables[0] != "Person" {
		t.Errorf("Incorrect order, first result should be %s but was %s", "Person", tables[0])
	}
	if tables[1] != "Actor" {
		t.Errorf("Incorrect order, second result should be %s but was %s", "Actor", tables[1])
	}
	if tables[2] != "Director" {
		t.Errorf("Incorrect order, third result should be %s but was %s", "Director", tables[2])
	}
	if tables[3] != "Movie" {
		t.Errorf("Incorrect order, fourth result should be %s but was %s", "Movie", tables[3])
	}
	if tables[4] != "CastMember" {
		t.Errorf("Incorrect order, fifth result should be %s but was %s", "CastMember", tables[4])
	}
}

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

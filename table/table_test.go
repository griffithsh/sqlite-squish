package table

import (
	"testing"

	"github.com/griffithsh/sql-squish/statement"
)

func TestAddMultipleCreates(t *testing.T) {
	table := Table{Name: "Test"}

	if err := table.Add(statement.Statement{Verb: statement.Create}); err != nil {
		t.Error("unexpected Error on initial create")
	}
	if err := table.Add(statement.Statement{Verb: statement.Insert}); err != nil {
		t.Error("unexpected Error on first insert")
	}
	if err := table.Add(statement.Statement{Verb: statement.Create}); err == nil {
		t.Error("missing error on subsequent create")
	}
}

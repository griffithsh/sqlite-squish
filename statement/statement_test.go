package statement

import (
	"sort"
	"testing"
)

func TestSortCreateBeforeInsert(t *testing.T) {
	s := Statements{
		{Verb: Insert},
		{Verb: Create},
	}

	sort.Sort(s)

	if s[0].Verb != Create {
		t.Errorf("first element should have been %s but was %s", Create, s[0].Verb)
	}
	if s[1].Verb != Insert {
		t.Errorf("second element should have been %s but was %s", Insert, s[0].Verb)
	}
}

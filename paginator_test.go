package paginator

import (
	"testing"
)

func TestFieldsRequired(t *testing.T) {
	/*
		https://graphql.org/swapi-graphql
	*/
	totalCount := 87
	first := 20
	last := 20
	var tests = []struct {
		first  *int
		last   *int
		after  string
		before string
		skip   int
		total  int
		limit  int
		offset int
		from   int
		to     int
	}{
		{&first, nil, "", "", 0, totalCount, 20, 0, 0, 19},
		{nil, &last, "", "", 0, totalCount, 20, 66, 67, 87},
		{&first, nil, EncodeCursor("cursor", 9), "", 0, totalCount, 20, 9, 10, 29},
	}

	for i, test := range tests {
		paginator, err := NewPaginator("cursor", test.first, test.last, &test.after, &test.before, &test.skip, test.total)
		if err != nil {
			t.Errorf("Got Error on NewPaginator(%q): %s", test, err)
		} else {
			if paginator.Limit() != test.limit {
				t.Errorf("Got Error on case %d limit doest not match %d : %d", i, test.limit, paginator.Limit())
			}
			if paginator.Offset() != test.offset {
				t.Errorf("Got Error on case %d offset doest not match %d : %d", i, test.offset, paginator.Offset())
			}
			if paginator.From() != test.from {
				t.Errorf("Got Error on case %d from doest not match %d : %d", i, test.from, paginator.From())
			}
			if paginator.To() != test.to {
				t.Errorf("Got Error on case %d to doest not match %d : %d", i, test.to, paginator.To())
			}
		}
	}
}

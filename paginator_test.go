package paginator

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Test struct {
	gorm.Model
	Name      string
	FirstName string
	LastName  string
	Email     string
}

func TestPage(t *testing.T) {
	/*
		https://graphql.org/swapi-graphql
		allPeople
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
		{&first, nil, "", "", 0, totalCount, first, 0, 0, 19},
		{nil, &last, "", "", 0, totalCount, last, 67, 67, 86},
		{&first, nil, EncodeCursor("cursor", 9), "", 0, totalCount, 20, 10, 10, 29},
		{nil, &last, "", EncodeCursor("cursor", 50), 0, totalCount, 20, 30, 30, 49},
		{&first, nil, EncodeCursor("cursor", 10), EncodeCursor("cursor", 15), 0, totalCount, 4, 11, 11, 14},
		{&first, nil, EncodeCursor("cursor", 10), EncodeCursor("cursor", 50), 0, totalCount, 20, 11, 11, 30},
		{nil, &last, EncodeCursor("cursor", 10), EncodeCursor("cursor", 15), 0, totalCount, 4, 11, 11, 14},
		{nil, &last, EncodeCursor("cursor", 10), EncodeCursor("cursor", 50), 0, totalCount, 20, 30, 30, 49},
		{nil, &last, EncodeCursor("cursor", 50), EncodeCursor("cursor", 10), 0, totalCount, 0, 51, 51, 9},
		{&first, nil, EncodeCursor("cursor", 50), EncodeCursor("cursor", 10), 0, totalCount, 0, 51, 51, 9},
	}

	for i, test := range tests {
		paginator, err := NewPaginator("cursor", test.first, test.last, &test.after, &test.before, &test.skip, test.total)
		if err != nil {
			t.Errorf("Got Error on NewPaginator(%v): %s", test, err)
		} else {
			if paginator.Limit() != test.limit {
				t.Errorf("Got Error on case %d Limit doest not match %d : %d", i, test.limit, paginator.Limit())
			}
			if paginator.Offset() != test.offset {
				t.Errorf("Got Error on case %d Offset doest not match %d : %d", i, test.offset, paginator.Offset())
			}
			if paginator.From() != test.from {
				t.Errorf("Got Error on case %d From doest not match %d : %d", i, test.from, paginator.From())
			}
			if paginator.To() != test.to {
				t.Errorf("Got Error on case %d To doest not match %d : %d", i, test.to, paginator.To())
			}
		}
	}
}

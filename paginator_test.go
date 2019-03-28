package paginator

import (
	"testing"
)

func TestFieldsRequired(t *testing.T) {
	/*
		case 1:
		$paginator = new Paginator($first: 100, $last: null, $after: null, $before: null, $total: 1000);
		$result: select * from testing limit 100

		case 2:
		$paginator = new Paginator($first: 0, $last: 100, $after: null, $before: null, $total: 1000);
		$result: select * from testing limit 900, 1000

		case 3
		$paginator = new Paginator($first: 100, $last: null, $after: 100, $before: null, $total: 1000);
		$result: select * from testing limit 101, 201

		case 4
		$paginator = new Paginator($first: null, $last: 100, $after: null, $before: 100, $total: 1000);
		$result: select * from testing limit 0, 99

		case 5
		$paginator = new Paginator($first: 100, $last: null, $after: 50, $before: 80, $total: 1000);
		$result: select * from testing limit 51, 79

		case 6
		$paginator = new Paginator($first: 100, $last: null, $after: 50, $before: 180, $total: 1000);
		$result: select * from testing limit 51, 151

		case 7
		$paginator = new Paginator($first: null, $last: 100, $after: 50, $before: 80, $total: 1000);
		$result: select * from testing limit 819, 919

		case 8
		$paginator = new Paginator($first: null, $last: 100, $after: 880, $before: 80, $total: 1000);
		$result: select * from testing limit 880, 919
	*/
	number100 := 100
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
		{&number100, nil, "", "", 0, 1000, 100, 0, 0, 99},
		{nil, &number100, "", "", 0, 1000, 100, 900, 900, 1000},
		{&number100, nil, EncodeCursor("cursor", 100), "", 0, 1000, 100, 0, 0, 99},
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

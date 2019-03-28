package paginator

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
)

type Paginator struct {
	cursorPrefix string
	first        *int
	last         *int
	after        *string
	before       *string
	total        int
	limit        int
	offset       int
	from         int
	to           int
}

func NewPaginator(cursorPrefix string, first *int, last *int, after *string, before *string, skip *int, total int) (*Paginator, error) {
	var p = Paginator{
		cursorPrefix: cursorPrefix,
		first:        first,
		last:         last,
		after:        after,
		before:       before,
		total:        total,
		from:         0,
		to:           0,
		limit:        0,
		offset:       0,
	}
	var err error
	if first != nil && last != nil {
		return nil, errors.New("Passing both `first` and `last` is not supported.")
	}

	if first != nil && *first > 100 {
		return nil, errors.New("Requesting " + strconv.Itoa(*first) + " records on the connection exceeds the `first` limit of 100 records.")
	}

	if last != nil && *last > 100 {
		return nil, errors.New("Requesting " + strconv.Itoa(*last) + " records on the connection exceeds the `last` limit of 100 records.")
	}

	if after != nil && *after != "" {
		p.from, err = DecodeCursor(p.cursorPrefix, *after)
		if err != nil {
			return nil, errors.New("`" + *after + "` does not appear to be a valid cursor.")
		}
		p.from++
		p.from = Min(p.from, total-1)
	}

	if before != nil && *before != "" {
		p.to, err = DecodeCursor(p.cursorPrefix, *before)
		if err != nil {
			return nil, errors.New("`" + *before + "` does not appear to be a valid cursor.")
		}
		p.to--
		p.to = Min(p.to, 0)
	}

	if first != nil {
		p.to = Min(p.from+*first, total-1)
	}

	if last != nil {
		p.from = Max(p.to-*last, 0)
	}

	if skip != nil {
		p.from = p.from - *skip
	}

	p.offset = p.from
	if p.to >= p.from {
		p.limit = p.to - p.from
	}

	return &p, nil
}

func (p Paginator) From() int {
	return p.from
}

func (p Paginator) To() int {
	return p.to
}

func (p Paginator) Limit() int {
	return p.limit
}

func (p Paginator) Offset() int {
	return p.offset
}

func (p Paginator) HasNextPage() bool {
	if p.total > p.to {
		return true
	}
	return false
}

func (p Paginator) HasPreviousPage() bool {
	if p.from > 0 {
		return true
	}
	return false
}

// EncodeCursor the cursot position in base64
func EncodeCursor(cursorPrefix string, index int) string {
	return base64.StdEncoding.EncodeToString([]byte(cursorPrefix + strconv.Itoa(index)))
}

// DecodeCursor decodes the base 64 encoded cursor and resturns the integer
func DecodeCursor(cursorPrefix string, cursor string) (int, error) {
	c, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(strings.TrimPrefix(string(c), cursorPrefix))
	if err != nil {
		return 0, err
	}

	return i, nil
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

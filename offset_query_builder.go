package mysql_query_builder

import "strconv"

type offsetQueryBuilder struct {
	offset string
}

func (qb *offsetQueryBuilder) Offset(offset int) {
	qb.offset = strconv.Itoa(offset)
}

func (qb *offsetQueryBuilder) getOffsetPart() string {
	if qb.offset == "" {
		return ""
	}
	return "OFFSET " + qb.offset
}

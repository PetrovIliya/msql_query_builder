package msql_query_builder

import "strconv"

type limitQueryBuilder struct {
	limit string
}

func (qb *limitQueryBuilder) Limit(limit int) {
	qb.limit = strconv.Itoa(limit)
}

func (qb *limitQueryBuilder) getLimitPart() string {
	if qb.limit == "" {
		return ""
	}
	return "LIMIT " + qb.limit
}

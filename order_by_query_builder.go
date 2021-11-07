package msql_query_builder

type orderByQueryBuilder struct {
	orderBy string
}

func (qb *orderByQueryBuilder) OrderBy(orderBy string) {
	qb.orderBy = orderBy
}

func (qb *orderByQueryBuilder) getOrderByPart() string {
	if qb.orderBy == "" {
		return ""
	}
	return "ORDER BY " + qb.orderBy
}

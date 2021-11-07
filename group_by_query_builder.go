package msql_query_builder

type groupByQueryBuilder struct {
	groupBy string
}

func (qb *groupByQueryBuilder) GroupBy(groupBy string) {
	qb.groupBy = groupBy
}

func (qb *groupByQueryBuilder) getGroupByPart() string {
	if qb.groupBy == "" {
		return ""
	}
	return "GROUP BY " + qb.groupBy
}

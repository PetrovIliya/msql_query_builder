package msql_query_builder

type DeleteQueryBuilder struct {
	baseQueryBuilder
	table     string
	deleteStr string
}

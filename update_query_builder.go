package msql_query_builder

type UpdateQueryBuilder struct {
	baseQueryBuilder
	table      string
	setStrings []string
}

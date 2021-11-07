package msql_query_builder

func Select(table string, alias string, selectStr string) *selectQueryBuilder {
	return &selectQueryBuilder{table: table, selectStr: selectStr, alias: alias}
}

func UnionSelect(table string, alias string, selectStr string) *unionSelectQueryBuilder {
	return &unionSelectQueryBuilder{table: table, selectStr: selectStr, alias: alias}
}

func Update(table string, alias string) *updateQueryBuilder {
	return &updateQueryBuilder{table: table, alias: alias}
}

func Delete(table string, alias string, deleteStr string) *deleteQueryBuilder {
	return &deleteQueryBuilder{table: table, deleteStr: deleteStr, alias: alias}
}

func InsertInto(table string, insertString string) *insertQueryBuilder {
	return &insertQueryBuilder{table: table, insertString: insertString, valuesUnionSubQuery: nil, valuesSelectSubQuery: nil}
}

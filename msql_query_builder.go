package msql_query_builder

func Select(table string, alias string, selectStr string) *SelectQueryBuilder {
	return &SelectQueryBuilder{table: table, selectStr: selectStr, alias: alias}
}

func UnionSelect(table string, alias string, selectStr string) *UnionSelectQueryBuilder {
	return &UnionSelectQueryBuilder{table: table, selectStr: selectStr, alias: alias}
}

func Update(table string, alias string) *UpdateQueryBuilder {
	return &UpdateQueryBuilder{table: table, alias: alias}
}

func Delete(table string, alias string, deleteStr string) *DeleteQueryBuilder {
	return &DeleteQueryBuilder{table: table, deleteStr: deleteStr, alias: alias}
}

func InsertInto(table string, insertString string) *InsertQueryBuilder {
	return &InsertQueryBuilder{table: table, insertString: insertString, valuesUnionSubQuery: nil, valuesSelectSubQuery: nil}
}

package msql_query_builder

func Select(table string, alias string, selectStr string) *SelectQueryBuilder {
	return &SelectQueryBuilder{table: table, selectStr: selectStr, alias: alias}
}

func UnionSelect(table string, alias string, selectStr string) *UnionSelectQueryBuilder {
	return &UnionSelectQueryBuilder{table: table, selectStr: selectStr, alias: alias}
}

func Update(table string) *UpdateQueryBuilder {
	return &UpdateQueryBuilder{table: table}
}

func Delete(table string, deleteStr string) *DeleteQueryBuilder {
	return &DeleteQueryBuilder{table: table, deleteStr: deleteStr}
}

func Insert(table string, values [][]string) *InsertQueryBuilder {
	return &InsertQueryBuilder{table: table, values: values}
}

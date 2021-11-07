package msql_query_builder

import "errors"

type insertQueryBuilder struct {
	table                string
	insertString         string
	values               [][]string
	ignore               bool
	valuesSelectSubQuery *selectQueryBuilder
	valuesUnionSubQuery  *unionSelectQueryBuilder
}

func (qb *insertQueryBuilder) GetSql() (string, error) {
	if err := qb.validate(); err != nil {
		return "", err
	}

	valuesPart, err := qb.getValuesPart()
	if err != nil {
		return "", err
	}
	return escape(qb.getInsertPart() + " " + valuesPart), nil
}

func (qb *insertQueryBuilder) Ignore(ignore bool) {
	qb.ignore = ignore
}

func (qb *insertQueryBuilder) Values(values [][]string) {
	qb.valuesUnionSubQuery = nil
	qb.valuesSelectSubQuery = nil
	qb.values = values
}

func (qb *insertQueryBuilder) ValuesQuery(query *selectQueryBuilder) {
	qb.valuesUnionSubQuery = nil
	qb.values = make([][]string, 0)
	qb.valuesSelectSubQuery = query
}

func (qb *insertQueryBuilder) ValuesUnionQuery(query *unionSelectQueryBuilder) {
	qb.valuesSelectSubQuery = nil
	qb.values = make([][]string, 0)
	qb.valuesUnionSubQuery = query
}

func (qb *insertQueryBuilder) getInsertPart() string {
	ignoreStr := ""
	if qb.ignore {
		ignoreStr = "IGNORE"
	}
	return "INSERT " + ignoreStr + " INTO `" + qb.table + "` " + qb.insertString
}

func (qb *insertQueryBuilder) getValuesPart() (string, error) {
	if len(qb.values) > 0 {
		return getValuesString(qb.values), nil
	} else if qb.valuesSelectSubQuery != nil {
		sql, err := qb.valuesSelectSubQuery.GetSql()
		if err != nil {
			return "", err
		}
		return sql, nil
	} else {
		sql, err := qb.valuesUnionSubQuery.GetSql()
		if err != nil {
			return "", err
		}
		return sql, nil
	}
}

func (qb *insertQueryBuilder) validate() error {
	if qb.insertString == "" {
		return errors.New("'insertString' param can not be empty")
	}

	if len(qb.values) == 0 && qb.valuesSelectSubQuery == nil && qb.valuesUnionSubQuery == nil {
		return errors.New("no values")
	}

	if len(qb.values) != 0 {
		for i := 0; i < len(qb.values); i++ {
			if len(qb.values[i]) == 0 {
				return errors.New("one of values empty")
			}
		}
	}

	if qb.table == "" {
		return errors.New("'table' param can not be empty")
	}

	return nil
}

func getValuesString(values [][]string) string {
	valuesStr := ""
	for i := 0; i < len(values); i++ {
		valuesStr += "(" + values[i][0]
		for j := 1; j < len(values[i]); j++ {
			value := values[i][j]
			valuesStr += ", " + value
		}
		valuesStr += "), "
	}

	return valuesStr[:len(valuesStr)-3]
}

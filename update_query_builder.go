package msql_query_builder

import "errors"

type setter struct {
	filed string
	value string
}

type UpdateQueryBuilder struct {
	joinQueryBuilder
	whereQueryBuilder
	orderByQueryBuilder
	limitQueryBuilder
	table   string
	alias   string
	setters []setter
}

func (qb *UpdateQueryBuilder) GetSql() (string, error) {
	err := qb.validate()
	if err != nil {
		return "", err
	}

	sql := escape(qb.getUpdatePart() + " " +
		qb.getJoinsPart() + " " +
		qb.getWherePart() + " " +
		qb.getOrderByPart() + " " +
		qb.getLimitPart() + " ",
	)

	return sql, nil
}

func (qb *UpdateQueryBuilder) Set(field string, value string) {
	qb.setters = append(qb.setters, setter{filed: field, value: value})
}

func (qb *UpdateQueryBuilder) SetSubQuery(field string, subQuery *SelectQueryBuilder) error {
	sql, err := subQuery.GetSql()
	if err != nil {
		return err
	}
	qb.setters = append(qb.setters, setter{filed: field, value: "(" + sql + ")"})
	return nil
}

func (qb *UpdateQueryBuilder) getUpdatePart() string {
	return "UPDATE " + qb.table + " " + qb.alias
}

func (qb *UpdateQueryBuilder) getSetPart() string {
	firstSetter := qb.setters[0]
	setStr := "SET " + firstSetter.filed + " = " + firstSetter.value

	for i := 1; i < len(qb.setters); i++ {
		currSetter := qb.setters[i]
		setStr += ", " + currSetter.filed + " = " + currSetter.value
	}

	return setStr
}

func (qb *UpdateQueryBuilder) validate() error {
	err := qb.whereQueryBuilder.validate()
	if err != nil {
		return err
	}

	err = qb.joinQueryBuilder.validate()
	if err != nil {
		return err
	}

	if qb.table == "" {
		return errors.New("'table' param can not be empty")
	}

	if len(qb.setters) == 0 {
		return errors.New("no setters")
	}

	return nil
}

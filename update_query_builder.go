package mysql_query_builder

import "errors"

type setter struct {
	filed string
	value string
}

type updateQueryBuilder struct {
	joinQueryBuilder
	whereQueryBuilder
	orderByQueryBuilder
	limitQueryBuilder
	table   string
	alias   string
	setters []setter
}

func (qb *updateQueryBuilder) GetSql() (string, error) {
	err := qb.validate()
	if err != nil {
		return "", err
	}

	sql := escape(qb.getUpdatePart() + " " +
		qb.getJoinsPart() + " " +
		qb.getSetPart() + " " +
		qb.getWherePart() + " " +
		qb.getOrderByPart() + " " +
		qb.getLimitPart() + " ",
	)

	return sql, nil
}

func (qb *updateQueryBuilder) Set(field string, value string) {
	qb.setters = append(qb.setters, setter{filed: field, value: value})
}

func (qb *updateQueryBuilder) SetSubQuery(field string, subQuery *selectQueryBuilder) error {
	sql, err := subQuery.GetSql()
	if err != nil {
		return err
	}
	qb.setters = append(qb.setters, setter{filed: field, value: "(" + sql + ")"})
	return nil
}

func (qb *updateQueryBuilder) getUpdatePart() string {
	return "UPDATE `" + qb.table + "` " + qb.alias
}

func (qb *updateQueryBuilder) getSetPart() string {
	firstSetter := qb.setters[0]
	setStr := "SET " + firstSetter.filed + " = '" + firstSetter.value + "'"

	for i := 1; i < len(qb.setters); i++ {
		currSetter := qb.setters[i]
		setStr += ", " + currSetter.filed + " = '" + currSetter.value + "'"
	}

	return setStr
}

func (qb *updateQueryBuilder) validate() error {
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

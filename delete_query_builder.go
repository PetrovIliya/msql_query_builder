package msql_query_builder

import "errors"

type DeleteQueryBuilder struct {
	joinQueryBuilder
	whereQueryBuilder
	orderByQueryBuilder
	limitQueryBuilder
	offsetQueryBuilder
	table     string
	alias     string
	deleteStr string
}

func (qb *DeleteQueryBuilder) GetSql() (string, error) {
	if err := qb.validate(); err != nil {
		return "", err
	}

	sql := escape(qb.getDeletePart() + " " +
		qb.getJoinsPart() + " " +
		qb.getWherePart() + " " +
		qb.getOrderByPart() + " " +
		qb.getLimitPart() + " " +
		qb.getOffsetPart() + " ",
	)

	return sql, nil
}

func (qb *DeleteQueryBuilder) getDeletePart() string {

	return "DELETE " + qb.deleteStr + " FROM " + qb.table + " " + qb.alias
}

func (qb DeleteQueryBuilder) validate() error {
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

	return nil
}

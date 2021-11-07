package mysql_query_builder

import (
	"errors"
)

type selectQueryBuilder struct {
	joinQueryBuilder
	whereQueryBuilder
	offsetQueryBuilder
	groupByQueryBuilder
	orderByQueryBuilder
	limitQueryBuilder
	table     string
	alias     string
	selectStr string
	orderBy   string
}

func (qb *selectQueryBuilder) GetSql() (string, error) {
	if err := qb.validate(); err != nil {
		return "", err
	}

	sql := escape(qb.getSelectPart() + " " +
		qb.getFromPart() + " " +
		qb.getWherePart() + " " +
		qb.getGroupByPart() + " " +
		qb.getOrderByPart() + " " +
		qb.getLimitPart() + " " +
		qb.getOffsetPart() + " ",
	)

	return sql, nil
}

func (qb *selectQueryBuilder) getSelectPart() string {
	if qb.selectStr == "" {
		qb.selectStr = "*"
	}
	return "SELECT " + qb.selectStr
}

func (qb *selectQueryBuilder) getFromPart() string {
	return "FROM" + " `" + qb.table + "` " + qb.alias + " " + qb.getJoinsPart()
}

func (qb *selectQueryBuilder) validate() error {
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

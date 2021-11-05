package msql_query_builder

import (
	"errors"
	"strconv"
)

type SelectQueryBuilder struct {
	baseQueryBuilder
	table     string
	alias     string
	selectStr string
	groupBy   string
	offset    string
	orderBy   string
	limit     string
}

func (qb *SelectQueryBuilder) GetSql() (string, error) {
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

func (qb *SelectQueryBuilder) GroupBy(groupBy string) {
	qb.groupBy = groupBy
}

func (qb *SelectQueryBuilder) Offset(offset int) {
	qb.offset = strconv.Itoa(offset)
}

func (qb *SelectQueryBuilder) Limit(limit int) {
	qb.limit = strconv.Itoa(limit)
}

func (qb *SelectQueryBuilder) OrderBy(orderBy string) {
	qb.orderBy = orderBy
}

func (qb *SelectQueryBuilder) getSelectPart() string {
	if qb.selectStr == "" {
		qb.selectStr = "*"
	}
	return "SELECT " + qb.selectStr
}

func (qb *SelectQueryBuilder) getFromPart() string {
	return "FROM" + " `" + qb.table + "` " + qb.alias + " " + qb.getJoinsPart()
}

func (qb *SelectQueryBuilder) getGroupByPart() string {
	if qb.groupBy == "" {
		return ""
	}
	return "GROUP BY " + qb.groupBy
}

func (qb *SelectQueryBuilder) getOffsetPart() string {
	if qb.offset == "" {
		return ""
	}
	return "OFFSET " + qb.offset
}

func (qb *SelectQueryBuilder) getLimitPart() string {
	if qb.limit == "" {
		return ""
	}
	return "LIMIT " + qb.limit
}

func (qb *SelectQueryBuilder) getOrderByPart() string {
	if qb.orderBy == "" {
		return ""
	}
	return "ORDER BY " + qb.orderBy
}

func (qb *SelectQueryBuilder) validate() error {
	err := qb.baseQueryBuilder.validate()
	if err != nil {
		return err
	}

	if qb.table == "" {
		return errors.New("'table' param is empty")
	}

	return nil
}

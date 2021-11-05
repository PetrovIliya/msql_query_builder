package msql_query_builder

import (
	"errors"
)

type UnionSelectQueryBuilder struct {
	baseQueryBuilder
	table     string
	alias     string
	selectStr string
	groupBy   string
	unions    []*UnionSelectQueryBuilder
}

func (qb *UnionSelectQueryBuilder) GetSql() (string, error) {
	if err := qb.validate(); err != nil {
		return "", err
	}

	unionPart, err := qb.getUnionPart()
	if err != nil {
		return "", err
	}

	sql := escape(qb.getSelectPart() + " " +
		qb.getFromPart() + " " +
		qb.getWherePart() + " " +
		qb.getGroupByPart() + " " +
		unionPart + " ",
	)
	return sql, nil
}

func (qb *UnionSelectQueryBuilder) GroupBy(groupBy string) {
	qb.groupBy = groupBy
}

func (qb *UnionSelectQueryBuilder) Union(union *UnionSelectQueryBuilder) {
	qb.unions = append(qb.unions, union)
}

func (qb *UnionSelectQueryBuilder) getSelectPart() string {
	if qb.selectStr == "" {
		qb.selectStr = "*"
	}
	return "SELECT " + qb.selectStr
}

func (qb *UnionSelectQueryBuilder) getUnionPart() (string, error) {
	unionPart := ""
	for i := 0; i < len(qb.unions); i++ {
		sql, err := qb.unions[i].GetSql()
		if err != nil {
			return "", err
		}
		unionPart += " UNION " + sql
	}

	return unionPart, nil
}

func (qb *UnionSelectQueryBuilder) getFromPart() string {
	return "FROM" + " `" + qb.table + "` " + qb.alias + " " + qb.getJoinsPart()
}

func (qb *UnionSelectQueryBuilder) getGroupByPart() string {
	if qb.groupBy == "" {
		return ""
	}
	return "GROUP BY " + qb.groupBy
}

func (qb *UnionSelectQueryBuilder) validate() error {
	err := qb.baseQueryBuilder.validate()
	if err != nil {
		return err
	}

	if qb.table == "" {
		return errors.New("'table' param is empty")
	}

	for i := 0; i < len(qb.unions); i++ {
		if qb.unions[i].selectStr != qb.selectStr {
			return errors.New("different union select strings")
		}
	}

	return nil
}

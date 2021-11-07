package msql_query_builder

import (
	"errors"
)

type UnionSelectQueryBuilder struct {
	joinQueryBuilder
	whereQueryBuilder
	groupByQueryBuilder
	table     string
	alias     string
	selectStr string
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

func (qb *UnionSelectQueryBuilder) validate() error {
	err := qb.whereQueryBuilder.validate()
	if err != nil {
		return err
	}

	err = qb.joinQueryBuilder.validate()
	if err != nil {
		return err
	}

	for i := 0; i < len(qb.unions); i++ {
		if qb.unions[i].selectStr != qb.selectStr {
			return errors.New("different union select strings")
		}
	}

	if qb.table == "" {
		return errors.New("'table' param can not be empty")
	}

	return nil
}

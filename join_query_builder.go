package mysql_query_builder

import "errors"

const innerJoinType = 1
const leftJoinType = 2
const outerJoinType = 3
const crossJoinType = 4

var joinPrefixMap = map[int]string{
	innerJoinType: "INNER",
	leftJoinType:  "LEFT",
	crossJoinType: "CROSS",
	outerJoinType: "OUTER",
}

type join struct {
	table     string
	alias     string
	condition string
	joinType  int
}
type joinQueryBuilder struct {
	joins []join
}

func (qb *joinQueryBuilder) InnerJoin(table string, alias string, condition string) {
	qb.joins = append(qb.joins, join{joinType: innerJoinType, table: table, alias: alias, condition: condition})
}

func (qb *joinQueryBuilder) LeftJoin(table string, alias string, condition string) {
	qb.joins = append(qb.joins, join{joinType: leftJoinType, table: table, alias: alias, condition: condition})
}

func (qb *joinQueryBuilder) CrossJoin(table string, alias string, condition string) {
	qb.joins = append(qb.joins, join{joinType: crossJoinType, table: table, alias: alias, condition: condition})
}

func (qb *joinQueryBuilder) OuterJoin(table string, alias string, condition string) {
	qb.joins = append(qb.joins, join{joinType: outerJoinType, table: table, alias: alias, condition: condition})
}

func (qb *joinQueryBuilder) getJoinsPart() string {
	joinPart := ""

	for i := 0; i < len(qb.joins); i++ {
		currJoin := qb.joins[i]
		joinPrefix := getJoinPrefix(currJoin.joinType)
		currJoinStr := joinPrefix + " JOIN `" + currJoin.table + "` " + currJoin.alias + " ON " + currJoin.condition + " "
		joinPart += currJoinStr
	}

	return joinPart
}

func (qb *joinQueryBuilder) validate() error {
	for i := 0; i < len(qb.joins); i++ {
		currJoinCondition := qb.joins[i].condition
		if currJoinCondition == "" {
			return errors.New("join condition can not be empty")
		}
	}

	return nil
}

func getJoinPrefix(joinType int) string {
	return joinPrefixMap[joinType]
}

package msql_query_builder

import (
	"errors"
)

const innerJoinType = 1
const leftJoinType = 2
const outerJoinType = 3
const crossJoinType = 4

const andWhereType = 1
const orWhereType = 2

var joinPrefixMap = map[int]string{
	innerJoinType: "INNER",
	leftJoinType:  "LEFT",
	crossJoinType: "CROSS",
	outerJoinType: "OUTER",
}

var wherePrefixMap = map[int]string{
	andWhereType: "AND",
	orWhereType:  "OR",
}

type join struct {
	table     string
	alias     string
	condition string
	joinType  int
}

type whereInCondition struct {
	filed     string
	values    []string
	whereType int
}

type whereCondition struct {
	condition string
	whereType int
}

type baseQueryBuilder struct {
	joins             []join
	whereConditions   []whereCondition
	whereInConditions []whereInCondition
}

func (qb *baseQueryBuilder) InnerJoin(table string, alias string, condition string) {
	qb.joins = append(qb.joins, join{joinType: innerJoinType, table: table, alias: alias, condition: condition})
}

func (qb *baseQueryBuilder) LeftJoin(table string, alias string, condition string) {
	qb.joins = append(qb.joins, join{joinType: leftJoinType, table: table, alias: alias, condition: condition})
}

func (qb *baseQueryBuilder) CrossJoin(table string, alias string, condition string) {
	qb.joins = append(qb.joins, join{joinType: crossJoinType, table: table, alias: alias, condition: condition})
}

func (qb *baseQueryBuilder) OuterJoin(table string, alias string, condition string) {
	qb.joins = append(qb.joins, join{joinType: outerJoinType, table: table, alias: alias, condition: condition})
}

func (qb *baseQueryBuilder) AndWhere(condition string) {
	qb.whereConditions = append(qb.whereConditions, whereCondition{whereType: andWhereType, condition: condition})
}

func (qb *baseQueryBuilder) OrWhere(condition string) {
	qb.whereConditions = append(qb.whereConditions, whereCondition{whereType: orWhereType, condition: condition})
}

func (qb *baseQueryBuilder) AndWhereIn(filed string, values []string) {
	qb.whereInConditions = append(qb.whereInConditions, whereInCondition{whereType: andWhereType, filed: filed, values: values})
}

func (qb *baseQueryBuilder) OrWhereIn(filed string, values []string) {
	qb.whereInConditions = append(qb.whereInConditions, whereInCondition{whereType: orWhereType, filed: filed, values: values})
}

func (qb *baseQueryBuilder) getJoinsPart() string {
	joinPart := ""

	for i := 0; i < len(qb.joins); i++ {
		currJoin := qb.joins[i]
		joinPrefix := getJoinPrefix(currJoin.joinType)
		currJoinStr := joinPrefix + " JOIN " + currJoin.table + " " + currJoin.alias + " ON " + currJoin.condition + " "
		joinPart += currJoinStr
	}

	return joinPart
}

func (qb *baseQueryBuilder) getWherePart() string {
	var wherePart string
	whereInConditionStartIndex := 0
	if len(qb.whereConditions) > 0 {
		wherePart = "WHERE " + qb.whereConditions[0].condition + " "
	} else if len(qb.whereInConditions) > 0 {
		whereInConditionStartIndex = 1
		firstWhereInCondition := qb.whereInConditions[0]
		wherePart = "WHERE " + getWhereInStr(firstWhereInCondition.filed, firstWhereInCondition.values) + " "
	} else {
		return ""
	}

	for i := 1; i < len(qb.whereConditions); i++ {
		currWhereCondition := qb.whereConditions[i]
		currWhereConditionStr := getWherePrefix(currWhereCondition.whereType) + " " + currWhereCondition.condition + " "
		wherePart += currWhereConditionStr
	}

	for i := whereInConditionStartIndex; i < len(qb.whereInConditions); i++ {
		currWhereInCondition := qb.whereInConditions[i]
		currWhereInConditionStr := getWherePrefix(currWhereInCondition.whereType) + " " +
			getWhereInStr(currWhereInCondition.filed, currWhereInCondition.values) + " "
		wherePart += currWhereInConditionStr
	}

	return wherePart
}

func (qb *baseQueryBuilder) validate() error {
	for i := 1; i < len(qb.whereInConditions); i++ {
		currWhereInCondition := qb.whereInConditions[i]
		if len(currWhereInCondition.values) == 0 {
			return errors.New("where in values should not be empty")
		}
	}

	return nil
}

func getJoinPrefix(joinType int) string {
	return joinPrefixMap[joinType]
}

func getWherePrefix(whereType int) string {
	return wherePrefixMap[whereType]
}

func getWhereInStr(field string, values []string) string {
	firstValue := values[0]
	valuesStr := field + " IN('" + firstValue + "'"

	for i := 1; i < len(values); i++ {
		valuesStr = valuesStr + ", '" + values[i] + "'"
	}
	valuesStr = valuesStr + ")"

	return valuesStr
}

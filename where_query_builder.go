package msql_query_builder

import "errors"

const andWhereType = 1
const orWhereType = 2

var wherePrefixMap = map[int]string{
	andWhereType: "AND",
	orWhereType:  "OR",
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

type whereQueryBuilder struct {
	whereConditions   []whereCondition
	whereInConditions []whereInCondition
}

func (qb *whereQueryBuilder) AndWhere(condition string) {
	qb.whereConditions = append(qb.whereConditions, whereCondition{whereType: andWhereType, condition: condition})
}

func (qb *whereQueryBuilder) OrWhere(condition string) {
	qb.whereConditions = append(qb.whereConditions, whereCondition{whereType: orWhereType, condition: condition})
}

func (qb *whereQueryBuilder) AndWhereIn(filed string, values []string) {
	qb.whereInConditions = append(qb.whereInConditions, whereInCondition{whereType: andWhereType, filed: filed, values: values})
}

func (qb *whereQueryBuilder) OrWhereIn(filed string, values []string) {
	qb.whereInConditions = append(qb.whereInConditions, whereInCondition{whereType: orWhereType, filed: filed, values: values})
}

func (qb *whereQueryBuilder) getWherePart() string {
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

func (qb *whereQueryBuilder) validate() error {
	for i := 1; i < len(qb.whereInConditions); i++ {
		currWhereInCondition := qb.whereInConditions[i]
		if len(currWhereInCondition.values) == 0 {
			return errors.New("where in values should not be empty")
		}
	}

	return nil
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

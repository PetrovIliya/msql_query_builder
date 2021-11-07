# MySQL Query Builder for Golang

Query Builder makes it easier to build queries for MySQL databases.

So, if your application relies heavily on MySQL database and using the repository pattern is too expensive for performance, then you should use query services in which you need to build optimized queries. This query builder will help you build large and readable queries

# Installation

1. Add the dependency `github.com/PetrovIliya/mysql_query_builder v1.0.1beta` to your go.mod file
2. In the required file add import `qb "github.com/PetrovIliya/mysql_query_builder"`

# Usage

## Select

    selectQb := qb.Select("order", "o", "mi.*")
	selectQb.InnerJoin("menu_item", "mi", "mi.order_id = o.order_id")
	
	orderIds := make([]string, 2)
	orderIds[0] = "1"
	orderIds[1] = "2"
	
	selectQb.AndWhereIn("o.order_id", orderIds)
    selectQb.Limit(1)
    selectQb.Offset(2)
    selectQb.OrderBy("o.order_id DESC")
    sqlStr, err := selectQb.GetSql()

## Union select

    unionSubQueryQb := qb.UnionSelect("order", "o2", "o2.cost")
    unionSubQueryQb.AndWhere("o2.order_id = 2")

	selectQb := qb.UnionSelect("order", "o", "o.cost")
	selectQb.AndWhere("o.order_id = 1")
	selectQb.Union(unionSubQueryQb)

	sqlStr, err := selectQb.GetSql()

## Update

	updateQb := qb.Update("order", "o")
	updateQb.Set("o.cost", "970")
	updateQb.AndWhere("order_id = 1")
	sqlStr, err := updateQb.GetSql()

## Update by sub query

    updateQb := qb.Update("order", "o")
	updateQb.Set("o.cost", "(SELECT o2.cost FROM `order` o2 WHERE order_id = 2)")
	updateQb.AndWhere("order_id = 1")
	sqlStr, err := updateQb.GetSql()

## Insert by query

    selectQb := qb.Select("order", "o", "order_id, cost")
    insertQb := qb.InsertInto("order", "(order_id, cost)")
    insertQb.Ignore(true)
    insertQb.ValuesQuery(selectQb)
    sqlStr, err := insertQb.GetSql()

## Insert by union query
    unionSbQb := qb.UnionSelect("order", "o2", "order_id, cost")
	selectQb := qb.UnionSelect("order", "o", "order_id, cost")
	selectQb.Union(unionSbQb)

	insertQb := qb.InsertInto("order", "(order_id, cost)")
	insertQb.Ignore(true)
	insertQb.ValuesUnionQuery(selectQb)

	sqlStr, err := insertQb.GetSql()

## Insert values

	insertQb := qb.InsertInto("order", "(order_id, cost)")
	insertQb.Ignore(true)
	a := make([]string, 2)
	a[0] = "1"
	a[1] = "3"
	b := make([][]string, 1)
	b[0] = a 
	insertQb.Values(b)
	
	sql, err := insertQb.GetSql()

Result: INSERT IGNORE INTO \`order\` (order_id, cost) VALUES (\'1\', \'3\')

## Delete

    deleteSq := qb.Delete("menu_item", "mi", "mi.*")
	deleteSq.InnerJoin("order", "o", "o.order_id = mi.order_id")
	deleteSq.AndWhere("o.order_id = 1")
    sqlStr, err := deleteSq.GetSql()

# PS


I like the Doctrine query builder style and PHP arrays, so I transferred my Doctrine experience to this query builder.

The usability of using go slices as parameters to some query builder methods is questionable. If this is the case let me know, and I will try to fix it for the v1.0.0 release.

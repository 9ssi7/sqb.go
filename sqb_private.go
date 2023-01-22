package sqb_go

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *sqb) optimizeSelect(sel string) {
	if s.sel == "*" {
		s.sel = sel
	} else {
		s.sel += ", " + sel
	}
}

func (s *sqb) optSqlFuncWithAlias(sqlFunc string, column string, alias string) {
	s.optimizeSelect(sqlFunc + "(" + column + ") AS " + alias)
}

func (s *sqb) optJoin(table string, field1 string, field2 string, joinType string, operator ...string) {
	on := field1
	tb := s.prefix + table
	if len(operator) > 0 {
		on += " " + operator[0] + " " + field2
	} else {
		on += " = " + field2
	}
	if joinType != "" {
		joinType = joinType + " JOIN"
	}
	if s.join == "" {
		s.join = joinType + " " + tb + " ON " + on
	} else {
		s.join += " " + joinType + " " + tb + " ON " + on
	}
}

func (s *sqb) optWhere(where string, operator string, value interface{}, whType string) {
	wh := where + " " + operator + " " + s.escape(fmt.Sprintf("%v", value))
	s.setWhereGrouped(wh, whType)
}

func (s *sqb) setWhere(where string, whType string) {
	opt := " " + whType + " "
	if s.where == "" {
		s.where = where
	} else {
		s.where += opt + where
	}
}

func (s *sqb) setWhereGrouped(where string, whType string) {
	wh := where
	if s.grouped {
		wh = "(" + where
		s.grouped = false
	}
	s.setWhere(wh, whType)
}

func (s *sqb) escape(str string) string {
	if _, err := strconv.Atoi(str); err == nil {
		return str
	}
	return "'" + str + "'"
}

func (s *sqb) escapeInterfaceArr(arr []interface{}) string {
	var strArr []string
	for _, v := range arr {
		strArr = append(strArr, s.escapeInterface(v))
	}
	return strings.Join(strArr, ", ")
}

func (s *sqb) escapeInterface(str interface{}) string {
	return s.escape(fmt.Sprintf("%v", str))
}

func (s *sqb) optWhereNull(where string, whType string) {
	wh := where + " " + whType + " NULL"
	s.setWhere(wh, whType)
}

func (s *sqb) optWhereIn(where string, values []interface{}, inType string, whType string) {
	wh := ""
	if inType != "" {
		wh = where + " " + inType + " IN (" + s.escapeInterfaceArr(values) + ")"
	} else {
		wh = where + " IN (" + s.escapeInterfaceArr(values) + ")"
	}
	s.setWhereGrouped(wh, whType)
}

func (s *sqb) optWhereFindInSet(column string, values []interface{}, findType string, whType string) {
	wh := "FIND_IN_SET(" + column + ", " + s.escapeInterfaceArr(values) + ")"
	if findType != "" {
		wh = findType + " " + wh
	}
	s.setWhereGrouped(wh, whType)
}

func (s *sqb) optWhereBetween(column string, value1 interface{}, value2 interface{}, betweenType string, whType string) {
	wh := ""
	if betweenType != "" {
		wh = "(" + column + " " + betweenType + " BETWEEN " + s.escape(fmt.Sprintf("%v", value1)) + " AND " + s.escape(fmt.Sprintf("%v", value2))
	} else {
		wh = "(" + column + " BETWEEN " + s.escape(fmt.Sprintf("%v", value1)) + " AND " + s.escape(fmt.Sprintf("%v", value2))
	}
	wh += ")"
	s.setWhereGrouped(wh, whType)
}

func (s *sqb) optLike(column string, value interface{}, likeType string, andOr string) {
	like := s.escape(fmt.Sprintf("%v", value))
	wh := ""
	if likeType == "" {
		wh = column + " LIKE " + like
	} else {
		wh = column + " " + likeType + " LIKE " + like
	}
	s.setWhereGrouped(wh, andOr)
}

func (s *sqb) optTableAction(action string) string {
	return action + " TABLE " + s.from
}

func (s *sqb) build(action string) string {
	sql := action + " " + s.sel + " FROM " + s.prefix + s.from
	if s.join != "" {
		sql += " " + s.join
	}
	if s.where != "" {
		sql += " WHERE " + s.where
	}
	if s.groupBy != "" {
		sql += " GROUP BY " + s.groupBy
	}
	if s.having != "" {
		sql += " HAVING " + s.having
	}
	if s.orderBy != "" {
		sql += " ORDER BY " + s.orderBy
	}
	if s.limit != "" {
		sql += " LIMIT " + s.limit
	}
	if s.offset != "" {
		sql += " OFFSET " + s.offset
	}
	return sql
}

func (s *sqb) buildInsert(m *M) string {
	sql := "INSERT INTO " + s.prefix + s.from + " ("
	var cols []string
	var values []interface{}
	for k, v := range *m {
		cols = append(cols, k)
		values = append(values, v)
	}
	sql += strings.Join(cols, ", ")
	sql += ") VALUES ("
	sql += s.escapeInterfaceArr(values)
	sql += ")"
	return sql
}

func (s *sqb) buildInsertMany(m []*M) string {
	sql := "INSERT INTO " + s.prefix + s.from + " ("
	var cols []string
	var values []string
	for _, v := range m {
		var cols2 []string
		var values2 []interface{}
		for k, v2 := range *v {
			cols2 = append(cols2, k)
			values2 = append(values2, v2)
		}
		cols = cols2
		values = append(values, "("+s.escapeInterfaceArr(values2)+")")
	}
	sql += strings.Join(cols, ", ")
	sql += ") VALUES "
	sql += strings.Join(values, ", ")
	return sql
}

func (s *sqb) buildUpdate(m *M) string {
	sql := "UPDATE " + s.prefix + s.from + " SET "
	var cols []string
	for k, v := range *m {
		cols = append(cols, k+" = "+s.escapeInterface(v))
	}
	sql += strings.Join(cols, ", ")
	if s.where != "" {
		sql += " WHERE " + s.where
	}
	return sql
}

func (s *sqb) buildUpdateMany(m []*M) string {
	sql := "UPDATE " + s.prefix + s.from + " SET "
	var cols []string
	for _, v := range m {
		var cols2 []string
		for k2, v2 := range *v {
			cols2 = append(cols2, k2+" = "+s.escapeInterface(v2))
		}
		cols = append(cols, strings.Join(cols2, ", "))
	}
	sql += strings.Join(cols, ", ")
	if s.where != "" {
		sql += " WHERE " + s.where
	}
	return sql
}

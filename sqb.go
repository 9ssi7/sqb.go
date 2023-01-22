package sqb_go

import (
	"fmt"
	"strings"
)

type sqb struct {
	sel       string
	from      string
	where     string
	limit     string
	offset    string
	join      string
	orderBy   string
	groupBy   string
	having    string
	grouped   bool
	query     string
	prefix    string
	operators []string
}

func New() Sqb {
	return &sqb{
		sel:       "*",
		prefix:    "",
		operators: []string{"=", "!=", "<", ">", "<=", ">=", "<>"},
	}
}

func (s *sqb) Table(table ...string) Sqb {
	s.from = s.prefix + strings.Join(table, ", ")
	return s
}

func (s *sqb) Select(columns ...string) Sqb {
	s.optimizeSelect(strings.Join(columns, ", "))
	return s
}

func (s *sqb) GroupConcat(column string, alias string) Sqb {
	s.optSqlFuncWithAlias("GROUP_CONCAT", column, alias)
	return s
}

func (s *sqb) Least(column string, alias string) Sqb {
	s.optSqlFuncWithAlias("LEAST", column, alias)
	return s
}

func (s *sqb) Max(column string, alias string) Sqb {
	s.optSqlFuncWithAlias("MAX", column, alias)
	return s
}

func (s *sqb) Min(column string, alias string) Sqb {
	s.optSqlFuncWithAlias("MIN", column, alias)
	return s
}

func (s *sqb) Sum(column string, alias string) Sqb {
	s.optSqlFuncWithAlias("SUM", column, alias)
	return s
}

func (s *sqb) Count(column string, alias string) Sqb {
	s.optSqlFuncWithAlias("COUNT", column, alias)
	return s
}

func (s *sqb) Avg(column string, alias string) Sqb {
	s.optSqlFuncWithAlias("AVG", column, alias)
	return s
}

func (s *sqb) Join(table string, field1 string, field2 string, joinType string, operator ...string) Sqb {
	s.optJoin(table, field1, field2, joinType, operator...)
	return s
}

func (s *sqb) InnerJoin(table string, field1 string, field2 string, operator ...string) Sqb {
	s.optJoin(table, field1, field2, "INNER", operator...)
	return s
}

func (s *sqb) LeftJoin(table string, field1 string, field2 string, operator ...string) Sqb {
	s.optJoin(table, field1, field2, "LEFT", operator...)
	return s
}

func (s *sqb) RightJoin(table string, field1 string, field2 string, operator ...string) Sqb {
	s.optJoin(table, field1, field2, "RIGHT", operator...)
	return s
}

func (s *sqb) FullOuterJoin(table string, field1 string, field2 string, operator ...string) Sqb {
	s.optJoin(table, field1, field2, "FULL OUTER", operator...)
	return s
}

func (s *sqb) LeftOuterJoin(table string, field1 string, field2 string, operator ...string) Sqb {
	s.optJoin(table, field1, field2, "LEFT OUTER", operator...)
	return s
}

func (s *sqb) RightOuterJoin(table string, field1 string, field2 string, operator ...string) Sqb {
	s.optJoin(table, field1, field2, "RIGHT OUTER", operator...)
	return s
}

func (s *sqb) Where(column string, operator string, value interface{}) Sqb {
	s.optWhere(column, operator, value, "AND")
	return s
}

func (s *sqb) OrWhere(column string, operator string, value interface{}) Sqb {
	s.optWhere(column, operator, value, "OR")
	return s
}

func (s *sqb) NotWhere(column string, operator string, value interface{}) Sqb {
	s.optWhere(column, operator, value, "")
	return s
}

func (s *sqb) OrNotWhere(column string, operator string, value interface{}) Sqb {
	s.optWhere(column, operator, value, "OR NOT")
	return s
}

func (s *sqb) WhereNull(column string) Sqb {
	s.optWhereNull(column, "IS")
	return s
}

func (s *sqb) WhereNotNull(column string) Sqb {
	s.optWhereNull(column, "IS NOT")
	return s
}

func (s *sqb) Grouped(f func(Sqb) Sqb) Sqb {
	s.grouped = true
	f(s)
	s.where += ")"
	return s
}

func (s *sqb) In(column string, values ...interface{}) Sqb {
	if len(values) == 0 {
		return s
	}
	s.optWhereIn(column, values, "", "")
	return s
}

func (s *sqb) NotIn(column string, values ...interface{}) Sqb {
	if len(values) == 0 {
		return s
	}
	s.optWhereIn(column, values, "NOT", "NOT")
	return s
}

func (s *sqb) OrIn(column string, values ...interface{}) Sqb {
	if len(values) == 0 {
		return s
	}
	s.optWhereIn(column, values, "", "OR")
	return s
}

func (s *sqb) OrNotIn(column string, values ...interface{}) Sqb {
	if len(values) == 0 {
		return s
	}
	s.optWhereIn(column, values, "NOT", "OR")
	return s
}

func (s *sqb) FindInSet(column string, values ...interface{}) Sqb {
	if len(values) == 0 {
		return s
	}
	s.optWhereFindInSet(column, values, "", "")
	return s
}

func (s *sqb) NotFindInSet(column string, values ...interface{}) Sqb {
	if len(values) == 0 {
		return s
	}
	s.optWhereFindInSet(column, values, "NOT", "AND")
	return s
}

func (s *sqb) OrFindInSet(column string, values ...interface{}) Sqb {
	if len(values) == 0 {
		return s
	}
	s.optWhereFindInSet(column, values, "", "OR")
	return s
}

func (s *sqb) OrNotFindInSet(column string, values ...interface{}) Sqb {
	if len(values) == 0 {
		return s
	}
	s.optWhereFindInSet(column, values, "NOT", "OR")
	return s
}

func (s *sqb) Between(column string, value1 interface{}, value2 interface{}) Sqb {
	s.optWhereBetween(column, value1, value2, "", "")
	return s
}

func (s *sqb) NotBetween(column string, value1 interface{}, value2 interface{}) Sqb {
	s.optWhereBetween(column, value1, value2, "NOT", "AND")
	return s
}

func (s *sqb) OrBetween(column string, value1 interface{}, value2 interface{}) Sqb {
	s.optWhereBetween(column, value1, value2, "", "OR")
	return s
}

func (s *sqb) OrNotBetween(column string, value1 interface{}, value2 interface{}) Sqb {
	s.optWhereBetween(column, value1, value2, "NOT", "OR")
	return s
}

func (s *sqb) Like(column string, value interface{}) Sqb {
	s.optLike(column, value, "", "AND")
	return s
}

func (s *sqb) NotLike(column string, value interface{}) Sqb {
	s.optLike(column, value, "NOT", "AND")
	return s
}

func (s *sqb) OrLike(column string, value interface{}) Sqb {
	s.optLike(column, value, "", "OR")
	return s
}

func (s *sqb) OrNotLike(column string, value interface{}) Sqb {
	s.optLike(column, value, "NOT", "OR")
	return s
}

func (s *sqb) Limit(limit int, end ...int) Sqb {
	if len(end) > 0 {
		s.limit = fmt.Sprintf("%d, %d", limit, end[0])
	} else {
		s.limit = fmt.Sprintf("%d", limit)
	}
	return s
}

func (s *sqb) Offset(offset int) Sqb {
	s.offset = fmt.Sprintf("%d", offset)
	return s
}

func (s *sqb) Pagination(page int, perPage int) Sqb {
	s.limit = fmt.Sprintf("%d, %d", (page-1)*perPage, perPage)
	return s
}

func (s *sqb) OrderBy(column string, direction Order) Sqb {
	if column != "" {
		s.orderBy = fmt.Sprintf("%s %s", column, direction)
	}
	return s
}

func (s *sqb) GroupBy(column ...string) Sqb {
	if len(column) > 0 {
		s.groupBy = strings.Join(column, ", ")
	}
	return s
}

func (s *sqb) Having(column string, operator string, value interface{}) Sqb {
	s.having = fmt.Sprintf("%s %s %s", column, operator, s.escape(fmt.Sprintf("%v", value)))
	return s
}

func (s *sqb) FromQuery(query string, bindings ...interface{}) Sqb {
	if len(bindings) == 0 {
		s.query = query
		return s
	}
	q := strings.Split(query, "?")
	if len(q) != len(bindings)+1 {
		return s
	}
	for i, v := range bindings {
		q[i] += s.escape(fmt.Sprintf("%v", v))
	}
	s.query = strings.Join(q, "")
	return s
}

func (s *sqb) Reset() Sqb {
	s.sel = "*"
	s.groupBy = ""
	s.having = ""
	s.orderBy = ""
	s.limit = ""
	s.offset = ""
	s.where = ""
	s.prefix = ""
	s.from = ""
	return s
}

func (s *sqb) Get() string {
	s.limit = "1"
	return s.GetAll()
}

func (s *sqb) GetAll() string {
	s.query = s.build("SELECT")
	s.Reset()
	return s.query
}

func (s *sqb) Insert(m *M) string {
	s.query = s.buildInsert(m)
	s.Reset()
	return s.query
}

func (s *sqb) InsertMany(m []*M) string {
	s.query = s.buildInsertMany(m)
	s.Reset()
	return s.query
}

func (s *sqb) Update(m *M) string {
	s.query = s.buildUpdate(m)
	s.Reset()
	return s.query
}

func (s *sqb) UpdateMany(m []*M) string {
	s.query = s.buildUpdateMany(m)
	s.Reset()
	return s.query
}

func (s *sqb) Delete() string {
	s.query = s.build("DELETE")
	s.Reset()
	return s.query
}

func (s *sqb) Analyze() string {
	s.query = s.optTableAction("ANALYZE")
	s.Reset()
	return s.query
}

func (s *sqb) Check() string {
	s.query = s.optTableAction("CHECK")
	s.Reset()
	return s.query
}

func (s *sqb) Checksum() string {
	s.query = s.optTableAction("CHECKSUM")
	s.Reset()
	return s.query
}

func (s *sqb) Optimize() string {
	s.query = s.optTableAction("OPTIMIZE")
	s.Reset()
	return s.query
}

func (s *sqb) Repair() string {
	s.query = s.optTableAction("REPAIR")
	s.Reset()
	return s.query
}

func (s *sqb) GetQuery() string {
	return s.query
}

func (s *sqb) Build() string {
	return s.build("SELECT")
}

var (
	QB = New()
)

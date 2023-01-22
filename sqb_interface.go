package sqb_go

type Sqb interface {
	Table(table ...string) Sqb
	Select(columns ...string) Sqb
	GroupConcat(column string, alias string) Sqb
	Least(column string, alias string) Sqb
	Max(column string, alias string) Sqb
	Min(column string, alias string) Sqb
	Sum(column string, alias string) Sqb
	Count(column string, alias string) Sqb
	Avg(column string, alias string) Sqb
	Join(table string, field1 string, field2 string, joinType string, operator ...string) Sqb
	InnerJoin(table string, field1 string, field2 string, operator ...string) Sqb
	LeftJoin(table string, field1 string, field2 string, operator ...string) Sqb
	RightJoin(table string, field1 string, field2 string, operator ...string) Sqb
	FullOuterJoin(table string, field1 string, field2 string, operator ...string) Sqb
	LeftOuterJoin(table string, field1 string, field2 string, operator ...string) Sqb
	RightOuterJoin(table string, field1 string, field2 string, operator ...string) Sqb
	Where(column string, operator string, value interface{}) Sqb
	OrWhere(column string, operator string, value interface{}) Sqb
	NotWhere(column string, operator string, value interface{}) Sqb
	OrNotWhere(column string, operator string, value interface{}) Sqb
	WhereNull(column string) Sqb
	WhereNotNull(column string) Sqb
	Grouped(func(Sqb) Sqb) Sqb
	In(column string, values ...interface{}) Sqb
	NotIn(column string, values ...interface{}) Sqb
	OrIn(column string, values ...interface{}) Sqb
	OrNotIn(column string, values ...interface{}) Sqb
	FindInSet(column string, values ...interface{}) Sqb
	NotFindInSet(column string, values ...interface{}) Sqb
	OrFindInSet(column string, values ...interface{}) Sqb
	OrNotFindInSet(column string, values ...interface{}) Sqb
	Between(column string, value1 interface{}, value2 interface{}) Sqb
	NotBetween(column string, value1 interface{}, value2 interface{}) Sqb
	OrBetween(column string, value1 interface{}, value2 interface{}) Sqb
	OrNotBetween(column string, value1 interface{}, value2 interface{}) Sqb
	Like(column string, value interface{}) Sqb
	NotLike(column string, value interface{}) Sqb
	OrLike(column string, value interface{}) Sqb
	OrNotLike(column string, value interface{}) Sqb
	Limit(limit int, end ...int) Sqb
	Offset(offset int) Sqb
	Pagination(page int, perPage int) Sqb
	OrderBy(column string, direction Order) Sqb
	GroupBy(column ...string) Sqb
	Having(column string, operator string, value interface{}) Sqb
	FromQuery(query string, bindings ...interface{}) Sqb
	Reset() Sqb
	Get() string
	GetAll() string
	Insert(m *M) string
	InsertMany(m []*M) string
	Update(m *M) string
	UpdateMany(m []*M) string
	Delete() string
	Analyze() string
	Check() string
	Checksum() string
	Optimize() string
	Repair() string
	GetQuery() string
	Build() string
}

package sqb_go

import (
	"testing"
)

func Test_sqb_Analyze(t *testing.T) {
	tests := []struct {
		name  string
		table string
		want  string
	}{
		{
			name:  "Test Analyse",
			table: "Test",
			want:  "ANALYZE TABLE Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Analyze(); got != tt.want {
				t.Errorf("Analyse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Avg(t *testing.T) {
	type args struct {
		column string
		alias  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Avg",
			args: args{
				column: "Test",
				alias:  "Test2",
			},
			want: "SELECT AVG(Test) AS Test2 FROM ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Avg(tt.args.column, tt.args.alias).Build(); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Between(t *testing.T) {
	type args struct {
		column string
		value1 interface{}
		value2 interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Between",
			args: args{
				column: "Test",
				value1: 1,
				value2: 2,
			},
			want: "SELECT * FROM Test WHERE (Test BETWEEN 1 AND 2)",
		},
		{
			name: "Test Between",
			args: args{
				column: "DatetimeTest",
				value1: "2020-01-01",
				value2: "2020-01-02",
			},
			want: "SELECT * FROM Test WHERE (DatetimeTest BETWEEN '2020-01-01' AND '2020-01-02')",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").Between(tt.args.column, tt.args.value1, tt.args.value2).Build(); got != tt.want {
				t.Errorf("Between() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Check(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test Check",
			want: "CHECK TABLE Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").Check(); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Checksum(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test Checksum",
			want: "CHECKSUM TABLE Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").Checksum(); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Count(t *testing.T) {
	type args struct {
		column string
		alias  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Count",
			args: args{
				column: "Test",
				alias:  "Test2",
			},
			want: "SELECT COUNT(Test) AS Test2 FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").Count(tt.args.column, tt.args.alias).Build(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Delete(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test Delete",
			want: "DELETE * FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").Delete(); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_FindInSet(t *testing.T) {
	type args struct {
		column string
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test FindInSet",
			args: args{
				column: "Test",
				values: []interface{}{1, 2},
			},
			want: "SELECT * FROM Test WHERE FIND_IN_SET(Test, 1, 2)",
		},
		{
			name: "Test FindInSet",
			args: args{
				column: "Test",
				values: []interface{}{"1", "2", "3"},
			},
			want: "SELECT * FROM Test WHERE FIND_IN_SET(Test, 1, 2, 3)",
		},
		{
			name: "Test FindInSet with empty values",
			args: args{
				column: "Test",
				values: []interface{}{},
			},
			want: "SELECT * FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").FindInSet(tt.args.column, tt.args.values...).Build(); got != tt.want {
				t.Errorf("FindInSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_FromQuery(t *testing.T) {
	type args struct {
		query    string
		bindings []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test FromQuery",
			args: args{
				query:    "SELECT * FROM Test",
				bindings: []interface{}{},
			},
			want: "SELECT * FROM Test",
		},
		{
			name: "Test FromQuery",
			args: args{
				query:    "SELECT * FROM Test WHERE Test = ?",
				bindings: []interface{}{"Test"},
			},
			want: "SELECT * FROM Test WHERE Test = 'Test'",
		},
		{
			name: "Test FromQuery with not match params count",
			args: args{
				query:    "SELECT * FROM Test WHERE Test = ?",
				bindings: []interface{}{"Test", "Test2"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").FromQuery(tt.args.query, tt.args.bindings...).GetQuery(); got != tt.want {
				t.Errorf("FromQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_FullOuterJoin(t *testing.T) {
	type args struct {
		table  string
		field1 string
		field2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test FullOuterJoin",
			args: args{
				table:  "Test",
				field1: "Test.TestId",
				field2: "Test.TestIdd",
			},
			want: "SELECT * FROM Test FULL OUTER JOIN Test ON Test.TestId = Test.TestIdd",
		},
		{
			name: "Test FullOuterJoin another table",
			args: args{
				table:  "Test2",
				field1: "Test.Test2Id",
				field2: "Test2.TestId",
			},
			want: "SELECT * FROM Test FULL OUTER JOIN Test2 ON Test.Test2Id = Test2.TestId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").FullOuterJoin(tt.args.table, tt.args.field1, tt.args.field2).Build(); got != tt.want {
				t.Errorf("FullOuterJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Get(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test Get",
			want: "SELECT * FROM Test LIMIT 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").Get(); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_GetAll(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test GetAll",
			want: "SELECT * FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table("Test").GetAll(); got != tt.want {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_GetQuery(t *testing.T) {

	tests := []struct {
		name  string
		want  string
		table string
	}{
		{
			name:  "Test GetQuery",
			table: "Test",
			want:  "SELECT * FROM Test LIMIT 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.Table(tt.table).Get()
			if got := s.GetQuery(); got != tt.want {
				t.Errorf("GetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_GroupBy(t *testing.T) {

	type args struct {
		column []string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test GroupBy",
			table: "Test",
			args: args{
				column: []string{"Test"},
			},
			want: "SELECT * FROM Test GROUP BY Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).GroupBy(tt.args.column...).Build(); got != tt.want {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_GroupConcat(t *testing.T) {
	type args struct {
		column string
		alias  string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test GroupConcat",
			table: "Test",
			args: args{
				column: "Test",
				alias:  "Test",
			},
			want: "SELECT GROUP_CONCAT(Test) AS Test FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).GroupConcat(tt.args.column, tt.args.alias).Build(); got != tt.want {
				t.Errorf("GroupConcat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Grouped(t *testing.T) {
	type args struct {
		f func(Sqb) Sqb
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Grouped",
			table: "Test",
			args: args{
				f: func(s Sqb) Sqb {
					return s.Where("Can", "=", "Test").Where("ECan", "=", "Test")
				},
			},
			want: "SELECT * FROM Test WHERE (Can = 'Test' AND ECan = 'Test')",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Grouped(tt.args.f).Build(); got != tt.want {
				t.Errorf("Grouped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Having(t *testing.T) {
	type args struct {
		column   string
		operator string
		value    interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Having",
			table: "Test",
			args: args{
				column:   "Test",
				operator: "=",
				value:    "Test",
			},
			want: "SELECT * FROM Test HAVING Test = 'Test'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Having(tt.args.column, tt.args.operator, tt.args.value).Build(); got != tt.want {
				t.Errorf("Having() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_In(t *testing.T) {
	type args struct {
		column string
		values []interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test In",
			table: "Test",
			args: args{
				column: "Test",
				values: []interface{}{"Test"},
			},
			want: "SELECT * FROM Test WHERE Test IN ('Test')",
		},
		{
			name:  "Test In with multiple values",
			table: "Test",
			args: args{
				column: "Test",
				values: []interface{}{"Test", "Test2"},
			},
			want: "SELECT * FROM Test WHERE Test IN ('Test', 'Test2')",
		},
		{
			name:  "Test In with empty values",
			table: "Test",
			args: args{
				column: "Test",
				values: []interface{}{},
			},
			want: "SELECT * FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).In(tt.args.column, tt.args.values...).Build(); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_InnerJoin(t *testing.T) {
	type args struct {
		table  string
		field1 string
		field2 string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test InnerJoin",
			table: "Test",
			args: args{
				table:  "Test",
				field1: "Test.TestId",
				field2: "Test.TestAnotherId",
			},
			want: "SELECT * FROM Test INNER JOIN Test ON Test.TestId = Test.TestAnotherId",
		},
		{
			name:  "Test InnerJoin with another table",
			table: "Test",
			args: args{
				table:  "Test2",
				field1: "Test.TestId",
				field2: "Test2.TestAnotherId",
			},
			want: "SELECT * FROM Test INNER JOIN Test2 ON Test.TestId = Test2.TestAnotherId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).InnerJoin(tt.args.table, tt.args.field1, tt.args.field2).Build(); got != tt.want {
				t.Errorf("InnerJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Insert(t *testing.T) {
	type values *M
	tests := []struct {
		name   string
		values values
		table  string
		want   string
	}{
		{
			name:  "Test Insert",
			table: "Test",
			want:  "INSERT INTO Test (Test) VALUES ('Test')",
			values: &M{
				"Test": "Test",
			},
		},
		{
			name:  "Test Insert with multiple values",
			table: "Test",
			want:  "INSERT INTO Test (Test, Test2) VALUES ('Test', 'Test2')",
			values: &M{
				"Test":  "Test",
				"Test2": "Test2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Insert(tt.values); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_InsertMany(t *testing.T) {
	type values []*M
	tests := []struct {
		name   string
		values values
		table  string
		want   string
	}{
		{
			name:  "Test Insert",
			table: "Test",
			values: []*M{
				{
					"Test": "Test",
				},
				{
					"Test": "Test2",
				},
			},
			want: "INSERT INTO Test (Test) VALUES ('Test'), ('Test2')",
		},
		{
			name:  "Test Insert with multiple values",
			table: "Test",
			values: []*M{
				{
					"Test": "Test",
				},
			},
			want: "INSERT INTO Test (Test) VALUES ('Test')",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).InsertMany(tt.values); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Join(t *testing.T) {
	type args struct {
		table    string
		field1   string
		field2   string
		joinType string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Join",
			table: "Test",
			args: args{
				table:    "Test",
				field1:   "Test.TestId",
				field2:   "Test.TestAnotherId",
				joinType: "INNER",
			},
			want: "SELECT * FROM Test INNER JOIN Test ON Test.TestId = Test.TestAnotherId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Join(tt.args.table, tt.args.field1, tt.args.field2, tt.args.joinType).Build(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("run multiple joins", func(t *testing.T) {
		s := New()
		if got := s.Table("Test").Join("Test", "Test.TestId", "Test.TestAnotherId", "INNER").Join("Test", "Test.TestId", "Test.TestAnotherId", "INNER").Build(); got != "SELECT * FROM Test INNER JOIN Test ON Test.TestId = Test.TestAnotherId INNER JOIN Test ON Test.TestId = Test.TestAnotherId" {
			t.Errorf("Join() = %v, want %v", got, "SELECT * FROM Test INNER JOIN Test ON Test.TestId = Test.TestAnotherId INNER JOIN Test ON Test.TestId = Test.TestAnotherId")
		}
	})
	t.Run("run with another operator", func(t *testing.T) {
		s := New()
		if got := s.Table("Test").Join("Test", "Test.TestId", "Test.TestAnotherId", "INNER", ">").Build(); got != "SELECT * FROM Test INNER JOIN Test ON Test.TestId > Test.TestAnotherId" {
			t.Errorf("Join() = %v, want %v", got, "SELECT * FROM Test INNER JOIN Test ON Test.TestId > Test.TestAnotherId")
		}
	})
}

func Test_sqb_Least(t *testing.T) {
	type args struct {
		column string
		alias  string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Least",
			table: "Test",
			args: args{
				column: "Test",
				alias:  "Test",
			},
			want: "SELECT LEAST(Test) AS Test FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Least(tt.args.column, tt.args.alias).Build(); got != tt.want {
				t.Errorf("Least() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_LeftJoin(t *testing.T) {

	type args struct {
		table  string
		field1 string
		field2 string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Left Join",
			table: "Test",
			args: args{
				table:  "Test",
				field1: "Test.TestId",
				field2: "Test.TestAnotherId",
			},
			want: "SELECT * FROM Test LEFT JOIN Test ON Test.TestId = Test.TestAnotherId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).LeftJoin(tt.args.table, tt.args.field1, tt.args.field2).Build(); got != tt.want {
				t.Errorf("LeftJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_LeftOuterJoin(t *testing.T) {
	type args struct {
		table  string
		field1 string
		field2 string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Left Outer Join",
			table: "Test",
			args: args{
				table:  "Test",
				field1: "Test.TestId",
				field2: "Test.TestAnotherId",
			},
			want: "SELECT * FROM Test LEFT OUTER JOIN Test ON Test.TestId = Test.TestAnotherId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).LeftOuterJoin(tt.args.table, tt.args.field1, tt.args.field2).Build(); got != tt.want {
				t.Errorf("LeftOuterJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Like(t *testing.T) {
	type args struct {
		column string
		value  interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Like",
			table: "Test",
			args: args{
				column: "Test",
				value:  "Test",
			},
			want: "SELECT * FROM Test WHERE Test LIKE 'Test'",
		},
		{
			name:  "Test Like with %",
			table: "Test",
			args: args{
				column: "Test",
				value:  "%Test%",
			},
			want: "SELECT * FROM Test WHERE Test LIKE '%Test%'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Like(tt.args.column, tt.args.value).Build(); got != tt.want {
				t.Errorf("Like() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Limit(t *testing.T) {

	type args struct {
		limit int
		end   []int
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Limit",
			table: "Test",
			args: args{
				limit: 10,
			},
			want: "SELECT * FROM Test LIMIT 10",
		},
		{
			name:  "Test Limit with end",
			table: "Test",
			args: args{
				limit: 10,
				end:   []int{10},
			},
			want: "SELECT * FROM Test LIMIT 10, 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Limit(tt.args.limit, tt.args.end...).Build(); got != tt.want {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Max(t *testing.T) {
	type args struct {
		column string
		alias  string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Max",
			table: "Test",
			args: args{
				column: "Test",
				alias:  "Test",
			},
			want: "SELECT MAX(Test) AS Test FROM Test",
		},
		{
			name:  "Test Max with alias",
			table: "Test",
			args: args{
				column: "Test",
				alias:  "Test2",
			},
			want: "SELECT MAX(Test) AS Test2 FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Max(tt.args.column, tt.args.alias).Build(); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Min(t *testing.T) {
	type args struct {
		column string
		alias  string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Min",
			table: "Test",
			args: args{
				column: "Test",
				alias:  "Test",
			},
			want: "SELECT MIN(Test) AS Test FROM Test",
		},
		{
			name:  "Test Min with alias",
			table: "Test",
			args: args{
				column: "Test",
				alias:  "Test2",
			},
			want: "SELECT MIN(Test) AS Test2 FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Min(tt.args.column, tt.args.alias).Build(); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_NotBetween(t *testing.T) {

	type args struct {
		column string
		value1 interface{}
		value2 interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test NotBetween",
			table: "Test",
			args: args{
				column: "Test",
				value1: 1,
				value2: 2,
			},
			want: "SELECT * FROM Test WHERE (Test NOT BETWEEN 1 AND 2)",
		},
		{
			name:  "Test NotBetween with datetime",
			table: "Test",
			args: args{
				column: "Test",
				value1: "2020-01-01 00:00:00",
				value2: "2023-01-01 00:00:00",
			},
			want: "SELECT * FROM Test WHERE (Test NOT BETWEEN '2020-01-01 00:00:00' AND '2023-01-01 00:00:00')",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).NotBetween(tt.args.column, tt.args.value1, tt.args.value2).Build(); got != tt.want {
				t.Errorf("NotBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_NotFindInSet(t *testing.T) {
	type args struct {
		column string
		values []interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test NotFindInSet",
			table: "Test",
			args: args{
				column: "Test",
				values: []interface{}{1, 2},
			},
			want: "SELECT * FROM Test WHERE NOT FIND_IN_SET(Test, 1, 2)",
		},
		{
			name:  "empty values not effected by NotFindInSet",
			table: "Test",
			args: args{
				column: "Test",
				values: []interface{}{},
			},
			want: "SELECT * FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).NotFindInSet(tt.args.column, tt.args.values...).Build(); got != tt.want {
				t.Errorf("NotFindInSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_NotIn(t *testing.T) {
	type args struct {
		column string
		values []interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test NotIn",
			table: "Test",
			args: args{
				column: "Test",
				values: []interface{}{1, 2},
			},
			want: "SELECT * FROM Test WHERE Test NOT IN (1, 2)",
		},
		{
			name:  "Test NotIn with empty values",
			table: "Test",
			args: args{
				column: "Test",
				values: []interface{}{},
			},
			want: "SELECT * FROM Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).NotIn(tt.args.column, tt.args.values...).Build(); got != tt.want {
				t.Errorf("NotIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_NotLike(t *testing.T) {
	type args struct {
		column string
		value  interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test NotLike",
			table: "Test",
			args: args{
				column: "Test",
				value:  "Test",
			},
			want: "SELECT * FROM Test WHERE Test NOT LIKE 'Test'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).NotLike(tt.args.column, tt.args.value).Build(); got != tt.want {
				t.Errorf("NotLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_NotWhere(t *testing.T) {
	type args struct {
		column   string
		operator string
		value    interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test NotWhere",
			table: "Test",
			args: args{
				column:   "Test",
				operator: "!=",
				value:    1,
			},
			want: "SELECT * FROM Test WHERE Test != 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).NotWhere(tt.args.column, tt.args.operator, tt.args.value).Build(); got != tt.want {
				t.Errorf("NotWhere() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Offset(t *testing.T) {

	type args struct {
		offset int
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test Offset",
			table: "Test",
			args: args{
				offset: 1,
			},
			want: "SELECT * FROM Test OFFSET 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Offset(tt.args.offset).Build(); got != tt.want {
				t.Errorf("Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Optimize(t *testing.T) {

	tests := []struct {
		name  string
		table string
		want  string
	}{
		{
			name:  "Test Optimize",
			table: "Test",
			want:  "OPTIMIZE TABLE Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Optimize(); got != tt.want {
				t.Errorf("Optimize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrBetween(t *testing.T) {

	type args struct {
		column string
		value1 interface{}
		value2 interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "Test OrBetween",
			table: "Test",
			args: args{
				column: "Test",
				value1: 1,
				value2: 2,
			},
			want: "SELECT * FROM Test WHERE Type = 2 OR (Test BETWEEN 1 AND 2)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where("Type", "=", "2").OrBetween(tt.args.column, tt.args.value1, tt.args.value2).Build(); got != tt.want {
				t.Errorf("OrBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrFindInSet(t *testing.T) {

	type args struct {
		column string
		values []interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "test",
			table: "test",
			args: args{
				column: "test",
				values: []interface{}{"test"},
			},
			want: "SELECT * FROM test WHERE Type = 2 OR FIND_IN_SET(test, 'test')",
		},
		{
			name:  "empty values not effected by OrFindInSet",
			table: "test",
			args: args{
				column: "test",
				values: []interface{}{},
			},
			want: "SELECT * FROM test WHERE Type = 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where("Type", "=", "2").OrFindInSet(tt.args.column, tt.args.values...).Build(); got != tt.want {
				t.Errorf("OrFindInSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrIn(t *testing.T) {

	type args struct {
		column string
		values []interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "test",
			table: "test",
			args: args{
				column: "Type",
				values: []interface{}{1, 2, 3},
			},
			want: "SELECT * FROM test WHERE Type = 2 OR Type IN (1, 2, 3)",
		},
		{
			name:  "test with empty values",
			table: "test",
			args: args{
				column: "Type",
				values: []interface{}{},
			},
			want: "SELECT * FROM test WHERE Type = 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where("Type", "=", "2").OrIn(tt.args.column, tt.args.values...).Build(); got != tt.want {
				t.Errorf("OrIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrLike(t *testing.T) {

	type args struct {
		column string
		value  interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "test",
			table: "test",
			args: args{
				column: "Test",
				value:  "test",
			},
			want: "SELECT * FROM test WHERE Type = 2 OR Test LIKE 'test'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where("Type", "=", "2").OrLike(tt.args.column, tt.args.value).Build(); got != tt.want {
				t.Errorf("OrLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrNotBetween(t *testing.T) {

	type args struct {
		column string
		value1 interface{}
		value2 interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "test",
			table: "test",
			args: args{
				column: "id",
				value1: 1,
				value2: 2,
			},
			want: "SELECT * FROM test WHERE Type = 2 OR (id NOT BETWEEN 1 AND 2)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where("Type", "=", "2").OrNotBetween(tt.args.column, tt.args.value1, tt.args.value2).Build(); got != tt.want {
				t.Errorf("OrNotBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrNotFindInSet(t *testing.T) {

	type args struct {
		column  string
		values  []interface{}
		column2 string
		values2 []interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "test",
			table: "test",
			args: args{
				column:  "test",
				values:  []interface{}{"test"},
				column2: "test2",
				values2: []interface{}{"test2"},
			},
			want: "SELECT * FROM test WHERE FIND_IN_SET(test, 'test') OR NOT FIND_IN_SET(test2, 'test2')",
		},
		{
			name:  "empty values not effected by OrNotFindInSet",
			table: "test",
			args: args{
				column:  "test",
				values:  []interface{}{"test"},
				column2: "test2",
				values2: []interface{}{},
			},
			want: "SELECT * FROM test WHERE FIND_IN_SET(test, 'test')",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).FindInSet(tt.args.column, tt.args.values...).OrNotFindInSet(tt.args.column2, tt.args.values2...).Build(); got != tt.want {
				t.Errorf("OrNotFindInSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrNotIn(t *testing.T) {

	type args struct {
		column  string
		values  []interface{}
		column2 string
		values2 []interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "OrNotIn",
			table: "users",
			args: args{
				column:  "id",
				values:  []interface{}{1, 2, 3},
				column2: "name",
				values2: []interface{}{"a", "b", "c"},
			},
			want: "SELECT * FROM users WHERE id IN (1, 2, 3) OR name NOT IN ('a', 'b', 'c')",
		},
		{
			name:  "OrNotIn with empty values",
			table: "users",
			args: args{
				column:  "id",
				values:  []interface{}{1, 2, 3},
				column2: "name",
				values2: []interface{}{},
			},
			want: "SELECT * FROM users WHERE id IN (1, 2, 3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).In(tt.args.column, tt.args.values...).OrNotIn(tt.args.column2, tt.args.values2...).Build(); got != tt.want {
				t.Errorf("OrNotIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrNotLike(t *testing.T) {

	type args struct {
		column  string
		value   interface{}
		column2 string
		value2  interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "OrNotLike",
			table: "users",
			args: args{
				column:  "name",
				value:   "test",
				column2: "name",
				value2:  "test2",
			},
			want: "SELECT * FROM users WHERE name LIKE 'test2' OR name NOT LIKE 'test'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Like(tt.args.column2, tt.args.value2).OrNotLike(tt.args.column, tt.args.value).Build(); got != tt.want {
				t.Errorf("OrNotLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrNotWhere(t *testing.T) {

	type args struct {
		column    string
		operator  string
		value     interface{}
		column2   string
		operator2 string
		value2    interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "OrNotWhere",
			table: "users",
			args: args{
				column:    "name",
				operator:  "=",
				value:     "test",
				column2:   "name",
				operator2: "=",
				value2:    "test2",
			},
			want: "SELECT * FROM users WHERE name = 'test' OR NOT name = 'test2'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where(tt.args.column, tt.args.operator, tt.args.value).OrNotWhere(tt.args.column2, tt.args.operator2, tt.args.value2).Build(); got != tt.want {
				t.Errorf("OrNotWhere() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrWhere(t *testing.T) {

	type args struct {
		column   string
		operator string
		column2  string
		value2   interface{}
		value    interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "or where",
			table: "users",
			args: args{
				column:   "id",
				column2:  "id",
				operator: "=",
				value:    1,
				value2:   2,
			},
			want: "SELECT * FROM users WHERE id = 2 OR id = 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).OrWhere(tt.args.column2, tt.args.operator, tt.args.value2).OrWhere(tt.args.column, tt.args.operator, tt.args.value).Build(); got != tt.want {
				t.Errorf("OrWhere() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_OrderBy(t *testing.T) {

	type args struct {
		column    string
		direction Order
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "order by",
			table: "users",
			args: args{
				column:    "id",
				direction: Orders.ASC,
			},
			want: "SELECT * FROM users ORDER BY id ASC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).OrderBy(tt.args.column, tt.args.direction).Build(); got != tt.want {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Pagination(t *testing.T) {

	type args struct {
		page    int
		perPage int
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "pagination",
			table: "users",
			args: args{
				page:    1,
				perPage: 10,
			},
			want: "SELECT * FROM users LIMIT 0, 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Pagination(tt.args.page, tt.args.perPage).Build(); got != tt.want {
				t.Errorf("Pagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Repair(t *testing.T) {

	tests := []struct {
		name  string
		table string
		want  string
	}{
		{
			name:  "repair",
			table: "users",
			want:  "REPAIR TABLE users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Repair(); got != tt.want {
				t.Errorf("Repair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_RightJoin(t *testing.T) {

	type args struct {
		table  string
		field1 string
		field2 string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "right join",
			table: "users",
			args: args{
				table:  "posts",
				field1: "users.id",
				field2: "posts.user_id",
			},
			want: "SELECT * FROM users RIGHT JOIN posts ON users.id = posts.user_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).RightJoin(tt.args.table, tt.args.field1, tt.args.field2).Build(); got != tt.want {
				t.Errorf("RightJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_RightOuterJoin(t *testing.T) {

	type args struct {
		table  string
		field1 string
		field2 string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "right outer join",
			table: "users",
			args: args{
				table:  "posts",
				field1: "users.id",
				field2: "posts.user_id",
			},
			want: "SELECT * FROM users RIGHT OUTER JOIN posts ON users.id = posts.user_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).RightOuterJoin(tt.args.table, tt.args.field1, tt.args.field2).Build(); got != tt.want {
				t.Errorf("RightOuterJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Select(t *testing.T) {

	type args struct {
		columns []string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "select",
			table: "users",
			args: args{
				columns: []string{"id", "name"},
			},
			want: "SELECT id, name FROM users",
		},
		{
			name:  "select with multiple columns",
			table: "users",
			args: args{
				columns: []string{"id", "name", "email"},
			},
			want: "SELECT id, name, email FROM users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Select(tt.args.columns...).Build(); got != tt.want {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("select with retry call", func(t *testing.T) {
		s := New()
		if got := s.Table("users").Select("id", "name").Select("email").Build(); got != "SELECT id, name, email FROM users" {
			t.Errorf("Select() = %v, want %v", got, "SELECT id, name, email FROM users")
		}
	})
}

func Test_sqb_Sum(t *testing.T) {

	type args struct {
		column string
		alias  string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "sum",
			table: "users",
			args: args{
				column: "id",
				alias:  "total",
			},
			want: "SELECT SUM(id) AS total FROM users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Sum(tt.args.column, tt.args.alias).Build(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Table(t *testing.T) {

	type args struct {
		table []string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "table",
			table: "users",
			args: args{
				table: []string{"users", "posts"},
			},
			want: "SELECT * FROM users, posts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Table(tt.args.table...).Build(); got != tt.want {
				t.Errorf("Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Update(t *testing.T) {
	tests := []struct {
		name   string
		values *M
		want   string
		table  string
	}{
		{
			name: "update",
			values: &M{
				"name": "test",
			},
			want:  "UPDATE users SET name = 'test' WHERE Type = 3",
			table: "users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where("Type", "=", "3").Update(tt.values); got != tt.want {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_UpdateMany(t *testing.T) {
	tests := []struct {
		name   string
		values []*M
		want   string
		table  string
	}{
		{
			name: "update",
			values: []*M{
				{
					"name": "test",
				},
				{
					"name": "test2",
				},
			},
			want:  "UPDATE users SET name = 'test', name = 'test2' WHERE Type = 3",
			table: "users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where("Type", "=", "3").UpdateMany(tt.values); got != tt.want {
				t.Errorf("UpdateMany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_Where(t *testing.T) {

	type args struct {
		column   string
		operator string
		value    interface{}
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "where",
			table: "users",
			args: args{
				column:   "id",
				operator: "=",
				value:    1,
			},
			want: "SELECT * FROM users WHERE id = 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).Where(tt.args.column, tt.args.operator, tt.args.value).Build(); got != tt.want {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_WhereNotNull(t *testing.T) {

	type args struct {
		column string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "where not null",
			table: "users",
			args: args{
				column: "name",
			},
			want: "SELECT * FROM users WHERE name IS NOT NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).WhereNotNull(tt.args.column).Build(); got != tt.want {
				t.Errorf("WhereNotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqb_WhereNull(t *testing.T) {
	type args struct {
		column string
	}
	tests := []struct {
		name  string
		table string
		args  args
		want  string
	}{
		{
			name:  "where null",
			table: "users",
			args: args{
				column: "name",
			},
			want: "SELECT * FROM users WHERE name IS NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			if got := s.Table(tt.table).WhereNull(tt.args.column).Build(); got != tt.want {
				t.Errorf("WhereNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

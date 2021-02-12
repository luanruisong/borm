package sqlbuilder

type (
	Table interface {
		TableName() string
	}

	Sql interface {
		Sql() string
		Args() []interface{}
	}

	SqlBuilder interface {
		Sql
		From(tableName string) SqlBuilder
		Set(key string, value interface{}) SqlBuilder
		Where(sql string, value ...interface{}) SqlBuilder
	}

	Selector interface {
		Sql
		From(tableName string) Selector
		Where(sql string, value ...interface{}) Selector
		Select(...string) Selector
		OrderBy(string) Selector
		GroupBy(string) Selector
		Limit(int64) Selector
		Offset(int64) Selector
	}
)

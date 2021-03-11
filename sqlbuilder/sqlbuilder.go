package sqlbuilder

type (
	Table interface {
		TableName() string
	}

	Sql interface {
		Sql() string
		Args() []interface{}
	}

	InsertBuilder interface {
		Sql
		From(tableName string) InsertBuilder
		Set(key string, value interface{}) InsertBuilder
	}

	DeleteBuilder interface {
		Sql
		From(tableName string) DeleteBuilder
		Where(sql string, value ...interface{}) DeleteBuilder
		And(sql string, value ...interface{}) DeleteBuilder
		Or(sql string, value ...interface{}) DeleteBuilder
	}

	UpdateBuilder interface {
		Sql
		From(tableName string) UpdateBuilder
		Set(key string, value interface{}) UpdateBuilder
		Where(sql string, value ...interface{}) UpdateBuilder
		And(sql string, value ...interface{}) UpdateBuilder
	}

	Selector interface {
		Sql
		Where(sql string, value ...interface{}) Selector
		And(sql string, value ...interface{}) Selector
		Or(sql string, value ...interface{}) Selector
		From(tableName string) Selector
		Select(...string) Selector
		OrderBy(string) Selector
		GroupBy(string) Selector
		Limit(int64) Selector
		Offset(int64) Selector
	}
)

const (
	DbwTag         = "dbw"
	DbwGreaterThan = "gt"
	DbwLessThan    = "lt"
	DbwEqual       = "eq"
)

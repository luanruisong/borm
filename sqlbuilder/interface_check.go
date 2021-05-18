package sqlbuilder

var (
	_ DeleteBuilder = &deleteBuilder{}
	_ Selector      = &selectBuilder{}
	_ UpdateBuilder = &updateBuilder{}
	_ InsertBuilder = &insertBuilder{}
)

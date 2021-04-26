package iterator

type IErr string

func (ie IErr) Error() string {
	return string(ie)
}

const (
	ErrNotPtr     = IErr("Arg Must Be Ptr")
	ErrNotSlice   = IErr("Arg Must Be Slice")
	ErrNoMoreRows = IErr("No More Rows")
)

// Auto-generated - DO NOT EDIT

package pg_query

type FunctionParameterMode uint

const (
	/* the assigned enum values appear in pg_proc, don't change 'em! */
	FUNC_PARAM_IN = iota
	FUNC_PARAM_OUT
	FUNC_PARAM_INOUT
	FUNC_PARAM_VARIADIC
	FUNC_PARAM_TABLE
)
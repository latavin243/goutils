package respcode

var (
	CodeOK = RespCode{0, "OK"}

	// below are some codes as example

	// internal code
	CodeInternalErr = RespCode{10000, "Internal error"}

	// input parameters code
	CodeParamErr = RespCode{20000, "Parameter error"}

	// business code
	CodeBusinessErr = RespCode{30000, "Business error"}
)

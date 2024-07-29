package respcode

var (
	CodeOK = RespCode{0, "OK"}

	// below are some codes as example

	// internal code
	CodeInternalErr = RespCode{10000, "Internal error"} // common/unclassified internal error

	// input parameters code
	CodeParamErr      = RespCode{20000, "Parameter error"} // common/unclassified parameter error
	CodeNoAccess      = RespCode{20001, "No access"}
	CodeInvalidPaging = RespCode{20002, "Invalid paging"}
	CodeRateLimited   = RespCode{20003, "Rate limited"}

	// business code
	CodeBusinessErr = RespCode{30000, "Business error"} // common/unclassified business error
	CodeNoData      = RespCode{30001, "No data"}

	// external service code
	CodeExternalErr = RespCode{40000, "External service error"} // common/unclassified external service error

	// external DB code
	CodeDBQueryErr = RespCode{50001, "DB query error"}
)

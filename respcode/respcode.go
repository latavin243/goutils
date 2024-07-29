package respcode

type RespCode struct {
	code    uint32
	message string
}

func (c *RespCode) ToError(errDetail string) *RespError {
	return &RespError{c.code, c.message, errDetail}
}

type RespError struct {
	code      uint32
	message   string
	errDetail string
}

func (e *RespError) Error() string {
	return e.errDetail
}

type RespPageInfo struct {
	Total       uint64 `json:"total"`
	PageNum     uint64 `json:"page_num"`
	PageSize    uint64 `json:"page_size"`
	HasNextPage bool   `json:"has_next_page"`
}

func NewRespPageInfo(total, pageNum, pageSize uint64) *RespPageInfo {
	hasNextPage := (pageNum*pageSize < total)
	return &RespPageInfo{
		Total:       total,
		PageNum:     pageNum,
		PageSize:    pageSize,
		HasNextPage: hasNextPage,
	}
}

func RespErr(respCode RespCode) map[string]interface{} {
	return map[string]interface{}{
		"code":    respCode.code,
		"message": respCode.message,
	}
}

func RespErrWithDetail(respCode RespCode, detail string) map[string]interface{} {
	return map[string]interface{}{
		"code":    respCode.code,
		"message": respCode.message,
		"detail":  detail,
	}
}

func RespErrWithData(respCode RespCode, data interface{}) (resp map[string]interface{}) {
	return map[string]interface{}{
		"code":    respCode.code,
		"message": respCode.message,
		"result":  data,
	}
}

// RespOKContent attaches data to a success response map
func RespOKContent(data interface{}) (resp map[string]interface{}) {
	return map[string]interface{}{
		"code":    CodeOK.code,
		"message": CodeOK.message,
		"result":  data,
	}
}

// RespOKContentWithPage attaches data and page info to a success response map
func RespOKContentWithPage(data interface{}, page *RespPageInfo) (resp map[string]interface{}) {
	return map[string]interface{}{
		"code":    CodeOK.code,
		"message": CodeOK.message,
		"result": map[string]interface{}{
			"page": page,
			"data": data,
		},
	}
}

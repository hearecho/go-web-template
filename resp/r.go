package resp

/**
返回结构体
*/

type R struct {
	Code int         `json:"status_code"`
	Path string      `json:"path"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r R) Ok() R {
	r.Code = SUCCESS
	r.Msg = MsgFlags[SUCCESS]
	return r
}
func (r R) Error() R {
	r.Code = ERROR
	r.Msg = MsgFlags[ERROR]
	return r
}

func (r R) SetData(data interface{}) R {
	r.Data = data
	return r
}

func (r R) SetStatus(code int) R {
	r.Code = code
	r.Msg = MsgFlags[code]
	return r
}

func (r R) SetPath(path string) R {
	r.Path = path
	return r
}


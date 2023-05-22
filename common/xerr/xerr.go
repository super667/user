package xerr

const (
	OK = 200

	NotFound = 500
)

func New(code int64, msg string) xerr {
	return xerr{code: code, msg: msg}
}

type xerr struct {
	code int64  `json:"code"`
	msg  string `json:"msg"`
}

func (x *xerr) Code() int64 {
	return x.code
}

func (x *xerr) Error() string {
	return x.msg
}

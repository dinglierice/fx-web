package consts

type CommonError struct {
	Code int
	Msg  string
	Data any
}

var (
	Success                    = CommonError{Code: 0, Msg: "success"}
	ParamError                 = CommonError{Code: 10001, Msg: "param error"}
	SystemError                = CommonError{Code: 10002, Msg: "system error"}
	NotFound                   = CommonError{Code: 10003, Msg: "not found"}
	Forbidden                  = CommonError{Code: 10004, Msg: "forbidden"}
	Unauthorized               = CommonError{Code: 10005, Msg: "unauthorized"}
	Conflict                   = CommonError{Code: 10006, Msg: "conflict"}
	BadRequest                 = CommonError{Code: 10007, Msg: "bad request"}
	InternalServerError        = CommonError{Code: 10008, Msg: "internal server error"}
	NotFoundError              = CommonError{Code: 10009, Msg: "not found"}
	BadGateway                 = CommonError{Code: 10010, Msg: "bad gateway"}
	GatewayTimeout             = CommonError{Code: 10011, Msg: "gateway timeout"}
	ServiceUnavailable         = CommonError{Code: 10012, Msg: "service unavailable"}
	MethodNotAllowed           = CommonError{Code: 10013, Msg: "method not allowed"}
	NotImplemented             = CommonError{Code: 10014, Msg: "not implemented"}
	RequestTimeout             = CommonError{Code: 10015, Msg: "request timeout"}
	TooManyRequests            = CommonError{Code: 10016, Msg: "too many requests"}
	UnprocessableEntity        = CommonError{Code: 10017, Msg: "unprocessable entity"}
	UnsupportedMediaType       = CommonError{Code: 10018, Msg: "unsupported media type"}
	UnavailableForLegalReasons = CommonError{Code: 10019, Msg: "unavailable for legal reasons"}
	NotAcceptable              = CommonError{Code: 10020, Msg: "not acceptable"}
	NotExtracted               = CommonError{Code: 10021, Msg: "not extracted"}
	NotModified                = CommonError{Code: 10022, Msg: "not modified"}
	NotAcceptableError         = CommonError{Code: 10023, Msg: "not acceptable"}
)

func (e CommonError) Error() string {
	return e.Msg
}

func (e CommonError) WithData(data any) CommonError {
	e.Data = data
	return e
}

package bizresponse

var (
	ErrUnknown          = NewCodeErrorWithInit(5001001, "The system is busy. Please try again later")
	ErrInvalidArgs      = NewCodeErrorWithInit(5001002, "Wrong parameters. Let's try again in a different way")
	ErrInternalFailed   = NewCodeErrorWithInit(5001003, "The server is straying, please try again later")
	ErrClientCancel     = NewCodeErrorWithInit(5001004, "The service connection is down, please try again later")
	ErrDeadlineExceed   = NewCodeErrorWithInit(5001005, "The server is busy. Please try again later")
	ErrRequestOutOfTime = NewCodeErrorWithInit(5001006, "Access expired, please try again")
)

var (
	pool = make(map[string]*CodeError)
)

/*
各服务错误码请按照以下定义的区段进行定义：
5001*** -> System
5002*** -> User
8001*** -> Account
*/

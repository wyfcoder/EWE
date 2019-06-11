package DealWrongs

const (
	TIMEFORCOOKIE = 100 //cookie的存活时间

	//空访问
	NullErrorInformation = "NULL ."
	NullErrorSolve       = "Please back or login ."

	ErrorCoderForSystem    = -1 //系统错误码
	SystemErrorInformation = "System error, sorry ."
	SystemErrorSolve       = "Please login."

	ErrorCodeForLogin     = 0 //登陆错误码
	LoginErrorInformation = "Your account does not exist or your password is wrong."
	LoginErrorSolve       = "Please check your input and login." +
		"If you forget your password , you can select forget to find back your password ."

	ErrorCodeForRegister     = 1 //注册错误码
	RegisterErrorInformation = ""
	RegisterErrorSolve       = "Please modify registration information to ensure your name, account unique."

	ErrorCodeForForget     = 2 //找回密码错误
	ForgetErrorInformation = "The account information is wrong."
	ForgetErrorSolve       = "Please modify the account information and try again."

	//连接超时
	ErrorCodeForTimeOut = 3
	TimeOutErrorInformation = "Connection timed out."
	TimeOutErrorSolve       = "Please login again."
	//被连续登陆
	ErrorCodeForLoginAgain = 4
	LoginAgainInformation = "Your account is logged in again."
	LoginAgainErrorSolve = "You can find back your account or just login."
)

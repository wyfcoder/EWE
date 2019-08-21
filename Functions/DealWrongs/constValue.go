package DealWrongs

const (
	TIMEFORCOOKIE = 100 //cookie的存活时间

	//空访问
	NullErrorInformation = "NULL ."
	NullErrorSolve       = "Please back or login ."

	ErrorCodeForSystem    = -1 //系统错误码
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

	//标签与文件的后缀不符合
	ErrorCodeForTag=5
	TagErrorInformation = "The file suffix you choose is different from the lable you choose."
	TagErrorSolve       = "You can check your file again."

	//重复文件错误
	ErrorCodeForFileName=6
	FileNameErrorInformation = "The file's name is used."
	FileNameErrorSolve       = "You can rename your file."

	//长传文件过大错误
	ErrorCodeForFileSize = 7
	FileSizeErrorInformation = "The file you upload is too larger."
	FileSizeErrorSolve ="Reduce the size of the file."

	//上传文件内容错误
	ErrorCodeForFileContent =8
	FileContentErrorInformation ="The content of file you upload is wrong."
	FileContentErrorSolve ="Check the file's content."

	//文件个数错误
	ErrorCodeForFileNumber =9
	FileNumberErrorInformation ="The Number of file you upload is more than one."
	FileNumberErrorSolve ="Upload file one by one."

	//文件为空
	ErrorCodeForFileNull =10
	FileNullErrorInformation = "The file is Null."
	FileNullErrorSolve ="Please check your file."

	//绘图文件不一致
	ErrorCodeForDrawTag = 11
	DrawTagErrorInformation = "Tags of files are not same."
	DrawTagErrorSolve = "Please select files again and make sure their tags are same."

	//绘制文件内容错误
	ErrorCodeForFileContents = 12
	FileContentsErrorInformation = "The file's content you select is wrong."
	FileContentsErrorSolve = "Check the file's content."

)

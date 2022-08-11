package errcode

var (
	ErrorGetUserInfoFail        = NewError(20010001, "获取用户信息失败")
	ErrorSignUpFail             = NewError(20020002, "注册失败")
	ErrorDuplicateUserWithEmail = NewError(20022003, "邮件已被注册")
	ErrorHashPasswordFail       = NewError(20020003, "密码哈希失败")
	ErrorInvalidPassword        = NewError(20020004, "无效的密码")
	ErrorTokenGenerateFail      = NewError(20020005, "生成token失败")
	ErrorSignOutFail            = NewError(20020006, "登出失败")
	ErrorSignInFail             = NewError(20002007, "登录失败")
	ErrorUploadFileFail         = NewError(20030001, "上传文件失败")
)

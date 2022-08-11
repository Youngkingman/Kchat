package errcode

var (
	// ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	// ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	// ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	// ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	// ErrorCountTagFail   = NewError(20010005, "统计标签失败")

	// ErrorGetArticleFail    = NewError(20020001, "获取单个文章失败")
	// ErrorGetArticlesFail   = NewError(20020002, "获取多个文章失败")
	// ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	// ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	// ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")
	ErrorGetUserInfoFail   = NewError(20010001, "获取用户信息失败")
	ErrorSignUpFail        = NewError(20020002, "注册失败")
	ErrorHashPasswordFail  = NewError(20020003, "密码哈希失败")
	ErrorInvalidPassword   = NewError(20020003, "无效的密码")
	ErrorTokenGenerateFail = NewError(20020004, "生成token失败")
	ErrorSignOutFail       = NewError(20020005, "登出失败")
	ErrorUploadFileFail    = NewError(20030001, "上传文件失败")
)

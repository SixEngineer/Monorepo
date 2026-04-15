package myerror

const (
	SuccessMessage = "ok"
)

const (
	ErrorCodeOK = 1000 // 成功

	ErrorCodeJsonFormatInvalid = 1001 // JSON 格式有误

	ErrorCodeTokenUploadFailed = 1002 // 令牌上传失败

	ErrorCodeProviderRegisterFailed = 1003 // Provider 注册失败

	ErrorCodeParameterInvalid = 1004 // 参数无效

	ErrorCodeProviderDeleteFailed = 1005 // Provider 删除失败

	ErrorCodeProviderUpdateFailed = 1006 // Provider 更新失败

	ErrorCodeProviderGetFailed = 1007 // Provider 获取失败

	ErrorCodeProviderListFailed = 1008 // Provider 列表获取失败

	ErrorCodeQuotaQueryFailed = 1009 // Quota 查询失败

	ErrorCodeQuotaSyncFailed = 1010 // Quota 同步失败

)

package myerror

const (
	SuccessMessage = "ok"
)

const (
	ErrorCodeOK = 1000 // 成功

	ErrorCodeJsonFormatInvalid = 1001 // JSON 格式有误

	ErrorCodeTokenUploadFailed = 1002 // 令牌上传失败

	ErrorCodeAccessTokenExpired       = 1001 // access token 过期
	ErrorCodeAccessTokenInvalid       = 1002 // access token 无效
	ErrorCodeRefreshTokenExpired      = 1003 // refresh token 过期
	ErrorCodeRefreshTokenInvalid      = 1004 // refresh token 无效
	ErrorCodeAccessTokenUploadFailed  = 1005 // access token 上传失败
	ErrorCodeRefreshTokenUploadFailed = 1006 // refresh token 上传失败

	ErrorCodeUploadExceeded = 1007 // 上传文件大小超过容量剩余

)
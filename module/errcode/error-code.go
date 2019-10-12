package errcode

//go:generate go-easy generate error-code --type ResponseErrorCode
type ResponseErrorCode uint64

const (
	// 输入的两次密码不一致
	ResponseErrorCodeRePassword ResponseErrorCode = iota
)

package mcode


const (

	SUCCESS int =  2000
	ERROR int = 5000
	INVALID_PARAMS int = 4000
	ERROR_EXIST_TAG int = 10001
	ERROR_NOT_EXIST_TAG int = 10002
	ERROR_NOT_EXIST_APP_PYPE int = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL int = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT int = 20002
	ERROR_AUTH_TOKEN int = 20003
	ERROR_AUTH int = 20004
	ERROR_CELERY_TIOMEOUT int = 40051
	ERROR_NOT_FOUND int = 42000
	METHOD_NOT_ALLOWED int = 403
	NO_ROUTE int = 404
	SERVICE_FRAMEWORK_ERROR int = 5555
)
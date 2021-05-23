package errmsg

// Database
const (
	ErrDBQuery          = "err_db_query"
	ErrDBGet            = "err_db_get"
	ErrDBInsert         = "err_db_insert"
	ErrDBUpdate         = "err_db_update"
	ErrDBDelete         = "err_db_delete"
	ErrDBCount          = "err_db_count"
	ErrDBNotFound       = "err_db_not_found"
	ErrDBCheckDuplicate = "err_db_check_duplicate"
	ErrDBDuplicate      = "err_db_duplicate"
)

// System
const (
	ServerError      = "server_error"
	UnknownError     = "unknown_error"
	APINotFound      = "api_not_found"
	Unauthorized     = "unauthorized"
	PermissionDenied = "permission_denied"
)

// User
const (
	AccountNotFound   = "account_not_found"
	IncorrectPassword = "incorrect_password"
	ErrLogout         = "err_logout"
)

// File
const (
	ErrFileUpload = "err_file_upload"
	ErrFileGet    = "err_file_get"
)

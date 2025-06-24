package grpcPb

import "github.com/whoami00911/Audit-web-app-service/internal/domain"

var (
	actions = map[string]LogRequest_Actions{
		"SignUp":      LogRequest_SignUp,
		"SignIn":      LogRequest_SignIn,
		"Logout":      LogRequest_Logout,
		"Upload":      LogRequest_Upload,
		"GetFile":     LogRequest_Get_file,
		"GetFiles":    LogRequest_Get_files,
		"DeleteFile":  LogRequest_Delete_file,
		"DeleteFiles": LogRequest_Delete_files,
	}

	methods = map[string]LogRequest_Methods{
		"GET":    LogRequest_GET,
		"PUT":    LogRequest_PUT,
		"POST":   LogRequest_POST,
		"DELETE": LogRequest_DELETE,
	}
)

func ToPbAction(action string) (LogRequest_Actions, error) {
	value, ok := actions[action]
	if !ok {
		return -1, domain.ErrNoAction
	}

	return value, nil
}

func ToPbMethod(method string) (LogRequest_Methods, error) {
	value, ok := methods[method]
	if !ok {
		return -1, domain.ErrNoAction
	}

	return value, nil
}

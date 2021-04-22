package errcodes

import "net/http"

const (
	InvalidAPIKeys                = "INVALID_API_KEYS"
	CanNotRetrieveProviderDetails = "CANNOT_RETRIEVE_PROVIDER_DETAILS"
	ErrorSendingTestSms           = "ERROR_SENDING_TEST_SMS"
	ErrorSendingTestPush          = "ERROR_SENDING_TEST_PUSH"
)

func HttpStatusCodeByErrCode(code string) int {
	if status, ok := statusCodes[code]; ok {
		return status
	}
	panic("code is not present")
}

var statusCodes = map[string]int{
	InvalidAPIKeys:                http.StatusFailedDependency,
	CanNotRetrieveProviderDetails: http.StatusInternalServerError,
}

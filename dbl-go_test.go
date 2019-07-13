package dbl_go

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"testing"
)

func getTestingApi() DBLApi {
	envVariableName := "DBL_ACCESS_TOKEN"
	accessToken := os.Getenv(envVariableName)
	if accessToken == "" {
		panic(fmt.Sprintf("you have to set %s environment variable to continue testing", envVariableName))
	}

	return NewDBLApi(accessToken)
}

func TestDBLApi_getRequestURL(t *testing.T) {
	endpoint := "/test"
	expectedResult := baseURL + endpoint
	apiObject := NewDBLApi("")

	actualResult := apiObject.getRequestURL(endpoint)
	if actualResult != expectedResult {
		t.Fatalf("request url does not match expectations (\"%s\" != \"%s\")", expectedResult, actualResult)
		return
	}
}

func TestDBLApi_getBaseRequest(t *testing.T) {
	fakeAccessToken := "x"

	api := NewDBLApi(fakeAccessToken)
	request := api.getBaseRequest()

	if request == nil {
		t.Fatal("getBaseRequest() returned null")
	}

	headers := request.Header
	headerAccessToken := headers.Get("Authorization")

	if headerAccessToken != fakeAccessToken {
		t.Fatalf("authentication header is incorrect (%s)", headerAccessToken)
	}
}

func TestDBLApi_PostStatsSimple(t *testing.T) {
	api := getTestingApi()
	err := api.PostStatsSimple(576)
	if err != nil {
		t.Error(errors.Wrap(err, "posting bot stats returned error:"))
	}
}

package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func Function_test(t *testing.T) {
	headers := http.Header{"Authorization": []string{"ApiKey 512"}}
	got, err := GetAPIKey(headers)
	want := "512"
	var want_err error = nil
	//want_err = 

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	if !reflect.DeepEqual(err, want_err) {
		t.Fatalf("expected: %v, got: %v | Detection of lacking authorization doesn't work.", want, got)
	}
}

func Test_no_auth_error(t *testing.T) {
	headers := http.Header{}
	got, err := GetAPIKey(headers)
	want := ""
	want_err := ErrNoAuthHeaderIncluded
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	if !reflect.DeepEqual(err, want_err) {
		t.Fatalf("expected: %v, got: %v | Detection of lacking authorization doesn't work.", want, got)
	}

}

func Test_malformed_error(t *testing.T) {
	headers := http.Header{"Authorization": []string{"afdsf"}}
	got, err := GetAPIKey(headers)
	want := ""
	want_err := errors.New("malformed authorization header")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	if !reflect.DeepEqual(err, want_err) {
		t.Fatalf("expected: %v, got: %v | Detection of lacking authorization doesn't work.", want, got)
	}

}

/*
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
*/

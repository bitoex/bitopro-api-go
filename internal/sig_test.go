package internal

import (
	"testing"
)

func Test_getSig(t *testing.T) {
	if getSig("bitopro", "eyJpZGVudGl0eSI6ImhjbWxpbmpAZ21haWwuY29tIiwibm9uY2UiOjE1NTQzODA5MDkxMzF9") != "01a85a9083db47c20da7196380598f3feacd3c76a9077aaf7ffaf08ce0091abf65b61778792607b010921adfe1c2941a" {
		t.Error("not match")
	}
}

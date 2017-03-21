package stringutil

import "testing"

func Test_Reverse(t *testing.T) {
	if Reverse("apple") == "elppa" {
		t.Log("Passed")
	} else {
		t.Error("Failed")
	}
}

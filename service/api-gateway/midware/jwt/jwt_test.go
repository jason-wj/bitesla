package jwt

import "testing"

func TestValidate(t *testing.T) {
	validData("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiNmMxMTQ3MDQtOWFhYS00NzU2LWE5ZWMtNGMyMGFmNzIyMWE4IiwiYXBwX3NlY3JldCI6ImExN2VmMzI2YjZiMWU0ZjRhY2UyMmNkMTcyZDFhZmYzIiwiZXhwIjoxOTA1MTI1MzQ5LCJpc3MiOiJhaWNoYWluIn0.scBO5rj5moQdyfHPF-ElYUGjTzZaCvRAWDgcklwgSN", "1")
}

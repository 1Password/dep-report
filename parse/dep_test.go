package parse

import (
	"testing"
)

func TestReadGopkg(t *testing.T) {
	filepath := "../Gopkg.lock"

	_, err := ReadGopkg(filepath)
	if err != nil {
		t.Errorf("unable to read Gopkg, %v",err)
	}
}

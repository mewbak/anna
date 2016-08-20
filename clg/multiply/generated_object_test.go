package multiply

// This file is generated by the CLG generator. Don't edit it manually. The CLG
// generator is invoked by go generate. For more information about the usage of
// the CLG generator check https://github.com/xh3b4sd/clggen or have a look at
// the network package. There is the go generate statement to invoke clggen.

import (
	"testing"
)

func Test_CLG_GetID(t *testing.T) {
	firstCLG := MustNew()
	secondCLG := MustNew()

	if firstCLG.GetID() == secondCLG.GetID() {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_CLG_GetType(t *testing.T) {
	newCLG := MustNew()
	objectType := newCLG.GetType()

	if objectType != ObjectType {
		t.Fatal("expected", ObjectType, "got", objectType)
	}
}
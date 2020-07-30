package uidgen

import (
	"reflect"
	"testing"
)

func TestUIDGen(t *testing.T) {
	guid := New(0, VARCHAR26)
	uid := guid.UID()
	p, ok := guid.Parse(uid.String())
	if !ok {
		t.Errorf("Generated UID %s failed to parse", uid)
	}
	if !reflect.DeepEqual(uid, p) {
		t.Errorf("Generated UID %s not equal with parse", uid)
	}
}

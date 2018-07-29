package recwatch

import (
	"github.com/fsnotify/fsnotify"
	"testing"
)

func TestEvent(t *testing.T) {
	fe := &fsnotify.Event{}
	e := &Event{}
	if fe.String() != e.String() {
		t.Error("String mismatch")
	}
}

package booker

import (
	"testing"
)

func TestBookAPosition(t *testing.T) {
	booker := New()
	booker.Start()
	bookChannel, releaseChannel := booker.Channel()
	resChannel := make(chan bool)
	bookChannel <- Book{0, resChannel}

	if !<-resChannel {
		t.Error("No booking realised for position 0")
	}

	releaseChannel <- Release{0, resChannel}
	if !<-resChannel {
		t.Error("Release failed for position 0")
	}
}

// More tests to come

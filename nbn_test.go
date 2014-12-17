package nbn

import (
	"log"
	"testing"
)

func TestRolloutInfo(t *testing.T) {
	_, err := RolloutInfo(-12.376362, 130.894135)

	if err != nil {
		log.Fatal(err)
	}
}

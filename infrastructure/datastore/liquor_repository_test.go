package datastore

import (
	"testing"
)

func TestCreateFatal(t *testing.T) {
	t.Error("error")
}
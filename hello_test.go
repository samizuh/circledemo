package main

import (
	"testing"
)
  
func TestUpper(t *testing.T) {
	incoming := "Test String"
	expected := "TEST STRINGx"
	actual := ToUpper(incoming)
  
	if actual != expected {
	  t.Errorf("Expected the string of %s to be %s but instead got %s!", incoming, expected, actual)
	}
}

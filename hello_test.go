package main

import (
	"testing"
)
  
func TestSum(t *testing.T) {
	incoming := "Test String"
	expected := "TEST STRING"
	actual := toUpper(incoming)
  
	if actual != expected {
	  t.Errorf("Expected the string of %s to be %s but instead got %s!", incoming, expected, actual)
	}
}
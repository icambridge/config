package config

import (
	"testing"
)

func Test_GetString_Returns_StringValue(t *testing.T) {
	
	c, err := NewConfig("example/test.conf")
	
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	
	s := c.GetString("hello")
	
	if expected := "world"; s != expected {
		t.Errorf("Expected %q, got %q", expected, s)
		return
	} 
}

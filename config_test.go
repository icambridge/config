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
	
	s, err := c.GetString("hello")

	if err != nil {
		t.Errorf("%v", err)
		return
	}
	
	if expected := "world"; s != expected {
		t.Errorf("Expected %q, got %q", expected, s)
		return
	} 
}

func Test_GetInt_Returns_IntValue(t *testing.T) {

	c, err := NewConfig("example/test.conf")

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	s, err := c.GetInt("int")

	if err != nil {
		t.Errorf("%v", err)
		return
	}
	
	if expected := 5; s != expected {
		t.Errorf("Expected %q, got %q", expected, s)
		return
	}
}

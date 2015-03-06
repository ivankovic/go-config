package config

import (
	"testing"
)

func TestGetFromRealFile(t *testing.T) {
	c := Load("test_config.conf")

	for ti, tt := range []struct {
		section       string
		name          string
		default_value string
		expected      string
	}{
		{"global", "string", "FAIL", "global works"},
		{"section", "string", "FAIL", "section works"},
		{"global", "nonexisting", "DEFAULT", "DEFAULT"},
		{"section", "nonexisting", "DEFAULT", "DEFAULT"},
		{"section", "novalue", "DEFAULT", "DEFAULT"},
		{"no-such-section", "nope", "DEFAULT", "DEFAULT"},
		{"global", "unicode", "FAIL", "Работает с Unicode тоже!"},
	} {
		result := c.Get(tt.section, tt.name, tt.default_value)
		if result != tt.expected {
			t.Errorf("(%d) Get(%s, %s, %s) = %s expected %s", ti, tt.section, tt.name, tt.default_value, result, tt.expected)
		}
	}
}

func TestSectionFromRealFile(t *testing.T) {
	c := Load("test_config.conf")

	s, err := c.GetSection("section")
	if err != nil {
		t.Errorf("GetSection('section') failed: %s", err)
		t.FailNow()
	}

	for ti, tt := range []struct {
		name          string
		default_value string
		expected      string
	}{
		{"string", "FAIL", "section works"},
		{"nonexisting", "DEFAULT", "DEFAULT"},
	} {
		result := s.Get(tt.name, tt.default_value)
		if result != tt.expected {
			t.Errorf("(%d) Get(%s, %s) = %s expected %s", ti, tt.name, tt.default_value, result, tt.expected)
		}
	}

	s, err = c.GetSection("no-such-section")
	if err == nil {
		t.Errorf("GetSection('no-such-section') did not fail!")
		t.FailNow()
	}
}

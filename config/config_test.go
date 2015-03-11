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

func MockedConfig() Config {
	return Config{
		Sections: map[string]Section{
			"s1": Section{
				Values: map[string]string{
					"k": "v",
					"i": "14",
				},
			},
		},
	}
}

func TestGet(t *testing.T) {
	c := MockedConfig()

	for ti, tt := range []struct {
		section       string
		name          string
		default_value string
		expected      string
	}{
		{"global", "a", "OK", "OK"},
		{"s1", "k", "FAIL", "v"},
	} {
		result := c.Get(tt.section, tt.name, tt.default_value)
		if result != tt.expected {
			t.Errorf("(%d) Get(%s, %s, %s) = %s expected %s", ti, tt.section, tt.name, tt.default_value, result, tt.expected)
		}
	}
}

func TestGetInt(t *testing.T) {
	c := MockedConfig()

	for ti, tt := range []struct {
		section       string
		name          string
		default_value int
		expected      int
	}{
		{"s1", "i", 9999, 14},
		{"s1", "in", 9999, 9999},
	} {
		result := c.GetInt(tt.section, tt.name, tt.default_value)
		if result != tt.expected {
			t.Errorf("(%d) Get(%s, %s, %d) = %d expected %d", ti, tt.section, tt.name, tt.default_value, result, tt.expected)
		}
	}
}

package config

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

type Section struct {
	values map[string]string
}

type Config struct {
	sections map[string]Section
}

func addEntry(x string, section string, c *Config) string {
	x = strings.TrimSpace(x)

	if x == "" || x[0] == '#' {
		return section
	}

	if x[0] == '[' {
		section = x[1 : len(x)-1]
	} else {
		var name, value string

		split := strings.SplitN(x, "=", 2)

		name = strings.TrimSpace(split[0])
		if len(split) > 1 {
			value = strings.TrimSpace(split[1])
		}

		s, ok := (*c).sections[section]
		if !ok {
			s = Section{values: make(map[string]string)}
			(*c).sections[section] = s
		}
		s.values[name] = value
	}

	return section
}

func Load(path string) Config {
	c := Config{sections: make(map[string]Section)}
	log.Printf("Trying to read config from %s", path)

	f, err := os.Open(path)
	if err != nil {
		log.Printf("Unable to open %s. Using defaults.", path)
		return c
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	section := "global"

	for s.Scan() {
		section = addEntry(s.Text(), section, &c)
	}

	if err := s.Err(); err != nil {
		log.Printf("Error reading config file '%s': %s", path, err)
	}

	return c
}

func (c Config) Get(section, name, default_value string) string {
	s, ok := c.sections[section]
	if ok {
		return s.Get(name, default_value)
	}
	return default_value
}

func (c Config) GetSection(section string) (Section, error) {
	s, ok := c.sections[section]
	if !ok {
		return Section{}, errors.New("No such section!")
	}
	return s, nil
}

func (s Section) Get(name, default_value string) string {
	v, ok := s.values[name]
	if ok {
		if v != "" {
			return v
		}
	}
	return default_value
}
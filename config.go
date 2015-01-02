package config

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	m map[string]string
}

func (c Config) GetString(key string) string {
	return c.m[key]
}

func NewConfig(filename string) (Config, error) {
	c := Config{}
	m := map[string]string{}
	f, err := os.Open(filename)
	
	if err != nil {
		return c, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, "=")
		key := parts[0]
		value := parts[1]
		m[key] = value
	}
	c.m = m
	return c, nil
}

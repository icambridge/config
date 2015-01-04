package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	m map[string]string
}

func (c Config) GetString(key string) (string, error) {
	return c.m[key], nil
}

func (c Config) GetInt(key string) (int, error) {
	s, err := c.GetString(key)
	if err != nil {
		return -1, err
	}
	i, err := strconv.Atoi(s)
	return i, err
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

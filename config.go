package pino

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config holds the configuration that Pino expects
type Config struct {
	IRC            IRCConfig                   `yaml:"IRC"`
	Slack          SlackConfig                 `yaml:"Slack"`
	ChannelMapping map[SlackChannel]IRCChannel `yaml:"ChannelMapping"`
}

// IRCChannel is the name of an IRC channel, like "#CAA"
type IRCChannel string

// IRCChannelKey is an optional password for an IRC channel
type IRCChannelKey string

// SlackChannel is the name of a Slack channel, like "#CAA-on-Slack"
type SlackChannel string

// IRCConfig define the IRC-specific config
type IRCConfig struct {
	Nickname string                       `yaml:"Nickname"`
	Name     string                       `yaml:"Name"`
	Server   string                       `yaml:"Server"`
	IsSSL    bool                         `yaml:"IsSSL"`
	Channels map[IRCChannel]IRCChannelKey `yaml:"Channels"`
}

// SlackConfig defines the Slack-specific config
type SlackConfig struct {
	Owner    string                  `yaml:"Owner"`
	Token    string                  `yaml:"Token"`
	Channels map[SlackChannel]string `yaml:"Channels"`
}

// LoadConfig returns the Config parsed from the given config file path
func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("Unable to read config file from %v: %v", path, err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("Unable to parse YAML from config file %v: %v", path, err)
	}

	return config, nil
}

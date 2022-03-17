package config

import (
)


type Config struct {
	Listen []string       `yaml:"listen"`
	Peers map[string]Peer `yaml:"peers"`
}

type Peer struct {
	Address				string `yaml:"address"`
	Port     			int16  `yaml:"port"`
	Interval 			int    `yaml:"interval"`  			// target interval in ms
	DetectionMultiplier int    `yaml:"detectionMultiplier"`
}

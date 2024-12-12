package config

type AgentConfig struct {
    ID         string `yaml:"id"`
    ServerAddr string `yaml:"server_addr"`
    Labels     map[string]string `yaml:"labels"`
    
    Metrics struct {
        Enable   bool   `yaml:"enable"`
        Interval string `yaml:"interval"`
    } `yaml:"metrics"`
    
    Probe struct {
        Types []string `yaml:"types"`
    } `yaml:"probe"`
} 
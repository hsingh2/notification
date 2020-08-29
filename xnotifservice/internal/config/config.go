package config

import (
	"time"
)

//Config specifies the config required for the service
type Config struct {
	InfoApp struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Attributes  struct {
			DisplayName string `yaml:"display-name"`
			Parent      string `yaml:"parent"`
			Type        string `yaml:"type"`
		} `yaml:"attributes"`
	} `yaml:"info.app"`

	InfoBuild struct {
		Artifact      string    `yaml:"artifact"`
		BuildDateTime time.Time `yaml:"buildDateTime"`
		BuildNumber   string    `yaml:"buildNumber"`
		Release       string    `yaml:"release"`
		Group         string    `yaml:"group"`
		Name          string    `yaml:"name"`
		Version       string    `yaml:"version"`
	} `yaml:"info.build"`

	Server struct {
		Port        int    `yaml:"port"`
		ContextPath string `yaml:"context-path"`
		Enabled     bool   `yaml:"enabled"`
	} `yaml:"server"`

	SpringDataCassandra struct {
		KeyspaceName string `yaml:"keyspace-name"`
		Enabled      bool   `yaml:"enabled"`
	} `yaml:"spring.data.cassandra"`

	SpringDatasource struct {
		Name    string `yaml:"name"`
		Enabled bool   `yaml:"enabled"`
	} `yaml:"spring.datasource"`

	SpringCloudConsul struct {
		Enabled bool `yaml:"enabled"`
		Config  struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"config"`
		Discovery struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"discovery"`
	} `yaml:"spring.cloud.consul"`

	SpringCloudVault struct {
		Enabled bool `yaml:"enabled"`
		Generic struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"generic"`
	} `yaml:"spring.cloud.vault"`

	SpringCloudStreamKafkaBinder struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"spring.cloud.stream.kafka.binder"`

	SpringRedis struct {
		Enable bool `yaml:"enable"`
	} `yaml:"spring.redis"`

	Stats struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"stats"`

	Swagger struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"swagger"`

	Trace struct {
		Enabled  bool `yaml:"enabled"`
		Reporter struct {
			Name string `yaml:"name"`
		} `yaml:"reporter"`
	} `yaml:"trace"`
}

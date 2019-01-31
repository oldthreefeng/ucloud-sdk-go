package external

import (
	"fmt"
	"time"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// ConfigProvider is the provider to store and provide config/credential instance
type ConfigProvider interface {
	Credential() *auth.Credential

	Config() *ucloud.Config
}

// config will read configuration
type config struct {
	// Named profile
	Profile              string
	SharedConfigFile     string
	SharedCredentialFile string

	// Credential configuration
	PublicKey  string
	PrivateKey string

	// Client configuration
	ProjectId string
	Zone      string
	Region    string
	BaseUrl   string
	Timeout   time.Duration
}

func newConfig() *config {
	return &config{}
}

func (c *config) loadEnv() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	err = c.merge(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (c *config) loadFile() error {
	cfgFile := c.SharedConfigFile
	if len(cfgFile) == 0 {
		cfgFile = DefaultSharedConfigFile()
	}

	credFile := c.SharedCredentialFile
	if len(credFile) == 0 {
		credFile = DefaultSharedCredentialsFile()
	}

	cfg, err := loadSharedConfigFile(cfgFile, credFile, c.Profile)
	if err != nil {
		return err
	}

	err = c.merge(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (c *config) merge(other *config) error {
	if other == nil {
		return nil
	}

	c.Profile = other.Profile
	c.SharedConfigFile = other.SharedConfigFile
	c.SharedCredentialFile = other.SharedCredentialFile
	c.PublicKey = other.PublicKey
	c.PrivateKey = other.PrivateKey
	c.ProjectId = other.ProjectId
	c.Zone = other.Zone
	c.Region = other.Region
	c.BaseUrl = other.BaseUrl
	c.Timeout = other.Timeout
	return nil
}

// Config is the configuration of ucloud client
func (c *config) Config() *ucloud.Config {
	cfg := ucloud.NewConfig()
	setStringify(&cfg.ProjectId, c.ProjectId)
	setStringify(&cfg.Zone, c.Zone)
	setStringify(&cfg.Region, c.Region)
	setStringify(&cfg.BaseUrl, c.BaseUrl)

	if c.Timeout != time.Duration(0) {
		cfg.Timeout = c.Timeout
	}
	return &cfg
}

// Credential is the configuration of ucloud authorization information
func (c *config) Credential() *auth.Credential {
	cred := auth.NewCredential()
	setStringify(&cred.PublicKey, c.PublicKey)
	setStringify(&cred.PrivateKey, c.PrivateKey)
	return &cred
}

// LoadDefaultUCloudConfig is the default loader to load config
func LoadDefaultUCloudConfig() (ConfigProvider, error) {
	cfg := newConfig()

	if err := cfg.loadEnv(); err != nil {
		return nil, fmt.Errorf("error on loading env, %s", err)
	}

	if err := cfg.loadFile(); err != nil {
		return nil, fmt.Errorf("error on loading shared config file, %s", err)
	}

	return cfg, nil
}

func setStringify(p *string, s string) {
	if p != nil && len(s) != 0 {
		*p = s
	}
}
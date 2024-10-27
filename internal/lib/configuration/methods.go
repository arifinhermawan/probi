package configuration

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/secret"
	internalContext "github.com/arifinhermawan/probi/internal/lib/context"
	"gopkg.in/yaml.v2"
)

func (c *Configuration) GetConfig() *AppConfig {
	ctx := internalContext.DefaultContext()
	c.doLoadConfigOnce.Do(func() {
		cfg, err := c.loadConfig(ctx)
		if err != nil {
			log.Error(ctx, nil, err, "[GetConfig] c.loadConfig() got error")
		}

		c.config = cfg
	})

	return &c.config
}

func (c *Configuration) loadConfig(ctx context.Context) (AppConfig, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	filePath := fmt.Sprintf("etc/files/config.%s.yaml", env)
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Error(ctx, nil, err, "[LoadConfig] os.ReadFile() got error")
		return AppConfig{}, err
	}

	var config AppConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Error(ctx, nil, err, "[LoadConfig] yaml.Unmarshal() got error")
		return AppConfig{}, err
	}

	secretConfigPath := fmt.Sprintf("etc/credentials/secret.%s.json", env)
	secret, err := secret.GetSecret(secretConfigPath)
	if err != nil {
		log.Error(ctx, nil, err, "[LoadConfig] secret.GetSecret() got error")
		return AppConfig{}, err
	}

	var secretConfig AppConfig
	err = json.Unmarshal(secret, &secretConfig)
	if err != nil {
		log.Error(ctx, nil, err, "[LoadConfig] json.Unmarshal() got error")
		return AppConfig{}, err
	}

	config = addSecretConfig(config, secretConfig)

	return config, nil
}

func addSecretConfig(config AppConfig, secretConfig AppConfig) AppConfig {
	config.Database = secretConfig.Database
	config.Hash = secretConfig.Hash
	config.NewRelic = secretConfig.NewRelic
	config.Redis = secretConfig.Redis

	return config
}

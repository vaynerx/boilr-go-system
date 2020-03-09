package app

import (
	bytes "bytes"
	context "context"
	base64 "encoding/base64"
	strings "strings"

	viper "github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// Config struct contains the application parameters
type Config struct {
	api    *SkaffoldAPIData    // Skaffold API definitions
	system *SkaffoldSystemData // Skaffold system definitions
	client *SkaffoldClientData // Skaffold client definitions
	app    *SkaffoldAppData    // Skaffold application definitions
}

func NewConfig(ctx context.Context) *Config {
	return &Config{}
}

func (app *App) InitConfig(ctx context.Context) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	*app.config = config

	return nil
}

type remoteConfig struct {
	remoteConfigProvider      string // remote configuration source ("consul", "etcd", "envvar")
	remoteConfigEndpoint      string // remote configuration URL (ip:port)
	remoteConfigPath          string // remote configuration path where to search fo the configuration file ("/config/dummy")
	remoteConfigSecretKeyring string // path to the openpgp secret keyring used to decript the remote configuration data ("/etc/dummy/configkey.gpg")
	remoteConfigData          string // base64 encoded JSON configuration data to be used with the "envvar" provider
}

// isEmpty returns true if all the fields are empty strings
func (rcfg remoteConfig) isEmpty() bool {
	return rcfg.remoteConfigProvider == "" && rcfg.remoteConfigEndpoint == "" && rcfg.remoteConfigPath == "" && rcfg.remoteConfigSecretKeyring == ""
}

// getConfig returns the configuration parameters
func getConfig() (config Config, err error) {
	cfg, rcfg, err := getLocalConfig()
	if err != nil {
		return config, err
	}
	return getRemoteConfig(cfg, rcfg)
}

// getRemoteConfig returns the remote & local configuration parameters
func _getConfig() (config Config, err error) {
	cfg, rcfg, err := getLocalConfig()
}

// getLocalConfig returns the local configuration parameters
func getLocalConfig() (cfg Config, rcfg remoteConfig, err error) {

	viper.Reset()

	viper.SetDefault("app.hydra.host", SkaffoldAppHydraHost)
	viper.SetDefault("app.hydra.port", SkaffoldAppHydraPort)
	viper.SetDefault("app.hydra.address", SkaffoldAppHydraAddress)
	viper.SetDefault("app.hydra.adminHost", SkaffoldAppHydraAdminHost)
	viper.SetDefault("app.hydra.adminPort", SkaffoldAppHydraAdminPort)
	viper.SetDefault("app.hydra.adminAddress", SkaffoldAppHydraAdminAddress)

	viper.SetDefault("app.kafka.host", SkaffoldAppKafkaHost)
	viper.SetDefault("app.kafka.port", SkaffoldAppKafkaPort)
	viper.SetDefault("app.kafka.address", SkaffoldAppKafkaAddress)

	viper.SetDefault("app.prisma.host", SkaffoldAppPrismaHost)
	viper.SetDefault("app.prisma.port", SkaffoldAppPrismaPort)
	viper.SetDefault("app.prisma.address", SkaffoldAppPrismaAddress)
	viper.SetDefault("app.prisma.secret", SkaffoldAppPrismaSecret)

	// configuration type
	viper.SetConfigType("json")

	// add local configuration paths
	for _, cpath := range ConfigPath {
		viper.AddConfigPath(cpath)
	}

	// name of the global configuration file without extension
	viper.SetConfigName("config")

	// Find and read the global configuration file (if any)
	err = viper.ReadInConfig()
	if err != nil {
		return cfg, rcfg, err
	}

	// read configuration parameters
	cfg = getViper()

	// support environment variables for the remote configuration
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(EnvironmentVariablesPrefix, "-", "_", -1)) // will be uppercased automatically
	envVar := []string{
		"remoteConfigProvider",
		"remoteConfigEndpoint",
		"remoteConfigPath",
		"remoteConfigSecretKeyring",
		"remoteConfigData",
	}
	for _, ev := range envVar {
		_ = viper.BindEnv(ev)
	}

	rcfg = remoteConfig{
		remoteConfigProvider:      viper.GetString("remoteConfigProvider"),
		remoteConfigEndpoint:      viper.GetString("remoteConfigEndpoint"),
		remoteConfigPath:          viper.GetString("remoteConfigPath"),
		remoteConfigSecretKeyring: viper.GetString("remoteConfigSecretKeyring"),
		remoteConfigData:          viper.GetString("remoteConfigData"),
	}

	return cfg, rcfg, nil
}

// getRemoteConfig returns the remote configuration parameters
func getRemoteConfig(cfg Config, rcfg remoteConfig) (Config, error) {

	if rcfg.isEmpty() {
		return cfg, nil
	}

	viper.Reset()

	viper.SetDefault("app.postgresql.host", cfg.app.postgresql.Host)
	viper.SetDefault("app.postgresql.port", cfg.app.postgresql.Port)
	viper.SetDefault("app.postgresql.address", cfg.app.postgresql.Address)
	viper.SetDefault("app.postgresql.username", cfg.app.postgresql.Username)
	viper.SetDefault("app.postgresql.password", cfg.app.postgresql.Password)

	viper.SetDefault("app.kafka.host", cfg.app.kafka.Host)
	viper.SetDefault("app.kafka.port", cfg.app.kafka.Port)
	viper.SetDefault("app.kafka.address", cfg.app.kafka.Address)

	viper.SetDefault("app.prisma.host", cfg.app.prisma.Host)
	viper.SetDefault("app.prisma.port", cfg.app.prisma.Port)
	viper.SetDefault("app.prisma.address", cfg.app.prisma.Address)
	viper.SetDefault("app.prisma.secret", cfg.app.prisma.Secret)

	// configuration type
	viper.SetConfigType("json")

	var err error

	if rcfg.remoteConfigProvider == "envvar" {
		var data []byte
		data, err = base64.StdEncoding.DecodeString(rcfg.remoteConfigData)
		if err == nil {
			err = viper.ReadConfig(bytes.NewReader(data))
		}
	} else {
		// add remote configuration provider
		if rcfg.remoteConfigSecretKeyring == "" {
			err = viper.AddRemoteProvider(rcfg.remoteConfigProvider, rcfg.remoteConfigEndpoint, rcfg.remoteConfigPath)
		} else {
			err = viper.AddSecureRemoteProvider(rcfg.remoteConfigProvider, rcfg.remoteConfigEndpoint, rcfg.remoteConfigPath, rcfg.remoteConfigSecretKeyring)
		}
		if err == nil {
			// try to read the remote configuration (if any)
			err = viper.ReadRemoteConfig()
		}
	}

	if err != nil {
		return cfg, err
	}

	// read configuration parameters
	return getViper(), nil
}

// getViper reads the config params via Viper
func getViper() Config {
	return Config{
		app: &SkaffoldAppData{
			hydra: &SkaffoldAppHydra{
				Host:         viper.GetString("app.hydra.host"),
				Port:         viper.GetString("app.hydra.port"),
				Address:      viper.GetString("app.hydra.address"),
				AdminHost:    viper.GetString("app.hydra.adminHost"),
				AdminPort:    viper.GetString("app.hydra.adminPort"),
				AdminAddress: viper.GetString("app.hydra.adminAddress"),
			},
			kafka: &SkaffoldAppKafka{
				Host:    viper.GetString("app.kafka.host"),
				Port:    viper.GetString("app.kafka.port"),
				Address: viper.GetString("app.kafka.address"),
			},
			prisma: &SkaffoldAppPrisma{
				Host:    viper.GetString("app.prisma.host"),
				Port:    viper.GetString("app.prisma.port"),
				Address: viper.GetString("app.prisma.address"),
				Secret:  viper.GetString("app.prisma.secret"),
			},
		},
	}
}

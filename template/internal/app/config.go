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

// getLocalConfig returns the local configuration parameters
func getLocalConfig() (cfg Config, rcfg remoteConfig, err error) {

	viper.Reset()

	// set default remote configuration values
	viper.SetDefault("remoteConfigProvider", RemoteConfigProvider)
	viper.SetDefault("remoteConfigEndpoint", RemoteConfigEndpoint)
	viper.SetDefault("remoteConfigPath", RemoteConfigPath)
	viper.SetDefault("remoteConfigSecretKeyring", RemoteConfigSecretKeyring)

	viper.SetDefault("api.system.host", SkaffoldAPIMainHost)
	viper.SetDefault("api.system.port", SkaffoldAPIMainPort)
	viper.SetDefault("api.system.address", SkaffoldAPIMainAddress)

	viper.SetDefault("system.account.host", SkaffoldSystemAccountHost)
	viper.SetDefault("system.account.port", SkaffoldSystemAccountPort)
	viper.SetDefault("system.account.address", SkaffoldSystemAccountAddress)

	viper.SetDefault("system.auth.host", SkaffoldSystemAuthHost)
	viper.SetDefault("system.auth.port", SkaffoldSystemAuthPort)
	viper.SetDefault("system.auth.address", SkaffoldSystemAuthAddress)

	viper.SetDefault("system.media.host", SkaffoldSystemMediaHost)
	viper.SetDefault("system.media.port", SkaffoldSystemMediaPort)
	viper.SetDefault("system.media.address", SkaffoldSystemMediaAddress)

	viper.SetDefault("system.youtube.host", SkaffoldSystemYoutubeHost)
	viper.SetDefault("system.youtube.port", SkaffoldSystemYoutubePort)
	viper.SetDefault("system.youtube.address", SkaffoldSystemYoutubeAddress)

	viper.SetDefault("client.auth.host", SkaffoldClientAuthHost)
	viper.SetDefault("client.auth.port", SkaffoldClientAuthPort)
	viper.SetDefault("client.auth.address", SkaffoldClientAuthAddress)

	viper.SetDefault("client.main.host", SkaffoldClientMainHost)
	viper.SetDefault("client.main.port", SkaffoldClientMainPort)
	viper.SetDefault("client.main.address", SkaffoldClientMainAddress)

	viper.SetDefault("app.hydra.host", SkaffoldAppHydraHost)
	viper.SetDefault("app.hydra.port", SkaffoldAppHydraPort)
	viper.SetDefault("app.hydra.address", SkaffoldAppHydraAddress)
	viper.SetDefault("app.hydra.adminHost", SkaffoldAppHydraAdminHost)
	viper.SetDefault("app.hydra.adminPort", SkaffoldAppHydraAdminPort)
	viper.SetDefault("app.hydra.adminAddress", SkaffoldAppHydraAdminAddress)

	viper.SetDefault("app.postgresql.host", SkaffoldAppPostgresqlHost)
	viper.SetDefault("app.postgresql.port", SkaffoldAppPostgresqlPort)
	viper.SetDefault("app.postgresql.address", SkaffoldAppPostgresqlAddress)
	viper.SetDefault("app.postgresql.username", SkaffoldAppPostgresqlUsername)
	viper.SetDefault("app.postgresql.password", SkaffoldAppPostgresqlPassword)

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
	viper.SetConfigName("global")

	// Find and read the global configuration file (if any)
	err = viper.ReadInConfig()
	if err != nil {
		return cfg, rcfg, err
	}

	// name of the local configuration file without extension
	viper.SetConfigName("local")

	// Find and merge the local configuration file (if any)
	err = viper.MergeInConfig()
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

	// set default configuration values
	viper.SetDefault("api.system.host", cfg.api.main.Host)
	viper.SetDefault("api.system.port", cfg.api.main.Port)
	viper.SetDefault("api.system.address", cfg.api.main.Address)

	viper.SetDefault("system.account.host", cfg.system.account.Host)
	viper.SetDefault("system.account.port", cfg.system.account.Port)
	viper.SetDefault("system.account.address", cfg.system.account.Address)

	viper.SetDefault("system.auth.host", cfg.system.auth.Host)
	viper.SetDefault("system.auth.port", cfg.system.auth.Port)
	viper.SetDefault("system.auth.address", cfg.system.auth.Address)

	viper.SetDefault("system.media.host", cfg.system.media.Host)
	viper.SetDefault("system.media.port", cfg.system.media.Port)
	viper.SetDefault("system.media.address", cfg.system.media.Address)

	viper.SetDefault("system.youtube.host", cfg.system.youtube.Host)
	viper.SetDefault("system.youtube.port", cfg.system.youtube.Port)
	viper.SetDefault("system.youtube.address", cfg.system.youtube.Address)

	viper.SetDefault("client.auth.host", cfg.client.auth.Host)
	viper.SetDefault("client.auth.port", cfg.client.auth.Port)
	viper.SetDefault("client.auth.address", cfg.client.auth.Address)

	viper.SetDefault("client.main.host", cfg.client.main.Host)
	viper.SetDefault("client.main.port", cfg.client.main.Port)
	viper.SetDefault("client.main.address", cfg.client.main.Address)

	viper.SetDefault("app.hydra.host", cfg.app.hydra.Host)
	viper.SetDefault("app.hydra.port", cfg.app.hydra.Port)
	viper.SetDefault("app.hydra.address", cfg.app.hydra.Address)
	viper.SetDefault("app.hydra.adminHost", cfg.app.hydra.AdminHost)
	viper.SetDefault("app.hydra.adminPort", cfg.app.hydra.AdminPort)
	viper.SetDefault("app.hydra.adminAddress", cfg.app.hydra.AdminAddress)

	viper.SetDefault("app.postgresql.host", cfg.app.postgresql.Host)
	viper.SetDefault("app.postgresql.port", cfg.app.postgresql.Port)
	viper.SetDefault("app.postgresql.address", cfg.app.postgresql.Address)
	viper.SetDefault("app.postgresql.username", cfg.app.postgresql.Username)
	viper.SetDefault("app.postgresql.password", cfg.app.postgresql.Password)

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
		api: &SkaffoldAPIData{
			main: &SkaffoldAPIMain{
				Host:    viper.GetString("api.main.host"),
				Port:    viper.GetString("api.main.port"),
				Address: viper.GetString("api.main.address"),
			},
		},
		system: &SkaffoldSystemData{
			account: &SkaffoldSystemAccount{
				Host:    viper.GetString("system.account.host"),
				Port:    viper.GetString("system.account.port"),
				Address: viper.GetString("system.account.address"),
			},
			auth: &SkaffoldSystemAuth{
				Host:    viper.GetString("system.auth.host"),
				Port:    viper.GetString("system.auth.port"),
				Address: viper.GetString("system.auth.address"),
			},
			media: &SkaffoldSystemMedia{
				Host:    viper.GetString("system.media.host"),
				Port:    viper.GetString("system.media.port"),
				Address: viper.GetString("system.media.address"),
			},
			youtube: &SkaffoldSystemYoutube{
				Host:    viper.GetString("system.youtube.host"),
				Port:    viper.GetString("system.youtube.port"),
				Address: viper.GetString("system.youtube.address"),
			},
		},
		client: &SkaffoldClientData{
			auth: &SkaffoldClientAuth{
				Host:    viper.GetString("client.auth.host"),
				Port:    viper.GetString("client.auth.port"),
				Address: viper.GetString("client.auth.address"),
			},
			main: &SkaffoldClientMain{
				Host:    viper.GetString("client.main.host"),
				Port:    viper.GetString("client.main.port"),
				Address: viper.GetString("client.main.address"),
			},
		},
		app: &SkaffoldAppData{
			hydra: &SkaffoldAppHydra{
				Host:         viper.GetString("app.hydra.host"),
				Port:         viper.GetString("app.hydra.port"),
				Address:      viper.GetString("app.hydra.address"),
				AdminHost:    viper.GetString("app.hydra.adminHost"),
				AdminPort:    viper.GetString("app.hydra.adminPort"),
				AdminAddress: viper.GetString("app.hydra.adminAddress"),
			},
			postgresql: &SkaffoldAppPostgresql{
				Host:     viper.GetString("app.postgresql.host"),
				Port:     viper.GetString("app.postgresql.port"),
				Address:  viper.GetString("app.postgresql.address"),
				Username: viper.GetString("app.postgresql.username"),
				Password: viper.GetString("app.postgresql.password"),
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

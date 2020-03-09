package app

// ProgramName defines this application name
const ProgramName = "{{ Project }}"

// ProgramVersion set this application version
// This is supposed to be automatically populated by the Makefile using the value from the VERSION file
// (-ldflags '-X main.ProgramVersion=$(shell cat VERSION)')
var ProgramVersion = "0.0.0"

// ProgramRelease contains this program release number (or build number)
// This is automatically populated by the Makefile using the value from the RELEASE file
// (-ldflags '-X main.ProgramRelease=$(shell cat RELEASE)')
var ProgramRelease = "0"

// ConfigPath list the paths where to look for configuration files (in order)
var ConfigPath = [...]string{
	"./",
	"config/",
	"$HOME/." + ProgramName + "/",
}

// RemoteConfigProvider is the remote configuration source ("consul", "etcd")
const RemoteConfigProvider = ""

// RemoteConfigEndpoint is the remote configuration URL (ip:port)
const RemoteConfigEndpoint = ""

// RemoteConfigPath is the remote configuration path where to search fo the configuration file ("/config/{{ Project }}")
const RemoteConfigPath = ""

// RemoteConfigSecretKeyring is the path to the openpgp secret keyring used to decript the remote configuration data ("/etc/{{ Project }}/configkey.gpg")
const RemoteConfigSecretKeyring = "" // #nosec

// EnvironmentVariablesPrefix prefix to add to the configuration environment variables
const EnvironmentVariablesPrefix = "{{ toUpper Project }}"

// ServerShutdownTimeout timeout in seconds before forcing the server to close
const ServerShutdownTimeout = 10

// ----------

// SkaffoldAppHydraHost is the network host for the public Hydra service.
const SkaffoldAppHydraHost = "app-hydra"

// SkaffoldAppHydraPort is the network port for the public Hydra service.
const SkaffoldAppHydraPort = "60200"

// SkaffoldAppHydraAddress is the network address for the public Hydra service (host:port).
const SkaffoldAppHydraAddress = "app-hydra:60200"

// SkaffoldAppHydraAdminHost is the network host for the admin Hydra service.
const SkaffoldAppHydraAdminHost = "app-hydra"

// SkaffoldAppHydraAdminPort is the network port for the admin Hydra service.
const SkaffoldAppHydraAdminPort = "60201"

// SkaffoldAppHydraAdminAddress is the network address for the admin Hydra service (host:port).
const SkaffoldAppHydraAdminAddress = "app-hydra:60221"

// ----------

// SkaffoldAppKafkaHost is the network host for the public Kafka service.
const SkaffoldAppKafkaHost = "app-kafka"

// SkaffoldAppKafkaPort is the network port for the public Kafka service.
const SkaffoldAppKafkaPort = "60220"

// SkaffoldAppKafkaAddress is the network address for the public Kafka service (host:port).
const SkaffoldAppKafkaAddress = "app-kafka:60220"

// ----------

// SkaffoldAppPrismalHost is the network host for the public Prismal service.
const SkaffoldAppPrismaHost = "app-prisma"

// SkaffoldAppPrismalPort is the network port for the public Prismal service.
const SkaffoldAppPrismaPort = "60220"

// SkaffoldAppPrismalAddress is the network address for the public Prismal service (host:port).
const SkaffoldAppPrismaAddress = "app-prisma:60220"

// SkaffoldAppPrismalSecret is the secret used to authenticate requests for connection to the ORM.
const SkaffoldAppPrismaSecret = "pri"

package app

// SkaffoldAPIData defines all available `api` entities.
type SkaffoldAPIData struct {
	main *SkaffoldAPIMain
}

// SkaffoldAPIMain defines configuration settings for the `main` API.
type SkaffoldAPIMain struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemData defines all available `system` entities.
type SkaffoldSystemData struct {
	account *SkaffoldSystemAccount
	auth    *SkaffoldSystemAuth
	media   *SkaffoldSystemMedia
	youtube *SkaffoldSystemYoutube
}

// SkaffoldSystemAccount defines configuration settings for the `account` system.
type SkaffoldSystemAccount struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemAPI defines configuration settings for the `api` system.
type SkaffoldSystemAPI struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemAuth defines configuration settings for the `auth` system.
type SkaffoldSystemAuth struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemMedia defines configuration settings for the `media` system.
type SkaffoldSystemMedia struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldSystemYoutube defines configuration settings for the `youtube` system.
type SkaffoldSystemYoutube struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldClientData defines all available `client` entities.
type SkaffoldClientData struct {
	auth *SkaffoldClientAuth
	main *SkaffoldClientMain
}

// SkaffoldClientAuth defines configuration settings for the `auth` client.
type SkaffoldClientAuth struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldClientMain defines configuration settings for the `main` client.
type SkaffoldClientMain struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldAppData defines all available `app` entities.
type SkaffoldAppData struct {
	hydra      *SkaffoldAppHydra
	postgresql *SkaffoldAppPostgresql
	prisma     *SkaffoldAppPrisma
}

// SkaffoldAppHydra defines configuration settings for the `hydra` application.
type SkaffoldAppHydra struct {
	Host         string
	Port         string
	Address      string
	AdminHost    string
	AdminPort    string
	AdminAddress string
}

// SkaffoldAppPostgresql defines configuration settings for the `postgresql` application.
type SkaffoldAppPostgresql struct {
	Host     string
	Port     string
	Address  string
	Username string
	Password string
}

// SkaffoldAppPrisma defines configuration settings for the `prisma` application.
type SkaffoldAppPrisma struct {
	Host    string
	Port    string
	Address string
	Secret  string
}

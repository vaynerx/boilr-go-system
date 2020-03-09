package app

import (
	bytes "bytes"
	fmt "fmt"
)

// SkaffoldAPIMain defines configuration settings for the `main` API.
type SkaffoldAPIMain struct {
	Host    string
	Port    string
	Address string
}

// SkaffoldAppData defines all available `app` entities.
type SkaffoldAppData struct {
	hydra  *SkaffoldAppHydra
	prisma *SkaffoldAppPrisma
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

// SkaffoldAppPrisma defines configuration settings for the `prisma` application.
type SkaffoldAppPrisma struct {
	Host    string
	Port    string
	Address string
	Secret  string
}

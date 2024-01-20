package models

import "github.com/niko0xdev/gqlgen/integration/server/remote_api"

type Viewer struct {
	User *remote_api.User
}

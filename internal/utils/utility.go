package utils

import (
	"os"

	"github.com/ritik-patel05/image-player/internal/constants"
)

func IsProductionEnv() bool {
	return os.Getenv(constants.ACTIVE_ENV) == constants.PRODUCTION
}

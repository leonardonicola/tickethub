package utils

import "github.com/leonardonicola/tickethub/config"

type UploadedFile struct {
	Identifier string
}

var (
	logger = config.NewLogger()
)

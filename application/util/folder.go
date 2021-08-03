package util

import (
	"os"
	"path/filepath"
)

// GetApplicationRootPath -
func GetApplicationRootPath() string {
	if os.Getenv("ENVIROMENT") == "developer" {
		return GetApplicationRootPathWeb()
	}

	return GetApplicationRootPathLinux()
}

// ObterCaminhoDiretorioUpload -
func ObterCaminhoDiretorioUpload() string {
	return filepath.Join(GetApplicationRootPath(), "upload")
}

// GetApplicationRootPath -
func GetApplicationRootPathWeb() string {
	pathRoot, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return pathRoot
}

// GetApplicationRootPathLinux
func GetApplicationRootPathLinux() string {
	pathRoot, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	return pathRoot
}

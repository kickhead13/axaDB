package axa_security

import (
	"os"
)

func GetEnvironmentVarValue(varName string) string{
	return os.Getenv(varName)
}
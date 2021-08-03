package tests

import (
	"os"
	"testing"
)

const erroPadrao = "O valor esperado %s, mas o resultado encontrado foi %s."

func TestCarregarConfiguracoes(t *testing.T) {
	t.Parallel() // Teste em paralelo

	if os.Getenv("ENVIROMENT") != "developer" && os.Getenv("ENVIROMENT") != "production" {
		t.Errorf(erroPadrao, "developer ou production", os.Getenv("ENVIROMENT"))
	}
}

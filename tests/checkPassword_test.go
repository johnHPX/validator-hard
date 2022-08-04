package tests

import (
	"testing"

	"github.com/johnHPX/validator-hard/pkg/validator"
)

func TestCheckPassword(t *testing.T) {
	// teste positivo
	t.Run("Criando Uma Senha", func(t *testing.T) {
		validator := validator.NewValidator()

		password := "12345678"
		sizePassword := 255

		hashPassword, err := validator.CheckPassword(sizePassword, password, "", "create")
		if err != nil {
			t.Errorf("Erro ao criar Senha! %e", err)
		}

		t.Logf("\n\n\n %s \n\n\n", hashPassword)

	})

	t.Run("Validando uma Senha", func(t *testing.T) {
		validator := validator.NewValidator()

		password := "12345678"
		sizePassword := 255
		hashPassword := "$2a$10$59NbGfNDLbTg6Ds.FkdeWOMx6g9aQ.0mQ7B1dMooGJkttNldQ771y"

		_, err := validator.CheckPassword(sizePassword, password, hashPassword, "compare")
		if err != nil {
			t.Errorf("senha incorreta! %e", err)
		}
	})

	// teste negativo
	t.Run("Criando Uma Senha", func(t *testing.T) {
		validator := validator.NewValidator()

		password := ""
		sizePassword := 10

		hashPassword, err := validator.CheckPassword(sizePassword, password, "", "create")
		if err != nil {
			t.Errorf("Erro ao criar Senha! %e", err)
		}

		t.Logf("\n\n\n %s \n\n\n", hashPassword)

	})

	t.Run("Validando uma Senha", func(t *testing.T) {
		validator := validator.NewValidator()

		password := "12345678asldjk"
		sizePassword := 255
		hashPassword := "$2a$10$59NbGfNDLbTg6Ds.FkdeWOMx6g9aQ.0mQ7B1dMooGJkttNldQ771y"

		_, err := validator.CheckPassword(sizePassword, password, hashPassword, "compare")
		if err != nil {
			t.Errorf("senha incorreta! %e", err)
		}
	})
}

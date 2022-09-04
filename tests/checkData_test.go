package tests

import (
	"testing"

	"github.com/johnHPX/validator-hard/pkg/validator"
)

func TestCheckData(t *testing.T) {
	// teste positivo
	t.Run("Validar um Valor correto", func(t *testing.T) {
		nome := "steve"
		sizeNome := 255
		email := "steve@gmail.com"
		sizeEmail := 255
		cpf := "61868307301"
		sizeCpf := 11

		idade := 15
		sizeIdade := 200

		validator := validator.NewValidator()

		_, err := validator.CheckAnyData("nome", sizeNome, nome, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}
		_, err = validator.CheckAnyData("email", sizeEmail, email, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}

		_, err = validator.CheckAnyData("cpf", sizeCpf, cpf, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}

		_, err = validator.CheckAnyData("idade", sizeIdade, idade, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}

	})
	//teste negativo
	t.Run("Validar um Valor negativo", func(t *testing.T) {
		nome := "steve"
		sizeNome := 3
		email := "asas"
		sizeEmail := 255
		cpf := "87667898701"
		sizeCpf := 11

		idade := 0
		sizeIdade := 3

		casada := false // casada?

		validator := validator.NewValidator()

		_, err := validator.CheckAnyData("nome", sizeNome, nome, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}
		_, err = validator.CheckAnyData("email", sizeEmail, email, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}

		_, err = validator.CheckAnyData("cpf", sizeCpf, cpf, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}

		_, err = validator.CheckAnyData("idade", sizeIdade, idade, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}

		_, err = validator.CheckAnyData("casada", 0, casada, true)
		if err != nil {
			t.Errorf("Erro ao verificar campo! %s", err)
		}

	})
}

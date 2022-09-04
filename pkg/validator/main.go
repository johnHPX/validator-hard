package validator

import (
	"reflect"

	"github.com/johnHPX/validator-hard/internal/infra"
)

type Validator interface {
	// CheckData validates any value from a request and if the value is a normal string, the function returns a new value.
	CheckAnyData(nameData string, sizeData int, data interface{}, required bool) (interface{}, error)
	// CheckPassword create a password by generating a hash and compare a password saved in the database with a password passed in the request. create vs compare. create: generate an encrypted password and return it. compare: compare passwordData with hashPassword, check if they are the same password
	CheckPassword(sizePassword int, passawordData, hashPassword, operation string) (string, error)
}

type validator struct{}

// CheckData validates any value from a request and if the value is a normal string, the function returns a new value.
func (s *validator) CheckAnyData(nameData string, sizeData int, data interface{}, required bool) (interface{}, error) {

	if reflect.ValueOf(data).Kind() == reflect.String {
		dataFormated, err := infra.FormatedInput(data.(string))
		if err != nil {
			return nil, err
		}

		if required {
			err = infra.CheckIsEmpty(dataFormated, nameData)
			if err != nil {
				return nil, err
			}
		}

		err = infra.CheckLen(sizeData, dataFormated)
		if err != nil {
			return nil, err
		}

		switch nameData {
		case "email":
			err = infra.CheckEmail(dataFormated)
			if err != nil {
				return nil, err
			}
		case "cpf":
			err = infra.CheckCPF(dataFormated)
			if err != nil {
				return nil, err
			}
		case "cpnj":
			err = infra.CheckCNPJ(dataFormated)
			if err != nil {
				return nil, err
			}
		case "cep":
			err = infra.CheckCEP(dataFormated)
			if err != nil {
				return nil, err
			}
		}

		return dataFormated, nil
	} else if reflect.ValueOf(data).Kind() == reflect.Int {
		if required {
			err := infra.CheckIsEmpty(data, nameData)
			if err != nil {
				return nil, err
			}
		}

		err := infra.CheckLen(sizeData, data)
		if err != nil {
			return nil, err
		}

		return data, nil
	} else if reflect.ValueOf(data).Kind() == reflect.Bool {
		if required {
			err := infra.CheckIsEmpty(data, nameData)
			if err != nil {
				return nil, err
			}
		}

		return data, nil
	}

	return nil, nil

}

// CheckPassword create a password by generating a hash and compare a password saved in the database with a password passed in the request. create vs compare. create: generate an encrypted password and return it. compare: compare passwordData with hashPassword, check if they are the same password
func (s *validator) CheckPassword(sizePassword int, passawordData, hashPassword, operation string) (string, error) {
	if operation == "create" {
		dataNew, err := infra.FormatedInput(passawordData)
		if err != nil {
			return "", err
		}

		err = infra.CheckIsEmpty(dataNew, "password")
		if err != nil {
			return "", err
		}

		err = infra.CheckLen(sizePassword, dataNew)
		if err != nil {
			return "", err
		}

		hash, err := infra.CheckPassword(dataNew)
		if err != nil {
			return "", err
		}

		return hash, nil

	} else if operation == "compare" {
		err := infra.CheckIsEmpty(passawordData, "password")
		if err != nil {
			return "", err
		}

		err = infra.CheckLen(sizePassword, passawordData)
		if err != nil {
			return "", err
		}

		err = infra.CheckHash(hashPassword, passawordData)
		if err != nil {
			return "", err
		}

		return "", nil
	}

	return "", nil
}

func NewValidator() Validator {
	return &validator{}
}

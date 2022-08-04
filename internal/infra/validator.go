package infra

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/paemuri/brdoc"
	"golang.org/x/crypto/bcrypt"
)

func CheckIsEmpty(input interface{}, nameInput string) error {

	if reflect.ValueOf(input).Kind() == reflect.String {
		str, ok := input.(string)
		s := strings.TrimSpace(str)
		if !ok || s == "" {
			return errors.New(fmt.Sprintf("%v é obrigatório e não pode está em branco", nameInput))
		}

		return nil
	} else if reflect.ValueOf(input).Kind() == reflect.Int {
		_, ok := input.(int)
		if !ok {
			return errors.New(fmt.Sprintf("%v é obrigatório e não pode está em branco", nameInput))
		}

		return nil
	} else if reflect.ValueOf(input).Kind() == reflect.Bool {
		_, ok := input.(bool)
		if !ok {
			return errors.New(fmt.Sprintf("%v é obrigatório e não pode está em branco", nameInput))
		}

		return nil
	}

	return errors.New("Tipo não Permitido!")
}
func CheckLen(size int, input interface{}) error {
	if size < len(input.(string)) {
		return errors.New("O tamanho excede o permitido")
	}
	return nil
}
func CheckEmail(email string) error {
	if err := checkmail.ValidateFormat(email); err != nil {
		return errors.New("O e-mail inserido é invalido. Para o valor ser valido é preciso está nesse padrão `exemplo@gmail.com` ")
	}
	return nil
}
func CheckPassword(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashByte), err
}
func CheckHash(pswWithHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(pswWithHash), []byte(password))
}

func CheckCPF(cpf string) error {
	okCpf := brdoc.IsCPF(cpf)
	if okCpf {
		return nil
	}

	return errors.New("CPF invalido. Para ser valido é preciso está nesse padrão `XXX.XXX.XXX-XX`")
}

func CheckCNPJ(cnpj string) error {
	okCnpj := brdoc.IsCNPJ(cnpj)
	if okCnpj {
		return nil
	}
	return errors.New("CNPJ invalido. Para ser valido é preciso está nesse padrão `XX.XXX.XXX/0001-XX`")
}

func CheckCEP(cep string) error {
	okCnpj := brdoc.IsCEP(cep)
	if okCnpj {
		return nil
	}
	return errors.New("CEP invalido. Para ser valido é preciso está nesse padrão `XXXXX-XXX`")
}

func FormatedInput(input string) (string, error) {
	return strings.TrimSpace(input), nil
}

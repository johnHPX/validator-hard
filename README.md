# Validator
Validator is a package for validating API requests.

## Como importar

Como é um módulo público basta apenas adicionar ao go.mod com "go get".

```
  go get github.com/johnHPX/validator-hard
```

Depois, é só rodar o comando go mod tidy para fazer o download das dependencias.

```
  go mod tidy
```

E por fim, importar o módulo e utilizar.

```
  import (
      "github.com/johnHPX/validator-hard/pkg/validator"
  )
  

  func main(){
    validator := validator.NewValidator()
  }
```

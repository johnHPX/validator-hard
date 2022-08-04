# validator-hard
validator is a package for validating API requests.

modo de usar

como é um módulo privado, é necessário especificar para o Go que é privado.

'''
  export GOPRIVATE=github.com/johnHPX/validator-hard
'''

depois apenas adicionar ao módulo local com o go get

'''
  go get github.com/johnHPX/validator-hard
'''

ademais, rodar o comando go mod tidy para fazer o download das dependencias.

'''
  go mod tidy
'''

e por fim, importar o módulo e utilizar.

'''
  import (
      "github.com/johnHPX/validator-hard/pkg/validator"
  )
  
  
  func main(){
    validator := validator.NewValidator()
  }
''

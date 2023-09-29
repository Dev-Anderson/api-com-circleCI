## 💻 **Sobre o projeto**

O projeto foi desenvolvido pensando em utilizar o Github Actions, foi adicionado o seguinte código para rodar o github actions. 
Dentro da pasta raiz deverá adicionar a seguinte configuração
Adicionar um pasta com o nome ".github/workflows", dentro da pasta adicionar o arquivo com o nome "go.yml", dentro do arquivo o seguinte código:   

     name: Go
    
    on:
      push:
        branches: [ "develop" ]
      pull_request:
        branches: [ "develop" ]
    
    jobs:
    
      build:
        runs-on: ubuntu-latest
        steps:
        - uses: actions/checkout@v3
    
        - name: Set up Go
          uses: actions/setup-go@v4
          with:
            go-version: '1.21'
    
        - name: Build
          run: go build -v ./...
    
        - name: Test
          run: go test -v ./...

Em todos os commits na branch "develop" deverá rodar fazer o build da aplicação e depois do teste da aplicação. 
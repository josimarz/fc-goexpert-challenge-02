# Go Expert

## Desafio Multithreading

Para executar o programa, através do terminal acesse o diretório raiz deste projeto e execute o comando a seguir:

```sh
go run cmd/cep/main.go {CEP}
```

O termo `{CEP}` deve ser substituído pelo CEP que será pesquisado. Exemplo:

```sh
go run cmd/cep/main.go 89023-600
```

Se a consulta for bem sucedida, os dados da resposta serão exibidas no terminal no seguinte formato:

```json
{
  "Provider": "provider",
  "Response": "object"
}
```

O campo `Provider` exibirá o nome do provedor que entregou a resposta primeiro. Os possíveis valores são:
* Via CEP
* Api CEP

O campo `Response` exibirá os dados da resposta do provedor. Quando o provedor for "Via CEP", o campo `Response` terá o seguinte formato:

```json
{
    "cep": "string",
    "logradouro": "string",
    "complemento": "string",
    "bairro": "string",
    "localidade": "string",
    "uf": "string",
    "ibge": "string",
    "gia": "string",
    "ddd": "string",
    "siafi": "string"
}
```

No caso de uma resposta do provedor "Api CEP", o campo `Response` será composto pela seguinte estrutura:

```json
{
    "code": "string",
    "state": "string",
    "city": "string",
    "district": "string",
    "address": "string",
    "status": "number",
    "ok": "boolean",
    "statusText": "string"
}
```

Em caso de *timeout*, será exibida a mensagem "timeout" no terminal.
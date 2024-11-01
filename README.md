# Oração Bandas
## Projeto usado no grupo homônimo ao projeto
#Ferramentas:
- Go
- Postgres
- Arquitetura modular
- Tailwind

# Variáveis de ambientes
É necessário definir a variável de ambiente `PROJECT_ENV`. 

## PROJECT_ENV == development

Caso esteja em um ambiente de desenvolvimento, a aplicação espera que haja um arquivo `env.json`.

## PROJECT_ENV == production

Em caso de ambiente de produção, se espera a definicão das variáveis no host.

## Chaves das variáveis
As chaves esperadas são as seguintes:

### Desenvolvimento:
No arquivo `env.example.json` haverá uma estrutura de exemplo para as variáveis. De qualquer forma, as variáveis esperadas são:
- Database
  - host
  - user
  - password
  - dbname
  - db_port
  - ssl
- System
  - system_port

 ### Produção
 - DB_HOST
 - DB_USER
 - DB_PASSWORD
 - DB_NAME
 - DB_PORT
 - DB_SSL
 - SYSTEM_PORT

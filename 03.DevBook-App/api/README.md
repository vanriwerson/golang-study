## Rodando Banco de Dados pelo Docker
> Essa API se conecta com um banco de dados mySQL, cujo serviço está configurado no
> arquivo `docker-compose.yml` e pode ser inicializado com o comando `docker-compose up`.

  - Após iniciar o serviço, deve-se rodar o script `sql/migrations.sql` no seu SGBD para
  criar o esquema e as tabelas necessárias à API.

# APRENDENDO GOLANG
> Atividades realizadas ao longo do curso "Aprenda Golang do Zero!
> Desenvolva uma APLICAÇÃO COMPLETA!" ministrado pelo professor 
> Otávio Augusto Gallego e disponível na Udemy.

## [Instalando Golang](https://go.dev/doc/install) no Linux

1. Faça o [download](https://go.dev/dl/) do instalador de acordo com o seu sistema operacional.
    - DICA: prefira uma versão LTS (Long-term Support), como [esta](https://go.dev/dl/go1.18.4.linux-amd64.tar.gz). 
2. Remova instalações anteriores do Go (se houver) de sua máquina:
    - `sudo rm -rf /usr/local/go`. 
3. Extraia o arquivo de instalação para a pasta `/usr/local/`.
    - `sudo tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz`
4. Crie uma variável de ambiente PATH com o valor `/usr/local/go/bin`
    - Obs1.: Para isso, adicione a seguinte linha de código ao arquivo `$HOME/.profile`: `export PATH=$PATH:/usr/local/go/bin`
    - Obs2.: Essa alteração só será aplicada após reiniciar o computador.
5. Verifique se a instalação foi feita com sucesso:
    - `go version` deve retornar `go version go1.18.4 linux/amd64`
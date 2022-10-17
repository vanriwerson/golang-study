# APRENDENDO GOLANG
> Atividades realizadas ao longo do curso "Aprenda Golang do Zero!
> Desenvolva uma APLICAÇÃO COMPLETA!" ministrado pelo professor 
> Otávio Augusto Gallego e disponível na Udemy, complementadas por
> minhas pesquisas e rotinas autodidatas.

## [Instalando Golang](https://go.dev/doc/install) no Linux

1. Faça o [download](https://go.dev/dl/) do instalador de acordo com o seu sistema operacional.
    - DICA: prefira uma versão LTS (Long-term Support), como [esta](https://go.dev/dl/go1.18.4.linux-amd64.tar.gz). 

2. Remova instalações anteriores da Golang (se houver) de sua máquina:
    - `sudo rm -rf /usr/local/go`.

3. Extraia o arquivo de instalação para a pasta `/usr/local/`.
    - `sudo tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz`

4. Crie uma variável de ambiente PATH com o valor `/usr/local/go/bin`
    - Obs1.: Para isso, adicione a seguinte linha de código ao arquivo `$HOME/.profile`: `export PATH=$PATH:/usr/local/go/bin`
    - Obs2.: Se você utiliza o terminal Oh My Zsh, também deve adicionar a variável PATH ao arquivo .zshrc.
    - Obs3.: Essa alteração só será aplicada após reiniciar o computador.

5. Verifique se a instalação foi feita com sucesso:
    - O comando `go version` deve retornar algo parecido com `go version go1.18.4 linux/amd64`

## 01.Fundamentals

> Guarda exemplos basilares da linguagem Go. O principal objetivo aqui é saber
> quais estruturas de dados existem e qual a sintaxe utilizada para implementá-las.

## 02.Aplied-Fundamentals
> Guarda exemplos da aplicabilidade da Golang no Desenvolvimento Web, integração com 
> Banco de Dados mySQL e um CRUD (Create, Read, Update, Delete) simples.
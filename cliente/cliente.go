package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	arguments := os.Args     // Pega os argumentos da linha de comando
	if len(arguments) == 1 { // Se não inserir o valor da porta, exibe a mensagem abaixo
		fmt.Print("[!] Insira o endereço e a porta primeiro. ex: localhost:8080")
		return
	}

	// Usa os argumentos e se conecta ao servidor host:porta
	c, err := net.Dial("tcp", arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	fmt.Print("[!] Bem vindo(a)! Pressione ENTER para acessar o Menu \n")
	for {
		reader := bufio.NewReader(os.Stdin) // Prepara o buffer de leitura
		fmt.Print("--> ")
		text, _ := reader.ReadString('\n') // Le o texto digitado
		fmt.Fprintf(c, text+"\n")          // Faz o envio do texto pro servidor

		message, _ := bufio.NewReader(c).ReadString('\n') // Aguarda resposta do servidor
		fmt.Print("Servidor >> " + message)
		// Se a resposta for SAIR, fecha a conexão e o cliente
		if strings.ToUpper(strings.TrimSpace(string(text))) == "SAIR" {
			fmt.Println("[!] Sessão finalizada")
			return
		}
	}
}

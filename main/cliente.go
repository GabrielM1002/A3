package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// Usa os argumentos e se conecta ao servidor host:port
	c, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	fmt.Print("Escreva sua mensagem.")
	fmt.Println(" Digite sair para cancelar")

	for {

		// Prepara o buffer de leitura
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")

		// Le um texto do teclado
		text, _ := reader.ReadString('\n')

		// Envia o texto pela conexão
		fmt.Fprintf(c, text+"\n")

		// Aguarda resposta do servidor
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print(message)

		// Se a resposta for EXIT, fecha a conexão e o cliente
		if strings.ToUpper(strings.TrimSpace(string(text))) == "SAIR" {
			fmt.Println("Sessão finalizada")
			return
		}
	}
}

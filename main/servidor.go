package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	fmt.Println("[Servidor Online]")
	l, err := net.Listen("tcp", "localhost:8080") // Inicia a conexão TCP na porta
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close() // Ao final, fecha a conexão.
	fmt.Println("Aguardando conexões...")

	c, err := l.Accept() // Fica aguardando um cliente se conectar
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Cliente conectado ao servidor!")

	for { // Laço eterno
		// Recebe informações no buffer de leitura
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		// Se for EXIT, fecha o sevidor, mas fecha a conexão antes
		if strings.ToUpper(strings.TrimSpace(string(netData))) == "SAIR" {
			fmt.Println("Cliente finalizou sua sessão.")
			return
		}

		// Mostra a mensagem na tela e envia de volta o horário
		fmt.Print("Cliente > ", string(netData))
		t := time.Now()
		myTime := "Recebido em: " + t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime)) // Escreve no buffer de escrita para o cliente
	}
}

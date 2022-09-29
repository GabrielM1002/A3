package main

import (
	"bufio"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Função main
func main() {
	fmt.Println("[Servidor Online]")
	l, err := net.Listen("tcp", "localhost:8080") // Inicia a conexão TCP na porta
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ao final, fecha a conexão.
	defer l.Close()

	fmt.Println("Aguardando conexões...")

	// Fica aguardando um cliente se conectar
	c, err := l.Accept()
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

		// Se for SAIR, fecha o sevidor, mas fecha a conexão antes
		if strings.ToUpper(strings.TrimSpace(string(netData))) == "SAIR" {
			fmt.Println("Cliente finalizou sua sessão.")
			return
		}

		// Converter o texto em int (netData em intVar)
		strVar := strings.ToUpper(strings.TrimSpace(string(netData)))
		intVar, err := strconv.Atoi(strVar)
		fmt.Println(intVar, err, reflect.TypeOf(intVar))

		// Mostra a mensagem na tela e envia de volta o horário
		fmt.Print("Cliente > ", string(strVar))
		t := time.Now()
		myTime := "Recebido em: " + t.Format(time.RFC3339) + "\n"

		// Escreve no buffer de escrita para o cliente
		c.Write([]byte(myTime))
	}
}

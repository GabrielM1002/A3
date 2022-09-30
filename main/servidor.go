package main

import (
	"bufio"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
)

// Função main
func main() {

	var partCount int = 0

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

	// Laço eterno
	// Laço eterno
	// Laço eterno
	// Laço eterno
	for {

		// Recebe informações no buffer de leitura
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// Limpa o netData
		strVar := strings.ToUpper(strings.TrimSpace(string(netData)))

		// SAIR: fecha o sevidor, mas fecha a conexão antes
		if strVar == "SAIR" {
			fmt.Println("Cliente finalizou sua sessão.")
			return
		}

		// VOLTAR: volta para a frase inicial
		if strVar == "VOLTAR" {
			partCount = 0
		}

		// Converter mensagem do cliente em int
		intVar, err := strconv.Atoi(strVar)
		fmt.Println(intVar, err, reflect.TypeOf(intVar))

		//partCounts
		//partCounts
		//partCounts
		//partCounts
		// Contadores que simulam um switch
		// Parte1: MULTIPLICA POR 2
		if partCount == 1 {

			// Conta
			result := intVar * 2

			// Converte resultado para string
			s2 := strconv.Itoa(result)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite VOLTAR para voltar.\n"))

			// Parte2: DIVIDE POR 2
		} else if partCount == 2 {

			// Conta
			result := intVar / 2

			// Converte resultado para string
			s1 := strconv.FormatInt(int64(result), 10)
			s2 := strconv.Itoa(result)
			fmt.Printf("%v, %v\n", s1, s2)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite VOLTAR para voltar.\n"))

			// Parte3: MULTIPLICA POR ELE MESMO
		} else if partCount == 3 {

			// Conta
			result := intVar * intVar

			// Converte resultado para string
			s1 := strconv.FormatInt(int64(result), 10)
			s2 := strconv.Itoa(result)
			fmt.Printf("%v, %v\n", s1, s2)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite VOLTAR para voltar.\n"))

			// Parte0: Escolhe o que fazer
		} else if partCount == 0 {

			if strVar == "1" {

				message := "Você escolheu MULTIPLICAR POR 2. Digite o número a ser multiplicado por dois:  \n"
				partCount = 1
				c.Write([]byte(message))

			} else if strVar == "2" {

				message := "Você escolheu DIVIDIR POR 2. Digite o número a ser dividido por dois:  \n"
				partCount = 2
				c.Write([]byte(message))

			} else if strVar == "3" {

				message := "Você escolheu MULTIPLICAR POR ELE MESMO. Digite o número a ser multiplicado por ele mesmo:  \n"
				partCount = 3
				c.Write([]byte(message))

			} else {

				message := "Digite 1 para MULTIPLICAR POR 2; Digite 2 para DIVIDIR POR 2; Digite 3 para MULTIPLICAR POR ELE MESMO." + "\n"
				c.Write([]byte(message))
			}

			// Parte0: Printa frase inicial com as opções
		} else if partCount == 0 {

			// Printa mensagem no servidor
			fmt.Print("Cliente > ", string(strVar))

			// Mensagem para o cliente
			start := "Digite 1 para MULTIPLICAR POR 2; Digite 2 para DIVIDIR POR 2; Digite 3 para MULTIPLICAR POR ELE MESMO." + "\n"
			c.Write([]byte(start))
		}

	}
}

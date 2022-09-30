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

	for { // Laço eterno

		// Recebe informações no buffer de leitura
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// Limpa o netData
		strVar := strings.ToUpper(strings.TrimSpace(string(netData)))

		// Se for SAIR, fecha o sevidor, mas fecha a conexão antes
		if strVar == "SAIR" {
			fmt.Println("Cliente finalizou sua sessão.")
			return
		}

		if strVar == "VOLTAR" {
			partCount = 0
		}

		// Converter o texto em int (strVar em intVar)
		intVar, err := strconv.Atoi(strVar)
		fmt.Println(intVar, err, reflect.TypeOf(intVar))

		if partCount == 1 {

			// Conta
			result := intVar * 2

			// Converte para string
			s1 := strconv.FormatInt(int64(result), 10)
			s2 := strconv.Itoa(result)
			fmt.Printf("%v, %v\n", s1, s2)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite VOLTAR para voltar.\n"))

		} else if partCount == 2 {

			// Conta
			result := intVar / 2

			// Converte para string
			s1 := strconv.FormatInt(int64(result), 10)
			s2 := strconv.Itoa(result)
			fmt.Printf("%v, %v\n", s1, s2)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite VOLTAR para voltar.\n"))

		} else if partCount == 3 {

			// Conta
			result := intVar * intVar

			// Converte para string
			s1 := strconv.FormatInt(int64(result), 10)
			s2 := strconv.Itoa(result)
			fmt.Printf("%v, %v\n", s1, s2)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite VOLTAR para voltar.\n"))

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
		}

		// Mostra a mensagem na tela e envia de volta o horário
		fmt.Print("Cliente > ", string(strVar))
		myTime := "Digite 1 para MULTIPLICAR POR 2; Digite 2 para DIVIDIR POR 2; Digite 3 para MULTIPLICAR POR ELE MESMO." + "\n"

		// Escreve no buffer de escrita para o cliente
		c.Write([]byte(myTime))
	}
}

func divideByTwo(num1 int) (int, error) {
	return num1 / 2, nil
}

func multiplyByTwo(num1 int) (int, error) {
	return num1 * 2, nil
}

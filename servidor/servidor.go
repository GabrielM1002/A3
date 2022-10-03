package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func filho(c net.Conn) {
	var (
		num1, num2, num3, num4, partCount int = 0, 0, 0, 0, 0 //instanciando as variaveis matemáticas
	)
	addr := c.RemoteAddr()                                       // Guarda o endereço do cliente
	fmt.Println("[+] Usuário", addr, " conectado com sucesso! ") //Informa o ID do cliente conectado no servidor
	defer fmt.Println("[-] Usuário", addr, " desconectado.")     //Informa o ID do cliente desconectado no servidor
	defer c.Close()                                              // Finaliza conexão no fim

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
			fmt.Println("Servidor finalizado")
			return
		}

		// MENU: volta para o menu
		if strVar == "MENU" {
			partCount = 0
		}

		// Converter mensagem do cliente em int
		intVar, err := strconv.Atoi(strVar)
		fmt.Println(intVar, err, reflect.TypeOf(intVar))

		// partCounts
		// Contadores que simulam um switch
		// Parte1: MULTIPLICA POR 2
		if partCount == 1 {

			// Conta
			result := intVar * 2

			// Converte resultado para string
			s2 := strconv.Itoa(result)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite MENU para voltar.\n"))

			// Parte2: DIVIDE POR 2
		} else if partCount == 2 {

			// Conta
			result := intVar / 2

			// Converte resultado para string 
			s1 := strconv.FormatInt(int64(result), 10)
			s2 := strconv.Itoa(result)
			fmt.Printf("%v, %v\n", s1, s2)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite MENU para voltar.\n"))

			// Parte3: MULTIPLICA POR ELE MESMO
		} else if partCount == 3 {

			// Conta
			result := intVar * intVar

			// Converte resultado para string
			s1 := strconv.FormatInt(int64(result), 10)
			s2 := strconv.Itoa(result)
			fmt.Printf("%v, %v\n", s1, s2)

			// Printa no cliente
			c.Write([]byte(" Resultado: " + s2 + ". Digite MENU para voltar.\n"))

			// Parte4: Digite o segundo número
		} else if partCount == 4 {

			// salva o numero
			num1 = intVar

			// Printa no cliente
			c.Write([]byte(" Digite o Segundo número:  \n"))
			partCount = 5

			// Parte0: Escolhe o que fazer
		} else if partCount == 5 {

			// salva o numero
			num2 = intVar

			// Printa no cliente
			c.Write([]byte(" Digite o Terceiro número:  \n"))

			partCount = 6

		} else if partCount == 6 {

			// salva o numero
			num3 = intVar

			// Conta
			num4 = (num1 + num2 + num3) / 3

			// Converte para string
			s1 := strconv.FormatInt(int64(num4), 10)
			s2 := strconv.Itoa(num4)
			fmt.Printf("%v, %v\n", s1, s2)

			// Printa no cliente
			c.Write([]byte(" A Média entre os números é: " + s1 + "\n"))
			partCount = 0

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

			} else if strVar == "4" {

				message := "Você escolheu MÉDIA ENTRE TRÊS NÚMEROS. Digite o primeiro número:  \n"
				c.Write([]byte(message))
				partCount = 4

			} else {

				message := "Digite 1 para MULTIPLICAR POR 2; Digite 2 para DIVIDIR POR 2; Digite 3 para MULTIPLICAR POR ELE MESMO; Digite 4 para MÉDIA DE 3 NÚMEROS. Ou SAIR para finalizar." + "\n"
				c.Write([]byte(message))
			}

			// Parte0: Printa frase inicial com as opções, caso nenhuma tenha sido escolhida.
		} else if partCount == 0 {

			// Printa mensagem no servidor
			fmt.Print("Cliente > ", string(strVar))

			// Mensagem para o cliente
			start := "Digite 1 para MULTIPLICAR POR 2; Digite 2 para DIVIDIR POR 2; Digite 3 para MULTIPLICAR POR ELE MESMO." + "\n"
			c.Write([]byte(start))
		}

	}
}

// Função main
func main() {
	arguments := os.Args     // Pega a porta como argumento da linha de comando
	if len(arguments) == 1 { // Se não for passado a porta, envia a mensagem abaixo:
		fmt.Println("[!] Digite o número da porta para inicializar.")
		return
	}
	l, err := net.Listen("tcp", ":"+arguments[1]) // Configura uma porta TCP
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close() // Finaliza a conexão no final
	fmt.Println("[!] Servidor inicializado na porta:", arguments[1])

	for {
		c, err := l.Accept() // Aguarda conexão de outro cliente simultâneo
		if err != nil {
			fmt.Println(err)
			return
		}
		go filho(c) // Passa o gerenciamento da conexão para a gorotutine "filho"
	}
}

package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Chatbot iniciado.")
	fmt.Printf("%d documentos carregados.\n", len(documents))
	fmt.Println("Digite 'sair' para encerrar.")

	for {
		fmt.Print("\nVocê: ")

		if !scanner.Scan() {
			break
		}

		question := strings.TrimSpace(scanner.Text())

		if question == "" {
			continue
		}

		if shouldExit(question) {
			fmt.Println("Chatbot encerrado.")
			break
		}

		response, hasResponse := AgentResponse(question)

		if !hasResponse {
			fmt.Println("Bot: Não encontrei informações relacionadas à pergunta.")
			continue
		}

		fmt.Printf("Bot: %s\n", response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
	}
}

func shouldExit(input string) bool {
	input = strings.ToLower(strings.TrimSpace(input))

	return input == "sair" ||
		input == "exit" ||
		input == "quit" ||
		input == "encerrar"
}

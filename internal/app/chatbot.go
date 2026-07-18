package app

import (
	"ai-agent/internal/answer"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/memory"
	"ai-agent/internal/search"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	documents, err := knowledge.LoadDocuments(knowledgeBasePath)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	engine := search.NewEngine(documents, minimumSimilarity)

	session := memory.NewSession()
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

		searchResult := engine.Search(question, session, maximumSearchResults)

		if !searchResult.Found {
			fmt.Println("Bot: Não encontrei informações relacionadas à pergunta.")
			continue
		}

		plan := answer.BuildPlan(searchResult.Tokens, searchResult.Intent, searchResult.Results)

		template := answer.SelectTemplateForPlan(plan, session)

		response := answer.RenderTemplate(template, plan)

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

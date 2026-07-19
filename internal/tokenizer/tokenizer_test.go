package tokenizer

import "testing"

func TestTokenizeForAnalysisKeepsQuestionAndPronounSignals(t *testing.T) {
	tokens := TokenizeForAnalysis("Qual é o email dele?")
	expected := []string{"qual", "é", "o", "email", "dele"}

	if len(tokens) != len(expected) {
		t.Fatalf("TokenizeForAnalysis() returned %v, want %v", tokens, expected)
	}

	for index, token := range tokens {
		if token != expected[index] {
			t.Fatalf("TokenizeForAnalysis()[%d] = %q, want %q", index, token, expected[index])
		}
	}
}

func TestTokenizeKeepsSearchTokensFiltered(t *testing.T) {
	tokens := Tokenize("Qual é o email dele?")
	expected := []string{"email"}

	if len(tokens) != len(expected) {
		t.Fatalf("Tokenize() returned %v, want %v", tokens, expected)
	}

	for index, token := range tokens {
		if token != expected[index] {
			t.Fatalf("Tokenize()[%d] = %q, want %q", index, token, expected[index])
		}
	}
}

package answer

import (
	"ai-agent/internal/nlp"
	"strings"
	"testing"
)

func TestEducationTemplatesRenderCompleteFactsDirectly(t *testing.T) {
	plan := Plan{
		Intent:         nlp.IntentEducation,
		Language:       nlp.LanguagePortuguese,
		Facts:          []string{"João Paulo tem formação prevista em Inteligência Artificial pela FIAP entre agosto de 2026 e junho de 2028."},
		FormattedFacts: "João Paulo tem formação prevista em Inteligência Artificial pela FIAP entre agosto de 2026 e junho de 2028.",
		DetailLevel:    DetailMedium,
	}

	for range 10 {
		template := SelectTemplateForPlan(plan)
		response := RenderTemplate(template, plan)

		if strings.Count(response, "FIAP") != 1 {
			t.Fatalf("RenderTemplate() = %q, want FIAP mentioned exactly once", response)
		}

		if strings.Contains(response, "FIAP segue estudando") {
			t.Fatalf("RenderTemplate() = %q, contains incoherent education template", response)
		}
	}
}

func TestFormatFactsPreservesProperNounsAfterConnectors(t *testing.T) {
	response := FormatFacts(
		[]string{
			"João Paulo trabalha com sistemas financeiros.",
			"Auronix demonstra consistência transacional.",
		},
		nlp.LanguagePortuguese,
	)

	if !strings.Contains(response, "Auronix") {
		t.Fatalf("FormatFacts() = %q, want proper noun Auronix preserved", response)
	}

	if strings.Contains(response, "auronix demonstra") {
		t.Fatalf("FormatFacts() = %q, lowercased a proper noun after connector", response)
	}
}

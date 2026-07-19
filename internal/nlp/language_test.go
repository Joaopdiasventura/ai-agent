package nlp

import "testing"

func TestDetectLanguageUsesWeightedSignals(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected Language
	}{
		{
			name:     "portuguese phrase wins over ambiguous me token",
			tokens:   []string{"me", "fale", "sobre", "auronix"},
			expected: LanguagePortuguese,
		},
		{
			name:     "english phrase is detected",
			tokens:   []string{"tell", "me", "about", "auronix"},
			expected: LanguageEnglish,
		},
		{
			name:     "tie defaults to portuguese",
			tokens:   []string{"me", "portfolio"},
			expected: LanguagePortuguese,
		},
		{
			name:     "missing signals defaults to portuguese",
			tokens:   []string{"auronix", "joão"},
			expected: LanguagePortuguese,
		},
		{
			name:     "portuguese email question wins over english email token",
			tokens:   []string{"qual", "o", "email", "dele"},
			expected: LanguagePortuguese,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if language := DetectLanguage(test.tokens); language != test.expected {
				t.Fatalf("DetectLanguage() = %q, want %q", language, test.expected)
			}
		})
	}
}

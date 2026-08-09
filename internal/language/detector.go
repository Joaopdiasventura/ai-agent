package language

import (
	"ai-agent/internal/domain"
	"math"
)

type Detection struct {
	Language   domain.Language
	Portuguese float64
	English    float64
	Confidence float64
}

var portugueseMarkers = map[string]float64{
	"qual":          2,
	"quais":         2,
	"quem":          2,
	"onde":          2,
	"quando":        2,
	"como":          1.5,
	"porque":        1.5,
	"ele":           2,
	"ela":           2,
	"dele":          2,
	"dela":          2,
	"trabalhou":     2,
	"trabalha":      2,
	"trabalho":      1,
	"usou":          2,
	"usa":           1.5,
	"utilizou":      2,
	"utiliza":       1.5,
	"desenvolveu":   2,
	"implementou":   2,
	"criou":         2,
	"projeto":       1.5,
	"projetos":      1.5,
	"experiencia":   2,
	"profissional":  1,
	"concorrencia":  2,
	"desempenho":    1.5,
	"performance":   0.5,
	"lideranca":     2,
	"arquitetura":   1.5,
	"formacao":      2,
	"educacao":      2,
	"certificacao":  2,
	"certificacoes": 2,
	"habilidade":    1.5,
	"habilidades":   1.5,
	"tecnologia":    1,
	"tecnologias":   1,
	"empresa":       1,
	"empresas":      1,
	"cargo":         1.5,
	"atual":         1,
	"atualmente":    1.5,
	"fale":          2,
	"conte":         2,
	"possui":        1.5,
	"sabe":          1.5,
	"conhece":       1.5,
	"brasil":        1,
	"portugues":     2,
	"ingles":        1,
	"curriculo":     2,
	"financeiro":    1,
	"financeiros":   1,
	"banco":         1,
	"dados":         1,
}

var englishMarkers = map[string]float64{
	"what":           2,
	"which":          2,
	"who":            2,
	"where":          2,
	"when":           2,
	"how":            1.5,
	"why":            1.5,
	"he":             2,
	"his":            2,
	"him":            2,
	"worked":         2,
	"works":          1.5,
	"work":           1,
	"used":           2,
	"uses":           1.5,
	"use":            1,
	"developed":      2,
	"implemented":    2,
	"created":        2,
	"project":        1.5,
	"projects":       1.5,
	"experience":     2,
	"professional":   1,
	"concurrency":    2,
	"performance":    0.5,
	"leadership":     2,
	"architecture":   1.5,
	"education":      2,
	"certification":  2,
	"certifications": 2,
	"skill":          1.5,
	"skills":         1.5,
	"technology":     1,
	"technologies":   1,
	"company":        1,
	"companies":      1,
	"role":           1.5,
	"current":        1,
	"currently":      1.5,
	"tell":           2,
	"know":           1.5,
	"knows":          1.5,
	"does":           1,
	"can":            1,
	"brazil":         1,
	"english":        2,
	"portuguese":     1,
	"resume":         2,
	"financial":      1,
	"bank":           1,
	"data":           1,
}

func DetectLanguage(value string) domain.Language {
	return DetectLanguageDetailed(value).Language
}

func DetectLanguageDetailed(value string) Detection {
	normalized := Normalize(value)
	tokens := TokenizeNormalized(normalized)

	var portuguese float64
	var english float64

	for _, token := range tokens {
		if score, found := portugueseMarkers[token]; found {
			portuguese += score
		}

		if score, found := englishMarkers[token]; found {
			english += score
		}
	}

	if containsPortugueseDiacritic(value) {
		portuguese++
	}

	result := Detection{
		Portuguese: portuguese,
		English:    english,
	}

	if portuguese == 0 && english == 0 {
		result.Language = domain.LanguageUnknown
		return result
	}

	difference := math.Abs(portuguese - english)
	total := portuguese + english

	if difference < 0.5 {
		result.Language = domain.LanguageUnknown
		result.Confidence = 0
		return result
	}

	result.Confidence = difference / total

	if portuguese > english {
		result.Language = domain.LanguagePortuguese
		return result
	}

	result.Language = domain.LanguageEnglish

	return result

}

func containsPortugueseDiacritic(value string) bool {
	for _, current := range value {
		switch current {
		case 'á', 'à', 'â', 'ã',
			'é', 'ê',
			'í',
			'ó', 'ô', 'õ',
			'ú',
			'ç',
			'Á', 'À', 'Â', 'Ã',
			'É', 'Ê',
			'Í',
			'Ó', 'Ô', 'Õ',
			'Ú',
			'Ç':
			return true
		}
	}

	return false
}

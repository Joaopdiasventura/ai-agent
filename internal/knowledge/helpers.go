package knowledge

import "ai-agent/internal/domain"

func localized(pt string, en string) domain.LocalizedText {
	return domain.LocalizedText{
		PT: pt,
		EN: en,
	}
}

func aliasMap(
	common []string,
	pt []string,
	en []string,
) map[domain.Language][]string {
	result := make(map[domain.Language][]string)

	if len(common) > 0 {
		result[domain.LanguageUnknown] = common
	}

	if len(pt) > 0 {
		result[domain.LanguagePortuguese] = pt
	}

	if len(en) > 0 {
		result[domain.LanguageEnglish] = en
	}

	return result
}

func concepts(values ...string) []domain.ConceptID {
	result := make([]domain.ConceptID, 0, len(values))

	for _, value := range values {
		result = append(result, domain.ConceptID(value))
	}

	return result
}

func currentPeriod(year int, month int) *domain.Period {
	start := domain.YearMonth{
		Year:  year,
		Month: month,
	}

	return &domain.Period{
		Start:   &start,
		Current: true,
	}
}

func closedPeriod(
	startYear int,
	startMonth int,
	endYear int,
	endMonth int,
) *domain.Period {
	start := domain.YearMonth{
		Year:  startYear,
		Month: startMonth,
	}

	end := domain.YearMonth{
		Year:  endYear,
		Month: endMonth,
	}

	return &domain.Period{
		Start: &start,
		End:   &end,
	}
}

func entityFact(
	id string,
	subject domain.EntityID,
	predicate domain.Relation,
	object domain.EntityID,
	category domain.FactCategory,
	factConcepts []domain.ConceptID,
	context []domain.EntityID,
	statement domain.LocalizedText,
	importance float64,
	period *domain.Period,
) domain.Fact {
	return domain.Fact{
		ID:         domain.FactID(id),
		Subject:    subject,
		Predicate:  predicate,
		Object:     domain.EntityObject(object),
		Category:   category,
		Concepts:   factConcepts,
		Context:    context,
		Statement:  statement,
		Importance: importance,
		Period:     period,
		Source:     SourceCV,
	}
}

func textFact(
	id string,
	subject domain.EntityID,
	predicate domain.Relation,
	value domain.LocalizedText,
	category domain.FactCategory,
	factConcepts []domain.ConceptID,
	context []domain.EntityID,
	statement domain.LocalizedText,
	importance float64,
	period *domain.Period,
) domain.Fact {
	return domain.Fact{
		ID:         domain.FactID(id),
		Subject:    subject,
		Predicate:  predicate,
		Object:     domain.TextObject(value),
		Category:   category,
		Concepts:   factConcepts,
		Context:    context,
		Statement:  statement,
		Importance: importance,
		Period:     period,
		Source:     SourceCV,
	}
}

func numberFact(
	id string,
	subject domain.EntityID,
	predicate domain.Relation,
	value float64,
	unit string,
	operator domain.NumberOperator,
	category domain.FactCategory,
	factConcepts []domain.ConceptID,
	context []domain.EntityID,
	statement domain.LocalizedText,
	importance float64,
	period *domain.Period,
) domain.Fact {
	return domain.Fact{
		ID:        domain.FactID(id),
		Subject:   subject,
		Predicate: predicate,
		Object: domain.QualifiedNumberObject(
			value,
			unit,
			operator,
		),
		Category:   category,
		Concepts:   factConcepts,
		Context:    context,
		Statement:  statement,
		Importance: importance,
		Period:     period,
		Source:     SourceCV,
	}
}

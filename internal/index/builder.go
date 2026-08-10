package index

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/language"
	"ai-agent/internal/ontology"
	"fmt"
	"strconv"
)

func buildDocument(
	fact domain.Fact,
	targetLanguage domain.Language,
	base *knowledge.Knowledge,
	currentOntology *ontology.Ontology,
) (Document, error) {
	fields := make(map[Field]FieldData)

	addTokens(
		fields,
		FieldStatement,
		language.Tokenize(
			fact.Statement.For(targetLanguage),
		),
	)

	subject, found := base.Entity(fact.Subject)

	if !found {
		return Document{}, fmt.Errorf(
			"unknown subject entity %s",
			fact.Subject,
		)
	}

	addTokens(
		fields,
		FieldSubject,
		entityTokens(subject, targetLanguage),
	)

	objectTokens, err := factObjectTokens(fact.Object, targetLanguage, base)

	if err != nil {
		return Document{}, err
	}

	addTokens(
		fields,
		FieldObject,
		objectTokens,
	)

	for _, conceptID := range fact.Concepts {
		currentConcept, found := currentOntology.Concept(conceptID)

		if !found {
			return Document{}, fmt.Errorf(
				"unknown concept %s",
				conceptID,
			)
		}

		addTokens(
			fields,
			FieldConcept,
			conceptTokens(
				currentConcept,
				targetLanguage,
			),
		)
	}

	for _, contextID := range fact.Context {
		contextEntity, found := base.Entity(contextID)

		if !found {
			return Document{}, fmt.Errorf(
				"unknown context entity %s",
				contextID,
			)
		}

		addTokens(
			fields,
			FieldContext,
			entityTokens(
				contextEntity,
				targetLanguage,
			),
		)
	}

	addTokens(
		fields,
		FieldPredicate,
		language.Tokenize(string(fact.Predicate)),
	)

	addTokens(
		fields,
		FieldCategory,
		language.Tokenize(string(fact.Category)),
	)

	totalLegth := 0

	for _, field := range fields {
		totalLegth += field.Length
	}

	return Document{
		FactID:   fact.ID,
		Language: targetLanguage,
		Fields:   fields,
		Length:   totalLegth,
	}, nil
}

func factObjectTokens(
	object domain.FactObject,
	targetLanguage domain.Language,
	base *knowledge.Knowledge,
) ([]string, error) {
	switch object.Kind {
	case domain.FactObjectEntity:
		entity, found := base.Entity(object.EntityID)

		if !found {
			return nil, fmt.Errorf(
				"unknown object entity %s",
				object.EntityID,
			)
		}

		return entityTokens(entity, targetLanguage), nil

	case domain.FactObjectText:
		return language.Tokenize(object.Text.For(targetLanguage)), nil

	case domain.FactObjectNumber:
		value := strconv.FormatFloat(object.Number, 'f', -1, 64)
		result := language.Tokenize(value)

		if object.Unit != "" {
			result = append(result, language.Tokenize(object.Unit)...)
		}

		return result, nil

	case domain.FactObjectBoolean:
		return []string{
			strconv.FormatBool(object.Boolean),
		}, nil

	default:
		return nil, fmt.Errorf(
			"unsupported fact object kind %s",
			object.Kind,
		)
	}
}

func entityTokens(
	entity domain.Entity,
	targetLanguage domain.Language,
) []string {
	values := entity.AllAliases(targetLanguage)

	return uniqueTokensFromValues(values)
}

func conceptTokens(
	currentConcept domain.Concept,
	targetLanguage domain.Language,
) []string {
	values := currentConcept.AllAliases(targetLanguage)

	return uniqueTokensFromValues(values)
}

func uniqueTokensFromValues(values []string) []string {
	result := make([]string, 0)
	seen := make(map[string]struct{})

	for _, value := range values {
		for _, token := range language.Tokenize(value) {
			if token == "" {
				continue
			}

			if _, exists := seen[token]; exists {
				continue
			}

			seen[token] = struct{}{}
			result = append(result, token)
		}
	}

	return result
}

func addTokens(
	fields map[Field]FieldData,
	field Field,
	tokens []string,
) {
	if len(tokens) == 0 {
		return
	}

	data, exists := fields[field]

	if !exists {
		data = FieldData{
			Terms: make(map[string]int),
		}
	}

	for _, token := range tokens {
		if token == "" {
			continue
		}

		data.Terms[token]++
		data.Length++
	}

	fields[field] = data
}

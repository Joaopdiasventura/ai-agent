package index

import (
	"ai-agent/internal/domain"
)

type Field string

const (
	FieldStatement Field = "statement"
	FieldSubject   Field = "subject"
	FieldObject    Field = "object"
	FieldConcept   Field = "concept"
	FieldContext   Field = "context"
	FieldPredicate Field = "predicate"
	FieldCategory  Field = "category"
)

var AllFields = []Field{
	FieldStatement,
	FieldSubject,
	FieldObject,
	FieldConcept,
	FieldContext,
	FieldPredicate,
	FieldCategory,
}

type Posting struct {
	FactID    domain.FactID
	Field     Field
	Frequency int
}

type FieldData struct {
	Length int
	Terms  map[string]int
}

func (f FieldData) Frequency(term string) int {
	return f.Terms[term]
}

type Document struct {
	FactID   domain.FactID
	Language domain.Language
	Fields   map[Field]FieldData
	Length   int
}

func (d Document) Field(
	field Field,
) (FieldData, bool) {
	value, found := d.Fields[field]
	return value, found
}

func (d Document) TermFrequency(
	field Field,
	term string,
) int {
	data, found := d.Fields[field]

	if !found {
		return 0
	}

	return data.Terms[term]
}

func (d Document) TotalTermFrequency(
	term string,
) int {
	total := 0

	for _, field := range d.Fields {
		total += field.Terms[term]
	}

	return total
}

type Statistics struct {
	DocumentCount         int
	AverageDocumentLength float64
	AverageFieldLength    map[Field]float64
}

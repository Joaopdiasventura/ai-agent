package domain

import (
	"errors"
	"fmt"
)

type FactID string

type FactCategory string

const (
	FactCategoryUnknown       FactCategory = "unknown"
	FactCategoryProfile       FactCategory = "profile"
	FactCategorySkill         FactCategory = "skill"
	FactCategoryExperience    FactCategory = "experience"
	FactCategoryProject       FactCategory = "project"
	FactCategoryEducation     FactCategory = "education"
	FactCategoryCertification FactCategory = "certification"
	FactCategoryLanguage      FactCategory = "language"
	FactCategoryContact       FactCategory = "contact"
	FactCategoryAchievement   FactCategory = "achievement"
)

type FactObjectKind string

const (
	FactObjectUnknown FactObjectKind = "unknown"
	FactObjectEntity  FactObjectKind = "entity"
	FactObjectText    FactObjectKind = "text"
	FactObjectNumber  FactObjectKind = "number"
	FactObjectBoolean FactObjectKind = "boolean"
)

type FactObject struct {
	Kind     FactObjectKind
	EntityID EntityID
	Text     LocalizedText
	Number   float64
	Boolean  bool
	Unit     string
}

func EntityObject(entityID EntityID) FactObject {
	return FactObject{
		Kind:     FactObjectEntity,
		EntityID: entityID,
	}
}

func TextObject(text LocalizedText) FactObject {
	return FactObject{
		Kind: FactObjectText,
		Text: text,
	}
}

func NumberObject(value float64, unit string) FactObject {
	return FactObject{
		Kind:   FactObjectNumber,
		Number: value,
		Unit:   unit,
	}
}

func BooleanObject(value bool) FactObject {
	return FactObject{
		Kind:    FactObjectBoolean,
		Boolean: value,
	}
}

func (o FactObject) Valid() bool {
	switch o.Kind {
	case FactObjectEntity:
		return o.EntityID != ""
	case FactObjectText:
		return !o.Text.Empty()
	case FactObjectNumber:
		return true
	case FactObjectBoolean:
		return true
	default:
		return false
	}
}

type YearMonth struct {
	Year  int
	Month int
}

func (d YearMonth) Valid() bool {
	return d.Year > 0 && d.Month >= 1 && d.Month <= 12
}

func (d YearMonth) Before(other YearMonth) bool {
	if d.Year != other.Year {
		return d.Year < other.Year
	}

	return d.Month < other.Month
}

type Period struct {
	Start   *YearMonth
	End     *YearMonth
	Current bool
}

func (p Period) Validate() error {
	if p.Start != nil && !p.Start.Valid() {
		return errors.New("invalid start date")
	}

	if p.End != nil && !p.End.Valid() {
		return errors.New("invalid end date")
	}

	if p.Current && p.End != nil {
		return errors.New("current period cannot have an end date")
	}

	if p.Start != nil && p.End != nil && p.End.Before(*p.Start) {
		return errors.New("end date cannot be before start date")
	}

	return nil
}

type Fact struct {
	ID         FactID
	Subject    EntityID
	Predicate  Relation
	Object     FactObject
	Category   FactCategory
	Concepts   []ConceptID
	Statement  LocalizedText
	Period     *Period
	Importance float64
	Source     string
}

func (f Fact) Validate() error {
	if f.ID == "" {
		return errors.New("fact id is required")
	}

	if f.Subject == "" {
		return errors.New("fact subject is required")
	}

	if f.Predicate == "" || f.Predicate == RelationUnknown {
		return errors.New("fact predicate is required")
	}

	if !f.Object.Valid() {
		return errors.New("fact object is invalid")
	}

	if f.Statement.Empty() {
		return errors.New("fact statement is required")
	}

	if f.Importance < 0 || f.Importance > 1 {
		return fmt.Errorf("fact importance must be between 0 and 1: %f", f.Importance)
	}

	if f.Period != nil {
		if err := f.Period.Validate(); err != nil {
			return err
		}
	}

	return nil
}

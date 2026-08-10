package generation

import (
	"fmt"

	"ai-agent/internal/domain"
	"ai-agent/internal/planning"
)

type Source interface {
	Fact(
		id domain.FactID,
	) (domain.Fact, bool)

	Entity(
		id domain.EntityID,
	) (domain.Entity, bool)
}

type Material struct {
	Facts    map[domain.FactID]domain.Fact
	Entities map[domain.EntityID]domain.Entity
}

func NewMaterial() Material {
	return Material{
		Facts: make(
			map[domain.FactID]domain.Fact,
		),
		Entities: make(
			map[domain.EntityID]domain.Entity,
		),
	}
}

func Materialize(
	plan planning.Plan,
	source Source,
) (Material, error) {
	if source == nil {
		return Material{},
			fmt.Errorf(
				"generation source is required",
			)
	}

	result :=
		NewMaterial()

	for _, factID :=
		range plan.FactIDs() {
		fact, found :=
			source.Fact(factID)

		if !found {
			return Material{},
				fmt.Errorf(
					"planned fact %s was not found",
					factID,
				)
		}

		result.Facts[factID] =
			fact
	}

	for _, entityID :=
		range plan.EntityIDs() {
		entity, found :=
			source.Entity(entityID)

		if !found {
			return Material{},
				fmt.Errorf(
					"planned entity %s was not found",
					entityID,
				)
		}

		result.Entities[entityID] =
			entity
	}

	return result, nil
}

func (m Material) Validate(
	plan planning.Plan,
) error {
	for _, factID :=
		range plan.FactIDs() {
		if _, found :=
			m.Facts[factID]; !found {
			return fmt.Errorf(
				"material does not contain planned fact %s",
				factID,
			)
		}
	}

	for _, entityID :=
		range plan.EntityIDs() {
		if _, found :=
			m.Entities[
				entityID,
			]; !found {
			return fmt.Errorf(
				"material does not contain planned entity %s",
				entityID,
			)
		}
	}

	return nil
}

func (m Material) Fact(
	id domain.FactID,
) (domain.Fact, bool) {
	value, found :=
		m.Facts[id]

	return value, found
}

func (m Material) Entity(
	id domain.EntityID,
) (domain.Entity, bool) {
	value, found :=
		m.Entities[id]

	return value, found
}
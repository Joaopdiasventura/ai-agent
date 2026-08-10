package confidence

import "fmt"

type Config struct {
	ClaimWeights      map[SignalName]float64
	AbstentionWeights map[SignalName]float64
}

func DefaultConfig() Config {
	return Config{
		ClaimWeights: map[SignalName]float64{
			SignalQueryQuality:       0.12,
			SignalRetrievalAgreement: 0.12,
			SignalSeparation:         0.10,
			SignalEvidenceStrength:   0.20,
			SignalEvidenceDirectness: 0.14,
			SignalSemanticCoverage:   0.14,
			SignalPlanGrounding:      0.08,
			SignalAnswerGrounding:    0.10,
		},
		AbstentionWeights: map[SignalName]float64{
			SignalQueryQuality:    0.15,
			SignalEvidenceAbsence: 0.45,
			SignalPlanGrounding:   0.20,
			SignalAnswerGrounding: 0.20,
		},
	}
}

func (c Config) Validate() error {
	if err :=
		validateWeights(
			c.ClaimWeights,
		); err != nil {
		return fmt.Errorf(
			"invalid claim weights: %w",
			err,
		)
	}

	if err :=
		validateWeights(
			c.AbstentionWeights,
		); err != nil {
		return fmt.Errorf(
			"invalid abstention weights: %w",
			err,
		)
	}

	return nil
}

func validateWeights(
	values map[SignalName]float64,
) error {
	if len(values) == 0 {
		return fmt.Errorf(
			"weights are required",
		)
	}

	total := 0.0

	for name, weight := range values {
		if weight < 0 {
			return fmt.Errorf(
				"negative weight for %s",
				name,
			)
		}

		total += weight
	}

	if total <= 0 {
		return fmt.Errorf(
			"total weight must be positive",
		)
	}

	return nil
}

func signalWeight(
	values map[SignalName]float64,
	name SignalName,
) float64 {
	return values[name]
}

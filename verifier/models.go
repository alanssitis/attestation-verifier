package verifier

import (
	"os"

	"gopkg.in/yaml.v3"
)

// copied from go-sslib to use yaml tags
type Functionary struct {
	KeyType             string   `yaml:"keyType"`
	Scheme              string   `yaml:"scheme"`
	KeyPath             string   `yaml:"keyPath"`
}

type KeyVal struct {
	Public string `yaml:"public"`
}

type Constraint struct {
	Rule           string `yaml:"rule"`
	AllowIfNoClaim bool   `yaml:"allowIfNoClaim"`
	Warn           bool   `yaml:"warn"`
	Debug          string `yaml:"debug"`
}

type ExpectedStepPredicates struct {
	PredicateType      string       `yaml:"predicateType"`
	ExpectedAttributes []Constraint `yaml:"expectedAttributes"`
	Functionaries      []string     `yaml:"functionaries"`
	Threshold          int          `yaml:"threshold"`
}

type Step struct {
	Name               string                   `yaml:"name"`
	Command            string                   `yaml:"command"`
	ExpectedMaterials  []string                 `yaml:"expectedMaterials"`
	ExpectedProducts   []string                 `yaml:"expectedProducts"`
	ExpectedPredicates []ExpectedStepPredicates `yaml:"expectedPredicates"`
}

type ExpectedSubjectPredicates struct {
	PredicateType      string       `yaml:"predicateType"`
	ExpectedAttributes []Constraint `yaml:"expectedAttributes"`
	Functionaries      []string     `yaml:"functionaries"`
	Threshold          int          `yaml:"threshold"`
}

type Subject struct {
	Subject            []string                    `yaml:"subject"`
	ExpectedPredicates []ExpectedSubjectPredicates `yaml:"expectedPredicates"`
}

type Inspection struct {
	Name               string       `yaml:"name"`
	Command            string       `yaml:"command"`
	Predicates         []string     `yaml:"predicates"`
	ExpectedMaterials  []string     `yaml:"expectedMaterials"`
	ExpectedProducts   []string     `yaml:"expectedProducts"`
	ExpectedAttributes []Constraint `yaml:"expectedAttributes"`
}

type Layout struct {
	Expires       string                 `yaml:"expires"`
	Functionaries map[string]Functionary `yaml:"functionaries"`
	Steps         []*Step                `yaml:"steps"`
	Subjects      []*Subject             `yaml:"subjects"`
	Inspections   []*Inspection          `yaml:"inspections"`
}

func LoadLayout(path string) (*Layout, error) {
	layoutBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	layout := &Layout{}
	if err := yaml.Unmarshal(layoutBytes, layout); err != nil {
		return nil, err
	}

	return layout, nil
}

type AttestationIdentifier struct {
	PredicateType string
	Functionary   string
}

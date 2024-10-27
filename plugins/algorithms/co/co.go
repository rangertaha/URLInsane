package co

import (
	"github.com/rangertaha/urlinsane"
	"github.com/rangertaha/urlinsane/plugins/algorithms"
)

const CODE = "co"

type CharacterOmission struct {
	types []string
}

func (n *CharacterOmission) Id() string {
	return CODE
}
func (n *CharacterOmission) IsType(str string) bool {
	return algorithms.IsType(n.types, str)
}

func (n *CharacterOmission) Name() string {
	return "Character Omission"
}

func (n *CharacterOmission) Description() string {
	return "omitting a character from the domain"
}

func (n *CharacterOmission) Fields() []string {
	return []string{}
}

func (n *CharacterOmission) Headers() []string {
	return []string{}
}

func (n *CharacterOmission) Exec(in urlinsane.Typo) (out []urlinsane.Typo) {
	out = append(out, in)
	return
}

// Register the plugin
func init() {
	algorithms.Add(CODE, func() urlinsane.Algorithm {
		return &CharacterOmission{
			types: []string{algorithms.ENTITY, algorithms.DOMAIN},
		}
	})
}

// Copyright (C) 2024 Rangertaha
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package vs

import (
	"fmt"

	"github.com/rangertaha/urlinsane/internal"
	"github.com/rangertaha/urlinsane/internal/pkg/domain"
	"github.com/rangertaha/urlinsane/internal/plugins/algorithms"
	algo "github.com/rangertaha/urlinsane/pkg/typo"
)

const (
	CODE        = "vs"
	NAME        = "Vowel Swapping"
	DESCRIPTION = "Vowel Swapping replaces vowels in the domain name with other vowels to create variations."
)

type Algo struct {
	config    internal.Config
	languages []internal.Language
	keyboards []internal.Keyboard
	funcs     map[int]func(internal.Typo) []internal.Typo
}

func (n *Algo) Id() string {
	return CODE
}

func (n *Algo) Init(conf internal.Config) {
	n.funcs = make(map[int]func(internal.Typo) []internal.Typo)
	n.keyboards = conf.Keyboards()
	n.languages = conf.Languages()
	n.config = conf

	// Supported targets
	n.funcs[internal.DOMAIN] = n.domain
	n.funcs[internal.PACKAGE] = n.name
	n.funcs[internal.EMAIL] = n.email
	n.funcs[internal.NAME] = n.name
}

func (n *Algo) Name() string {
	return NAME
}
func (n *Algo) Description() string {
	return DESCRIPTION
}

func (n *Algo) Exec(typo internal.Typo) []internal.Typo {
	return n.funcs[n.config.Type()](typo)
}

func (n *Algo) domain(typo internal.Typo) (typos []internal.Typo) {
	sub, prefix, suffix := typo.Original().Domain()

	for _, language := range n.languages {
		for _, variant := range algo.VowelSwapping(prefix, language.Vowels()...) {
			if prefix != variant {
				d := domain.New(sub, variant, suffix)
				new := typo.Clone(d.String())
				typos = append(typos, new)
			}
		}
	}
	return
}

func (n *Algo) email(typo internal.Typo) (typos []internal.Typo) {
	username, domain := typo.Original().Email()
	for _, language := range n.languages {
		for _, variant := range algo.VowelSwapping(username, language.Vowels()...) {
			if username != variant {
				new := typo.Clone(fmt.Sprintf("%s@%s", variant, domain))
				typos = append(typos, new)
			}
		}
	}
	return
}

func (n *Algo) name(typo internal.Typo) (typos []internal.Typo) {
	name := n.config.Target().Name()
	for _, language := range n.languages {
		for _, variant := range algo.VowelSwapping(name, language.Vowels()...) {
			if name != variant {
				typos = append(typos, typo.Clone(variant))
			}
		}
	}
	return
}

// Register the plugin
func init() {
	algorithms.Add(CODE, func() internal.Algorithm {
		return &Algo{}
	})
}

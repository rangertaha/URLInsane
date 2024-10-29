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
package mds

// Missing Dashes typos are created by omitting a dash from the domain.
// For example, www.a-b-c.com becomes www.ab-c.com, www.a-bc.com, and ww.abc.com



import (
	"github.com/rangertaha/urlinsane"
	"github.com/rangertaha/urlinsane/plugins/algorithms"
	"github.com/rangertaha/urlinsane/utils/nlp"
)

const CODE = "mds"
// const (
// 	CODE        = ""
// 	NAME        = ""
// 	DESCRIPTION = ""
// )


type Algo struct {
	types []string
}

func (n *Algo) Id() string {
	return CODE
}
func (n *Algo) IsType(str string) bool {
	return algorithms.IsType(n.types, str)
}

func (n *Algo) Name() string {
	return "Missing Dashes"
}

func (n *Algo) Description() string {
	return "created by stripping all dashes from the name"
}

func (n *Algo) Exec(typo urlinsane.Typo) (typos []urlinsane.Typo) {
	for _, variant := range nlp.MissingCharFunc(typo.Original().Repr(), "-") {
		if typo.Original().Repr() != variant {
			typos = append(typos, typo.New(variant))
		}
	}
	return
}

// Register the plugin
func init() {
	algorithms.Add(CODE, func() urlinsane.Algorithm {
		return &Algo{
			types: []string{algorithms.ENTITY, algorithms.DOMAIN},
		}
	})
}

// // missingDashFunc typos are created by omitting a dash from the domain.
// // For example, www.a-b-c.com becomes www.ab-c.com, www.a-bc.com, and ww.abc.com
// func missingDashFunc(tc Result) (results []Result) {
// 	for _, str := range missingCharFunc(tc.Original.Domain, "-") {
// 		if tc.Original.Domain != str {
// 			dm := Domain{tc.Original.Subdomain, str, tc.Original.Suffix, Meta{}, false}
// 			results = append(results, Result{Original: tc.Original, Variant: dm, Typo: tc.Typo, Data: tc.Data})
// 		}
// 	}
// 	dm := Domain{tc.Original.Subdomain, strings.Replace(tc.Original.Domain, "-", "", -1), tc.Original.Suffix, Meta{}, false}
// 	results = append(results, Result{Original: tc.Original, Variant: dm, Typo: tc.Typo, Data: tc.Data})
// 	return results
// }

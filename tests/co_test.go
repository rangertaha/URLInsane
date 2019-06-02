// The MIT License (MIT)
//
// Copyright © 2019 Tal Hachi
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tests

import (
	"testing"

	"github.com/cybersectech-org/urlinsane"
)

type testpair struct {
	domains []string
	values  map[string]bool
	total   int
}

var languages = []string{"en", "iw", "es", "fa", "fi", "ar", "ru", "hy"}

var characterOmissionCases = []testpair{
	{[]string{"google.com"},
		map[string]bool{
			"oogle.com": true,
			"gogle.com": true,
			"goole.com": true,
			"googl.com": true,
			"googe.com": true,
		}, 5},
	{[]string{"example.com"},
		map[string]bool{
			"xample.com": true,
			"exmple.com": true,
			"eample.com": true,
			"examle.com": true,
			"exaple.com": true,
			"exampl.com": true,
			"exampe.com": true,
		}, 5},
}

func TestCharacterOmission(t *testing.T) {
	for _, lang := range languages {
		count := 0
		for _, tcase := range characterOmissionCases {
			conf := urlinsane.BasicConfig{
				Domains:     tcase.domains,
				Keyboards:   []string{lang},
				Typos:       []string{"co"},
				Funcs:       []string{""},
				Concurrency: 50,
				Format:      "text",
				Verbose:     false,
			}

			urli := urlinsane.New(conf.Config())

			out := urli.Stream()

			for r := range out {
				// fmt.Println(r)
				_, ok := tcase.values[r.Variant.String()]
				if !ok {
					t.Errorf("Failed variant: %v for domains: %v, language: %v, algorithm %v", r.Variant.String(), tcase.domains, lang, r.Typo.Name)
				}
				count++
			}
			// TODO: Apply dup filter and uncomment
			// if count != tcase.total {
			// 	t.Errorf("Failed total number of records should be %v not %v", tcase.total, count)
			// }
			count = 0
		}
	}
}

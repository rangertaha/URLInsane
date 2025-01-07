// Copyright 2024 Rangertaha. All Rights Reserved.
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
package table

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rangertaha/urlinsane/internal"
	"github.com/rangertaha/urlinsane/internal/pkg"
	"github.com/rangertaha/urlinsane/internal/plugins/outputs"
	log "github.com/sirupsen/logrus"
	"golang.org/x/term"
)

const (
	CODE        = "table"
	DESCRIPTION = "Pretty table output format with color"
)

type Plugin struct {
	table   table.Writer
	config  internal.Config
	domains []internal.Domain
	output  string
}

func (p *Plugin) Id() string {
	return CODE
}

func (p *Plugin) Description() string {
	return DESCRIPTION
}

func (p *Plugin) Init(conf internal.Config) {
	n.config = conf
	n.table = table.NewWriter()

	if width, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil {
		n.table.SetAllowedRowLength(width - 4)
	}

	n.table.SetOutputMirror(os.Stdout)
	n.table.AppendHeader(n.Header())
	n.table.AppendFooter(n.Header())
	n.Config()
}

func (p *Plugin) Read(in internal.Domain) {
	n.table.AppendRow(n.Row(in))
}

func (p *Plugin) Header() (row table.Row) {
	row = append(row, "LD")
	row = append(row, "TYPE")
	row = append(row, "TYPO")

	for _, info := range n.config.Collectors() {
		for _, header := range info.Headers() {
			// if n.Filter(header) {
			row = append(row, header)
			// }
		}
	}
	return
}

func (p *Plugin) Row(domain internal.Domain) (row table.Row) {
	n.domains = append(n.domains, domain)
	row = append(row, domain.Ld())
	if n.config.Verbose() {
		row = append(row, domain.Algorithm().Name())
	} else {
		row = append(row, strings.ToUpper(domain.Algorithm().Id()))
	}
	row = append(row, domain.String())

	for _, info := range n.config.Collectors() {
		for _, header := range info.Headers() {
			// if n.Filter(header) {
			meta := domain.Meta()
			if col, ok := meta[header]; ok {
				row = append(row, col)
			} else {
				row = append(row, "")
			}
			// }
		}
	}
	return
}

func (p *Plugin) Filter(header string) bool {
	header = strings.TrimSpace(header)
	header = strings.ToLower(header)
	for _, filter := range n.config.Filters() {
		filter = strings.TrimSpace(filter)
		filter = strings.ToLower(filter)
		if filter == header {
			return true
		}
	}
	return false
}

func (p *Plugin) Config() (row table.Row) {
	n.table.SetStyle(pkg.StyleDefault)

	// nameTransformer := text.Transformer(func(val interface{}) string {
	// 	if val.(string) == "MD" {
	// 		return text.Colors{text.BgBlack, text.FgGreen}.Sprint(val)
	// 	}
	// 	return fmt.Sprint(val)
	// })

	// n.table.SetRowPainter()

	n.table.SetColumnConfigs(ColumnConfig)
	return
}
func (p *Plugin) Progress(typo <-chan internal.Domain) <-chan internal.Domain {
	return typo
}

func (p *Plugin) Write() {
	n.output = n.table.Render()
}

func (p *Plugin) Summary(report map[string]string) {
	fmt.Println("")
	for k, v := range report {
		fmt.Printf("%s %s   ", k, v)
	}
	fmt.Println("")
}

func (p *Plugin) Save(fname string) {
	results := []byte(n.output)
	if err := os.WriteFile(fname, results, 0644); err != nil {
		log.Errorf("Error: %s", err)
	}

}

// Register the plugin
func init() {
	outputs.Add(CODE, func() internal.Output {
		return &Plugin{}
	})
}

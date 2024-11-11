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
package mx

import (
	"net"
	"strings"

	"github.com/rangertaha/urlinsane/internal"
	"github.com/rangertaha/urlinsane/internal/plugins/collectors"
)

const (
	ORDER       = 2
	CODE        = "mx"
	NAME        = "MX Records"
	DESCRIPTION = "DNS MX Records"
)

type Plugin struct {
	conf internal.Config
	db   internal.Database
}

func (n *Plugin) Id() string {
	return CODE
}

func (n *Plugin) Order() int {
	return ORDER
}

func (i *Plugin) Init(c internal.Config) {
	i.db = c.Database()
	i.conf = c
}

func (n *Plugin) Description() string {
	return DESCRIPTION
}

func (n *Plugin) Headers() []string {
	return []string{"MX"}
}

func (i *Plugin) Exec(domain internal.Domain, acc internal.Accumulator) (err error) {
	mx, _ := i.db.Read(domain.String(), "MX")
	if mx != "" {
		domain.SetMeta("MX", mx)
		domain.Live(true)
		acc.Add(domain)
		return
	}

	records, err := net.LookupMX(domain.String())
	if err == nil {
		var answers []string
		for _, mx := range records {
			answers = append(answers, mx.Host)
		}
		record := strings.Join(answers, " ")
		domain.SetMeta("MX", record)
		domain.Live(true)

		err = i.db.Write(record, domain.String(), "MX")
	}
	acc.Add(domain)
	return
}

func (i *Plugin) Close() {}

// Register the plugin
func init() {
	collectors.Add(CODE, func() internal.Collector {
		return &Plugin{}
	})
}

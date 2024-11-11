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
package geo

import (
	"embed"
	"fmt"
	"io"
	"strings"

	"github.com/rainycape/geoip"
	"github.com/rangertaha/urlinsane/internal"
)

//go:embed GeoLite2-City.mmdb
var dataFile embed.FS

type Plugin struct {
	conf internal.Config
	db   internal.Database
}

func (i *Plugin) Init(c internal.Config) {
	i.db = c.Database()
	i.conf = c
}

func (n *Plugin) Headers() []string {
	return []string{"GEO"}
}

func (n *Plugin) Exec(domain internal.Domain, acc internal.Accumulator) (err error) {
	var location string
	if ipv4, ok := domain.GetMeta("IPv4").(string); ok {
		for _, ip := range strings.Split(ipv4, " ") {
			if loc, _ := n.getGeo(ip); loc != nil {
				if loc.Country != nil {
					location = loc.Country.Name.String()
				}
				if loc.City != nil {
					location = fmt.Sprintf("%s, %s", location, loc.City.Name.String())
				}
				domain.SetMeta("GEO", location)
			}
		}
	}

	acc.Add(domain)
	return
}

func (n *Plugin) getGeo(ip string) (r *geoip.Record, err error) {
	file, err := dataFile.Open("GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	db, err := geoip.New(file.(io.ReadSeeker))
	if err != nil {
		return nil, err
	}

	return db.Lookup(ip)
}

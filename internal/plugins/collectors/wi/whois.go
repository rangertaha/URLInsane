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
package wi

import (
	"net"
	"strings"

	"github.com/rangertaha/urlinsane/internal"
	"github.com/rangertaha/urlinsane/internal/db"
	"github.com/rangertaha/urlinsane/internal/plugins/collectors"
	log "github.com/sirupsen/logrus"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"strings"

// 	"github.com/rainycape/geoip"
// 	log "github.com/sirupsen/logrus"
// )

// type WhoisRecord struct {
// 	Domain         *WhoisDomain `json:"domain,omitempty"`
// 	Registrar      *Contact     `json:"registrar,omitempty"`
// 	Registrant     *Contact     `json:"registrant,omitempty"`
// 	Administrative *Contact     `json:"administrative,omitempty"`
// 	Technical      *Contact     `json:"technical,omitempty"`
// 	Billing        *Contact     `json:"billing,omitempty"`
// }

// // Domain storing domain name info
// type WhoisDomain struct {
// 	ID                   string     `json:"id,omitempty"`
// 	Domain               string     `json:"domain,omitempty"`
// 	Punycode             string     `json:"punycode,omitempty"`
// 	Name                 string     `json:"name,omitempty"`
// 	Extension            string     `json:"extension,omitempty"`
// 	WhoisServer          string     `json:"whois_server,omitempty"`
// 	Status               []string   `json:"status,omitempty"`
// 	NameServers          []string   `json:"name_servers,omitempty"`
// 	DNSSec               bool       `json:"dnssec,omitempty"`
// 	CreatedDate          string     `json:"created_date,omitempty"`
// 	CreatedDateInTime    *time.Time `json:"created_date_in_time,omitempty"`
// 	UpdatedDate          string     `json:"updated_date,omitempty"`
// 	UpdatedDateInTime    *time.Time `json:"updated_date_in_time,omitempty"`
// 	ExpirationDate       string     `json:"expiration_date,omitempty"`
// 	ExpirationDateInTime *time.Time `json:"expiration_date_in_time,omitempty"`
// }

// // Contact storing domain contact info
// type Contact struct {
// 	ID           string `json:"id,omitempty"`
// 	Name         string `json:"name,omitempty"`
// 	Organization string `json:"organization,omitempty"`
// 	Street       string `json:"street,omitempty"`
// 	City         string `json:"city,omitempty"`
// 	Province     string `json:"province,omitempty"`
// 	PostalCode   string `json:"postal_code,omitempty"`
// 	Country      string `json:"country,omitempty"`
// 	Phone        string `json:"phone,omitempty"`
// 	PhoneExt     string `json:"phone_ext,omitempty"`
// 	Fax          string `json:"fax,omitempty"`
// 	FaxExt       string `json:"fax_ext,omitempty"`
// 	Email        string `json:"email,omitempty"`
// 	ReferralURL  string `json:"referral_url,omitempty"`
// }

type Plugin struct {
	collectors.Plugin
}

func (i *Plugin) Exec(domain *db.Domain) (vaiant *db.Domain, err error) {
	records, err := net.LookupTXT(domain.Name)
	if err != nil {
		log.Error("TXT Lookup: ", err)
	}
	for _, record := range records {
		record := strings.TrimSpace(record)
		record = strings.Trim(record, ".")
		domain.Dns = append(domain.Dns, &db.DnsRecord{Type: "TXT", Value: record})
	}
	return domain, err
}

// Register the plugin
func init() {
	var CODE = "wi"
	collectors.Add(CODE, func() internal.Collector {
		return &Plugin{
			Plugin: collectors.Plugin{
				Num:       2,
				Code:      CODE,
				Title:     "WhoIs Lookup",
				Summary:   "Domain registration lookup",
				DependsOn: []string{},
			},
		}
	})
}

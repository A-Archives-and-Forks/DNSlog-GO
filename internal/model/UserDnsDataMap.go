package model

import (
	"sync"

	"github.com/lanyi1998/DNSlog-GO/internal/config"
)

var UserDnsDataMap *DnsDataMap

type DnsDataMap struct {
	userDnsData       sync.Map
	Mu                sync.RWMutex
	TokenSubDomainMap map[string]string
	SubDomainTokenMap map[string]string
}

func NewDnsDataMap(conf config.UserConfig) *DnsDataMap {
	TokenSubDomainMap := make(map[string]string)
	SubDomainTokenMap := make(map[string]string)
	for token, subdomain := range conf {
		TokenSubDomainMap[token] = subdomain
		SubDomainTokenMap[subdomain] = token
	}
	return &DnsDataMap{
		TokenSubDomainMap: TokenSubDomainMap,
		SubDomainTokenMap: SubDomainTokenMap,
	}
}

func (d *DnsDataMap) Get(token string) []DnsInfo {
	value, ok := d.userDnsData.Load(token)
	if ok {
		return value.([]DnsInfo)
	}
	return []DnsInfo{}
}

func (d *DnsDataMap) Set(token string, data DnsInfo) {
	value, ok := d.userDnsData.Load(token)
	d.Mu.Lock()
	defer d.Mu.Unlock()
	if ok {
		d.userDnsData.Store(token, append(value.([]DnsInfo), data))
	} else {
		d.userDnsData.Store(token, []DnsInfo{data})
	}

}

func (d *DnsDataMap) Clear(token string) {
	d.userDnsData.Delete(token)
}
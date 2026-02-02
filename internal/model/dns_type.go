package model

type DnsInfo struct {
	Type       DnsType
	Subdomain  string
	Ipaddress  string
	IpLocation string
	Time       int64
	Request    string
}
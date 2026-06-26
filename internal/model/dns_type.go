package model

type DnsInfo struct {
	Type       DnsType
	Subdomain  string
	Ipaddress  string
	IpLocation string
	Request    string
	Time       int64
}
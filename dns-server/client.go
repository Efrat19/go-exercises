package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/miekg/dns"
	"net"
	"time"
)

func main() {
	m1 := new(dns.Msg)
	m1.Id = dns.Id()
	m1.RecursionDesired = true
	m1.Question = make([]dns.Question, 1)
	m1.Question[0] = dns.Question{"google.com.", dns.TypeMX, dns.ClassINET}
	c := new(dns.Client)
	laddr := net.UDPAddr{
		IP:   net.ParseIP("[::1]"),
		Port: 12345,
		Zone: "",
	}
	c.Dialer = &net.Dialer{
		Timeout:   200 * time.Millisecond,
		LocalAddr: &laddr,
	}
	in, rtt, err := c.Exchange(m1, "8.8.8.8:53")
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(rtt)
	spew.Dump("****")
	spew.Dump(in)
}

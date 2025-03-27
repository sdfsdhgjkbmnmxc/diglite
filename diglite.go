package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func main() {
	for _, arg := range os.Args[1:] {
		err := printDomainInfo(arg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var types = map[uint16]string{
	dns.TypeA:     "A",
	dns.TypeAAAA:  "AAAA",
	dns.TypeMX:    "MX",
	dns.TypeTXT:   "TXT",
	dns.TypeNS:    "NS",
	dns.TypeSOA:   "SOA",
	dns.TypeCNAME: "CNAME",
	dns.TypeSRV:   "SRV",
	dns.TypeCAA:   "CAA",
}

func printDomainInfo(domainName string) error {
	fmt.Println("Domain:", domainName)
	fmt.Println()

	client := &dns.Client{}
	dnsServer := "1.1.1.1:53"

	for t, name := range types {
		msg := new(dns.Msg)
		msg.SetQuestion(dns.Fqdn(domainName), t)

		resp, _, err := client.Exchange(msg, dnsServer)
		if err != nil {
			fmt.Printf("%s query error: %v\n", name, err)
			continue
		}

		if len(resp.Answer) == 0 {
			continue
		}

		fmt.Printf("%s records:\n", name)
		for _, a := range resp.Answer {
			fmt.Println("  ", a.String())
		}
		fmt.Println()
	}

	fmt.Println()

	return nil
}

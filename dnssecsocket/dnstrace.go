package dnssecsocket

// inspired from github.com/rs/dnstrace/main.go

import "fmt"
import "net"
// import "os"
import "strings"
import "time"

import "github.com/miekg/dns"
import "github.com/rs/dnstrace/client"

// import log "github.com/sirupsen/logrus"

// this is cool, but breaks the Windows build
// import "github.com/wercker/journalhook"

// import "github.com/davecgh/go-spew/spew"

const (
	cReset    = 0
	cBold     = 1
	cRed      = 31
	cGreen    = 32
	cYellow   = 33
	cBlue     = 34
	cMagenta  = 35
	cCyan     = 36
	cGray     = 37
	cDarkGray = 90
)

func colorize(s interface{}, color int, enabled bool) string {
	if !enabled {
		return fmt.Sprintf("%v", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", color, s)
}

func Dnstrace(hostname string, qtypestr string) []dns.RR {
	// color := flag.Bool("color", true, "Enable/disable colors")
	color := true

	qname := dns.Fqdn(hostname)
	// qtype := dns.TypeA
	qtype := dns.StringToType[qtypestr]

	col := func(s interface{}, c int) string {
		return colorize(s, c, color)
	}

	m := &dns.Msg{}
	m.SetQuestion(qname, qtype)
	// Set DNSSEC opt to better emulate the default queries from a nameserver.
	o := &dns.OPT{
		Hdr: dns.RR_Header{
			Name:   ".",
			Rrtype: dns.TypeOPT,
		},
	}
	o.SetDo()
	o.SetUDPSize(dns.DefaultMsgSize)
	m.Extra = append(m.Extra, o)

	c := client.New(1)
	c.Client.Timeout = 500 * time.Millisecond
	t := client.Tracer{
		GotIntermediaryResponse: func(i int, m *dns.Msg, rs client.Responses, rtype client.ResponseType) {
			fr := rs.Fastest()
			var r *dns.Msg
			if fr != nil {
				r = fr.Msg
			}
			qname := m.Question[0].Name
			qtype := dns.TypeToString[m.Question[0].Qtype]
			if i > 1 {
				log(args.VerboseDnssec)
			}
			log(args.VerboseDnssec, "%d - query %s %s", i, qtype, qname)
			if r != nil {
				log(args.VerboseDnssec, ": %s", strings.Replace(strings.Replace(r.MsgHdr.String(), ";; ", "", -1), "\n", ", ", -1))
			}
			log(args.VerboseDnssec)
			for _, pr := range rs {
				ln := 0
				if pr.Msg != nil {
					ln = pr.Msg.Len()
				}
				rtt := float64(pr.RTT) / float64(time.Millisecond)
				lrtt := "0ms (from cache)"
				if pr.Server.HasGlue {
					lrtt = "0ms (from glue)"
				} else if pr.Server.LookupRTT > 0 {
					lrtt = fmt.Sprintf("%.2fms", float64(pr.Server.LookupRTT)/float64(time.Millisecond))
				}
				log(args.VerboseDnssec, col("  - %d bytes in %.2fms + %s lookup on %s(%s)", cDarkGray), ln, rtt, lrtt, pr.Server.Name, pr.Addr)
				if pr.Err != nil {
					err := pr.Err
					if oerr, ok := err.(*net.OpError); ok {
						err = oerr.Err
					}
					log(args.VerboseDnssec, ": %v", col(err, cRed))
				}
				log(args.VerboseDnssec, "\n")
			}

			switch rtype {
			case client.ResponseTypeDelegation:
				var label string
				for _, rr := range r.Ns {
					if ns, ok := rr.(*dns.NS); ok {
						label = ns.Header().Name
						break
					}
				}
				_, ns := c.DCache.Get(label)
				for _, s := range ns {
					var glue string
					if s.HasGlue {
						glue = col("glue: "+strings.Join(s.Addrs, ","), cDarkGray)
					} else {
						glue = col("no glue", cYellow)
					}
					log(args.VerboseDnssec, "%s %d NS %s (%s)\n", label, s.TTL, s.Name, glue)
				}
			case client.ResponseTypeCNAME:
				for _, rr := range r.Answer {
					log(args.VerboseDnssec, rr)
				}
			}
		},
		FollowingCNAME: func(domain, target string) {
			log(args.VerboseDnssec, col("\n~ following CNAME %s -> %s\n", cBlue), domain, target)
		},
	}
	r, rtt, err := c.RecursiveQuery(m, t)
	if err != nil {
		log(args.VerboseDnssec, col("*** error: %v\n", cRed), err)
		return nil
	}

	log(args.VerboseDnssec)
	log(args.VerboseDnssec, col(";; Cold best path time: %s\n\n", cGray), rtt)
	for i, rr := range r.Answer {
		log(args.VerboseDnssec, "r.Answer =", i, rr, args.VerboseDnssec)
	}
	return r.Answer
	// for _, rr := range r.Answer {
	// 	return rr
	// }
	// return nil
}

func ResolveIPv6hostname(hostname string) *net.TCPAddr {
	dnsRR := Dnstrace(hostname, "AAAA")
	if (dnsRR == nil) {
		return nil
	}
	aaaa := dns.Field(dnsRR[1], 1)
	localTCPAddr, _  := net.ResolveTCPAddr("tcp", aaaa)
	return localTCPAddr
}

func UseJournalctl() {
	log(args.VerboseDnssec, "journalhook is disabled because it breaks the Windows build right now")
        // journalhook.Enable()
}

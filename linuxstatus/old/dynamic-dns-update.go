package main

/* 
	https://pkg.go.dev/github.com/miekg/dns#section-readme

DYNAMIC UPDATES

Dynamic updates reuses the DNS message format, but renames three of the sections. Question is Zone, Answer is Prerequisite, Authority is Update, only the Additional is not renamed. See RFC 2136 for the gory details.

You can set a rather complex set of rules for the existence of absence of certain resource records or names in a zone to specify if resource records should be added or removed. The table from RFC 2136 supplemented with the Go DNS function shows which functions exist to specify the prerequisites.

3.2.4 - Table Of Metavalues Used In Prerequisite Section

 CLASS    TYPE     RDATA    Meaning                    Function
 --------------------------------------------------------------
 ANY      ANY      empty    Name is in use             dns.NameUsed
 ANY      rrset    empty    RRset exists (value indep) dns.RRsetUsed
 NONE     ANY      empty    Name is not in use         dns.NameNotUsed
 NONE     rrset    empty    RRset does not exist       dns.RRsetNotUsed
 zone     rrset    rr       RRset exists (value dep)   dns.Used

The prerequisite section can also be left empty. If you have decided on the prerequisites you can tell what RRs should be added or deleted. The next table shows the options you have and what functions to call.

3.4.2.6 - Table Of Metavalues Used In Update Section

 CLASS    TYPE     RDATA    Meaning                     Function
 ---------------------------------------------------------------
 ANY      ANY      empty    Delete all RRsets from name dns.RemoveName
 ANY      rrset    empty    Delete an RRset             dns.RemoveRRset
 NONE     rrset    rr       Delete an RR from RRset     dns.Remove
 zone     rrset    rr       Add to an RRset             dns.Insert
*/

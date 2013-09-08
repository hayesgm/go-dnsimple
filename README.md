go-dnsimple
===========

Unofficial DNSimple API library in Go

# Installation

    import "github.com/hayesgm/go-dnsimple/dnsimple"

# Usage

First, you'll need to create a client with your authorization scheme.
    
    // Token Authorization
    cli := &dnsimple.Client{dnsimple.NewTokenAuth(email, token)}

    // Domain Authorization
    cli := &dnsimple.Client{dnsimple.NewDomainAuth(domain, token)}

Next, run API commands:

    var records []dnsimple.RecordObj
    records, err := cli.GetRecords("example.com", "www") // returns []RecordObj

    record dnsimple.RecordObj
    record, err = cli.CreateRecord("example.com", "admin", dnsimple.TXT_RECORD, "geoff@example.com", 360, 10)

    err = cli.DeleteRecord("example.com", "admin", record.Record.Id)

# TODOs

* Implement more functions from [API docs](http://developer.dnsimple.com/)
* Make sure error handling works as expected
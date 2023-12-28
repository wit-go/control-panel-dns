# control-panel-dns

A Control Panel to monitor DNS settings

Goals:

* Correctly care and handle IPv6 (IPv4 is dead)
* Update your hostname DNS when your IP address changes
* Run as a daemon
* When run in GUI, add status via systray

# Rational

With the advent of IPv6, it is finally possible again to have real hostnames for
your machines, desktops, laptops, vm's, etc. This control panel will poll for
changes, find out what the DNS entries are, then, if they are not correct, attempt
to update the DNS server.

## References

Useful links and other
external things which might be useful

* [DNS Resource Record Types](https://en.wikipedia.org/wiki/List_of_DNS_record_types)
* [WIT GO projects](http://go.wit.com/)
* [GOLANG GUI](https://go.wit.com/gui)
* [GO Style Guide](https://google.github.io/styleguide/go/index)

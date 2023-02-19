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

* [WIT GO projects](http://go.wit.org/)
* [GOLANG GUI](https://github.com/wit-go/gui)
* [GO Style Guide](https://google.github.io/styleguide/go/index)

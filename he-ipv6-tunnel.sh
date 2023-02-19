#!/bin/bash -x

## Tunnel ID: 818143
# Creation Date:Feb 12, 2023
# Description:
# IPv6 Tunnel Endpoints
# Server IPv4 Address:184.105.253.14
# Server IPv6 Address:2001:470:1f10:2a::1/64
# Client IPv4 Address:74.87.91.117
# Client IPv6 Address:2001:470:1f10:2a::2/64
# Routed IPv6 Prefixes
# Routed /64:2001:470:1f11:2a::/64
# Routed /48:Assign /48
# DNS Resolvers
# Anycast IPv6 Caching Nameserver:2001:470:20::2
# Anycast IPv4 Caching Nameserver:74.82.42.42
# DNS over HTTPS / DNS over TLS:ordns.he.net
# rDNS DelegationsEdit
# rDNS Delegated NS1:
# rDNS Delegated NS2:
# rDNS Delegated NS3:
# rDNS Delegated NS4:
# rDNS Delegated NS5:

# ifconfig sit0 up
# ifconfig sit0 inet6 tunnel ::184.105.253.14
# ifconfig sit1 up
# ifconfig sit1 inet6 add 2001:470:1f10:2a::2/64
# route -A inet6 add ::/0 dev sit1

if [ "$1" = "down" ]; then
	ip tunnel del he-ipv6
	rmmod sit
	exit
fi

if [ "$1" = "ping" ]; then
	ping -c 3 2001:470:1f10:13d::1
	exit
fi

modprobe ipv6
ip tunnel add he-ipv6 mode sit remote 184.105.253.14 local 40.132.180.131 ttl 255
ip link set he-ipv6 up
ip addr add 2001:470:1f10:13d::2/64 dev he-ipv6
ip route add ::/0 dev he-ipv6
ip -f inet6 addr
ifconfig he-ipv6 mtu 1460


# old attempt from the something or pabtz hotel
# modprobe ipv6
# ip tunnel add he-ipv6 mode sit remote 184.105.253.14 local 74.87.91.117 ttl 255
# ip link set he-ipv6 up
# ip addr add 2001:470:1f10:2a::2/64 dev he-ipv6
# ip route add ::/0 dev he-ipv6
# ip -f inet6 addr

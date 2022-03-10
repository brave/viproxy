VIProxy
=======

The VIProxy package implements a TCP proxy that translates between AF\_INET and
AF\_VSOCK connections.  The proxy takes as input two addresses, one being
AF\_INET and the other being AF\_VSOCK.  The proxy then starts a TCP listener on
the in-address and once it receives an incoming connection to the in-address, it
establishes a TCP connection to the out-addresses.  Once both connections are
established, the proxy copies data back and forth.

The [example](example) directory contains a simple example of how one would use
viproxy.

===========
networktest
===========

.. image:: https://api.travis-ci.org/numerodix/networktest.png?branch=master
    :target: https://travis-ci.org/numerodix/networktest

.. image:: https://ci.appveyor.com/api/projects/status/ojcojhf1wlc837ug/branch/master?svg=true
    :target: https://ci.appveyor.com/project/numerodix/networktest

.. image:: http://codecov.io/github/numerodix/networktest/coverage.svg?branch=master
    :target: https://codecov.io/github/numerodix/networktest

This tool answers the question: *am I connected to a (local) network and do I
have internet connectivity?*




How does it work?
=================

While the question *"can this computer reach the Internet?"* is quite easily
answered by trying to load ``http://www.google.com`` in a browser, common
network configurations have a number of parameters that may vary between one
case of *"it's not working"* and the next.

``networktest`` aims to be a tool that makes it easy to establish the key
network configuration characteristics and get a quick idea of what might be
wrong. It is a smoke test of your network setup.

The questions we ask are:

1. **Am I connected to any networks?** If so, what kind of network are they?
   Loopback interface? Link-local? Or is it a network that allows me to route
   packets to the Internet (the kind we want)?

2. **What ip addresses do I have on the networks I'm connected to?** These
   are addresses that should be reachable (ie. pingable) from myself.

3. **What gateways do I have?** A gateway (or router) is a host on my network
   that allows me to route packets to other networks.

4. **Can I reach Internet hosts by ip?** If so, my gateway is working.

5. **What DNS servers (nameservers) do I have?** DNS servers allow me to
   resolve Internet hostnames to ip addresses.

6. **Can I reach Internet hosts by hostname?** If so, I have "Internet
   connectivity" in the common sense.




Downloading binaries
====================

The easiest way to use ``networktest`` is to download a `binary release
<https://github.com/numerodix/networktest/releases>`_.

You will need to know whether you need the 32bit or the 64bit version,
depending on your operating system. Most operating systems these days are
64bit.

``networktest`` is written in Go and targets Linux, FreeBSD, OS X (Darwin), and
Windows.




Building from source
====================

Tools required to build from source:

* go >= 1.5
* git >= 1.8.1.6
* python >= 2.6
* make (optional)

If you have make installed just use that::
    
    $ make

If you don't (eg. Windows)::
    
    $ python build.py --build

You can find the binary in ``bin/havenet``.



Usage
=====


To detect IPv4 networking (working internet connection)::

    $ havenet
    Platform: Linux
     + Scanning for networks...
        <lo>              127.0.0.0 / 255.0.0.0        
        <docker0>        172.17.0.0 / 255.255.0.0      
        <eth0>          192.168.1.0 / 255.255.255.0    
        <wlan0>         192.168.1.0 / 255.255.255.0    
     + Detecting ips...
        <lo>              127.0.0.1 / 255.0.0.0        ping: 0.059 ms
        <docker0>        172.17.0.1 / 255.255.0.0      ping: 0.066 ms
        <eth0>          192.168.1.6 / 255.255.255.0    ping: 0.065 ms
        <wlan0>        192.168.1.10 / 255.255.255.0    ping: 0.052 ms
     + Detecting gateways...
        <eth0>          192.168.1.1  ping: 0.746 ms
         ip:            192.168.1.6
         ip:           192.168.1.10
     + Testing internet connection...
        b.root-servers.net.   192.228.79.201  ping: 177.8 ms
     + Detecting dns servers...
                8.8.8.8  ping: 41.79 ms
     + Testing internet dns...
             debian.org  ping: 111.9 ms
           facebook.com  ping: 133.5 ms
              gmail.com  ping: 39.29 ms
             google.com  ping: 83.51 ms
              yahoo.com  ping: 174.4 ms

To detect IPv6 networking (no internet connection)::

    $ havenet -6
     + Scanning for networks...
        <lo>                                            ::1 / 128  [scope: host]
        <eth0>                                       fe80:: /  64  [scope: link]
        <wlan0>                                      fe80:: /  64  [scope: link]
     + Detecting ips...
        <lo>                                            ::1 / 128  ping: 0.047 ms
        <eth0>                    fe80::16da:fae1:c9ea:a4b9 /  64  ping: N/A
        <wlan0>                   fe80::762f:fe64:b7c7:7b7a /  64  ping: N/A
     + Detecting gateways...
        none found

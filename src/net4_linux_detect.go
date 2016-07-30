package main

import "fmt"
import "net"
import "regexp"
import "strings"


type LinuxNetDetect4 struct {
    ctx AppContext
}


func NewLinuxNetDetect4(ctx AppContext) LinuxNetDetect4 {
    return LinuxNetDetect4{
        ctx: ctx,
    }
}


func (lnd LinuxNetDetect4) detectNetConn4() IPNetworkInfo {
    var info = IPNetworkInfo{}

    lnd.detectIpAddr4(&info)
    lnd.detectIpRoute4(&info)

    var und = NewUnixNetDetect4(lnd.ctx)
    und.detectNsHosts4(&info)

    return info
}


func (lnd LinuxNetDetect4) detectIpAddr4(info *IPNetworkInfo) {
    var mgr = ProcMgr("ip", "-4", "addr", "show")
    var res = mgr.run()

    // The command failed :(
    if res.err != nil {
        lnd.ctx.ft.printError("Failed to detect ipv4 network", res.err)
        return
    }

    // Extract the output
    lnd.parseIpAddr4(res.stdout, info)

    // Parsing failed :(
    lnd.ctx.ft.printErrors("Failed to parse ipv4 network info", info.Errs)
}

func (lnd LinuxNetDetect4) detectIpRoute4(info *IPNetworkInfo) {
    var mgr = ProcMgr("ip", "-4", "route", "show")
    var res = mgr.run()

    // The command failed :(
    if res.err != nil {
        lnd.ctx.ft.printError("Failed to detect ipv4 routes", res.err)
        return
    }

    // Extract the output
    lnd.parseIpRoute4(res.stdout, info)

    // Parsing failed :(
    lnd.ctx.ft.printErrors("Failed to parse ipv4 route info", info.Errs)
}


func (lnd LinuxNetDetect4) parseIpAddr4(stdout string, info *IPNetworkInfo) {
    // We will read line by line
    var lines = strings.Split(stdout, "\n")

    // Prepare regex objects
    rxIface := regexp.MustCompile("^[0-9]+: ([^ ]+):")
    rxInet := regexp.MustCompile("^[ ]{4}inet ([0-9.]+)[/]([0-9]+)")

    // Loop variables
    var iface = ""
    var ip = ""
    var maskBits = ""

    for _, line := range lines {
        if rxIface.MatchString(line) {
            iface = rxIface.FindStringSubmatch(line)[1]
        }

        if rxInet.MatchString(line) {
            ip = rxInet.FindStringSubmatch(line)[1]
            maskBits = rxInet.FindStringSubmatch(line)[2]

            var ipNet = fmt.Sprintf("%s/%s", ip, maskBits)
            var ipobj, ipnet, err = net.ParseCIDR(ipNet)

            // Parse failed
            if err != nil {
                info.Errs = append(info.Errs, err)
                continue
            }

            var netMasked = applyMask(&ipnet.IP, &ipnet.Mask)
            var mask = ipnetMaskAsIP(ipnet)

            // Populate info
            info.Nets = append(info.Nets, Network{
                Iface: Interface{Name: iface},
                Ip: netMasked,
            })
            info.Ips = append(info.Ips, IpAddr{
                Iface: Interface{Name: iface},
                Ip: ipobj,
                Mask: mask,
            })
        }
    }
}


func (lnd LinuxNetDetect4) parseIpRoute4(stdout string, info *IPNetworkInfo) {
    // We will read line by line
    var lines = strings.Split(stdout, "\n")

    // Prepare regex objects
    rxIface := regexp.MustCompile("^default via ([0-9.]+) dev ([^ ]+)")

    // loop variables
    var iface = ""
    var ip = ""

    for _, line := range lines {
        if rxIface.MatchString(line) {
            ip = rxIface.FindStringSubmatch(line)[1]
            iface = rxIface.FindStringSubmatch(line)[2]

            var ipobj = net.ParseIP(ip)

            // Populate info
            info.Gws = append(info.Gws, Gateway{
                Iface: Interface{Name: iface},
                Ip: ipobj,
            })
        }
    }
}

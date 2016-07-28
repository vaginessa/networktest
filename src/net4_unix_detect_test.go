package main

import "testing"


const resolvConfContent = `
# Dynamic resolv.conf(5) file for glibc resolver(3) generated by resolvconf(8)
#     DO NOT EDIT THIS FILE BY HAND -- YOUR CHANGES WILL BE OVERWRITTEN
nameserver 8.8.8.8
nameserver 8.8.4.4
`


func Test_unixParseResolvConf4(t *testing.T) {
    var info = IPNetworkInfo{}

    var ctx = TestAppContext()
    var detector = NewUnixNetDetect4(ctx)
    detector.parseResolvConf4(resolvConfContent, &info)
    info.normalize()

    // Errors
    assertIntEq(t, 0, len(info.Errs), "Errs does not match")

    // Ns hosts
    assertIntEq(t, 2, len(info.NsHosts), "wrong number of dns servers")

    assertStrEq(t, "8.8.8.8", info.NsHosts[0].Ip.String(), "Ip does not match")
    assertStrEq(t, "8.8.4.4", info.NsHosts[1].Ip.String(), "Ip does not match")
}

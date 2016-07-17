package main

import (
    "fmt"
    "net"
)


type Formatter struct {
    colorBrush ColorBrush
}


func (ft *Formatter) FormatPingTime(pingExec PingExecution) string {
    if pingExec.Error != nil {
        return ft.colorBrush.red("failed")
    }

    var time = fmt.Sprintf("%.3f", pingExec.Time)
    time = time[:5]  // four significant digits + decimal point
    time = time + " ms"
    return ft.colorBrush.green(time)
}


func (ft *Formatter) FormatHeader(msg string) string {
    var msgFmt = ft.colorBrush.yellow(fmt.Sprintf(" + %s...", msg))
    return msgFmt
}

func (ft *Formatter) FormatError(msg string) string {
    var msgFmt = ft.colorBrush.red(fmt.Sprintf("%s", msg))
    return msgFmt
}

func (ft *Formatter) FormatIfaceField(iface string) string {
    iface = fmt.Sprintf("<%s>", iface)
    var ifaceFmt = ft.colorBrush.magenta(fmt.Sprintf("%-10s", iface))
    return ifaceFmt
}

func (ft *Formatter) FormatHostField(host string) string {
    var hostFmt = ft.colorBrush.cyan(fmt.Sprintf("%s", host))
    return hostFmt
}

func (ft *Formatter) FormatLanIpField(ip string) string {
    var ipFmt = ft.colorBrush.bgreen(fmt.Sprintf("%15s", ip))
    return ipFmt
}

func (ft *Formatter) FormatIpField(ip string) string {
    var ipFmt = ft.colorBrush.green(fmt.Sprintf("%15s", ip))
    return ipFmt
}

func (ft *Formatter) FormatIp6Field(ip net.IP) string {
    var ipFmt = ft.colorBrush.green(fmt.Sprintf("%s", ip))
    return ipFmt
}

func (ft *Formatter) FormatMask6Field(mask net.IPMask) string {
    var ones, _ = mask.Size()
    var subnetFmt = ft.colorBrush.cyan(fmt.Sprintf("/ %d", ones))
    return subnetFmt
}

func (ft *Formatter) FormatScope6Field(scope string) string {
    var ipFmt = ft.colorBrush.yellow(fmt.Sprintf("[scope: %s]", scope))
    return ipFmt
}

func (ft *Formatter) FormatSubnetField(ip string) string {
    var subnetFmt = ft.colorBrush.cyan(fmt.Sprintf("/ %-15s", ip))
    return subnetFmt
}
package netutils

import (
	"net"
)

// LocalIPv4s Get the IPv4 address on a locally enabled network interface
func LocalIPv4s() ([]string, error) {
	var ips []string
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return ips, err
	}

	// 遍历每个接口，获取相关的IP地址
	for _, iface := range interfaces {
		//fmt.Printf("网络接口: %s\n", iface.Name)
		// 检查网络接口是否启用，未启用则跳过
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		// 获取接口的地址列表
		addrs, err := iface.Addrs()
		if err != nil {
			//fmt.Printf("获取接口 %s 地址失败: %v\n", iface.Name, err)
			continue
		}

		// 遍历每个地址
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 过滤掉回环地址（如127.0.0.1） || ip.IsLoopback()
			if ip == nil {
				continue
			}

			// 打印IPv4或IPv6地址
			if ip.To4() != nil {
				//fmt.Printf("\tIPv4: %s\n", ip.String())
				ips = append(ips, ip.String())
			} else if ip.To16() != nil {
				//fmt.Printf("\tIPv6: %s\n", ip.String())
			}
		}
	}
	return ips, err
}

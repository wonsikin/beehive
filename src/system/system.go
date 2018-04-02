package system

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

// Information represents system information of current server
type Information struct {
	HostName string
	IP       string
	LogPath  string
}

// NewInformation returns a new system information
func NewInformation(path string) (*Information, error) {
	info := &Information{}
	name, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	info.HostName = name

	ip, err := externalIP()
	if err != nil {
		return nil, err
	}
	info.IP = ip

	st, err := logPath(path)
	if err != nil {
		return nil, err
	}
	info.LogPath = st

	return info, nil
}

func logPath(path string) (string, error) {
	tpl := "^\\./.*"
	re, err := regexp.Compile(tpl)
	if err != nil {
		return "", err
	}

	if matched := re.MatchString(path); !matched {
		return path, nil
	}

	dir, err := currentDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", dir, path[1:]), nil
}

func currentDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

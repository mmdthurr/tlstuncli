package main

import (
	"flag"
	"mmd/tlstuncli/tunnel"
	"strconv"
)

func main() {
	raddr := flag.String("r", "127.0.0.1:443", "remote addr")
	stP := flag.Int("port", 5000, "starting port")
	v2P := flag.String("v2p", "1086", "v2ray port")
	connc := flag.Int("c", 10, "amount of connections")
	passwd := flag.String("passwd", "123456", "tunnel passwd")
	flag.Parse()

	// var clis = make(map[int]tunnel.Cli)

	for p := *stP; p < (*stP + *connc); p++ {

		go tunnel.Cli{
			RemoteAddr: *raddr,
			ExposePort: strconv.Itoa(p),
			Passwd:     *passwd,
			Bckp:       *v2P,
		}.StartCli()

	}

}

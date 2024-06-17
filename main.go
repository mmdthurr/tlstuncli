package main

import (
	"flag"
	"mmd/tlstuncli/tunnel"
	"strconv"
	"sync"
	"time"
)

func main() {
	raddr := flag.String("r", "127.0.0.1:443", "remote addr")
	stP := flag.Int("port", 5000, "starting port")
	v2P := flag.String("v2p", "1086", "v2ray port")
	connc := flag.Int("c", 10, "amount of connections")
	passwd := flag.String("passwd", "123456", "tunnel passwd")
	flag.Parse()

	// var clis = make(map[int]tunnel.Cli)
	var wg sync.WaitGroup
	for p := *stP; p < (*stP + *connc); p++ {
		wg.Add(1)
		go func(p int, remoteaddr, passwd, v2port string) {
			for {
				tunnel.Cli{
					RemoteAddr: remoteaddr,
					ExposePort: strconv.Itoa(p),
					Passwd:     passwd,
					Bckp:       v2port,
				}.StartCli()
				time.Sleep(2 * time.Second)
			}
		}(p, *raddr, *passwd, *v2P)

	}
	wg.Wait()

}

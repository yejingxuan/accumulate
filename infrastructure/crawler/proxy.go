package crawler

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func randomProxySwitcher() colly.ProxyFunc {
	switcher, _ := proxy.RoundRobinProxySwitcher(
		"http://175.7.199.55",
		"http://175.7.199.170",
		"http://175.7.199.198",
		/*"socks5://127.0.0.1:1337",
		"socks5://127.0.0.1:1338",
		"http://127.0.0.1:8080",*/
	)
	return switcher
}

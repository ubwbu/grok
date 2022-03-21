package main

import (
	"fmt"

	"github.com/ubwbu/grok"
)

func main() {
	de, errs := grok.DenormalizePatternsFromMap(grok.CopyDefalutPatterns())
	if len(errs) != 0 {
		fmt.Print(errs)
		return
	}
	g, err := grok.CompilePattern("%{COMMONAPACHELOG}", de)
	if err != nil {
		fmt.Print(err)
	}
	ret, err := g.Run(`127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
	if err != nil {
		fmt.Print(err)
	}
	for k, v := range ret {
		fmt.Printf("%+15s: %s\n", k, v)
	}
}

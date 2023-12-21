package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/katana/pkg/engine/standard"
	"github.com/projectdiscovery/katana/pkg/output"
	"github.com/projectdiscovery/katana/pkg/types"
)

func main() {

	options := &types.Options{
		MaxDepth:     1,               // Maximum depth to crawl
		FieldScope:   "rdn",           // Crawling Scope Field
		BodyReadSize: 2 * 1024 * 1024, // Maximum response size to read
		RateLimit:    150,             // Maximum requests to send per second
		Strategy:     "depth-first",   // Visit strategy (depth-first, breadth-first)
		OnResult: func(result output.Result) { // Callback function to execute for result
			gologger.Info().Msg(result.Request.URL)
		},
	}
	crawlerOptions, err := types.NewCrawlerOptions(options)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	defer crawlerOptions.Close()
	crawler, err := standard.New(crawlerOptions)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	defer crawler.Close()
	var input = "https://tesla.com"
	err = crawler.Crawl(input)
	if err != nil {
		gologger.Warning().Msgf("Could not crawl %s: %s", input, err.Error())
	}
}

# Onyxium Market Feed
Market feed is a service to fetch the latest market data. It currently polls the bittrex api every second for market data and stores in a mongo db.

# Getting started
Install go: https://golang.org/dl/
Make sure your $GOPATH is correct. 
`go get https://github.com/onyxium-strategies/onyxium-market-feed.git`
`cd $GOPATH/bin`
`./onyxium-market-feed`

# Development
Dependancy management during development is done with [golang/dep](https://golang.github.io/dep/docs/introduction.html).
**Note* that this is not intended for end users who are installing Go software - that's what `go get` does.

Install with homebrew: `brew install dep`
Pull all dependencies: `dep ensure`
Add new dependency to project: `dep ensure -add github.com/foo/bar`


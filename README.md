# Coinflow Market Feed
Market feed is a service to fetch the latest market data. It currently polls the bittrex api every second for market data and stores in a mongo db.

# Getting started
Install go: https://golang.org/dl/
Make sure your $GOPATH is correct. In order to clone a private bitbucket repo we need to enable default cloning with SSH instead of https.  
`git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"`  
`go get bitbucket.org/visa-startups/coinflow-market-feed`  
`cd $GOPATH/bin`  
`./coinflow-market-feed`

# Development
Dependancy management during development is done with [golang/dep](https://golang.github.io/dep/docs/introduction.html).
**Note* that this is not intended for end users who are installing Go software - that's what `go get` does.

Install with homebrew: `brew install dep`  
Pull all dependencies: `dep ensure`  
Add new dependency to project: `dep ensure -add github.com/foo/bar`  


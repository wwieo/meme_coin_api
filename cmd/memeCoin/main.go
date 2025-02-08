package main

import "meme_coin_api/service"

// @title		Meme Coin API
// @version	1.0
// @basePath	/meme_coin
// @schemes	http
func main() {
	service.MemeCoin().Run()
}

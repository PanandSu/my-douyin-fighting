package main

import "my-douyin-fighting/initial"

func main() {
	initial.Global()
	initial.Viper()
	initial.Mysql()
	initial.Redis()
	initial.Route()
}

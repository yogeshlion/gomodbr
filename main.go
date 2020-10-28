package main
import "github.com/hashicorp/consul/api"

func main(){
	addr:=127.0.0.0:8500
	config := consul.DefaultConfig()
	config.Address:=addr
}
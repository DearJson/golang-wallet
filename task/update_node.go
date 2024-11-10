package task

import (
	"gfast/app/common/service"
)

// UpdateNode 更新节点
func UpdateNode() {
	//node := g.Client().Timeout(3 * time.Second).GetContent("http://119.23.187.205/getNodes")
	//if node != "" {
	cache := service.Cache.New()
	cache.Set("bnb_rpc_url", "https://bsc-dataseed1.binance.org/", 0)
	//}
	cache.Set("tron_rpc_url", "grpc.trongrid.io:50051", 0)
	cache.Set("eth_rpc_url", "https://eth.llamarpc.com", 0)
	//tronNode := g.Client().Timeout(3 * time.Second).GetContent("http://119.23.187.205/getTronNodes")
	//if tronNode != "" {
	//	cache := service.Cache.New()
	//	cache.Set("tron_rpc_url", tronNode, 0)
	//}
	//
	//hecoNode := g.Client().Timeout(3 * time.Second).GetContent("http://119.23.187.205/getHecoNodes")
	//if tronNode != "" {
	//	cache := service.Cache.New()
	//	cache.Set("heco_rpc_url", hecoNode, 0)
	//}
	//
	//nacNode := g.Client().Timeout(3 * time.Second).GetContent("http://119.23.187.205/getNacNodes")
	//if nacNode != "" {
	//	cache := service.Cache.New()
	//	cache.Set("nac_rpc_url", nacNode, 0)
	//}
	//
	//cache := service.Cache.New()
	//cache.Set("wemix_rpc_url", "https://api.wemix.com/", 0)
}

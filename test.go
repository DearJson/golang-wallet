package main

import (
	"fmt"
	"gfast/library"
)

func main() {

	//privateKey, err := crypto.GenerateKey()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("fail")
	//}
	//address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//hash := sha3.NewLegacyKeccak256()
	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//hash.Write(publicKeyBytes[1:])
	//address = hexutil.Encode(hash.Sum(nil)[12:])
	//
	//privateKey111 := hexutil.Encode(privateKeyBytes)[2:]

	//fmt.Printf("address = %v, privateKye= %v", address, privateKey111)
	//
	fmt.Println(library.NewGoogleAuth().GetSecret())

	//tronPrivateKey := "b69d5e0c19291ae779b4c63f92f4b8228f7d6190058d3632e8606774d976f26e"
	//tronAddress := "TQ3TQvZRUusqHa669QCo3tsH7ntLq5nLT9"
	//
	//tronClient, _ := rpc.NewClient("grpc.trongrid.io:50051")
	//txId, _ := tronClient.TransferTrx(tronPrivateKey, tronAddress, "TXumrPcA9Avbd97bNM2uEUUh1d11EFRDC9", decimal.NewFromFloat(0.00001), "我的\n我的\n你的\n你的")
	//
	//fmt.Printf("交易hash = %v", txId)

}

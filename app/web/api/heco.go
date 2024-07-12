package api

import (
	"crypto/ecdsa"
	"gfast/app/common/global"
	cservice "gfast/app/common/service"
	"gfast/app/system/dao"
	"gfast/app/system/service"
	"gfast/library"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"golang.org/x/crypto/sha3"
	"math/rand"
	"strings"
	"time"
)

type heco struct {
	WebBase
}

var Heco = new(heco)

// GenerateAddress 生成币安地址
func (b *heco) GenerateAddress(r *ghttp.Request) {
	var req *dao.AddressAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	//判断一下该用户标识的bsc链是否存在
	exists, err := service.Address.GetInfoByUserId(r.GetCtx(), req.UserId, "heco")
	if exists != nil {
		result := make(map[string]string)
		result["user_id"] = exists.UserId
		result["address"] = exists.Address
		b.SusJsonExit(r, result)
	}

	//开始生成钱包地址
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		b.FailJsonExit(r, "生成失败")
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		b.FailJsonExit(r, "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	hash := sha3.NewLegacyKeccak256()
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash.Write(publicKeyBytes[1:])
	address = hexutil.Encode(hash.Sum(nil)[12:])

	req.MainChain = "heco"
	req.PrivateKey, _ = library.EncryptByAes(gconv.Bytes(hexutil.Encode(privateKeyBytes)[2:]))
	req.Address = address
	err = service.Address.Add(r.GetCtx(), req)
	if err != nil {
		b.FailJsonExit(r, err.Error())
	}
	result := make(map[string]string)
	result["user_id"] = req.UserId
	result["address"] = req.Address
	b.SusJsonExit(r, result)
}

// Withdraw 提现代币
func (b *heco) Withdraw(r *ghttp.Request) {
	var req *dao.WithdrawAddReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		b.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	if req.Sign != "TJH4uf6kmHPdQoOfALagrBBguvLgGqnUpH4UmKsbVBcjwfcHqmLfGGxEIAet" {
		b.FailJsonExit(r, "非法请求")
	}
	//检查币种是否存在
	coinInfo, _ := service.Currency.GetInfoByContractAddress(r.GetCtx(), req.ContractAddress, "heco")
	if coinInfo == nil {
		b.FailJsonExit(r, "暂未配置该币种,无法转账")
	}

	minAmount, _ := cservice.SysConfig.GetConfigByKey("sys.minWithdrawAudit")
	if req.Amount <= gconv.Float64(minAmount.ConfigValue) {
		req.Status = 2
	} else {
		req.Status = 1
	}
	req.MainChain = "heco"
	req.CoinToken = coinInfo.Name
	req.Address = strings.ToLower(req.Address)
	rand.Seed(time.Now().UnixNano())
	req.Nonce1 = gconv.String(rand.Int())
	req.HashKey = library.Md5Data(req.Address, req.ContractAddress, req.Amount, req.Status, req.Nonce1)
	err := service.Withdraw.Add(r.GetCtx(), req)
	if err != nil {
		b.FailJsonExit(r, err.Error())
	}
	b.SusJsonExit(r)
}

//SetAddress 设置监控地址
func (b *heco) SetAddress(r *ghttp.Request) {
	address := strings.ToLower(gconv.String(r.GetPost("address")))
	count, _ := g.Model("address").Where("main_chain", "heco").Where("address", address).FindCount()
	if count > 0 {
		b.FailJsonExit(r, "地址已存在")
	}
	g.Model("address").Data(g.Map{"main_chain": "heco", "user_id": gconv.String(r.GetPost("user_id")), "address": address}).Insert()
	cache := cservice.Cache.New()
	cache.Remove(global.BscUserAddressList)
	cache.Remove(global.TronUserAddressList)
	cache.Remove(global.HecoUserAddressList)
	b.SusJsonExit(r)
}

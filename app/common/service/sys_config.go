/*
* @desc:系统参数设置
* @company:NULL
* @Author: Json   wbjson@gmail.com
* @Date:   2021/7/5 18:00
 */

package service

import (
	"gfast/app/common/dao"
	"gfast/app/common/global"
	"gfast/app/common/model"
	"gfast/library"
	"strings"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type sysConfig struct {
}

var SysConfig = new(sysConfig)

func (s *sysConfig) SelectListByPage(req *model.SysConfigSearchReq) (total, page int, list []*model.SysConfig, err error) {
	m := dao.SysConfig.Ctx(req.Ctx)
	if req != nil {
		if req.ConfigName != "" {
			m = m.Where("config_name like ?", "%"+req.ConfigName+"%")
		}
		if req.ConfigType != "" {
			m = m.Where("config_type = ", gconv.Int(req.ConfigType))
		}
		if req.ConfigKey != "" {
			m = m.Where("config_key like ?", "%"+req.ConfigKey+"%")
		}
		if req.BeginTime != "" {
			m = m.Where("create_time >= ? ", req.BeginTime)
		}

		if req.EndTime != "" {
			m = m.Where("create_time<=?", req.EndTime)
		}
	}
	total, err = m.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
		return
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	page = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = model.PageSize
	}
	err = m.Page(page, req.PageSize).Order("config_id asc").Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	// 私钥字段脱敏，不返回给前端
	for _, item := range list {
		if strings.Contains(item.ConfigKey, "PrivateKey") && item.ConfigValue != "" {
			item.ConfigValue = "******"
		}
	}
	return
}

// CheckConfigKeyUniqueAll 验证参数键名是否存在
func (s *sysConfig) CheckConfigKeyUniqueAll(configKey string) error {
	entity, err := dao.SysConfig.Fields(dao.SysConfig.C.ConfigId).FindOne(dao.SysConfig.C.ConfigKey, configKey)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("校验数据失败")
	}
	if entity != nil {
		return gerror.New("参数键名已经存在")
	}
	return nil
}

// AddSave 添加操作
func (s *sysConfig) AddSave(req *model.SysConfigAddReq) (err error) {
	_, err = dao.SysConfig.Insert(req)
	return
}

func (s *sysConfig) GetById(id int) (data *model.SysConfig, err error) {
	err = dao.SysConfig.WherePri(id).Scan(&data)
	// 私钥字段脱敏，不返回给前端
	if data != nil && strings.Contains(data.ConfigKey, "PrivateKey") && data.ConfigValue != "" {
		data.ConfigValue = "******"
	}
	return
}

// CheckConfigKeyUnique 检查键是否已经存在
func (s *sysConfig) CheckConfigKeyUnique(configKey string, configId int64) error {
	entity, err := dao.SysConfig.Fields(dao.SysConfig.C.ConfigId).
		FindOne(dao.SysConfig.C.ConfigKey+"=? and "+dao.SysConfig.C.ConfigId+"!=?",
			configKey, configId)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("校验数据失败")
	}
	if entity != nil {
		return gerror.New("参数键名已经存在")
	}
	return nil
}

// EditSave 修改系统参数
func (s *sysConfig) EditSave(req *model.SysConfigEditReq) (err error) {
	// 私钥字段：如果前端传的是脱敏值"******"，说明未修改，跳过更新该字段
	if strings.Contains(req.ConfigKey, "PrivateKey") && req.ConfigValue == "******" {
		_, err = dao.SysConfig.FieldsEx(dao.SysConfig.C.ConfigId, dao.SysConfig.C.CreateBy, dao.SysConfig.C.ConfigValue).
			WherePri(req.ConfigId).Data(req).Update()
		return
	}
	if req.ConfigKey == "sys.bnbWithdrawAddressPrivateKey" || req.ConfigKey == "sys.bnbFeeAddressPrivateKey" ||
		req.ConfigKey == "sys.tronWithdrawAddressPrivateKey" || req.ConfigKey == "sys.tronFeeAddressPrivateKey" ||
		req.ConfigKey == "sys.hecoWithdrawAddressPrivateKey" || req.ConfigKey == "sys.hecoFeeAddressPrivateKey" ||
		req.ConfigKey == "sys.wemixWithdrawAddressPrivateKey" || req.ConfigKey == "sys.wemixFeeAddressPrivateKey" ||
		req.ConfigKey == "sys.ethWithdrawAddressPrivateKey" || req.ConfigKey == "sys.ethFeeAddressPrivateKey" {
		//如果私钥带有前面带有0x去掉
		if req.ConfigValue[0:2] == "0x" {
			req.ConfigValue = req.ConfigValue[2:]
		}
		req.ConfigValue, _ = library.EncryptByAes(gconv.Bytes(req.ConfigValue))
	}
	// Solana私钥加密存储（Ed25519私钥为纯hex，无0x前缀）
	if req.ConfigKey == "sys.solanaWithdrawAddressPrivateKey" || req.ConfigKey == "sys.solanaFeeAddressPrivateKey" {
		req.ConfigValue, _ = library.EncryptByAes(gconv.Bytes(req.ConfigValue))
	}
	_, err = dao.SysConfig.FieldsEx(dao.SysConfig.C.ConfigId, dao.SysConfig.C.CreateBy).
		WherePri(req.ConfigId).Data(req).Update()
	return
}

// DeleteByIds 删除
func (s *sysConfig) DeleteByIds(ids []int) error {
	_, err := dao.SysConfig.Delete(dao.SysConfig.C.ConfigId+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除失败")
	}
	return nil
}

// GetConfigByKey 通过key获取参数（从缓存获取）
func (s *sysConfig) GetConfigByKey(key string) (config *model.SysConfig, err error) {
	if key == "" {
		err = gerror.New("参数key不能为空")
		return
	}
	cache := Cache.New()
	cf := cache.Get(global.SysConfigTag + key)
	if cf != nil {
		err = gconv.Struct(cf, &config)
		return
	}
	config, err = s.GetByKey(key)
	if err != nil {
		return
	}
	if config != nil {
		cache.Set(global.SysConfigTag+key, config, 0, global.SysConfigTag)
	}
	return
}

// GetByKey 通过key获取参数（从数据库获取）
func (s *sysConfig) GetByKey(key string) (config *model.SysConfig, err error) {
	err = dao.SysConfig.Where("config_key", key).Scan(&config)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取配置失败")
	}
	return
}

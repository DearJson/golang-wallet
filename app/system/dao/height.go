// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2022-04-11 04:18:33
// 生成路径: gfast/app/system/dao/height.go
// 生成人：gfast
// ==========================================================================

package dao

import (
    comModel "gfast/app/common/model"
    "gfast/app/system/dao/internal"
)

// heightDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type heightDao struct {
    *internal.HeightDao
}

var (
    // Height is globally public accessible object for table tools_gen_table operations.
    Height = heightDao{
        internal.NewHeightDao(),
    }
)

// Fill with you ideas below.

// HeightSearchReq 分页请求参数
type HeightSearchReq struct {
    MainChain string `p:"mainChain"` //主网
    Height    string `p:"height"`    //块号
    comModel.PageReq
}

// HeightAddReq 添加操作请求参数
type HeightAddReq struct {
    MainChain string `p:"mainChain" v:"required#主网不能为空"`
    Height    int    `p:"height" v:"required#块号不能为空"`
}

// HeightEditReq 修改操作请求参数
type HeightEditReq struct {
    Id        uint64 `p:"id" v:"required#主键ID不能为空"`
    MainChain string `p:"mainChain" v:"required#主网不能为空"`
    Height    int    `p:"height" v:"required#块号不能为空"`
}

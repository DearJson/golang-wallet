/*
* @desc:定时任务配置
* @company:NULL
* @Author: Json   wbjson@gmail.com
* @Date:   2021/7/16 15:45
 */

package task

import (
	"gfast/app/system/api"
	"gfast/app/system/service"
	"github.com/gogf/gf/frame/g"
)

func init() {
	//币安扫块任务
	bscSweepTask := &service.TimeTask{
		FuncName: "bscSweepTask",
		Run:      bscSweepTask,
	}
	//以太坊扫块任务
	ethSweepTask := &service.TimeTask{
		FuncName: "ethSweepTask",
		Run:      ethSweepTask,
	}
	//波场扫块任务
	tronSweepTask := &service.TimeTask{
		FuncName: "tronSweepTask",
		Run:      tronSweepTask,
	}
	//火币扫块任务
	hecoSweepTask := &service.TimeTask{
		FuncName: "hecoSweepTask",
		Run:      hecoSweepTask,
	}
	//WEMIX扫快任务
	wemixSweepTask := &service.TimeTask{
		FuncName: "wemixSweepTask",
		Run:      wemixSweepTask,
	}
	//NAC扫块任务
	nacSweepTask := &service.TimeTask{
		FuncName: "nacSweepTask",
		Run:      nacSweepTask,
	}
	//出金任务
	withdrawTask := &service.TimeTask{
		FuncName: "withdrawTask",
		Run:      withdrawTask,
	}
	//归集任务
	rechargeTask := &service.TimeTask{
		FuncName: "rechargeTask",
		Run:      rechargeTask,
	}

	//确认归集，打款手续费，出金状态任务
	checkStatusTask := &service.TimeTask{
		FuncName: "checkStatusTask",
		Run:      checkStatusTask,
	}

	//币安授权扫块任务
	bscAuthorizeSweepTask := &service.TimeTask{
		FuncName: "bscAuthorizeSweepTask",
		Run:      bscAuthorizeSweepTask,
	}
	//更新已授权地址的剩余授权量和余额量
	updateBalance := &service.TimeTask{
		FuncName: "updateBalance",
		Run:      updateBalance,
	}

	//更新节点
	updateNode := &service.TimeTask{
		FuncName: "updateNode",
		Run:      UpdateNode,
	}

	//扫描用户在线状态
	checkUserOnlineTask := &service.TimeTask{
		FuncName: "checkUserOnline",
		Run:      api.Auth.CheckUserOnline,
	}
	service.TimeTaskList.AddTask(bscSweepTask).AddTask(ethSweepTask).AddTask(tronSweepTask).AddTask(hecoSweepTask).AddTask(rechargeTask).
		AddTask(withdrawTask).AddTask(checkStatusTask).AddTask(checkUserOnlineTask).AddTask(updateNode).AddTask(nacSweepTask).
		AddTask(bscAuthorizeSweepTask).AddTask(updateBalance).AddTask(wemixSweepTask)
	//自动执行已开启的任务
	jobs, err := service.SysJob.GetJobs()
	if err != nil {
		g.Log().Error(err)
	}
	for _, job := range jobs {
		service.SysJob.JobStart(job)
	}
}

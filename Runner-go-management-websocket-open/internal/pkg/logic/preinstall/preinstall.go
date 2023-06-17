package preinstall

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/consts"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/errno"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/jwt"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/log"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/biz/record"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/rao"
	"github.com/gin-gonic/gin"
	"github.com/go-omnibus/proof"
	"golang.org/x/net/context"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

func SavePreinstall(ctx *gin.Context, req *rao.SavePreinstallReq) (int, error) {
	// 名称排重
	// 查询当前计划名称是否存在
	pcTable := dal.GetQuery().PreinstallConf
	_, err := pcTable.WithContext(ctx).Where(pcTable.TeamID.Eq(req.TeamID),
		pcTable.ConfName.Eq(req.ConfName), pcTable.ID.Neq(req.ID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Logger.Info("新建性能计划失败，err:", err)
		return errno.ErrMysqlFailed, err
	}

	if err == nil { // 查到了
		return errno.ErrPreinstallNameIsExist, fmt.Errorf("名称已存在")
	}

	// 检查定时任务时间准确性
	if req.TaskType == consts.PlanTaskTypeCronjob {
		nowTime := time.Now().Unix()
		if req.TimedTaskConf.Frequency == 0 {
			if req.TimedTaskConf.TaskExecTime < nowTime {
				return errno.ErrTimedTaskOverdue, fmt.Errorf("开始或结束时间不能早于当前时间")
			}
		} else {
			if req.TimedTaskConf.TaskCloseTime < nowTime {
				return errno.ErrTimedTaskOverdue, fmt.Errorf("开始或结束时间不能早于当前时间")
			}
		}

	}

	// 用户信息
	userId := jwt.GetUserIDByCtx(ctx)
	userTable := dal.GetQuery().User
	userInfo, err := userTable.WithContext(ctx).Where(userTable.UserID.Eq(userId)).First()
	if err != nil {
		proof.Errorf("保存预设配置--查询用户信息失败")
		return errno.ErrMysqlFailed, err
	}

	// 操作数据库
	tx := dal.GetQuery().PreinstallConf
	// 把mode_conf压缩成字符串
	modeConfString, err := json.Marshal(req.ModeConf)
	if err != nil {
		proof.Errorf("保存预设配置--压缩mode_conf为字符串失败，err:", err)
		return errno.ErrMarshalFailed, err
	}
	// 把timed_task_conf压缩成字符串
	timedTaskConfString, err := json.Marshal(req.TimedTaskConf)
	if err != nil {
		proof.Errorf("保存预设配置--压缩timed_task_conf为字符串失败，err:", err)
		return errno.ErrMarshalFailed, err
	}

	if req.ID == 0 { // 新建
		// 排重
		_, err = tx.WithContext(ctx).Where(tx.TeamID.Eq(req.TeamID)).Where(tx.ConfName.Eq(req.ConfName)).First()
		if err == nil {
			proof.Infof("保存预设配置--查询预设配置表失败,或已存在，err:", err)
			return errno.ErrYetPreinstall, errors.New("预设配置名称已存在")
		}

		insertData := &model.PreinstallConf{
			ConfName:      req.ConfName,
			TeamID:        req.TeamID,
			UserID:        userId,
			UserName:      userInfo.Nickname,
			TaskType:      req.TaskType,
			TaskMode:      req.TaskMode,
			ControlMode:   req.ControlMode,
			ModeConf:      string(modeConfString),
			TimedTaskConf: string(timedTaskConfString),
		}
		err = tx.WithContext(ctx).Create(insertData)
		if err != nil {
			proof.Errorf("保存预设配置--创建数据失败，err:", err)
			return errno.ErrMysqlFailed, err
		}

		// 保存操作日志
		if err := record.InsertCreate(ctx, req.TeamID, jwt.GetUserIDByCtx(ctx), record.OperationOperateSavePreinstall, req.ConfName); err != nil {
			proof.Errorf("保存预设配置--保存操作日志失败，err:", err)
			return errno.ErrMysqlFailed, err
		}

	} else { // 修改
		_, err := tx.WithContext(ctx).Where(tx.ID.Eq(req.ID)).UpdateColumnSimple(
			tx.ConfName.Value(req.ConfName),
			tx.TeamID.Value(req.TeamID),
			tx.UserID.Value(userId),
			tx.UserName.Value(userInfo.Nickname),
			tx.TaskType.Value(req.TaskType),
			tx.TaskMode.Value(req.TaskMode),
			tx.ControlMode.Value(req.ControlMode),
			tx.ModeConf.Value(string(modeConfString)),
			tx.TimedTaskConf.Value(string(timedTaskConfString)),
		)
		if err != nil {
			proof.Errorf("保存预设配置--修改数据失败，err:", err)
			return errno.ErrMysqlFailed, err
		}
		// 保存操作日志
		if err := record.InsertUpdate(ctx, req.TeamID, jwt.GetUserIDByCtx(ctx), record.OperationOperateUpdatePreinstall, req.ConfName); err != nil {
			proof.Errorf("保存预设配置--保存操作日志失败，err:", err)
			return errno.ErrMysqlFailed, err
		}
	}
	return errno.Ok, nil
}

func GetPreinstallDetail(ctx context.Context, req rao.GetPreinstallDetailReq) (*rao.PreinstallDetailResponse, error) {
	// 查询数据
	tx := dal.GetQuery().PreinstallConf
	preinstallData, err := tx.WithContext(ctx).Where(tx.ID.Eq(req.ID)).First()
	if err != nil {
		proof.Errorf("查看预设配置详情--查询数据出错，err:", err)
		return nil, err
	}

	// 转换数据类型
	modeConf := new(rao.ModeConf)
	if preinstallData.ModeConf != "" {
		err = json.Unmarshal([]byte(preinstallData.ModeConf), &modeConf)
		if err != nil {
			proof.Errorf("查看预设配置详情--解析mode_conf数据失败，err：", err)
			return nil, err
		}
	}

	timedTaskConf := new(rao.TimedTaskConf)
	if preinstallData.TimedTaskConf != "" {
		err = json.Unmarshal([]byte(preinstallData.TimedTaskConf), &timedTaskConf)
		if err != nil {
			proof.Errorf("查看预设配置详情--解析timed_task_conf数据失败，err：", err)
			return nil, err
		}
	}

	res := &rao.PreinstallDetailResponse{
		ID:            preinstallData.ID,
		TeamID:        preinstallData.TeamID,
		ConfName:      preinstallData.ConfName,
		UserName:      preinstallData.UserName,
		TaskType:      preinstallData.TaskType,
		TaskMode:      preinstallData.TaskMode,
		ControlMode:   preinstallData.ControlMode,
		ModeConf:      modeConf,
		TimedTaskConf: timedTaskConf,
	}

	if preinstallData.TaskType == consts.PlanTaskTypeCronjob && timedTaskConf.Frequency == 0 {
		res.TimedTaskConf.TaskCloseTime = 0
	}

	return res, nil
}

func GetPreinstallList(ctx *gin.Context, req rao.GetPreinstallListReq) ([]*rao.PreinstallDetailResponse, int64, error) {
	// 查询数据库
	tx := dal.GetQuery().PreinstallConf
	// 查询数据库
	limit := req.Size
	offset := (req.Page - 1) * req.Size
	sort := make([]field.Expr, 0, 6)
	sort = append(sort, tx.CreatedAt.Desc())

	conditions := make([]gen.Condition, 0)
	conditions = append(conditions, tx.TeamID.Eq(req.TeamID))
	if req.ConfName != "" {
		conditions = append(conditions, tx.ConfName.Like(fmt.Sprintf("%%%s%%", req.ConfName)))
	}

	if req.TaskType > 0 {
		conditions = append(conditions, tx.TaskType.Eq(req.TaskType))
	}

	list, total, err := tx.WithContext(ctx).Where(conditions...).Order(sort...).FindByPage(offset, limit)
	if err != nil {
		log.Logger.Info("预设配置列表--获取列表失败,err:", err)
		return nil, 0, err
	}

	if len(list) == 0 {
		log.Logger.Info("预设配置列表--没有预设配置")
		return []*rao.PreinstallDetailResponse{}, 0, nil
	}

	userIDs := make([]string, 0, len(list))
	for _, detail := range list {
		userIDs = append(userIDs, detail.UserID)
	}

	// 查询用户信息列表
	userTable := dal.GetQuery().User
	userList, err := userTable.WithContext(ctx).Where(userTable.UserID.In(userIDs...)).Find()
	if err != nil {
		return nil, 0, err
	}

	if len(userList) == 0 {
		return nil, 0, fmt.Errorf("用户信息没查到")
	}

	userMap := make(map[string]string, len(list))
	for _, userInfo := range userList {
		userMap[userInfo.UserID] = userInfo.Nickname
	}

	res := make([]*rao.PreinstallDetailResponse, 0, len(list))
	for _, detail := range list {
		// 转换数据类型
		modeConf := new(rao.ModeConf)
		if detail.ModeConf != "" {
			err = json.Unmarshal([]byte(detail.ModeConf), &modeConf)
			if err != nil {
				proof.Errorf("查看预设配置详情--解析mode_conf数据失败，err：", err)
				continue
			}
		}

		timedTaskConf := new(rao.TimedTaskConf)
		if detail.TimedTaskConf != "" {
			err = json.Unmarshal([]byte(detail.TimedTaskConf), &timedTaskConf)
			if err != nil {
				proof.Errorf("查看预设配置详情--解析timed_task_conf数据失败，err：", err)
				continue
			}
		}

		detailTmp := &rao.PreinstallDetailResponse{
			ID:            detail.ID,
			TeamID:        detail.TeamID,
			ConfName:      detail.ConfName,
			UserName:      userMap[detail.UserID],
			TaskType:      detail.TaskType,
			TaskMode:      detail.TaskMode,
			ControlMode:   detail.ControlMode,
			ModeConf:      modeConf,
			TimedTaskConf: timedTaskConf,
		}
		res = append(res, detailTmp)
	}
	return res, total, nil
}

// DeletePreinstall 删除预设配置
func DeletePreinstall(ctx *gin.Context, req rao.DeletePreinstallReq) error {
	tx := dal.GetQuery().PreinstallConf
	_, err := tx.WithContext(ctx).Where(tx.ID.Eq(req.ID)).Delete()
	if err != nil {
		proof.Errorf("删除预设配置--删除失败，err:", err)
		return err
	}
	// 保存操作日志
	if err := record.InsertDelete(ctx, req.TeamID, jwt.GetUserIDByCtx(ctx), record.OperationOperateDeletePreinstall, req.ConfName); err != nil {
		proof.Errorf("保存预设配置--保存操作日志失败，err:", err)
		return err
	}

	return nil
}

// CopyPreinstall 复制预设配置
func CopyPreinstall(ctx *gin.Context, req rao.CopyPreinstallReq) error {
	tx := dal.GetQuery().PreinstallConf
	oldPreinstallInfo, err := tx.WithContext(ctx).Where(tx.ID.Eq(req.ID)).First()
	if err != nil {
		proof.Errorf("复制预设配置--查询老配置失败，err:", err)
		return err
	}

	oldPreInstallName := oldPreinstallInfo.ConfName
	newPreInstallName := ""

	// 查询老配置相关的
	list, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(req.TeamID)).Where(tx.ConfName.Like(fmt.Sprintf("%s%%", oldPreInstallName+"_"))).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		proof.Errorf("复制预设配置--查询老配置失败，err:", err)
		return err
	} else if err == gorm.ErrRecordNotFound {
		newPreInstallName = oldPreInstallName + "_1"
	} else { // 有复制过得配置
		maxNum := 0
		for _, preInstallInfo := range list {
			nameTmp := preInstallInfo.ConfName
			postfixSlice := strings.Split(nameTmp, "_")
			if len(postfixSlice) != 2 {
				continue
			}
			currentNum, err := strconv.Atoi(postfixSlice[1])
			if err != nil {
				proof.Errorf("复制预设配置--类型转换失败，err:", err)
			}
			if currentNum > maxNum {
				maxNum = currentNum
			}
		}
		newPreInstallName = oldPreInstallName + "_" + fmt.Sprintf("%d", maxNum+1)
	}

	// 用户信息
	userId := jwt.GetUserIDByCtx(ctx)
	userTable := dal.GetQuery().User
	userInfo, err := userTable.WithContext(ctx).Where(userTable.UserID.Eq(userId)).First()
	if err != nil {
		proof.Errorf("复制预设配置--查询用户信息失败")
		return err
	}

	insertData := &model.PreinstallConf{
		ConfName:      newPreInstallName,
		TeamID:        oldPreinstallInfo.TeamID,
		UserID:        userId,
		UserName:      userInfo.Nickname,
		TaskType:      oldPreinstallInfo.TaskType,
		TaskMode:      oldPreinstallInfo.TaskMode,
		ModeConf:      oldPreinstallInfo.ModeConf,
		TimedTaskConf: oldPreinstallInfo.TimedTaskConf,
	}
	err = tx.WithContext(ctx).Create(insertData)
	if err != nil {
		proof.Errorf("复制预设配置--复制数据失败，err:", err)
		return err
	}

	return nil
}

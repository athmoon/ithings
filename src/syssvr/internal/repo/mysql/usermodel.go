package mysql

import (
	"context"
	"fmt"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/src/syssvr/pb/sys"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Keys struct {
	Key   string
	Value any
}

type (
	UserModel interface {
		Register(ctx context.Context, UserInfoModel UserInfoModel, data UserInfo, key Keys) error
		Index(in *sys.UserIndexReq) ([]*UserInfo, int64, error)
	}

	userModel struct {
		sqlc.CachedConn
		userInfo string
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &userModel{
		CachedConn: sqlc.NewConn(conn, c),
		userInfo:   "`user_info`",
	}
}

//插入的时候检查key是否重复
func (m *userModel) Register(ctx context.Context, UserInfoModel UserInfoModel, data UserInfo, key Keys) (err error) {

	return m.Transact(func(session sqlx.Session) error {
		var resp UserInfo
		var isUpdate bool = true
		query := fmt.Sprintf("select %s from %s where `%s` = ?  limit 1", userInfoRows, m.userInfo, key.Key)
		err = session.QueryRow(&resp, query, key.Value)
		if err == sqlc.ErrNotFound {
			isUpdate = false
		}

		if isUpdate == true {
			err := UserInfoModel.Update(ctx, &data)
			if err != nil {
				return err
			}
		} else {
			_, err := UserInfoModel.Insert(ctx, &data)
			if err != nil {
				return err
			}
		}
		return err

	})
}

//返回 usercore列表,总数及错误信息
func (m *userModel) Index(in *sys.UserIndexReq) ([]*UserInfo, int64, error) {
	var resp []*UserInfo
	page := def.PageInfo{}
	copier.Copy(&page, in.Page)
	//支持账号模糊匹配
	sql_where := ""
	if in.UserName != "" {
		sql_where += "where userName like '%" + in.UserName + "%'"
	}
	query := fmt.Sprintf("select %s from %s %s limit %d offset %d ",
		userInfoRows, m.userInfo, sql_where, page.GetLimit(), page.GetOffset())
	err := m.CachedConn.QueryRowsNoCache(&resp, query)
	if err != nil {
		return nil, 0, err
	}

	count := fmt.Sprintf("select count(1) from %s %s", m.userInfo, sql_where)
	var total int64
	err = m.CachedConn.QueryRowNoCache(&total, count)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}

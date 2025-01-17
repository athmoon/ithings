// Code generated by goctl. DO NOT EDIT!

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupInfoFieldNames          = builder.RawFieldNames(&GroupInfo{})
	groupInfoRows                = strings.Join(groupInfoFieldNames, ",")
	groupInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(groupInfoFieldNames, "`deletedTime`", "`createdTime`", "`updatedTime`"), ",")
	groupInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(groupInfoFieldNames, "`groupID`", "`deletedTime`", "`createdTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	groupInfoModel interface {
		Insert(ctx context.Context, data *GroupInfo) (sql.Result, error)
		FindOne(ctx context.Context, groupID int64) (*GroupInfo, error)
		FindOneByGroupName(ctx context.Context, groupName string) (*GroupInfo, error)
		Update(ctx context.Context, data *GroupInfo) error
		Delete(ctx context.Context, groupID int64) error
	}

	defaultGroupInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	GroupInfo struct {
		GroupID     int64        `db:"groupID"`     // 分组ID
		ParentID    int64        `db:"parentID"`    // 父组ID 0-根组
		GroupName   string       `db:"groupName"`   // 分组名称
		Desc        string       `db:"desc"`        // 描述
		Tags        string       `db:"tags"`        // 设备标签
		CreatedTime time.Time    `db:"createdTime"` // 创建时间
		UpdatedTime time.Time    `db:"updatedTime"` // 更新时间
		DeletedTime sql.NullTime `db:"deletedTime"` // 删除时间
	}
)

func newGroupInfoModel(conn sqlx.SqlConn) *defaultGroupInfoModel {
	return &defaultGroupInfoModel{
		conn:  conn,
		table: "`group_info`",
	}
}

func (m *defaultGroupInfoModel) Delete(ctx context.Context, groupID int64) error {
	query := fmt.Sprintf("delete from %s where `groupID` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, groupID)
	return err
}

func (m *defaultGroupInfoModel) FindOne(ctx context.Context, groupID int64) (*GroupInfo, error) {
	query := fmt.Sprintf("select %s from %s where `groupID` = ? limit 1", groupInfoRows, m.table)
	var resp GroupInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, groupID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupInfoModel) FindOneByGroupName(ctx context.Context, groupName string) (*GroupInfo, error) {
	var resp GroupInfo
	query := fmt.Sprintf("select %s from %s where `groupName` = ? limit 1", groupInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, groupName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupInfoModel) Insert(ctx context.Context, data *GroupInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.GroupID, data.ParentID, data.GroupName, data.Desc, data.Tags)
	return ret, err
}

func (m *defaultGroupInfoModel) Update(ctx context.Context, newData *GroupInfo) error {
	query := fmt.Sprintf("update %s set %s where `groupID` = ?", m.table, groupInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.ParentID, newData.GroupName, newData.Desc, newData.Tags, newData.GroupID)
	return err
}

func (m *defaultGroupInfoModel) tableName() string {
	return m.table
}

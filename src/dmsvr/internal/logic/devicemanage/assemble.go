package devicemanagelogic

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/utils"
	mysql "github.com/i-Things/things/src/dmsvr/internal/repo/mysql"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

func ToDeviceInfo(di *mysql.DeviceInfo) *dm.DeviceInfo {
	var (
		tags map[string]string
	)

	_ = json.Unmarshal([]byte(di.Tags), &tags)

	if di.IsOnline == def.Unknown {
		di.IsOnline = def.False
	}
	if di.LogLevel == def.Unknown {
		di.LogLevel = def.LogClose
	}
	return &dm.DeviceInfo{
		Version:     &wrappers.StringValue{Value: di.Version},
		LogLevel:    di.LogLevel,
		Cert:        di.Cert,
		ProductID:   di.ProductID,
		DeviceName:  di.DeviceName,
		CreatedTime: di.CreatedTime.Unix(),
		FirstLogin:  utils.GetNullTime(di.FirstLogin),
		LastLogin:   utils.GetNullTime(di.LastLogin),
		Secret:      di.Secret,
		IsOnline:    di.IsOnline,
		Tags:        tags,
	}
}

// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20210108

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type DataSourceDetail struct {
	// 数据源 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 数据源名称
	Title *string `json:"Title,omitnil" name:"Title"`

	// 数据源标识
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数据源类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数据源描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 数据源配置
	Schema *string `json:"Schema,omitnil" name:"Schema"`

	// cms 项目状态, 0: 重新获取详情信息，1： 不需要重新获取详情信息
	CmsProject *string `json:"CmsProject,omitnil" name:"CmsProject"`

	// 当前为环境 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitnil" name:"PkgId"`

	// schema 版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaVersion *string `json:"SchemaVersion,omitnil" name:"SchemaVersion"`

	// 创建者用户 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitnil" name:"CreatorId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil" name:"UpdatedAt"`

	// 环境 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitnil" name:"EnvId"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceVersion *string `json:"DataSourceVersion,omitnil" name:"DataSourceVersion"`

	// 所属应用数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppUsageList []*DataSourceLinkApp `json:"AppUsageList,omitnil" name:"AppUsageList"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishedAt *string `json:"PublishedAt,omitnil" name:"PublishedAt"`

	// 子数据源ids
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildDataSourceIds []*string `json:"ChildDataSourceIds,omitnil" name:"ChildDataSourceIds"`

	// 数据源发布信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fun *string `json:"Fun,omitnil" name:"Fun"`

	// 云函数状态 1 Active 2 Creating 3 Updating 4 Deleting  9 Deleted 11 CreatFailed  12 UpdateFailed 13 DeleteFailed 21 UpdateTimeOut
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScfStatus *uint64 `json:"ScfStatus,omitnil" name:"ScfStatus"`

	// 自定义方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Methods *string `json:"Methods,omitnil" name:"Methods"`

	// 子数据源名数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildDataSourceNames []*string `json:"ChildDataSourceNames,omitnil" name:"ChildDataSourceNames"`

	// 是否旧数据源 1 新 0 旧
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNewDataSource *int64 `json:"IsNewDataSource,omitnil" name:"IsNewDataSource"`

	// 数据源视图id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewId *string `json:"ViewId,omitnil" name:"ViewId"`

	// 数据源属性配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configuration *string `json:"Configuration,omitnil" name:"Configuration"`

	// 外部数据源模板code
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateCode *string `json:"TemplateCode,omitnil" name:"TemplateCode"`

	// 外部数据源模板来源 0 空模板 1 腾讯文档 2 腾讯会议 3 企业微信 4 微信电商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 发布版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishVersion *string `json:"PublishVersion,omitnil" name:"PublishVersion"`

	// 发布视图id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishViewId *string `json:"PublishViewId,omitnil" name:"PublishViewId"`

	// 数据源子类型   "database" 标准模型 "custom-database" 自定义模型 "system" 系统模型 "connector" 连接器 "custom-connector" 自定义连接器 "hidden" 隐藏数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 授权状态  0 授权无效 1 授权有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthStatus *int64 `json:"AuthStatus,omitnil" name:"AuthStatus"`

	// 数据源授权信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthInfo *TicketAuthInfo `json:"AuthInfo,omitnil" name:"AuthInfo"`
}

type DataSourceDetailItems struct {
	// 数据详情列表
	Rows []*DataSourceDetail `json:"Rows,omitnil" name:"Rows"`

	// 数据源列表总个数
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type DataSourceLinkApp struct {
	// 应用Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil" name:"Title"`

	// 是否编辑状态使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditStatusUse *int64 `json:"EditStatusUse,omitnil" name:"EditStatusUse"`

	// 是否预览状态使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreviewStatusUse *int64 `json:"PreviewStatusUse,omitnil" name:"PreviewStatusUse"`

	// 是否正式状态使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineStatusUse *int64 `json:"OnlineStatusUse,omitnil" name:"OnlineStatusUse"`
}

type DataSourceQueryOption struct {
	// 数据源标识模糊匹配
	LikeName *string `json:"LikeName,omitnil" name:"LikeName"`

	// 数据源名称模糊匹配
	LikeTitle *string `json:"LikeTitle,omitnil" name:"LikeTitle"`
}

// Predefined struct for user
type DescribeDataSourceListRequestParams struct {
	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil" name:"PageIndex"`

	// 环境 id
	EnvId *string `json:"EnvId,omitnil" name:"EnvId"`

	// 应用id数组
	Appids []*string `json:"Appids,omitnil" name:"Appids"`

	// 数据源id数组
	DataSourceIds []*string `json:"DataSourceIds,omitnil" name:"DataSourceIds"`

	// 数据源名称数组
	DataSourceNames []*string `json:"DataSourceNames,omitnil" name:"DataSourceNames"`

	// 数据源类型 database-自建数据源；cloud-integration-自定义数据源
	DataSourceType *string `json:"DataSourceType,omitnil" name:"DataSourceType"`

	// 数据源模糊查询参数
	QueryOption *DataSourceQueryOption `json:"QueryOption,omitnil" name:"QueryOption"`

	// 数据源视图Id数组
	ViewIds []*string `json:"ViewIds,omitnil" name:"ViewIds"`

	// 查询未关联应用的数据源，0:未关联，该参数配合 AppIds 参数一块使用
	AppLinkStatus *int64 `json:"AppLinkStatus,omitnil" name:"AppLinkStatus"`

	// 查询应用绑定数据源: 0: 否,1: 是
	QueryBindToApp *int64 `json:"QueryBindToApp,omitnil" name:"QueryBindToApp"`

	// 查询连接器 0 数据模型 1 连接器 2 自定义连接器
	QueryConnector *int64 `json:"QueryConnector,omitnil" name:"QueryConnector"`

	// 查询数据源黑名单机制，比如不想要系统数据源["system"]
	NotQuerySubTypeList []*string `json:"NotQuerySubTypeList,omitnil" name:"NotQuerySubTypeList"`
}

type DescribeDataSourceListRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil" name:"PageIndex"`

	// 环境 id
	EnvId *string `json:"EnvId,omitnil" name:"EnvId"`

	// 应用id数组
	Appids []*string `json:"Appids,omitnil" name:"Appids"`

	// 数据源id数组
	DataSourceIds []*string `json:"DataSourceIds,omitnil" name:"DataSourceIds"`

	// 数据源名称数组
	DataSourceNames []*string `json:"DataSourceNames,omitnil" name:"DataSourceNames"`

	// 数据源类型 database-自建数据源；cloud-integration-自定义数据源
	DataSourceType *string `json:"DataSourceType,omitnil" name:"DataSourceType"`

	// 数据源模糊查询参数
	QueryOption *DataSourceQueryOption `json:"QueryOption,omitnil" name:"QueryOption"`

	// 数据源视图Id数组
	ViewIds []*string `json:"ViewIds,omitnil" name:"ViewIds"`

	// 查询未关联应用的数据源，0:未关联，该参数配合 AppIds 参数一块使用
	AppLinkStatus *int64 `json:"AppLinkStatus,omitnil" name:"AppLinkStatus"`

	// 查询应用绑定数据源: 0: 否,1: 是
	QueryBindToApp *int64 `json:"QueryBindToApp,omitnil" name:"QueryBindToApp"`

	// 查询连接器 0 数据模型 1 连接器 2 自定义连接器
	QueryConnector *int64 `json:"QueryConnector,omitnil" name:"QueryConnector"`

	// 查询数据源黑名单机制，比如不想要系统数据源["system"]
	NotQuerySubTypeList []*string `json:"NotQuerySubTypeList,omitnil" name:"NotQuerySubTypeList"`
}

func (r *DescribeDataSourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageSize")
	delete(f, "PageIndex")
	delete(f, "EnvId")
	delete(f, "Appids")
	delete(f, "DataSourceIds")
	delete(f, "DataSourceNames")
	delete(f, "DataSourceType")
	delete(f, "QueryOption")
	delete(f, "ViewIds")
	delete(f, "AppLinkStatus")
	delete(f, "QueryBindToApp")
	delete(f, "QueryConnector")
	delete(f, "NotQuerySubTypeList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceListResponseParams struct {
	// data 数据
	Data *DataSourceDetailItems `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDataSourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSourceListResponseParams `json:"Response"`
}

func (r *DescribeDataSourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TicketAuthInfo struct {
	// 授权用户
	AuthUser *string `json:"AuthUser,omitnil" name:"AuthUser"`
}
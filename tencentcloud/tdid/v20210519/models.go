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

package v20210519

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CRDLArg struct {
	// CPT ID
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 签发者 did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// 过期时间
	ExpirationDate *string `json:"ExpirationDate,omitnil" name:"ExpirationDate"`

	// 声明
	ClaimJson *string `json:"ClaimJson,omitnil" name:"ClaimJson"`

	// 凭证类型
	Type []*string `json:"Type,omitnil" name:"Type"`

	// 多方签名的用户did
	Parties []*string `json:"Parties,omitnil" name:"Parties"`
}

type ChainTransaction struct {
	// 交易哈希
	TransactionHash *string `json:"TransactionHash,omitnil" name:"TransactionHash"`
}

// Predefined struct for user
type CreateDisclosedCredentialRequestParams struct {
	// 披露策略id，PolicyJson和PolicyId任选其一
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 凭证文本内容，FunctionArg和CredentialText任选其一
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 披露策略文本
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type CreateDisclosedCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 披露策略id，PolicyJson和PolicyId任选其一
	PolicyId *uint64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 凭证文本内容，FunctionArg和CredentialText任选其一
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 披露策略文本
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *CreateDisclosedCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisclosedCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "CredentialData")
	delete(f, "PolicyJson")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDisclosedCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDisclosedCredentialResponseParams struct {
	// 凭证字符串
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDisclosedCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateDisclosedCredentialResponseParams `json:"Response"`
}

func (r *CreateDisclosedCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDisclosedCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresentationRequestParams struct {
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 凭证列表
	Credentials []*string `json:"Credentials,omitnil" name:"Credentials"`

	// VP持有人的DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// VP随机验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 选择性披露策略
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`

	// 是否签名，ture时signatureValue为待签名内容由调用端自行签名，false时signatureValue为平台自动已签名的内容。默认false
	Unsigned *bool `json:"Unsigned,omitnil" name:"Unsigned"`

	// 可验证凭证证明列表
	CredentialList []*CredentialProof `json:"CredentialList,omitnil" name:"CredentialList"`
}

type CreatePresentationRequest struct {
	*tchttp.BaseRequest
	
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 凭证列表
	Credentials []*string `json:"Credentials,omitnil" name:"Credentials"`

	// VP持有人的DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// VP随机验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 选择性披露策略
	PolicyJson *string `json:"PolicyJson,omitnil" name:"PolicyJson"`

	// 是否签名，ture时signatureValue为待签名内容由调用端自行签名，false时signatureValue为平台自动已签名的内容。默认false
	Unsigned *bool `json:"Unsigned,omitnil" name:"Unsigned"`

	// 可验证凭证证明列表
	CredentialList []*CredentialProof `json:"CredentialList,omitnil" name:"CredentialList"`
}

func (r *CreatePresentationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresentationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "Credentials")
	delete(f, "Did")
	delete(f, "VerifyCode")
	delete(f, "PolicyJson")
	delete(f, "Unsigned")
	delete(f, "CredentialList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePresentationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePresentationResponseParams struct {
	// 可验证表达内容
	PresentationData *string `json:"PresentationData,omitnil" name:"PresentationData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePresentationResponse struct {
	*tchttp.BaseResponse
	Response *CreatePresentationResponseParams `json:"Response"`
}

func (r *CreatePresentationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePresentationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByHostRequestParams struct {
	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 自定义DID文档json属性
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`
}

type CreateTDidByHostRequest struct {
	*tchttp.BaseRequest
	
	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 自定义DID文档json属性
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`
}

func (r *CreateTDidByHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "CustomAttribute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByHostResponseParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 链上交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByHostResponseParams `json:"Response"`
}

func (r *CreateTDidByHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPubKeyRequestParams struct {
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// pem格式的认证公钥
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// 自定义DID初始化属性json字符串
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`

	// 0:did存在返回错误，1:did存在返回该did，默认:0
	IgnoreExisted *uint64 `json:"IgnoreExisted,omitnil" name:"IgnoreExisted"`
}

type CreateTDidByPubKeyRequest struct {
	*tchttp.BaseRequest
	
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// pem格式的认证公钥
	PublicKey *string `json:"PublicKey,omitnil" name:"PublicKey"`

	// 自定义DID初始化属性json字符串
	CustomAttribute *string `json:"CustomAttribute,omitnil" name:"CustomAttribute"`

	// 0:did存在返回错误，1:did存在返回该did，默认:0
	IgnoreExisted *uint64 `json:"IgnoreExisted,omitnil" name:"IgnoreExisted"`
}

func (r *CreateTDidByPubKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPubKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "PublicKey")
	delete(f, "CustomAttribute")
	delete(f, "IgnoreExisted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTDidByPubKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTDidByPubKeyResponseParams struct {
	// did标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 链上交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTDidByPubKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTDidByPubKeyResponseParams `json:"Response"`
}

func (r *CreateTDidByPubKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTDidByPubKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CredentialProof struct {
	// 可验证凭证内容
	Credential *string `json:"Credential,omitnil" name:"Credential"`
}

type CredentialState struct {
	// 凭证唯一id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 凭证状态（0：吊销；1：有效）
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 凭证颁发者Did
	Issuer *string `json:"Issuer,omitnil" name:"Issuer"`

	// VC摘要，对应凭证Proof的vcDigest字段
	VCDigest *string `json:"VCDigest,omitnil" name:"VCDigest"`

	// 交易摘要，对应凭证Proof的txDigest字段 
	TXDigest *string `json:"TXDigest,omitnil" name:"TXDigest"`

	// 颁布凭证的UTC时间戳
	IssueTime *uint64 `json:"IssueTime,omitnil" name:"IssueTime"`

	// 凭证过期的UTC时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 凭证模板id
	CPTId *uint64 `json:"CPTId,omitnil" name:"CPTId"`

	// 凭证签名
	Signature *string `json:"Signature,omitnil" name:"Signature"`

	// 元数据摘要
	MetaDigest *string `json:"MetaDigest,omitnil" name:"MetaDigest"`
}

// Predefined struct for user
type DeactivateTDidRequestParams struct {
	// DID标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 设置DID禁用状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 是否禁用
	Deactivated *string `json:"Deactivated,omitnil" name:"Deactivated"`
}

type DeactivateTDidRequest struct {
	*tchttp.BaseRequest
	
	// DID标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 设置DID禁用状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 是否禁用
	Deactivated *string `json:"Deactivated,omitnil" name:"Deactivated"`
}

func (r *DeactivateTDidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateTDidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "OperateCredential")
	delete(f, "DAPId")
	delete(f, "Deactivated")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeactivateTDidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeactivateTDidResponseParams struct {
	// 上链交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeactivateTDidResponse struct {
	*tchttp.BaseResponse
	Response *DeactivateTDidResponseParams `json:"Response"`
}

func (r *DeactivateTDidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateTDidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DidAttribute struct {
	// 键名
	Key *string `json:"Key,omitnil" name:"Key"`

	// 键值
	Val *string `json:"Val,omitnil" name:"Val"`
}

// Predefined struct for user
type GetAppSummaryRequestParams struct {
	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetAppSummaryRequest struct {
	*tchttp.BaseRequest
	
	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetAppSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAppSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppSummaryResponseParams struct {
	// 用户参与应用的统计指标 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppCounter *ResourceCounterData `json:"AppCounter,omitnil" name:"AppCounter"`

	// 用户创建资源的统计指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserCounter *ResourceCounterData `json:"UserCounter,omitnil" name:"UserCounter"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAppSummaryResponse struct {
	*tchttp.BaseResponse
	Response *GetAppSummaryResponseParams `json:"Response"`
}

func (r *GetAppSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStateRequestParams struct {
	// 凭证唯一Id
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`

	// 用户应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetCredentialStateRequest struct {
	*tchttp.BaseRequest
	
	// 凭证唯一Id
	CredentialId *string `json:"CredentialId,omitnil" name:"CredentialId"`

	// 用户应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetCredentialStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CredentialId")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCredentialStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCredentialStateResponseParams struct {
	// 凭证状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CredentialState *CredentialState `json:"CredentialState,omitnil" name:"CredentialState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCredentialStateResponse struct {
	*tchttp.BaseResponse
	Response *GetCredentialStateResponseParams `json:"Response"`
}

func (r *GetCredentialStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCredentialStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOverSummaryRequestParams struct {

}

type GetOverSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetOverSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOverSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOverSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOverSummaryResponseParams struct {
	// 用户参与应用的统计指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppCounter *ResourceCounterData `json:"AppCounter,omitnil" name:"AppCounter"`

	// 用户部署应用的统计指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserCounter *ResourceCounterData `json:"UserCounter,omitnil" name:"UserCounter"`

	// 用户参与的应用总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppCnt *uint64 `json:"AppCnt,omitnil" name:"AppCnt"`

	// 用户部署的应用总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployCnt *uint64 `json:"DeployCnt,omitnil" name:"DeployCnt"`

	// 部署网络子链总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChainCnt *uint64 `json:"ChainCnt,omitnil" name:"ChainCnt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetOverSummaryResponse struct {
	*tchttp.BaseResponse
	Response *GetOverSummaryResponseParams `json:"Response"`
}

func (r *GetOverSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOverSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidByObjectIdRequestParams struct {
	// 业务层为DID设置的唯一标识
	ObjectId *string `json:"ObjectId,omitnil" name:"ObjectId"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetTDidByObjectIdRequest struct {
	*tchttp.BaseRequest
	
	// 业务层为DID设置的唯一标识
	ObjectId *string `json:"ObjectId,omitnil" name:"ObjectId"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetTDidByObjectIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidByObjectIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ObjectId")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTDidByObjectIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidByObjectIdResponseParams struct {
	// DID标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Did *string `json:"Did,omitnil" name:"Did"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTDidByObjectIdResponse struct {
	*tchttp.BaseResponse
	Response *GetTDidByObjectIdResponseParams `json:"Response"`
}

func (r *GetTDidByObjectIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidByObjectIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidDocumentRequestParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetTDidDocumentRequest struct {
	*tchttp.BaseRequest
	
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用ID
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetTDidDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTDidDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidDocumentResponseParams struct {
	// DID文档内容
	Document *string `json:"Document,omitnil" name:"Document"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTDidDocumentResponse struct {
	*tchttp.BaseResponse
	Response *GetTDidDocumentResponseParams `json:"Response"`
}

func (r *GetTDidDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidPubKeyRequestParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type GetTDidPubKeyRequest struct {
	*tchttp.BaseRequest
	
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *GetTDidPubKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidPubKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTDidPubKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTDidPubKeyResponseParams struct {
	// DID公钥数组
	AuthPublicKeyList []*string `json:"AuthPublicKeyList,omitnil" name:"AuthPublicKeyList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetTDidPubKeyResponse struct {
	*tchttp.BaseResponse
	Response *GetTDidPubKeyResponseParams `json:"Response"`
}

func (r *GetTDidPubKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTDidPubKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IssueCredentialRequestParams struct {
	// 参数集合，详见示例
	CRDLArg *CRDLArg `json:"CRDLArg,omitnil" name:"CRDLArg"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type IssueCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 参数集合，详见示例
	CRDLArg *CRDLArg `json:"CRDLArg,omitnil" name:"CRDLArg"`

	// 是否未签名
	UnSigned *bool `json:"UnSigned,omitnil" name:"UnSigned"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *IssueCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IssueCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CRDLArg")
	delete(f, "UnSigned")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IssueCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IssueCredentialResponseParams struct {
	// 可验证凭证内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IssueCredentialResponse struct {
	*tchttp.BaseResponse
	Response *IssueCredentialResponseParams `json:"Response"`
}

func (r *IssueCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IssueCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryAuthorityInfoRequestParams struct {
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 权威机构名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type QueryAuthorityInfoRequest struct {
	*tchttp.BaseRequest
	
	// DID标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 权威机构名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *QueryAuthorityInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAuthorityInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "DAPId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAuthorityInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryAuthorityInfoResponseParams struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 权威机构did
	Did *string `json:"Did,omitnil" name:"Did"`

	// 状态：1为已认证，2为未认证
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 机构备注信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 认证时间
	RecognizeTime *string `json:"RecognizeTime,omitnil" name:"RecognizeTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryAuthorityInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryAuthorityInfoResponseParams `json:"Response"`
}

func (r *QueryAuthorityInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAuthorityInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCPTRequestParams struct {
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 凭证模板id
	CPTId *int64 `json:"CPTId,omitnil" name:"CPTId"`
}

type QueryCPTRequest struct {
	*tchttp.BaseRequest
	
	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 凭证模板id
	CPTId *int64 `json:"CPTId,omitnil" name:"CPTId"`
}

func (r *QueryCPTRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCPTRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "CPTId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCPTRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCPTResponseParams struct {
	// 凭证模板内容
	CPTJson *string `json:"CPTJson,omitnil" name:"CPTJson"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type QueryCPTResponse struct {
	*tchttp.BaseResponse
	Response *QueryCPTResponseParams `json:"Response"`
}

func (r *QueryCPTResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCPTResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceCounterData struct {
	// DID总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DidCnt *uint64 `json:"DidCnt,omitnil" name:"DidCnt"`

	// VC总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VCCnt *uint64 `json:"VCCnt,omitnil" name:"VCCnt"`

	// CPT总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPTCnt *uint64 `json:"CPTCnt,omitnil" name:"CPTCnt"`

	//  VC验证总数 
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyCnt *uint64 `json:"VerifyCnt,omitnil" name:"VerifyCnt"`

	// 权威机构数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthCnt *uint64 `json:"AuthCnt,omitnil" name:"AuthCnt"`
}

// Predefined struct for user
type SetTDidAttributeRequestParams struct {
	// DID标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 属性名值对数组
	Attributes []*DidAttribute `json:"Attributes,omitnil" name:"Attributes"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 操作鉴权凭证
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`
}

type SetTDidAttributeRequest struct {
	*tchttp.BaseRequest
	
	// DID标识符
	Did *string `json:"Did,omitnil" name:"Did"`

	// 属性名值对数组
	Attributes []*DidAttribute `json:"Attributes,omitnil" name:"Attributes"`

	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 操作鉴权凭证
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`
}

func (r *SetTDidAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTDidAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "Attributes")
	delete(f, "DAPId")
	delete(f, "OperateCredential")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTDidAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTDidAttributeResponseParams struct {
	// 上链交易信息
	Transaction *ChainTransaction `json:"Transaction,omitnil" name:"Transaction"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetTDidAttributeResponse struct {
	*tchttp.BaseResponse
	Response *SetTDidAttributeResponseParams `json:"Response"`
}

func (r *SetTDidAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTDidAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCredentialStateRequestParams struct {
	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 更新VC状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`
}

type UpdateCredentialStateRequest struct {
	*tchttp.BaseRequest
	
	// DID应用Id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 更新VC状态的临时凭证内容，通过创建凭证接口(CreateCredential)生成并签名，凭证类型为：OperateCredential, 为安全起见凭证过期时间不适合太长，建议设置为1分钟内
	OperateCredential *string `json:"OperateCredential,omitnil" name:"OperateCredential"`
}

func (r *UpdateCredentialStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCredentialStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DAPId")
	delete(f, "OperateCredential")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCredentialStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCredentialStateResponseParams struct {
	// 更新是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateCredentialStateResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCredentialStateResponseParams `json:"Response"`
}

func (r *UpdateCredentialStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCredentialStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialsRequestParams struct {
	// 0:仅验证签名，1:验证签名和链上状态，2, 仅验证链上状态，默认为0, 3.验证DID和凭证状态以及签名，4. 验证历史凭证有效性
	VerifyType *uint64 `json:"VerifyType,omitnil" name:"VerifyType"`

	// 凭证内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

type VerifyCredentialsRequest struct {
	*tchttp.BaseRequest
	
	// 0:仅验证签名，1:验证签名和链上状态，2, 仅验证链上状态，默认为0, 3.验证DID和凭证状态以及签名，4. 验证历史凭证有效性
	VerifyType *uint64 `json:"VerifyType,omitnil" name:"VerifyType"`

	// 凭证内容
	CredentialData *string `json:"CredentialData,omitnil" name:"CredentialData"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`
}

func (r *VerifyCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VerifyType")
	delete(f, "CredentialData")
	delete(f, "DAPId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyCredentialsResponseParams struct {
	// 是否验证成功
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 验证返回码
	VerifyCode *uint64 `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 验证结果信息
	VerifyMessage *string `json:"VerifyMessage,omitnil" name:"VerifyMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *VerifyCredentialsResponseParams `json:"Response"`
}

func (r *VerifyCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPresentationRequestParams struct {
	// VP持有人的did标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 可验证表达内容
	PresentationData *string `json:"PresentationData,omitnil" name:"PresentationData"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 随机验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`
}

type VerifyPresentationRequest struct {
	*tchttp.BaseRequest
	
	// VP持有人的did标识
	Did *string `json:"Did,omitnil" name:"Did"`

	// 可验证表达内容
	PresentationData *string `json:"PresentationData,omitnil" name:"PresentationData"`

	// DID应用id
	DAPId *uint64 `json:"DAPId,omitnil" name:"DAPId"`

	// 随机验证码
	VerifyCode *string `json:"VerifyCode,omitnil" name:"VerifyCode"`
}

func (r *VerifyPresentationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPresentationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Did")
	delete(f, "PresentationData")
	delete(f, "DAPId")
	delete(f, "VerifyCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyPresentationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyPresentationResponseParams struct {
	// 是否验证成功
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 验证返回码
	VerifyCode *uint64 `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 验证消息
	VerifyMessage *string `json:"VerifyMessage,omitnil" name:"VerifyMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyPresentationResponse struct {
	*tchttp.BaseResponse
	Response *VerifyPresentationResponseParams `json:"Response"`
}

func (r *VerifyPresentationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyPresentationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
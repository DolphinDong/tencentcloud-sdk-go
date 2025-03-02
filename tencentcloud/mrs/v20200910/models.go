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

package v20200910

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AdmissionConditionBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type AdmissionDiagnosisBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Advice struct {
	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`
}

type AspectRatio struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *string `json:"Number,omitnil" name:"Number"`

	// 关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *string `json:"Relation,omitnil" name:"Relation"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type BaseInfo struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 标准值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type BaseItem struct {
	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原始文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化后值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// 四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type BaseItem2 struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原始文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化后值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type BaseItem3 struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原始文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化后值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`

	// 第几次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *int64 `json:"Order,omitnil" name:"Order"`
}

type BiopsyPart struct {
	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type BirthCert struct {
	// 新生儿信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeonatalInfo *NeonatalInfo `json:"NeonatalInfo,omitnil" name:"NeonatalInfo"`

	// 母亲信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MotherInfo *ParentInfo `json:"MotherInfo,omitnil" name:"MotherInfo"`

	// 父亲信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FatherInfo *ParentInfo `json:"FatherInfo,omitnil" name:"FatherInfo"`

	// 签发信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueInfo *IssueInfo `json:"IssueInfo,omitnil" name:"IssueInfo"`
}

type BirthPlaceBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Block struct {
	// 诊断信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Check []*Check `json:"Check,omitnil" name:"Check"`

	// 病理报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pathology []*PathologyReport `json:"Pathology,omitnil" name:"Pathology"`

	// 医学资料
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedDoc []*MedDoc `json:"MedDoc,omitnil" name:"MedDoc"`

	// 诊断证明
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagCert []*DiagCert `json:"DiagCert,omitnil" name:"DiagCert"`

	// 病案首页
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstPage []*FirstPage `json:"FirstPage,omitnil" name:"FirstPage"`

	// 检验报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicator []*Indicator `json:"Indicator,omitnil" name:"Indicator"`

	// 门诊病历信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalRecordInfo []*MedicalRecordInfo `json:"MedicalRecordInfo,omitnil" name:"MedicalRecordInfo"`

	// 出入院信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hospitalization []*Hospitalization `json:"Hospitalization,omitnil" name:"Hospitalization"`

	// 手术记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Surgery []*Surgery `json:"Surgery,omitnil" name:"Surgery"`

	// 处方单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prescription []*Prescription `json:"Prescription,omitnil" name:"Prescription"`

	// 免疫接种证明
	// 注意：此字段可能返回 null，表示取不到有效值。
	VaccineCertificate []*VaccineCertificate `json:"VaccineCertificate,omitnil" name:"VaccineCertificate"`

	// 心电图
	// 注意：此字段可能返回 null，表示取不到有效值。
	Electrocardiogram []*Electrocardiogram `json:"Electrocardiogram,omitnil" name:"Electrocardiogram"`

	// 病理报告v2
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologyV2 []*PathologyV2 `json:"PathologyV2,omitnil" name:"PathologyV2"`

	// 内窥镜报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endoscopy []*Endoscopy `json:"Endoscopy,omitnil" name:"Endoscopy"`

	// C14检验报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	C14 []*Indicator `json:"C14,omitnil" name:"C14"`

	// 体检结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exame []*Exame `json:"Exame,omitnil" name:"Exame"`

	// 出入院结构体
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedDocV2 []*DischargeInfoBlock `json:"MedDocV2,omitnil" name:"MedDocV2"`

	// 检验报告v3
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndicatorV3 []*IndicatorV3 `json:"IndicatorV3,omitnil" name:"IndicatorV3"`

	// 孕产报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Maternity []*Maternity `json:"Maternity,omitnil" name:"Maternity"`

	// 时间轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeline []*TimelineInformation `json:"Timeline,omitnil" name:"Timeline"`

	// 核酸报告结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Covid []*CovidItemsInfo `json:"Covid,omitnil" name:"Covid"`

	// 眼科报告结构体
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eye []*EyeItemsInfo `json:"Eye,omitnil" name:"Eye"`

	// 出生证明结构化信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthCert []*BirthCert `json:"BirthCert,omitnil" name:"BirthCert"`

	// 文本类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTypeListBlocks []*TextTypeListBlock `json:"TextTypeListBlocks,omitnil" name:"TextTypeListBlocks"`

	// 体检报告信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhysicalExamination *PhysicalExaminationV1 `json:"PhysicalExamination,omitnil" name:"PhysicalExamination"`
}

type BlockInfo struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitnil" name:"Positive"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size []*Size `json:"Size,omitnil" name:"Size"`
}

type BlockInfoV2 struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 疾病编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`
}

type BlockTitle struct {
	// name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// src
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type BloodPressureBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 舒张压
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormDiastolic *string `json:"NormDiastolic,omitnil" name:"NormDiastolic"`

	// 收缩压
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormSystolic *string `json:"NormSystolic,omitnil" name:"NormSystolic"`
}

type BloodPressureItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 项目原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *PhysicalBaseItem `json:"Item,omitnil" name:"Item"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PhysicalBaseItem `json:"Result,omitnil" name:"Result"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *PhysicalBaseItem `json:"Unit,omitnil" name:"Unit"`

	// 第几次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Times *PhysicalBaseItem `json:"Times,omitnil" name:"Times"`

	// 左右手臂
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *PhysicalBaseItem `json:"Location,omitnil" name:"Location"`
}

type BodyExaminationBlock struct {
	// 体温
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyTemperature *BodyTemperatureBlock `json:"BodyTemperature,omitnil" name:"BodyTemperature"`

	// 脉搏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pulse *BodyTemperatureBlock `json:"Pulse,omitnil" name:"Pulse"`

	// 呼吸
	// 注意：此字段可能返回 null，表示取不到有效值。
	Breathe *BodyTemperatureBlock `json:"Breathe,omitnil" name:"Breathe"`

	// 血压
	// 注意：此字段可能返回 null，表示取不到有效值。
	BloodPressure *BloodPressureBlock `json:"BloodPressure,omitnil" name:"BloodPressure"`
}

type BodyTemperatureBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Check struct {
	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *Desc `json:"Desc,omitnil" name:"Desc"`

	// 结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *Summary `json:"Summary,omitnil" name:"Summary"`

	// 检查报告块标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockTitle []*BlockTitle `json:"BlockTitle,omitnil" name:"BlockTitle"`
}

type ChestCircumferenceItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 项目原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *PhysicalBaseItem `json:"Item,omitnil" name:"Item"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PhysicalBaseItem `json:"Result,omitnil" name:"Result"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *PhysicalBaseItem `json:"Unit,omitnil" name:"Unit"`

	// 呼吸状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *PhysicalBaseItem `json:"State,omitnil" name:"State"`
}

type ChiefComplaintBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 单位输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 主诉详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*ChiefComplaintDetailBlock `json:"Detail,omitnil" name:"Detail"`
}

type ChiefComplaintDetailBlock struct {
	// 疾病名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseName *string `json:"DiseaseName,omitnil" name:"DiseaseName"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitnil" name:"Part"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil" name:"Time"`

	// 时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitnil" name:"TimeType"`
}

type ClinicalStaging struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Coord struct {
	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*Point `json:"Points,omitnil" name:"Points"`
}

type Coordinate struct {
	// 左上角x坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *int64 `json:"X,omitnil" name:"X"`

	// 左上角y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 宽度，单位像素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 高度，单位像素
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type CovidItem struct {
	// 采样时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleTime *BaseItem `json:"SampleTime,omitnil" name:"SampleTime"`

	// 检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestTime *BaseItem `json:"TestTime,omitnil" name:"TestTime"`

	// 检测机构
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestOrganization *BaseItem `json:"TestOrganization,omitnil" name:"TestOrganization"`

	// 检测结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestResult *BaseItem `json:"TestResult,omitnil" name:"TestResult"`

	// 健康码颜色
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeColor *BaseItem `json:"CodeColor,omitnil" name:"CodeColor"`
}

type CovidItemsInfo struct {
	// 核酸报告结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	CovidItems []*CovidItem `json:"CovidItems,omitnil" name:"CovidItems"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`
}

type DeathDateBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Desc struct {
	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 器官
	// 注意：此字段可能返回 null，表示取不到有效值。
	Organ []*Organ `json:"Organ,omitnil" name:"Organ"`

	// 结节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tuber []*TuberInfo `json:"Tuber,omitnil" name:"Tuber"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type DescInfo struct {
	// 描述段落文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *BaseInfo `json:"Text,omitnil" name:"Text"`

	// 描述段落详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infos []*DetailInformation `json:"Infos,omitnil" name:"Infos"`
}

type DetailInformation struct {
	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 组织大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	TissueSizes []*Size `json:"TissueSizes,omitnil" name:"TissueSizes"`

	// 结节大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	TuberSizes []*Size `json:"TuberSizes,omitnil" name:"TuberSizes"`

	// 肿瘤大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancerSizes []*Size `json:"CancerSizes,omitnil" name:"CancerSizes"`

	// 组织学等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyLevel *BaseInfo `json:"HistologyLevel,omitnil" name:"HistologyLevel"`

	// 组织学类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyType *HistologyTypeV2 `json:"HistologyType,omitnil" name:"HistologyType"`

	// 侵犯
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invasive []*InvasiveV2 `json:"Invasive,omitnil" name:"Invasive"`

	// pTNM
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *PTNM `json:"PTNM,omitnil" name:"PTNM"`

	// 浸润深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfiltrationDepth *BaseInfo `json:"InfiltrationDepth,omitnil" name:"InfiltrationDepth"`

	// 结节数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TuberNum *BaseInfo `json:"TuberNum,omitnil" name:"TuberNum"`

	// 钙化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Calcification *BaseInfo `json:"Calcification,omitnil" name:"Calcification"`

	// 坏死
	// 注意：此字段可能返回 null，表示取不到有效值。
	Necrosis *BaseInfo `json:"Necrosis,omitnil" name:"Necrosis"`

	// 异形
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormity *BaseInfo `json:"Abnormity,omitnil" name:"Abnormity"`

	// 断链
	// 注意：此字段可能返回 null，表示取不到有效值。
	Breaked *BaseInfo `json:"Breaked,omitnil" name:"Breaked"`
}

type DiagCert struct {
	// 建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advice *Advice `json:"Advice,omitnil" name:"Advice"`

	// 诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnosis []*DiagCertItem `json:"Diagnosis,omitnil" name:"Diagnosis"`
}

type DiagCertItem struct {
	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil" name:"Value"`
}

type DischargeConditionBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`
}

type DischargeDiagnosis struct {
	// 表格位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIndex *int64 `json:"TableIndex,omitnil" name:"TableIndex"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutDiagnosis *BlockInfo `json:"OutDiagnosis,omitnil" name:"OutDiagnosis"`

	// 疾病编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseCode *BlockInfo `json:"DiseaseCode,omitnil" name:"DiseaseCode"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	InStatus *BlockInfo `json:"InStatus,omitnil" name:"InStatus"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutStatus *BlockInfo `json:"OutStatus,omitnil" name:"OutStatus"`
}

type DischargeDiagnosisBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type DischargeInfoBlock struct {
	// 疾病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseHistory *DiseaseHistoryBlock `json:"DiseaseHistory,omitnil" name:"DiseaseHistory"`

	// 个人史
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonalHistory *PersonalHistoryBlock `json:"PersonalHistory,omitnil" name:"PersonalHistory"`

	// 药物史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrugHistory *DrugHistoryBlock `json:"DrugHistory,omitnil" name:"DrugHistory"`

	// 治疗相关
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentRecord *TreatmentRecordBlock `json:"TreatmentRecord,omitnil" name:"TreatmentRecord"`

	// 文本段落
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParagraphBlock *ParagraphBlock `json:"ParagraphBlock,omitnil" name:"ParagraphBlock"`
}

type DiseaseHistoryBlock struct {
	// 主要病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDiseaseHistory *MainDiseaseHistoryBlock `json:"MainDiseaseHistory,omitnil" name:"MainDiseaseHistory"`

	// 过敏史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllergyHistory *MainDiseaseHistoryBlock `json:"AllergyHistory,omitnil" name:"AllergyHistory"`

	// 注射史
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfectHistory *MainDiseaseHistoryBlock `json:"InfectHistory,omitnil" name:"InfectHistory"`

	// 手术史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryHistory *SurgeryHistoryBlock `json:"SurgeryHistory,omitnil" name:"SurgeryHistory"`

	// 输血史
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransfusionHistory *TransfusionHistoryBlock `json:"TransfusionHistory,omitnil" name:"TransfusionHistory"`
}

type DiseaseMedicalHistory struct {
	// 主病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDiseaseHistory *string `json:"MainDiseaseHistory,omitnil" name:"MainDiseaseHistory"`

	// 过敏史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllergyHistory *string `json:"AllergyHistory,omitnil" name:"AllergyHistory"`

	// 传染疾病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfectHistory *string `json:"InfectHistory,omitnil" name:"InfectHistory"`

	// 手术史
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationHistory *string `json:"OperationHistory,omitnil" name:"OperationHistory"`

	// 输血史
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransfusionHistory *string `json:"TransfusionHistory,omitnil" name:"TransfusionHistory"`
}

type DiseasePresentBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 归一化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`
}

type DosageBlock struct {
	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 单次计量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SingleMeasurement *string `json:"SingleMeasurement,omitnil" name:"SingleMeasurement"`

	// 频次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Frequency *string `json:"Frequency,omitnil" name:"Frequency"`

	// 给药途径
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrugDeliveryRoute *string `json:"DrugDeliveryRoute,omitnil" name:"DrugDeliveryRoute"`
}

type DrugHistoryBlock struct {
	// 药品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 药物列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrugList []*DrugListBlock `json:"DrugList,omitnil" name:"DrugList"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type DrugListBlock struct {
	// 通用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonName *string `json:"CommonName,omitnil" name:"CommonName"`

	// 商品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeName *string `json:"TradeName,omitnil" name:"TradeName"`

	// 用法用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dosage *DosageBlock `json:"Dosage,omitnil" name:"Dosage"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type EcgDescription struct {
	// 心率
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeartRate *EcgItem `json:"HeartRate,omitnil" name:"HeartRate"`

	// 心房率
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuricularRate *EcgItem `json:"AuricularRate,omitnil" name:"AuricularRate"`

	// 心室率
	// 注意：此字段可能返回 null，表示取不到有效值。
	VentricularRate *EcgItem `json:"VentricularRate,omitnil" name:"VentricularRate"`

	// 节律
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rhythm *EcgItem `json:"Rhythm,omitnil" name:"Rhythm"`

	// P波时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PDuration *EcgItem `json:"PDuration,omitnil" name:"PDuration"`

	// QRS时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	QrsDuration *EcgItem `json:"QrsDuration,omitnil" name:"QrsDuration"`

	// QRS电轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	QrsAxis *EcgItem `json:"QrsAxis,omitnil" name:"QrsAxis"`

	// P-R间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PRInterval *EcgItem `json:"PRInterval,omitnil" name:"PRInterval"`

	// P-P间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PPInterval *EcgItem `json:"PPInterval,omitnil" name:"PPInterval"`

	// R-R间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RRInterval *EcgItem `json:"RRInterval,omitnil" name:"RRInterval"`

	// P-J间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PJInterval *EcgItem `json:"PJInterval,omitnil" name:"PJInterval"`

	// Q-T间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	QTInterval *EcgItem `json:"QTInterval,omitnil" name:"QTInterval"`

	// qt/qtc间期
	// 注意：此字段可能返回 null，表示取不到有效值。
	QTCInterval *EcgItem `json:"QTCInterval,omitnil" name:"QTCInterval"`

	// RV5/SV1振幅
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rv5SV1Amplitude *EcgItem `json:"Rv5SV1Amplitude,omitnil" name:"Rv5SV1Amplitude"`

	// RV5+SV1振幅
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rv5AddSV1Amplitude *EcgItem `json:"Rv5AddSV1Amplitude,omitnil" name:"Rv5AddSV1Amplitude"`

	// PRT电轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	PRTAxis *EcgItem `json:"PRTAxis,omitnil" name:"PRTAxis"`

	// RV5振幅
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rv5Amplitude *EcgItem `json:"Rv5Amplitude,omitnil" name:"Rv5Amplitude"`

	// SV1振幅
	// 注意：此字段可能返回 null，表示取不到有效值。
	SV1Amplitude *EcgItem `json:"SV1Amplitude,omitnil" name:"SV1Amplitude"`

	// RV6/SV2
	// 注意：此字段可能返回 null，表示取不到有效值。
	RV6SV2 *EcgItem `json:"RV6SV2,omitnil" name:"RV6SV2"`

	// P/QRS/T电轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	PQRSTAxis *EcgItem `json:"PQRSTAxis,omitnil" name:"PQRSTAxis"`
}

type EcgDiagnosis struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil" name:"Value"`
}

type EcgItem struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`
}

type Elastic struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *string `json:"Score,omitnil" name:"Score"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type Electrocardiogram struct {
	// 心电图详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	EcgDescription *EcgDescription `json:"EcgDescription,omitnil" name:"EcgDescription"`

	// 心电图诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	EcgDiagnosis *EcgDiagnosis `json:"EcgDiagnosis,omitnil" name:"EcgDiagnosis"`
}

type Endoscopy struct {
	// 活检部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	BiopsyPart *BiopsyPart `json:"BiopsyPart,omitnil" name:"BiopsyPart"`

	// 可见描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *EndoscopyDesc `json:"Desc,omitnil" name:"Desc"`

	// 结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *Summary `json:"Summary,omitnil" name:"Summary"`
}

type EndoscopyDesc struct {
	// 描述内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 器官
	// 注意：此字段可能返回 null，表示取不到有效值。
	Organ []*EndoscopyOrgan `json:"Organ,omitnil" name:"Organ"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type EndoscopyOrgan struct {
	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 部位别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartAlias *string `json:"PartAlias,omitnil" name:"PartAlias"`

	// 症状描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymDescList []*BlockInfo `json:"SymDescList,omitnil" name:"SymDescList"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type Exame struct {
	// 结论段落
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverView []*ResultInfo `json:"OverView,omitnil" name:"OverView"`

	// 异常与建议段落
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormality []*ResultInfo `json:"Abnormality,omitnil" name:"Abnormality"`
}

type EyeChildItem struct {
	// 球镜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sph []*BaseItem3 `json:"Sph,omitnil" name:"Sph"`

	// 柱镜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cyl []*BaseItem3 `json:"Cyl,omitnil" name:"Cyl"`

	// 轴位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ax []*BaseItem3 `json:"Ax,omitnil" name:"Ax"`

	// 等效球镜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Se *BaseItem2 `json:"Se,omitnil" name:"Se"`
}

type EyeItem struct {
	// 左眼
	// 注意：此字段可能返回 null，表示取不到有效值。
	Left *EyeChildItem `json:"Left,omitnil" name:"Left"`

	// 右眼
	// 注意：此字段可能返回 null，表示取不到有效值。
	Right *EyeChildItem `json:"Right,omitnil" name:"Right"`

	// 瞳距
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pd *BaseItem2 `json:"Pd,omitnil" name:"Pd"`
}

type EyeItemsInfo struct {
	// 眼科报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	EyeItems *EyeItem `json:"EyeItems,omitnil" name:"EyeItems"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`
}

type FamilyHistoryBlock struct {
	// 家庭成员
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeHistory *RelativeHistoryBlock `json:"RelativeHistory,omitnil" name:"RelativeHistory"`

	// 家族肿瘤史
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeCancerHistory *RelativeCancerHistoryBlock `json:"RelativeCancerHistory,omitnil" name:"RelativeCancerHistory"`

	// 家族遗传史
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneticHistory *GeneticHistoryBlock `json:"GeneticHistory,omitnil" name:"GeneticHistory"`
}

type FamilyMedicalHistory struct {
	// 家族成员史
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeHistory *string `json:"RelativeHistory,omitnil" name:"RelativeHistory"`

	// 家族肿瘤史
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeCancerHistory *string `json:"RelativeCancerHistory,omitnil" name:"RelativeCancerHistory"`

	// 家族遗传史
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneticHistory *string `json:"GeneticHistory,omitnil" name:"GeneticHistory"`
}

type FertilityHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil" name:"State"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 妊娠次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PregCount *string `json:"PregCount,omitnil" name:"PregCount"`

	// 生产次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProduCount *string `json:"ProduCount,omitnil" name:"ProduCount"`
}

type Fetus struct {
	// 双顶径
	// 注意：此字段可能返回 null，表示取不到有效值。
	BPD *FieldInfo `json:"BPD,omitnil" name:"BPD"`

	// 腹前后径
	// 注意：此字段可能返回 null，表示取不到有效值。
	APTD *FieldInfo `json:"APTD,omitnil" name:"APTD"`

	// 腹左右径
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTD *FieldInfo `json:"TTD,omitnil" name:"TTD"`

	// 头臀径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CRL *FieldInfo `json:"CRL,omitnil" name:"CRL"`

	// 头围
	// 注意：此字段可能返回 null，表示取不到有效值。
	HC *FieldInfo `json:"HC,omitnil" name:"HC"`

	// 腹围
	// 注意：此字段可能返回 null，表示取不到有效值。
	AC *FieldInfo `json:"AC,omitnil" name:"AC"`

	// 股骨长
	// 注意：此字段可能返回 null，表示取不到有效值。
	FL *FieldInfo `json:"FL,omitnil" name:"FL"`

	// 肱骨长
	// 注意：此字段可能返回 null，表示取不到有效值。
	HL *FieldInfo `json:"HL,omitnil" name:"HL"`

	// 胎儿重量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *FieldInfo `json:"Weight,omitnil" name:"Weight"`

	// 颈项透明层
	// 注意：此字段可能返回 null，表示取不到有效值。
	NT *FieldInfo `json:"NT,omitnil" name:"NT"`

	// 脐动脉血流
	// 注意：此字段可能返回 null，表示取不到有效值。
	UmbilicalCord *FieldInfo `json:"UmbilicalCord,omitnil" name:"UmbilicalCord"`

	// 羊水最大深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterDeep *FieldInfo `json:"WaterDeep,omitnil" name:"WaterDeep"`

	// 羊水四象限测量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterQuad *FieldInfo `json:"WaterQuad,omitnil" name:"WaterQuad"`

	// 羊水指数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AFI *FieldInfo `json:"AFI,omitnil" name:"AFI"`

	// 胎心
	// 注意：此字段可能返回 null，表示取不到有效值。
	FHR *FieldInfo `json:"FHR,omitnil" name:"FHR"`

	// 胎动
	// 注意：此字段可能返回 null，表示取不到有效值。
	Movement *FieldInfo `json:"Movement,omitnil" name:"Movement"`

	// 胎数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Num *FieldInfo `json:"Num,omitnil" name:"Num"`

	// 胎位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *FieldInfo `json:"Position,omitnil" name:"Position"`

	// 是否活胎
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alive *FieldInfo `json:"Alive,omitnil" name:"Alive"`

	// 胎盘位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacentaLocation *FieldInfo `json:"PlacentaLocation,omitnil" name:"PlacentaLocation"`

	// 胎盘厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacentaThickness *FieldInfo `json:"PlacentaThickness,omitnil" name:"PlacentaThickness"`

	// 胎盘成熟度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacentaGrade *FieldInfo `json:"PlacentaGrade,omitnil" name:"PlacentaGrade"`

	// 妊娠时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	GestationTime *FieldInfo `json:"GestationTime,omitnil" name:"GestationTime"`

	// 妊娠周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	GestationPeriod *FieldInfo `json:"GestationPeriod,omitnil" name:"GestationPeriod"`

	// 绕颈
	// 注意：此字段可能返回 null，表示取不到有效值。
	AroundNeck *FieldInfo `json:"AroundNeck,omitnil" name:"AroundNeck"`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sym []*FieldInfo `json:"Sym,omitnil" name:"Sym"`

	// 原文内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`
}

type FieldInfo struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nums []*NumValue `json:"Nums,omitnil" name:"Nums"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`
}

type FirstPage struct {
	// 出入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis []*DischargeDiagnosis `json:"DischargeDiagnosis,omitnil" name:"DischargeDiagnosis"`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosis *BlockInfo `json:"PathologicalDiagnosis,omitnil" name:"PathologicalDiagnosis"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClinicalDiagnosis *BlockInfo `json:"ClinicalDiagnosis,omitnil" name:"ClinicalDiagnosis"`

	// 受伤中毒的外部原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DamagePoi *BlockInfoV2 `json:"DamagePoi,omitnil" name:"DamagePoi"`

	// 病案首页第二页
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fp2NdItems []*Fp2NdItem `json:"Fp2NdItems,omitnil" name:"Fp2NdItems"`
}

type Fp2NdItem struct {
	// 手术编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *BaseItem `json:"Code,omitnil" name:"Code"`

	// 手术名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *BaseItem `json:"Name,omitnil" name:"Name"`

	// 手术开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *BaseItem `json:"StartTime,omitnil" name:"StartTime"`

	// 手术结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *BaseItem `json:"EndTime,omitnil" name:"EndTime"`

	// 手术等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *BaseItem `json:"Level,omitnil" name:"Level"`

	// 手术类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *BaseItem `json:"Type,omitnil" name:"Type"`

	// 醉愈合方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealingGrade *BaseItem `json:"IncisionHealingGrade,omitnil" name:"IncisionHealingGrade"`

	// 麻醉方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnesthesiaMethod *BaseItem `json:"AnesthesiaMethod,omitnil" name:"AnesthesiaMethod"`
}

type GeneralExaminationBaseItem struct {
	// 生命体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	VitalSign *GeneralExaminationVitalSign `json:"VitalSign,omitnil" name:"VitalSign"`

	// 其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others *GeneralExaminationOthers `json:"Others,omitnil" name:"Others"`

	// 小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefSummary *GeneralExaminationBriefSummary `json:"BriefSummary,omitnil" name:"BriefSummary"`
}

type GeneralExaminationBriefSummary struct {
	// 一般检查小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type GeneralExaminationOthers struct {
	// 面容与表情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Countenance *KeyValueItem `json:"Countenance,omitnil" name:"Countenance"`

	// 精神状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	MentalStatus *KeyValueItem `json:"MentalStatus,omitnil" name:"MentalStatus"`

	// 发育及营养状况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevelopmentCondition *KeyValueItem `json:"DevelopmentCondition,omitnil" name:"DevelopmentCondition"`

	// 记忆力
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *KeyValueItem `json:"Memory,omitnil" name:"Memory"`

	// 臀围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hipline *ValueUnitItem `json:"Hipline,omitnil" name:"Hipline"`

	// 腰臀比
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaistHipRatio *ValueUnitItem `json:"WaistHipRatio,omitnil" name:"WaistHipRatio"`

	// 生活嗜好
	// 注意：此字段可能返回 null，表示取不到有效值。
	Addiction *KeyValueItem `json:"Addiction,omitnil" name:"Addiction"`

	// 生活能力评定
	// 注意：此字段可能返回 null，表示取不到有效值。
	AbilityOfLifeADL *KeyValueItem `json:"AbilityOfLifeADL,omitnil" name:"AbilityOfLifeADL"`

	// 一般检查其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others []*KeyValueItem `json:"Others,omitnil" name:"Others"`

	// 胸围
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChestCircumference *ChestCircumferenceItem `json:"ChestCircumference,omitnil" name:"ChestCircumference"`
}

type GeneralExaminationVitalSign struct {
	// 生命体征总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *ValueUnitItem `json:"Text,omitnil" name:"Text"`

	// 体温
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyTemperature *ValueUnitItem `json:"BodyTemperature,omitnil" name:"BodyTemperature"`

	// 脉率
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pulse *ValueUnitItem `json:"Pulse,omitnil" name:"Pulse"`

	// 心率
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeartRate *ValueUnitItem `json:"HeartRate,omitnil" name:"HeartRate"`

	// 呼吸频率
	// 注意：此字段可能返回 null，表示取不到有效值。
	BreathingRate *ValueUnitItem `json:"BreathingRate,omitnil" name:"BreathingRate"`

	// 身高
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyHeight *ValueUnitItem `json:"BodyHeight,omitnil" name:"BodyHeight"`

	// 体重
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyWeight *ValueUnitItem `json:"BodyWeight,omitnil" name:"BodyWeight"`

	// 体质指数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyMassIndex *ValueUnitItem `json:"BodyMassIndex,omitnil" name:"BodyMassIndex"`

	// 腰围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Waistline *ValueUnitItem `json:"Waistline,omitnil" name:"Waistline"`

	// 血压
	// 注意：此字段可能返回 null，表示取不到有效值。
	BloodPressure *GeneralExaminationVitalSignBloodPressure `json:"BloodPressure,omitnil" name:"BloodPressure"`
}

type GeneralExaminationVitalSignBloodPressure struct {
	// 血压
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *BloodPressureItem `json:"Text,omitnil" name:"Text"`

	// 收缩压/舒张压
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystolicDiastolicPressure []*BloodPressureItem `json:"SystolicDiastolicPressure,omitnil" name:"SystolicDiastolicPressure"`

	// 收缩压
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystolicPressure []*BloodPressureItem `json:"SystolicPressure,omitnil" name:"SystolicPressure"`

	// 舒张压
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiastolicPressure []*BloodPressureItem `json:"DiastolicPressure,omitnil" name:"DiastolicPressure"`
}

type GeneticHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 遗传列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneticList *string `json:"GeneticList,omitnil" name:"GeneticList"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type GynaecologyAdnexal struct {
	// 子宫附件总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type GynaecologyBaseItem struct {
	// 外阴
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vulva *GynaecologyVulva `json:"Vulva,omitnil" name:"Vulva"`

	// 阴道
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vagina *GynaecologyVagina `json:"Vagina,omitnil" name:"Vagina"`

	// 子宫颈
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cervix *GynaecologyCervix `json:"Cervix,omitnil" name:"Cervix"`

	// 子宫
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uterus *GynaecologyUterus `json:"Uterus,omitnil" name:"Uterus"`

	// 子宫附件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Adnexal *GynaecologyAdnexal `json:"Adnexal,omitnil" name:"Adnexal"`

	// 盆腔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PelvicCavity *GynaecologyPelvicCavity `json:"PelvicCavity,omitnil" name:"PelvicCavity"`

	// 妇科其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others []*KeyValueItem `json:"Others,omitnil" name:"Others"`

	// 月经史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualHistory *GynaecologyMenstrualHistory `json:"MenstrualHistory,omitnil" name:"MenstrualHistory"`

	// 小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefSummary *GynaecologyBriefSummary `json:"BriefSummary,omitnil" name:"BriefSummary"`
}

type GynaecologyBriefSummary struct {
	// 小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type GynaecologyCervix struct {
	// 子宫颈总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type GynaecologyMenstrualHistory struct {
	// 妇科月经史总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type GynaecologyPelvicCavity struct {
	// 盆腔总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type GynaecologyUterus struct {
	// 子宫总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type GynaecologyVagina struct {
	// 阴道总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type GynaecologyVulva struct {
	// 外阴总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type HandleParam struct {
	// ocr引擎
	OcrEngineType *int64 `json:"OcrEngineType,omitnil" name:"OcrEngineType"`

	// 是否返回分行文本内容
	IsReturnText *bool `json:"IsReturnText,omitnil" name:"IsReturnText"`

	// 顺时针旋转角度
	RotateTheAngle *float64 `json:"RotateTheAngle,omitnil" name:"RotateTheAngle"`

	// 自动适配方向,仅支持优图引擎
	AutoFitDirection *bool `json:"AutoFitDirection,omitnil" name:"AutoFitDirection"`

	// 坐标优化
	AutoOptimizeCoordinate *bool `json:"AutoOptimizeCoordinate,omitnil" name:"AutoOptimizeCoordinate"`

	// 是否开启图片压缩，开启时imageOriginalSize必须正确填写
	IsScale *bool `json:"IsScale,omitnil" name:"IsScale"`

	// 原始图片大小(单位byes),用来判断该图片是否需要压缩
	ImageOriginalSize *uint64 `json:"ImageOriginalSize,omitnil" name:"ImageOriginalSize"`

	// 采用后台默认值(2048Kb)
	ScaleTargetSize *uint64 `json:"ScaleTargetSize,omitnil" name:"ScaleTargetSize"`
}

type HearingItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 项目原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *PhysicalBaseItem `json:"Item,omitnil" name:"Item"`

	// 方位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *PhysicalBaseItem `json:"Location,omitnil" name:"Location"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PhysicalBaseItem `json:"Result,omitnil" name:"Result"`
}

type HistologyClass struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type HistologyLevel struct {
	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitnil" name:"Grade"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`
}

type HistologyType struct {
	// 浸润
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infiltration *string `json:"Infiltration,omitnil" name:"Infiltration"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type HistologyTypeV2 struct {
	// 浸润
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infiltration *string `json:"Infiltration,omitnil" name:"Infiltration"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化后的组织学类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type Hospitalization struct {
	// 入院时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionTime *string `json:"AdmissionTime,omitnil" name:"AdmissionTime"`

	// 出院时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeTime *string `json:"DischargeTime,omitnil" name:"DischargeTime"`

	// 住院天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDays *string `json:"AdmissionDays,omitnil" name:"AdmissionDays"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: AdmissionDignosis is deprecated.
	AdmissionDignosis *string `json:"AdmissionDignosis,omitnil" name:"AdmissionDignosis"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionCondition *string `json:"AdmissionCondition,omitnil" name:"AdmissionCondition"`

	// 诊疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnosisTreatment *string `json:"DiagnosisTreatment,omitnil" name:"DiagnosisTreatment"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis *string `json:"DischargeDiagnosis,omitnil" name:"DischargeDiagnosis"`

	// 出院医嘱
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeInstruction *string `json:"DischargeInstruction,omitnil" name:"DischargeInstruction"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDiagnosis *string `json:"AdmissionDiagnosis,omitnil" name:"AdmissionDiagnosis"`
}

type IHCBlock struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 具体值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *ValueBlock `json:"Value,omitnil" name:"Value"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type IHCInfo struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	Value *Value `json:"Value,omitnil" name:"Value"`
}

type IHCV2 struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// ihc归一化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// ihc详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *Value `json:"Value,omitnil" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type ImageInfo struct {
	// 图片id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 图片url
	Url *string `json:"Url,omitnil" name:"Url"`

	// 图片base64编码
	Base64 *string `json:"Base64,omitnil" name:"Base64"`
}

// Predefined struct for user
type ImageMaskAsyncGetResultRequestParams struct {
	// 异步任务ID
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`
}

type ImageMaskAsyncGetResultRequest struct {
	*tchttp.BaseRequest
	
	// 异步任务ID
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`
}

func (r *ImageMaskAsyncGetResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskAsyncGetResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageMaskAsyncGetResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageMaskAsyncGetResultResponseParams struct {
	// 脱敏后图片的base64编码
	MaskedImage *string `json:"MaskedImage,omitnil" name:"MaskedImage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImageMaskAsyncGetResultResponse struct {
	*tchttp.BaseResponse
	Response *ImageMaskAsyncGetResultResponseParams `json:"Response"`
}

func (r *ImageMaskAsyncGetResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskAsyncGetResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageMaskAsyncRequestParams struct {
	// 图片信息,目前只支持传图片base64
	Image *ImageInfo `json:"Image,omitnil" name:"Image"`

	// 图片脱敏选项, 不传默认都脱敏
	MaskFlag *ImageMaskFlags `json:"MaskFlag,omitnil" name:"MaskFlag"`
}

type ImageMaskAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 图片信息,目前只支持传图片base64
	Image *ImageInfo `json:"Image,omitnil" name:"Image"`

	// 图片脱敏选项, 不传默认都脱敏
	MaskFlag *ImageMaskFlags `json:"MaskFlag,omitnil" name:"MaskFlag"`
}

func (r *ImageMaskAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "MaskFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageMaskAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageMaskAsyncResponseParams struct {
	// 加密任务ID
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImageMaskAsyncResponse struct {
	*tchttp.BaseResponse
	Response *ImageMaskAsyncResponseParams `json:"Response"`
}

func (r *ImageMaskAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageMaskFlags struct {
	// 是否对医院信息进行脱敏
	HospitalFlag *bool `json:"HospitalFlag,omitnil" name:"HospitalFlag"`

	// 是否对医生信息进行脱敏
	DoctorFlag *bool `json:"DoctorFlag,omitnil" name:"DoctorFlag"`

	// 是否对患者信息进行脱敏
	PatientFlag *bool `json:"PatientFlag,omitnil" name:"PatientFlag"`

	// 是否对二维码信息进行脱敏
	BarFlag *bool `json:"BarFlag,omitnil" name:"BarFlag"`
}

// Predefined struct for user
type ImageMaskRequestParams struct {
	// 图片信息,目前只支持传图片base64
	Image *ImageInfo `json:"Image,omitnil" name:"Image"`

	// 图片脱敏选项, 不传默认都脱敏
	MaskFlag *ImageMaskFlags `json:"MaskFlag,omitnil" name:"MaskFlag"`
}

type ImageMaskRequest struct {
	*tchttp.BaseRequest
	
	// 图片信息,目前只支持传图片base64
	Image *ImageInfo `json:"Image,omitnil" name:"Image"`

	// 图片脱敏选项, 不传默认都脱敏
	MaskFlag *ImageMaskFlags `json:"MaskFlag,omitnil" name:"MaskFlag"`
}

func (r *ImageMaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "MaskFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageMaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageMaskResponseParams struct {
	// 脱敏后图片的Base64信息
	MaskedImage *string `json:"MaskedImage,omitnil" name:"MaskedImage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImageMaskResponse struct {
	*tchttp.BaseResponse
	Response *ImageMaskResponseParams `json:"Response"`
}

func (r *ImageMaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageMaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToClassRequestParams struct {
	// 图片列表，允许传入多张图片，支持传入图片的base64编码，暂不支持图片url
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitnil" name:"HandleParam"`

	// 不填，默认为0
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil" name:"UserType"`
}

type ImageToClassRequest struct {
	*tchttp.BaseRequest
	
	// 图片列表，允许传入多张图片，支持传入图片的base64编码，暂不支持图片url
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitnil" name:"HandleParam"`

	// 不填，默认为0
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil" name:"UserType"`
}

func (r *ImageToClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfoList")
	delete(f, "HandleParam")
	delete(f, "Type")
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToClassResponseParams struct {
	// 分类结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTypeList []*TextType `json:"TextTypeList,omitnil" name:"TextTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImageToClassResponse struct {
	*tchttp.BaseResponse
	Response *ImageToClassResponseParams `json:"Response"`
}

func (r *ImageToClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToObjectRequestParams struct {
	// 图片列表，允许传入多张图片，目前只支持传入图片base64编码，图片url暂不支持
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitnil" name:"HandleParam"`

	// 报告类型，目前支持11（检验报告），12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为 False，则 Type 字段不能为 0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitnil" name:"IsUsedClassify"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *int64 `json:"UserType,omitnil" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病历记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitnil" name:"ReportTypeVersion"`
}

type ImageToObjectRequest struct {
	*tchttp.BaseRequest
	
	// 图片列表，允许传入多张图片，目前只支持传入图片base64编码，图片url暂不支持
	ImageInfoList []*ImageInfo `json:"ImageInfoList,omitnil" name:"ImageInfoList"`

	// 图片处理参数
	HandleParam *HandleParam `json:"HandleParam,omitnil" name:"HandleParam"`

	// 报告类型，目前支持11（检验报告），12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为 False，则 Type 字段不能为 0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitnil" name:"IsUsedClassify"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *int64 `json:"UserType,omitnil" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病历记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitnil" name:"ReportTypeVersion"`
}

func (r *ImageToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfoList")
	delete(f, "HandleParam")
	delete(f, "Type")
	delete(f, "IsUsedClassify")
	delete(f, "UserType")
	delete(f, "ReportTypeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToObjectResponseParams struct {
	// 报告结构化结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Template *Template `json:"Template,omitnil" name:"Template"`

	// 多级分类结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTypeList []*TextType `json:"TextTypeList,omitnil" name:"TextTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImageToObjectResponse struct {
	*tchttp.BaseResponse
	Response *ImageToObjectResponseParams `json:"Response"`
}

func (r *ImageToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImmunohistochemistryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 免疫组化详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*IHCBlock `json:"Value,omitnil" name:"Value"`
}

type Indicator struct {
	// 检验指标项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicators []*IndicatorItem `json:"Indicators,omitnil" name:"Indicators"`

	// 检验报告块标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockTitle []*BlockTitle `json:"BlockTitle,omitnil" name:"BlockTitle"`
}

type IndicatorItem struct {
	// 英文缩写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 标准缩写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scode *string `json:"Scode,omitnil" name:"Scode"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 标准名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sname *string `json:"Sname,omitnil" name:"Sname"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitnil" name:"Result"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 参考范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *string `json:"Range,omitnil" name:"Range"`

	// 上下箭头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arrow *string `json:"Arrow,omitnil" name:"Arrow"`

	// 是否正常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Normal *bool `json:"Normal,omitnil" name:"Normal"`

	// 项目原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemString *string `json:"ItemString,omitnil" name:"ItemString"`

	// 指标项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 指标项坐标位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords *Coordinate `json:"Coords,omitnil" name:"Coords"`

	// 推测结果是否异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferNormal *string `json:"InferNormal,omitnil" name:"InferNormal"`
}

type IndicatorItemV2 struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *BaseItem `json:"Item,omitnil" name:"Item"`

	// 英文编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *BaseItem `json:"Code,omitnil" name:"Code"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *BaseItem `json:"Result,omitnil" name:"Result"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *BaseItem `json:"Unit,omitnil" name:"Unit"`

	// 参考范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Range *BaseItem `json:"Range,omitnil" name:"Range"`

	// 上下箭头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arrow *BaseItem `json:"Arrow,omitnil" name:"Arrow"`

	// 检测方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *BaseItem `json:"Method,omitnil" name:"Method"`

	// 结果是否异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Normal *bool `json:"Normal,omitnil" name:"Normal"`

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *int64 `json:"Order,omitnil" name:"Order"`

	// 推测结果是否异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferNormal *string `json:"InferNormal,omitnil" name:"InferNormal"`
}

type IndicatorV3 struct {
	// 检验报告V3结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: TableIndictors is deprecated.
	TableIndictors []*TableIndicators `json:"TableIndictors,omitnil" name:"TableIndictors"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`

	// 检验报告V3结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIndicators []*TableIndicators `json:"TableIndicators,omitnil" name:"TableIndicators"`
}

type InternalMedicineAbdomen struct {
	// 内科腹部小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 肝脏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Liver *InternalMedicineAbdomenLiver `json:"Liver,omitnil" name:"Liver"`

	// 胆囊
	// 注意：此字段可能返回 null，表示取不到有效值。
	GallBladder *InternalMedicineAbdomenGallBladder `json:"GallBladder,omitnil" name:"GallBladder"`

	// 胰腺
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pancreas *InternalMedicineAbdomenPancreas `json:"Pancreas,omitnil" name:"Pancreas"`

	// 脾脏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spleen *InternalMedicineAbdomenSpleen `json:"Spleen,omitnil" name:"Spleen"`

	// 肾脏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kidney *InternalMedicineAbdomenKidney `json:"Kidney,omitnil" name:"Kidney"`

	// 腹部其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others []*KeyValueItem `json:"Others,omitnil" name:"Others"`
}

type InternalMedicineAbdomenGallBladder struct {
	// 胆囊总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *KeyValueItem `json:"Src,omitnil" name:"Src"`

	// 胆囊大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *KeyValueItem `json:"Size,omitnil" name:"Size"`

	// 胆囊触诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Palpation *KeyValueItem `json:"Palpation,omitnil" name:"Palpation"`

	// 胆囊叩诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percussion *KeyValueItem `json:"Percussion,omitnil" name:"Percussion"`

	// 胆囊压痛
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tenderness *KeyValueItem `json:"Tenderness,omitnil" name:"Tenderness"`

	// 胆囊质地
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *KeyValueItem `json:"Consistency,omitnil" name:"Consistency"`
}

type InternalMedicineAbdomenKidney struct {
	// 肾脏总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *KeyValueItem `json:"Src,omitnil" name:"Src"`

	// 肾脏大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *KeyValueItem `json:"Size,omitnil" name:"Size"`

	// 肾脏触诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Palpation *KeyValueItem `json:"Palpation,omitnil" name:"Palpation"`

	// 肾脏叩诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percussion *KeyValueItem `json:"Percussion,omitnil" name:"Percussion"`

	// 肾脏压痛
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tenderness *KeyValueItem `json:"Tenderness,omitnil" name:"Tenderness"`

	// 肾脏质地
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *KeyValueItem `json:"Consistency,omitnil" name:"Consistency"`
}

type InternalMedicineAbdomenLiver struct {
	// 肝脏总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *KeyValueItem `json:"Src,omitnil" name:"Src"`

	// 肝脏大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *KeyValueItem `json:"Size,omitnil" name:"Size"`

	// 肝脏触诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Palpation *KeyValueItem `json:"Palpation,omitnil" name:"Palpation"`

	// 肝脏叩诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percussion *KeyValueItem `json:"Percussion,omitnil" name:"Percussion"`

	// 肝脏压痛
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tenderness *KeyValueItem `json:"Tenderness,omitnil" name:"Tenderness"`

	// 肝脏质地
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *KeyValueItem `json:"Consistency,omitnil" name:"Consistency"`
}

type InternalMedicineAbdomenPancreas struct {
	// 胰腺总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *KeyValueItem `json:"Src,omitnil" name:"Src"`

	// 胰腺大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *KeyValueItem `json:"Size,omitnil" name:"Size"`

	// 胰腺触诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Palpation *KeyValueItem `json:"Palpation,omitnil" name:"Palpation"`

	// 胰腺叩诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percussion *KeyValueItem `json:"Percussion,omitnil" name:"Percussion"`

	// 肝脏压痛
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tenderness *KeyValueItem `json:"Tenderness,omitnil" name:"Tenderness"`

	// 胰腺质地
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *KeyValueItem `json:"Consistency,omitnil" name:"Consistency"`
}

type InternalMedicineAbdomenSpleen struct {
	// 脾脏总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *KeyValueItem `json:"Src,omitnil" name:"Src"`

	// 脾脏大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *KeyValueItem `json:"Size,omitnil" name:"Size"`

	// 脾脏触诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Palpation *KeyValueItem `json:"Palpation,omitnil" name:"Palpation"`

	// 脾脏叩诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percussion *KeyValueItem `json:"Percussion,omitnil" name:"Percussion"`

	// 脾脏压痛
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tenderness *KeyValueItem `json:"Tenderness,omitnil" name:"Tenderness"`

	// 脾脏质地
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *KeyValueItem `json:"Consistency,omitnil" name:"Consistency"`
}

type InternalMedicineBaseItem struct {
	// 体检报告-内科-腹部
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abdomen *InternalMedicineAbdomen `json:"Abdomen,omitnil" name:"Abdomen"`

	// 体检报告-内科-心脏
	// 注意：此字段可能返回 null，表示取不到有效值。
	Heart *InternalMedicineHeart `json:"Heart,omitnil" name:"Heart"`

	// 体检报告-内科-血管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vessel *InternalMedicineVessel `json:"Vessel,omitnil" name:"Vessel"`

	// 体检报告-内科-呼吸系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	RespiratorySystem *InternalMedicineRespiratorySystem `json:"RespiratorySystem,omitnil" name:"RespiratorySystem"`

	// 体检报告-内科-内科其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others []*KeyValueItem `json:"Others,omitnil" name:"Others"`

	// 体检报告-内科-小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefSummary *InternalMedicineBriefSummary `json:"BriefSummary,omitnil" name:"BriefSummary"`
}

type InternalMedicineBriefSummary struct {
	// 内科小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type InternalMedicineHeart struct {
	// 心脏总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 心律
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeartRhythm *KeyValueItem `json:"HeartRhythm,omitnil" name:"HeartRhythm"`

	// 心率
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeartRate *ValueUnitItem `json:"HeartRate,omitnil" name:"HeartRate"`

	// 心脏听诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeartAuscultation *KeyValueItem `json:"HeartAuscultation,omitnil" name:"HeartAuscultation"`
}

type InternalMedicineRespiratorySystem struct {
	// 呼吸系统总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 胸廓
	// 注意：此字段可能返回 null，表示取不到有效值。
	Thoracic *KeyValueItem `json:"Thoracic,omitnil" name:"Thoracic"`

	// 痰量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sputum *KeyValueItem `json:"Sputum,omitnil" name:"Sputum"`

	// 肺部叩诊
	// 注意：此字段可能返回 null，表示取不到有效值。
	LungPercussion *KeyValueItem `json:"LungPercussion,omitnil" name:"LungPercussion"`

	// 肺部听诊其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	LungAuscultation []*KeyValueItem `json:"LungAuscultation,omitnil" name:"LungAuscultation"`
}

type InternalMedicineVessel struct {
	// 血管总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 血管杂音
	// 注意：此字段可能返回 null，表示取不到有效值。
	VascularMurmur *KeyValueItem `json:"VascularMurmur,omitnil" name:"VascularMurmur"`

	// 外周血管
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeripheralVessel *KeyValueItem `json:"PeripheralVessel,omitnil" name:"PeripheralVessel"`
}

type Invas struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitnil" name:"Positive"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`
}

type InvasiveV2 struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 阴性或阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitnil" name:"Positive"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type IssueInfo struct {
	// 编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertNumber *string `json:"CertNumber,omitnil" name:"CertNumber"`

	// 签发机构
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssuedAuthority *string `json:"IssuedAuthority,omitnil" name:"IssuedAuthority"`

	// 签发日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssuedDate *string `json:"IssuedDate,omitnil" name:"IssuedDate"`
}

type KeyValueItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 项目原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *PhysicalBaseItem `json:"Item,omitnil" name:"Item"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PhysicalBaseItem `json:"Result,omitnil" name:"Result"`
}

type LastMenstrualPeriodBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Lymph struct {
	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 转移数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferNum *int64 `json:"TransferNum,omitnil" name:"TransferNum"`
}

type LymphNode struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 转移数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferNum *int64 `json:"TransferNum,omitnil" name:"TransferNum"`

	// 淋巴结大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sizes []*int64 `json:"Sizes,omitnil" name:"Sizes"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type LymphTotal struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 转移数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferNum *string `json:"TransferNum,omitnil" name:"TransferNum"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *string `json:"Total,omitnil" name:"Total"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type MainDiseaseHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *bool `json:"State,omitnil" name:"State"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 否定列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Neglist *NeglistBlock `json:"Neglist,omitnil" name:"Neglist"`

	// 肯定列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Poslist *PoslistBlock `json:"Poslist,omitnil" name:"Poslist"`
}

type Maternity struct {
	// 描述部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *MaternityDesc `json:"Desc,omitnil" name:"Desc"`

	// 结论部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *MaternitySummary `json:"Summary,omitnil" name:"Summary"`

	// 报告原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrText *string `json:"OcrText,omitnil" name:"OcrText"`
}

type MaternityDesc struct {
	// 胎儿数据结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fetus []*Fetus `json:"Fetus,omitnil" name:"Fetus"`

	// 胎儿数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FetusNum *FieldInfo `json:"FetusNum,omitnil" name:"FetusNum"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type MaternitySummary struct {
	// 胎儿数据结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fetus []*Fetus `json:"Fetus,omitnil" name:"Fetus"`

	// 胎儿数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FetusNum *FieldInfo `json:"FetusNum,omitnil" name:"FetusNum"`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sym []*FieldInfo `json:"Sym,omitnil" name:"Sym"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type MedDoc struct {
	// 建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advice *Advice `json:"Advice,omitnil" name:"Advice"`

	// 诊断结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnosis []*DiagCertItem `json:"Diagnosis,omitnil" name:"Diagnosis"`

	// 疾病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseMedicalHistory *DiseaseMedicalHistory `json:"DiseaseMedicalHistory,omitnil" name:"DiseaseMedicalHistory"`

	// 个人史
	PersonalMedicalHistory *PersonalMedicalHistory `json:"PersonalMedicalHistory,omitnil" name:"PersonalMedicalHistory"`

	// 婚孕史
	ObstericalMedicalHistory *ObstericalMedicalHistory `json:"ObstericalMedicalHistory,omitnil" name:"ObstericalMedicalHistory"`

	// 家族史
	FamilyMedicalHistory *FamilyMedicalHistory `json:"FamilyMedicalHistory,omitnil" name:"FamilyMedicalHistory"`

	// 月经史
	MenstrualMedicalHistory *MenstrualMedicalHistory `json:"MenstrualMedicalHistory,omitnil" name:"MenstrualMedicalHistory"`

	// 诊疗记录
	TreatmentRecord *TreatmentRecord `json:"TreatmentRecord,omitnil" name:"TreatmentRecord"`
}

type MedicalRecordInfo struct {
	// 就诊日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnosisTime *string `json:"DiagnosisTime,omitnil" name:"DiagnosisTime"`

	// 就诊科室
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnosisDepartmentName *string `json:"DiagnosisDepartmentName,omitnil" name:"DiagnosisDepartmentName"`

	// 就诊医生
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnosisDoctorName *string `json:"DiagnosisDoctorName,omitnil" name:"DiagnosisDoctorName"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClinicalDiagnosis *string `json:"ClinicalDiagnosis,omitnil" name:"ClinicalDiagnosis"`

	// 主述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainNarration *string `json:"MainNarration,omitnil" name:"MainNarration"`

	// 体格检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhysicalExamination *string `json:"PhysicalExamination,omitnil" name:"PhysicalExamination"`

	// 检查结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectionFindings *string `json:"InspectionFindings,omitnil" name:"InspectionFindings"`

	// 治疗意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentOpinion *string `json:"TreatmentOpinion,omitnil" name:"TreatmentOpinion"`
}

type Medicine struct {
	// 药品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 商品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeName *string `json:"TradeName,omitnil" name:"TradeName"`

	// 厂商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Firm *string `json:"Firm,omitnil" name:"Firm"`

	// 医保类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil" name:"Category"`

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 最小包装数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinQuantity *string `json:"MinQuantity,omitnil" name:"MinQuantity"`

	// 最小制剂单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	DosageUnit *string `json:"DosageUnit,omitnil" name:"DosageUnit"`

	// 最小包装单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackingUnit *string `json:"PackingUnit,omitnil" name:"PackingUnit"`
}

type MenstrualFlowBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type MenstrualHistoryBlock struct {
	// 末次月经
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastMenstrualPeriod *LastMenstrualPeriodBlock `json:"LastMenstrualPeriod,omitnil" name:"LastMenstrualPeriod"`

	// 月经量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualFlow *MenstrualFlowBlock `json:"MenstrualFlow,omitnil" name:"MenstrualFlow"`

	// 初潮时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenarcheAge *LastMenstrualPeriodBlock `json:"MenarcheAge,omitnil" name:"MenarcheAge"`

	// 是否绝经
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstruationOrNot *MenstruationOrNotBlock `json:"MenstruationOrNot,omitnil" name:"MenstruationOrNot"`

	// 月经周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualCycles *LastMenstrualPeriodBlock `json:"MenstrualCycles,omitnil" name:"MenstrualCycles"`

	// 月经经期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualPeriod *MenstrualPeriodBlock `json:"MenstrualPeriod,omitnil" name:"MenstrualPeriod"`
}

type MenstrualHistoryDetailBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil" name:"State"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitnil" name:"TimeType"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type MenstrualMedicalHistory struct {
	// 末次月经时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastMenstrualPeriod *string `json:"LastMenstrualPeriod,omitnil" name:"LastMenstrualPeriod"`

	// 经量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualFlow *string `json:"MenstrualFlow,omitnil" name:"MenstrualFlow"`

	// 月经初潮年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenarcheAge *string `json:"MenarcheAge,omitnil" name:"MenarcheAge"`

	// 是否来月经
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstruationOrNot *string `json:"MenstruationOrNot,omitnil" name:"MenstruationOrNot"`

	// 月经周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualCycles *string `json:"MenstrualCycles,omitnil" name:"MenstrualCycles"`

	// 月经持续天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualPeriod *string `json:"MenstrualPeriod,omitnil" name:"MenstrualPeriod"`
}

type MenstrualPeriodBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type MenstruationOrNotBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitnil" name:"TimeType"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Molecular struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 基因名称标注化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分子病理详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *MolecularValue `json:"Value,omitnil" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type MolecularValue struct {
	// 外显子
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exon *string `json:"Exon,omitnil" name:"Exon"`

	// 点位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *string `json:"Position,omitnil" name:"Position"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 阳性或阴性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitnil" name:"Positive"`

	// 基因名称原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`
}

type Multiple struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type NeglistBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type NeonatalInfo struct {
	// 新生儿名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeonatalName *string `json:"NeonatalName,omitnil" name:"NeonatalName"`

	// 新生儿性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeonatalGender *string `json:"NeonatalGender,omitnil" name:"NeonatalGender"`

	// 出生身长
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthLength *string `json:"BirthLength,omitnil" name:"BirthLength"`

	// 出生体重
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthWeight *string `json:"BirthWeight,omitnil" name:"BirthWeight"`

	// 出生孕周
	// 注意：此字段可能返回 null，表示取不到有效值。
	GestationalAge *string `json:"GestationalAge,omitnil" name:"GestationalAge"`

	// 出生时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthTime *string `json:"BirthTime,omitnil" name:"BirthTime"`

	// 出生地点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *string `json:"BirthPlace,omitnil" name:"BirthPlace"`

	// 医疗机构
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalInstitutions *string `json:"MedicalInstitutions,omitnil" name:"MedicalInstitutions"`
}

type NormPart struct {
	// 部位值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitnil" name:"Part"`

	// 部位方向
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartDirection *string `json:"PartDirection,omitnil" name:"PartDirection"`

	// 组织值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tissue *string `json:"Tissue,omitnil" name:"Tissue"`

	// 组织方向
	// 注意：此字段可能返回 null，表示取不到有效值。
	TissueDirection *string `json:"TissueDirection,omitnil" name:"TissueDirection"`

	// 上级部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Upper *string `json:"Upper,omitnil" name:"Upper"`

	// 部位详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartDetail *PartDesc `json:"PartDetail,omitnil" name:"PartDetail"`
}

type NormSize struct {
	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number []*string `json:"Number,omitnil" name:"Number"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Impl *string `json:"Impl,omitnil" name:"Impl"`
}

type NumValue struct {
	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Num *string `json:"Num,omitnil" name:"Num"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`
}

type ObstericalMedicalHistory struct {
	// 婚史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarriageHistory *string `json:"MarriageHistory,omitnil" name:"MarriageHistory"`

	// 孕史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FertilityHistory *string `json:"FertilityHistory,omitnil" name:"FertilityHistory"`
}

type ObstetricalHistoryBlock struct {
	// 婚姻史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarriageHistory *MenstrualHistoryDetailBlock `json:"MarriageHistory,omitnil" name:"MarriageHistory"`

	// 婚育史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FertilityHistory *FertilityHistoryBlock `json:"FertilityHistory,omitnil" name:"FertilityHistory"`
}

type OphthalmologyBareEyeSight struct {
	// 左眼视力
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftEyeVisual *KeyValueItem `json:"LeftEyeVisual,omitnil" name:"LeftEyeVisual"`

	// 裸眼视力
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 右眼视力
	// 注意：此字段可能返回 null，表示取不到有效值。
	RightEyeVisual *KeyValueItem `json:"RightEyeVisual,omitnil" name:"RightEyeVisual"`
}

type OphthalmologyBaseItem struct {
	// 裸眼视力
	// 注意：此字段可能返回 null，表示取不到有效值。
	BareEyeSight *OphthalmologyBareEyeSight `json:"BareEyeSight,omitnil" name:"BareEyeSight"`

	// 矫正视力
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorrectedVisualAcuity *OphthalmologyCorrectedVisualAcuity `json:"CorrectedVisualAcuity,omitnil" name:"CorrectedVisualAcuity"`

	// 色觉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColourVision *OphthalmologyColourVision `json:"ColourVision,omitnil" name:"ColourVision"`

	// 眼底
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fundoscopy *OphthalmologyFundoscopy `json:"Fundoscopy,omitnil" name:"Fundoscopy"`

	// 眼科其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others []*KeyValueItem `json:"Others,omitnil" name:"Others"`

	// 眼科小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefSummary *OphthalmologyBriefSummary `json:"BriefSummary,omitnil" name:"BriefSummary"`
}

type OphthalmologyBriefSummary struct {
	// 眼科小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type OphthalmologyColourVision struct {
	// 色觉总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type OphthalmologyCorrectedVisualAcuity struct {
	// 左眼矫正视力
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftEyeVisual *KeyValueItem `json:"LeftEyeVisual,omitnil" name:"LeftEyeVisual"`

	// 矫正视力
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 右眼矫正视力
	// 注意：此字段可能返回 null，表示取不到有效值。
	RightEyeVisual *KeyValueItem `json:"RightEyeVisual,omitnil" name:"RightEyeVisual"`
}

type OphthalmologyFundoscopy struct {
	// 眼底检查总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type Organ struct {
	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size []*Size `json:"Size,omitnil" name:"Size"`

	// 包膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envelope *BlockInfo `json:"Envelope,omitnil" name:"Envelope"`

	// 边缘
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edge *BlockInfo `json:"Edge,omitnil" name:"Edge"`

	// 内部回声
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEcho *BlockInfo `json:"InnerEcho,omitnil" name:"InnerEcho"`

	// 腺体
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gland *BlockInfo `json:"Gland,omitnil" name:"Gland"`

	// 形状
	// 注意：此字段可能返回 null，表示取不到有效值。
	Shape *BlockInfo `json:"Shape,omitnil" name:"Shape"`

	// 厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Thickness *BlockInfo `json:"Thickness,omitnil" name:"Thickness"`

	// 形态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShapeAttr *BlockInfo `json:"ShapeAttr,omitnil" name:"ShapeAttr"`

	// 血液cdfi
	// 注意：此字段可能返回 null，表示取不到有效值。
	CDFI *BlockInfo `json:"CDFI,omitnil" name:"CDFI"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymDesc *BlockInfo `json:"SymDesc,omitnil" name:"SymDesc"`

	// 大小状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	SizeStatus *BlockInfo `json:"SizeStatus,omitnil" name:"SizeStatus"`

	// 轮廓
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outline *BlockInfo `json:"Outline,omitnil" name:"Outline"`

	// 结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Structure *BlockInfo `json:"Structure,omitnil" name:"Structure"`

	// 密度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Density *BlockInfo `json:"Density,omitnil" name:"Density"`

	// 血管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vas *BlockInfo `json:"Vas,omitnil" name:"Vas"`

	// 囊壁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cysticwall *BlockInfo `json:"Cysticwall,omitnil" name:"Cysticwall"`

	// 被膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capsule *BlockInfo `json:"Capsule,omitnil" name:"Capsule"`

	// 峡部厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IsthmusThicknese is deprecated.
	IsthmusThicknese *Size `json:"IsthmusThicknese,omitnil" name:"IsthmusThicknese"`

	// 内部回声分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoDistribution *BlockInfo `json:"InnerEchoDistribution,omitnil" name:"InnerEchoDistribution"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 透声度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transparent *BlockInfo `json:"Transparent,omitnil" name:"Transparent"`

	// MRI ADC
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriAdc *BlockInfo `json:"MriAdc,omitnil" name:"MriAdc"`

	// MRI DWI
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriDwi *BlockInfo `json:"MriDwi,omitnil" name:"MriDwi"`

	// MRI T1信号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriT1 *BlockInfo `json:"MriT1,omitnil" name:"MriT1"`

	// MRI T2信号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriT2 *BlockInfo `json:"MriT2,omitnil" name:"MriT2"`

	// CT HU值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtHu *BlockInfo `json:"CtHu,omitnil" name:"CtHu"`

	// SUmax值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suvmax *BlockInfo `json:"Suvmax,omitnil" name:"Suvmax"`

	// 代谢情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metabolism *BlockInfo `json:"Metabolism,omitnil" name:"Metabolism"`

	// 放射性摄取
	// 注意：此字段可能返回 null，表示取不到有效值。
	RadioactiveUptake *BlockInfo `json:"RadioactiveUptake,omitnil" name:"RadioactiveUptake"`

	// 淋巴结情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphEnlargement *BlockInfo `json:"LymphEnlargement,omitnil" name:"LymphEnlargement"`

	// 影像特征
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageFeature *BlockInfo `json:"ImageFeature,omitnil" name:"ImageFeature"`

	// 导管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duct *BlockInfo `json:"Duct,omitnil" name:"Duct"`

	// 趋势
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trend *BlockInfo `json:"Trend,omitnil" name:"Trend"`

	// 手术情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *BlockInfo `json:"Operation,omitnil" name:"Operation"`

	// 器官在报告图片中的坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`

	// 峡部厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsthmusThickness *Size `json:"IsthmusThickness,omitnil" name:"IsthmusThickness"`
}

type OtherInfo struct {
	// 麻醉方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Anesthesia *SurgeryAttr `json:"Anesthesia,omitnil" name:"Anesthesia"`

	// 术中出血
	// 注意：此字段可能返回 null，表示取不到有效值。
	BloodLoss *SurgeryAttr `json:"BloodLoss,omitnil" name:"BloodLoss"`

	// 输血
	// 注意：此字段可能返回 null，表示取不到有效值。
	BloodTransfusion *SurgeryAttr `json:"BloodTransfusion,omitnil" name:"BloodTransfusion"`

	// 手术用时
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *SurgeryAttr `json:"Duration,omitnil" name:"Duration"`

	// 手术开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *SurgeryAttr `json:"EndTime,omitnil" name:"EndTime"`

	// 手术结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *SurgeryAttr `json:"StartTime,omitnil" name:"StartTime"`
}

type OtolaryngologyBaseItem struct {
	// 耳朵
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ear *OtolaryngologyEar `json:"Ear,omitnil" name:"Ear"`

	// 鼻
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nose *OtolaryngologyNose `json:"Nose,omitnil" name:"Nose"`

	// 喉
	// 注意：此字段可能返回 null，表示取不到有效值。
	Larynx *OtolaryngologyLarynx `json:"Larynx,omitnil" name:"Larynx"`

	// 耳鼻喉其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others []*KeyValueItem `json:"Others,omitnil" name:"Others"`

	// 小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefSummary *OtolaryngologyBriefSummary `json:"BriefSummary,omitnil" name:"BriefSummary"`
}

type OtolaryngologyBriefSummary struct {
	// 耳鼻喉小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type OtolaryngologyEar struct {
	// 耳总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 听力
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hearing *HearingItem `json:"Hearing,omitnil" name:"Hearing"`
}

type OtolaryngologyLarynx struct {
	// 喉总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type OtolaryngologyNose struct {
	// 鼻总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type PTNM struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// pT
	// 注意：此字段可能返回 null，表示取不到有效值。
	PT *string `json:"PT,omitnil" name:"PT"`

	// pN
	// 注意：此字段可能返回 null，表示取不到有效值。
	PN *string `json:"PN,omitnil" name:"PN"`

	// pM
	// 注意：此字段可能返回 null，表示取不到有效值。
	PM *string `json:"PM,omitnil" name:"PM"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type PTNMBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// PTNM分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMM *string `json:"PTNMM,omitnil" name:"PTNMM"`

	// PTNM分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMN *string `json:"PTNMN,omitnil" name:"PTNMN"`

	// PTNM分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMT *string `json:"PTNMT,omitnil" name:"PTNMT"`
}

type ParagraphBlock struct {
	// 切口愈合情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealingText *string `json:"IncisionHealingText,omitnil" name:"IncisionHealingText"`

	// 辅助检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExaminationText *string `json:"AuxiliaryExaminationText,omitnil" name:"AuxiliaryExaminationText"`

	// 特殊检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialExamText *string `json:"SpecialExamText,omitnil" name:"SpecialExamText"`

	// 门诊诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutpatientDiagnosisText *string `json:"OutpatientDiagnosisText,omitnil" name:"OutpatientDiagnosisText"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionConditionText *string `json:"AdmissionConditionText,omitnil" name:"AdmissionConditionText"`

	// 诊疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcessText *string `json:"CheckAndTreatmentProcessText,omitnil" name:"CheckAndTreatmentProcessText"`

	// 体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSignsText *string `json:"SymptomsAndSignsText,omitnil" name:"SymptomsAndSignsText"`

	// 出院医嘱
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeInstructionsText *string `json:"DischargeInstructionsText,omitnil" name:"DischargeInstructionsText"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDiagnosisText *string `json:"AdmissionDiagnosisText,omitnil" name:"AdmissionDiagnosisText"`

	// 手术情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryConditionText *string `json:"SurgeryConditionText,omitnil" name:"SurgeryConditionText"`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosisText *string `json:"PathologicalDiagnosisText,omitnil" name:"PathologicalDiagnosisText"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeConditionText *string `json:"DischargeConditionText,omitnil" name:"DischargeConditionText"`

	// 检查记录
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecordText *string `json:"CheckRecordText,omitnil" name:"CheckRecordText"`

	// 主诉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaintText *string `json:"ChiefComplaintText,omitnil" name:"ChiefComplaintText"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosisText *string `json:"DischargeDiagnosisText,omitnil" name:"DischargeDiagnosisText"`

	// 既往史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDiseaseHistoryText *string `json:"MainDiseaseHistoryText,omitnil" name:"MainDiseaseHistoryText"`

	// 现病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseasePresentText *string `json:"DiseasePresentText,omitnil" name:"DiseasePresentText"`

	// 个人史
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonalHistoryText *string `json:"PersonalHistoryText,omitnil" name:"PersonalHistoryText"`

	// 月经史
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: MenstruallHistoryText is deprecated.
	MenstruallHistoryText *string `json:"MenstruallHistoryText,omitnil" name:"MenstruallHistoryText"`

	// 婚育史
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObstericalHistoryText *string `json:"ObstericalHistoryText,omitnil" name:"ObstericalHistoryText"`

	// 家族史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FamilyHistoryText *string `json:"FamilyHistoryText,omitnil" name:"FamilyHistoryText"`

	// 过敏史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllergyHistoryText *string `json:"AllergyHistoryText,omitnil" name:"AllergyHistoryText"`

	// 病史信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseHistoryText *string `json:"DiseaseHistoryText,omitnil" name:"DiseaseHistoryText"`

	// 其它诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherDiagnosisText *string `json:"OtherDiagnosisText,omitnil" name:"OtherDiagnosisText"`

	// 体格检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyExaminationText *string `json:"BodyExaminationText,omitnil" name:"BodyExaminationText"`

	// 专科检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialistExaminationText *string `json:"SpecialistExaminationText,omitnil" name:"SpecialistExaminationText"`

	// 治疗结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentResultText *string `json:"TreatmentResultText,omitnil" name:"TreatmentResultText"`

	// 月经史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualHistoryText *string `json:"MenstrualHistoryText,omitnil" name:"MenstrualHistoryText"`
}

type ParentInfo struct {
	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Age *string `json:"Age,omitnil" name:"Age"`

	// 证件号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCard *string `json:"IdCard,omitnil" name:"IdCard"`

	// 民族
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ethnicity *string `json:"Ethnicity,omitnil" name:"Ethnicity"`

	// 国籍
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nationality *string `json:"Nationality,omitnil" name:"Nationality"`

	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil" name:"Address"`
}

type Part struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormPart *NormPart `json:"NormPart,omitnil" name:"NormPart"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueBrief *string `json:"ValueBrief,omitnil" name:"ValueBrief"`
}

type PartDesc struct {
	// 主要部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainDir *string `json:"MainDir,omitnil" name:"MainDir"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitnil" name:"Part"`

	// 次要部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondaryDir *string `json:"SecondaryDir,omitnil" name:"SecondaryDir"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type PathologicalDiagnosisBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 病理详细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*PathologicalDiagnosisDetailBlock `json:"Detail,omitnil" name:"Detail"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type PathologicalDiagnosisDetailBlock struct {
	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitnil" name:"Part"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologicalType *string `json:"HistologicalType,omitnil" name:"HistologicalType"`

	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologicalGrade *string `json:"HistologicalGrade,omitnil" name:"HistologicalGrade"`
}

type PathologyReport struct {
	// 癌症部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancerPart *Part `json:"CancerPart,omitnil" name:"CancerPart"`

	// 癌症部位大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancerSize []*Size `json:"CancerSize,omitnil" name:"CancerSize"`

	// 描述文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescText *string `json:"DescText,omitnil" name:"DescText"`

	// 组织学等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyLevel *HistologyLevel `json:"HistologyLevel,omitnil" name:"HistologyLevel"`

	// 组织学类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistologyType *HistologyType `json:"HistologyType,omitnil" name:"HistologyType"`

	// IHC信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IHC []*IHCInfo `json:"IHC,omitnil" name:"IHC"`

	// 浸润深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfiltrationDepth *BlockInfo `json:"InfiltrationDepth,omitnil" name:"InfiltrationDepth"`

	// 肿瘤扩散
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invasive []*Invas `json:"Invasive,omitnil" name:"Invasive"`

	// 淋巴结
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphNodes []*Lymph `json:"LymphNodes,omitnil" name:"LymphNodes"`

	// PTNM信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *BlockInfo `json:"PTNM,omitnil" name:"PTNM"`

	// 病理报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalReportType *BlockInfo `json:"PathologicalReportType,omitnil" name:"PathologicalReportType"`

	// 报告原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportText *string `json:"ReportText,omitnil" name:"ReportText"`

	// 标本类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *BlockInfo `json:"SampleType,omitnil" name:"SampleType"`

	// 结论文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryText *string `json:"SummaryText,omitnil" name:"SummaryText"`
}

type PathologyV2 struct {
	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalReportType *Report `json:"PathologicalReportType,omitnil" name:"PathologicalReportType"`

	// 描述段落
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *DescInfo `json:"Desc,omitnil" name:"Desc"`

	// 诊断结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *SummaryInfo `json:"Summary,omitnil" name:"Summary"`

	// 报告全文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportText *string `json:"ReportText,omitnil" name:"ReportText"`

	// 淋巴结总计转移信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphTotal []*LymphTotal `json:"LymphTotal,omitnil" name:"LymphTotal"`

	// 单淋巴结转移信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphNodes []*LymphNode `json:"LymphNodes,omitnil" name:"LymphNodes"`

	// ihc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ihc []*IHCV2 `json:"Ihc,omitnil" name:"Ihc"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clinical *BaseInfo `json:"Clinical,omitnil" name:"Clinical"`

	// 是否癌前病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Precancer *HistologyClass `json:"Precancer,omitnil" name:"Precancer"`

	// 是否恶性肿瘤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Malignant *HistologyClass `json:"Malignant,omitnil" name:"Malignant"`

	// 是否良性肿瘤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Benigntumor *HistologyClass `json:"Benigntumor,omitnil" name:"Benigntumor"`

	// 送检材料
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *BaseInfo `json:"SampleType,omitnil" name:"SampleType"`

	// 淋巴结大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphSize []*Size `json:"LymphSize,omitnil" name:"LymphSize"`

	// 分子病理
	// 注意：此字段可能返回 null，表示取不到有效值。
	Molecular []*Molecular `json:"Molecular,omitnil" name:"Molecular"`
}

type PatientInfo struct {
	// 患者姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 患者性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sex *string `json:"Sex,omitnil" name:"Sex"`

	// 患者年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Age *string `json:"Age,omitnil" name:"Age"`

	// 患者手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 患者地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil" name:"Address"`

	// 患者身份证
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCard *string `json:"IdCard,omitnil" name:"IdCard"`

	// 健康卡号
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCardNo *string `json:"HealthCardNo,omitnil" name:"HealthCardNo"`

	// 社保卡号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SocialSecurityCardNo *string `json:"SocialSecurityCardNo,omitnil" name:"SocialSecurityCardNo"`

	// 出生日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Birthday *string `json:"Birthday,omitnil" name:"Birthday"`

	// 民族
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ethnicity *string `json:"Ethnicity,omitnil" name:"Ethnicity"`

	// 婚姻状况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Married *string `json:"Married,omitnil" name:"Married"`

	// 职业
	// 注意：此字段可能返回 null，表示取不到有效值。
	Profession *string `json:"Profession,omitnil" name:"Profession"`

	// 教育程度
	// 注意：此字段可能返回 null，表示取不到有效值。
	EducationBackground *string `json:"EducationBackground,omitnil" name:"EducationBackground"`

	// 国籍
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nationality *string `json:"Nationality,omitnil" name:"Nationality"`

	// 籍贯
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *string `json:"BirthPlace,omitnil" name:"BirthPlace"`

	// 医保类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalInsuranceType *string `json:"MedicalInsuranceType,omitnil" name:"MedicalInsuranceType"`

	// 标准化年龄
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgeNorm *string `json:"AgeNorm,omitnil" name:"AgeNorm"`

	// 民族。该字段已不再使用，请从Ethnicity取值
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Nation is deprecated.
	Nation *string `json:"Nation,omitnil" name:"Nation"`

	// 婚姻代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarriedCode *string `json:"MarriedCode,omitnil" name:"MarriedCode"`

	// 职业代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProfessionCode *string `json:"ProfessionCode,omitnil" name:"ProfessionCode"`

	// 居民医保代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalInsuranceTypeCode *string `json:"MedicalInsuranceTypeCode,omitnil" name:"MedicalInsuranceTypeCode"`

	// 床号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BedNo *string `json:"BedNo,omitnil" name:"BedNo"`
}

type PdfInfo struct {
	// pdf文件url链接(暂不支持)
	Url *string `json:"Url,omitnil" name:"Url"`

	// pdf文件base64编码字符串
	Base64 *string `json:"Base64,omitnil" name:"Base64"`
}

type PersonalHistoryBlock struct {
	// 出生地
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *BirthPlaceBlock `json:"BirthPlace,omitnil" name:"BirthPlace"`

	// 居住地
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivePlace *BirthPlaceBlock `json:"LivePlace,omitnil" name:"LivePlace"`

	// 职业
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *BirthPlaceBlock `json:"Job,omitnil" name:"Job"`

	// 吸烟
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmokeHistory *SmokeHistoryBlock `json:"SmokeHistory,omitnil" name:"SmokeHistory"`

	// 喝酒
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlcoholicHistory *SmokeHistoryBlock `json:"AlcoholicHistory,omitnil" name:"AlcoholicHistory"`

	// 月经史
	// 注意：此字段可能返回 null，表示取不到有效值。
	MenstrualHistory *MenstrualHistoryBlock `json:"MenstrualHistory,omitnil" name:"MenstrualHistory"`

	// 婚姻-生育史
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObstericalHistory *ObstetricalHistoryBlock `json:"ObstericalHistory,omitnil" name:"ObstericalHistory"`

	// 家族史
	// 注意：此字段可能返回 null，表示取不到有效值。
	FamilyHistory *FamilyHistoryBlock `json:"FamilyHistory,omitnil" name:"FamilyHistory"`
}

type PersonalMedicalHistory struct {
	// 出生史
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthPlace *string `json:"BirthPlace,omitnil" name:"BirthPlace"`

	// 居住史
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivePlace *string `json:"LivePlace,omitnil" name:"LivePlace"`

	// 工作史
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *string `json:"Job,omitnil" name:"Job"`

	// 吸烟史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmokeHistory *string `json:"SmokeHistory,omitnil" name:"SmokeHistory"`

	// 饮酒史
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlcoholicHistory *string `json:"AlcoholicHistory,omitnil" name:"AlcoholicHistory"`
}

type PhysicalBaseItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原始文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 归一化后值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 四点坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type PhysicalExamination struct {
	// 一般检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	GeneralExamination *GeneralExaminationBaseItem `json:"GeneralExamination,omitnil" name:"GeneralExamination"`

	// 内科
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternalMedicine *InternalMedicineBaseItem `json:"InternalMedicine,omitnil" name:"InternalMedicine"`

	// 外科
	// 注意：此字段可能返回 null，表示取不到有效值。
	Surgery *SurgeryBaseItem `json:"Surgery,omitnil" name:"Surgery"`

	// 口腔科
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stomatology *StomatologyBaseItem `json:"Stomatology,omitnil" name:"Stomatology"`

	// 眼科
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ophthalmology *OphthalmologyBaseItem `json:"Ophthalmology,omitnil" name:"Ophthalmology"`

	// 耳鼻喉科
	// 注意：此字段可能返回 null，表示取不到有效值。
	Otolaryngology *OtolaryngologyBaseItem `json:"Otolaryngology,omitnil" name:"Otolaryngology"`

	// 妇科
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gynaecology *GynaecologyBaseItem `json:"Gynaecology,omitnil" name:"Gynaecology"`

	// 未标准化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unclassified []*KeyValueItem `json:"Unclassified,omitnil" name:"Unclassified"`
}

type PhysicalExaminationV1 struct {
	// 体检报告信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhysicalExaminationMulti *PhysicalExamination `json:"PhysicalExaminationMulti,omitnil" name:"PhysicalExaminationMulti"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`
}

type Point struct {
	// x坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *int64 `json:"X,omitnil" name:"X"`

	// y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *int64 `json:"Y,omitnil" name:"Y"`
}

type PoslistBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Prescription struct {
	// 药品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicineList []*Medicine `json:"MedicineList,omitnil" name:"MedicineList"`
}

type Rectangle struct {
	// x坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *int64 `json:"X,omitnil" name:"X"`

	// y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	W *int64 `json:"W,omitnil" name:"W"`

	// 高
	// 注意：此字段可能返回 null，表示取不到有效值。
	H *int64 `json:"H,omitnil" name:"H"`
}

type RelapseDateBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 疾病名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseaseName *string `json:"DiseaseName,omitnil" name:"DiseaseName"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 归一化值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Norm *string `json:"Norm,omitnil" name:"Norm"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type RelativeCancerHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 肿瘤史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeCancerList *string `json:"RelativeCancerList,omitnil" name:"RelativeCancerList"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type RelativeHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 成员列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*RelativeHistoryDetailBlock `json:"Detail,omitnil" name:"Detail"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type RelativeHistoryDetailBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *string `json:"Relation,omitnil" name:"Relation"`

	// 死亡时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeOfDeath *string `json:"TimeOfDeath,omitnil" name:"TimeOfDeath"`

	// 时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitnil" name:"TimeType"`
}

type Report struct {
	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 原文对应坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type ReportInfo struct {
	// 医院名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hospital *string `json:"Hospital,omitnil" name:"Hospital"`

	// 科室名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentName *string `json:"DepartmentName,omitnil" name:"DepartmentName"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingTime *string `json:"BillingTime,omitnil" name:"BillingTime"`

	// 报告时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTime *string `json:"ReportTime,omitnil" name:"ReportTime"`

	// 检查时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectTime *string `json:"InspectTime,omitnil" name:"InspectTime"`

	// 检查号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckNum *string `json:"CheckNum,omitnil" name:"CheckNum"`

	// 影像号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageNum *string `json:"ImageNum,omitnil" name:"ImageNum"`

	// 放射号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RadiationNum *string `json:"RadiationNum,omitnil" name:"RadiationNum"`

	// 检验号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TestNum *string `json:"TestNum,omitnil" name:"TestNum"`

	// 门诊号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutpatientNum *string `json:"OutpatientNum,omitnil" name:"OutpatientNum"`

	// 病理号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologyNum *string `json:"PathologyNum,omitnil" name:"PathologyNum"`

	// 住院号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InHospitalNum *string `json:"InHospitalNum,omitnil" name:"InHospitalNum"`

	// 样本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleNum *string `json:"SampleNum,omitnil" name:"SampleNum"`

	// 标本种类
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *string `json:"SampleType,omitnil" name:"SampleType"`

	// 病历号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalRecordNum *string `json:"MedicalRecordNum,omitnil" name:"MedicalRecordNum"`

	// 报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportName *string `json:"ReportName,omitnil" name:"ReportName"`

	// 超声号
	// 注意：此字段可能返回 null，表示取不到有效值。
	UltraNum *string `json:"UltraNum,omitnil" name:"UltraNum"`

	// 临床诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnose *string `json:"Diagnose,omitnil" name:"Diagnose"`

	// 检查项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckItem *string `json:"CheckItem,omitnil" name:"CheckItem"`

	// 检查方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckMethod *string `json:"CheckMethod,omitnil" name:"CheckMethod"`

	// 诊断时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnoseTime *string `json:"DiagnoseTime,omitnil" name:"DiagnoseTime"`

	// 体检号
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckupNum *string `json:"HealthCheckupNum,omitnil" name:"HealthCheckupNum"`

	// 其它时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherTime *string `json:"OtherTime,omitnil" name:"OtherTime"`

	// 打印时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrintTime *string `json:"PrintTime,omitnil" name:"PrintTime"`

	// 未归类时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Times []*Time `json:"Times,omitnil" name:"Times"`

	// 床号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BedNo *string `json:"BedNo,omitnil" name:"BedNo"`
}

type ReportTypeVersion struct {
	// 检验报告
	ReportType *int64 `json:"ReportType,omitnil" name:"ReportType"`

	// 版本2
	Version *int64 `json:"Version,omitnil" name:"Version"`
}

type ResultInfo struct {
	// 段落文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *BaseInfo `json:"Text,omitnil" name:"Text"`

	// 结论详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*BaseInfo `json:"Items,omitnil" name:"Items"`
}

type Size struct {
	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 标准大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormSize *NormSize `json:"NormSize,omitnil" name:"NormSize"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type SmokeHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 时间单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 时间归一化
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeNorm *string `json:"TimeNorm,omitnil" name:"TimeNorm"`

	// 吸烟量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount *string `json:"Amount,omitnil" name:"Amount"`

	// 戒烟状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuitState *bool `json:"QuitState,omitnil" name:"QuitState"`

	// 是否吸烟
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *bool `json:"State,omitnil" name:"State"`

	// 对外输出值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type StomatologyBaseItem struct {
	// 龋齿
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToothDecay *StomatologyToothDecay `json:"ToothDecay,omitnil" name:"ToothDecay"`

	// 牙龈
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gingiva *StomatologyGingiva `json:"Gingiva,omitnil" name:"Gingiva"`

	// 牙周
	// 注意：此字段可能返回 null，表示取不到有效值。
	Periodontics *StomatologyPeriodontics `json:"Periodontics,omitnil" name:"Periodontics"`

	// 口腔其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others []*KeyValueItem `json:"Others,omitnil" name:"Others"`

	// 小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefSummary *StomatologyBriefSummary `json:"BriefSummary,omitnil" name:"BriefSummary"`
}

type StomatologyBriefSummary struct {
	// 口腔小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type StomatologyGingiva struct {
	// 牙龈总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type StomatologyPeriodontics struct {
	// 牙周总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type StomatologyToothDecay struct {
	// 龋齿总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type Summary struct {
	// 症状
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symptom []*SymptomInfo `json:"Symptom,omitnil" name:"Symptom"`

	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type SummaryInfo struct {
	// 诊断结论文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *BaseInfo `json:"Text,omitnil" name:"Text"`

	// 诊断结论详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infos []*DetailInformation `json:"Infos,omitnil" name:"Infos"`
}

type Surgery struct {
	// 手术史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryHistory *SurgeryHistory `json:"SurgeryHistory,omitnil" name:"SurgeryHistory"`

	// 其他信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherInfo *OtherInfo `json:"OtherInfo,omitnil" name:"OtherInfo"`
}

type SurgeryAnorectal struct {
	// 肛门直肠总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 直肠指检
	// 注意：此字段可能返回 null，表示取不到有效值。
	DigitalRectalExamination *KeyValueItem `json:"DigitalRectalExamination,omitnil" name:"DigitalRectalExamination"`

	// 痔疮
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hemorrhoid *KeyValueItem `json:"Hemorrhoid,omitnil" name:"Hemorrhoid"`
}

type SurgeryAttr struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type SurgeryBaseItem struct {
	// 体检报告-外科-头颈部
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadNeck *SurgeryHeadNeck `json:"HeadNeck,omitnil" name:"HeadNeck"`

	// 体检报告-外科-甲状腺
	// 注意：此字段可能返回 null，表示取不到有效值。
	Thyroid *SurgeryThyroid `json:"Thyroid,omitnil" name:"Thyroid"`

	// 体检报告-外科-乳房
	// 注意：此字段可能返回 null，表示取不到有效值。
	Breast *SurgeryBreast `json:"Breast,omitnil" name:"Breast"`

	// 体检报告-外科-浅表淋巴结
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphNode *SurgeryLymphNode `json:"LymphNode,omitnil" name:"LymphNode"`

	// 体检报告-外科-脊柱
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpinalExtremities *SurgerySpinalExtremities `json:"SpinalExtremities,omitnil" name:"SpinalExtremities"`

	// 体检报告-外科-皮肤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Skin *SurgerySkin `json:"Skin,omitnil" name:"Skin"`

	// 体检报告-外科-肛门直肠
	// 注意：此字段可能返回 null，表示取不到有效值。
	Anorectal *SurgeryAnorectal `json:"Anorectal,omitnil" name:"Anorectal"`

	// 体检报告-外科-泌尿生殖系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrogenitalSystem *SurgeryUrogenitalSystem `json:"UrogenitalSystem,omitnil" name:"UrogenitalSystem"`

	// 体检报告-外科-外科其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	Others []*KeyValueItem `json:"Others,omitnil" name:"Others"`

	// 体检报告-外科-小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefSummary *SurgeryBriefSummary `json:"BriefSummary,omitnil" name:"BriefSummary"`
}

type SurgeryBreast struct {
	// 乳房总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type SurgeryBriefSummary struct {
	// 外科小结
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type SurgeryConditionBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 手术历史
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryList []*SurgeryListBlock `json:"SurgeryList,omitnil" name:"SurgeryList"`

	// 对外输出值
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type SurgeryHeadNeck struct {
	// 头颈部总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type SurgeryHistory struct {
	// 手术名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryName *SurgeryAttr `json:"SurgeryName,omitnil" name:"SurgeryName"`

	// 手术日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryDate *SurgeryAttr `json:"SurgeryDate,omitnil" name:"SurgeryDate"`

	// 术前诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreoperativePathology *SurgeryAttr `json:"PreoperativePathology,omitnil" name:"PreoperativePathology"`

	// 术中诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntraoperativePathology *SurgeryAttr `json:"IntraoperativePathology,omitnil" name:"IntraoperativePathology"`

	// 术后诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostoperativePathology *SurgeryAttr `json:"PostoperativePathology,omitnil" name:"PostoperativePathology"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis *SurgeryAttr `json:"DischargeDiagnosis,omitnil" name:"DischargeDiagnosis"`
}

type SurgeryHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 手术列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Surgerylist []*SurgeryListBlock `json:"Surgerylist,omitnil" name:"Surgerylist"`
}

type SurgeryListBlock struct {
	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil" name:"Time"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeType *string `json:"TimeType,omitnil" name:"TimeType"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name []*string `json:"Name,omitnil" name:"Name"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *string `json:"Part,omitnil" name:"Part"`
}

type SurgeryLymphNode struct {
	// 浅表淋巴结总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type SurgerySkin struct {
	// 皮肤总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type SurgerySpinalExtremities struct {
	// 脊柱四肢总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 脊柱
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpinalColumn *KeyValueItem `json:"SpinalColumn,omitnil" name:"SpinalColumn"`

	// 四肢和关节
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimbJoint *KeyValueItem `json:"LimbJoint,omitnil" name:"LimbJoint"`

	// 平跛足
	// 注意：此字段可能返回 null，表示取不到有效值。
	Foot *KeyValueItem `json:"Foot,omitnil" name:"Foot"`

	// 骨骼
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bone *KeyValueItem `json:"Bone,omitnil" name:"Bone"`

	// 步态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gait *KeyValueItem `json:"Gait,omitnil" name:"Gait"`

	// 残疾或畸形
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deformity *KeyValueItem `json:"Deformity,omitnil" name:"Deformity"`
}

type SurgeryThyroid struct {
	// 甲状腺总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`
}

type SurgeryUrogenitalSystem struct {
	// 泌尿生殖系统总体描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *KeyValueItem `json:"Text,omitnil" name:"Text"`

	// 前列腺
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prostate *KeyValueItem `json:"Prostate,omitnil" name:"Prostate"`

	// 外生殖器（男性）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReproductiveOrgans *KeyValueItem `json:"ExternalReproductiveOrgans,omitnil" name:"ExternalReproductiveOrgans"`
}

type SymptomInfo struct {
	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *BlockInfo `json:"Grade,omitnil" name:"Grade"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symptom *BlockInfo `json:"Symptom,omitnil" name:"Symptom"`

	// 属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attrs []*BlockInfo `json:"Attrs,omitnil" name:"Attrs"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`
}

type TableIndicators struct {
	// 项目列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicators []*IndicatorItemV2 `json:"Indicators,omitnil" name:"Indicators"`

	// 采样标本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sample *BaseItem `json:"Sample,omitnil" name:"Sample"`
}

type Template struct {
	// 患者信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PatientInfo *PatientInfo `json:"PatientInfo,omitnil" name:"PatientInfo"`

	// 报告信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportInfo *ReportInfo `json:"ReportInfo,omitnil" name:"ReportInfo"`

	// 检查报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Check *Check `json:"Check,omitnil" name:"Check"`

	// 病理报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pathology *PathologyReport `json:"Pathology,omitnil" name:"Pathology"`

	// 出院报告，入院报告，门诊病历
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedDoc *MedDoc `json:"MedDoc,omitnil" name:"MedDoc"`

	// 诊断证明
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagCert *DiagCert `json:"DiagCert,omitnil" name:"DiagCert"`

	// 病案首页
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstPage *FirstPage `json:"FirstPage,omitnil" name:"FirstPage"`

	// 检验报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indicator *Indicator `json:"Indicator,omitnil" name:"Indicator"`

	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportType *string `json:"ReportType,omitnil" name:"ReportType"`

	// 门诊病历信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedicalRecordInfo *MedicalRecordInfo `json:"MedicalRecordInfo,omitnil" name:"MedicalRecordInfo"`

	// 出入院信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hospitalization *Hospitalization `json:"Hospitalization,omitnil" name:"Hospitalization"`

	// 手术记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Surgery *Surgery `json:"Surgery,omitnil" name:"Surgery"`

	// 心电图报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Electrocardiogram *Electrocardiogram `json:"Electrocardiogram,omitnil" name:"Electrocardiogram"`

	// 内窥镜报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endoscopy *Endoscopy `json:"Endoscopy,omitnil" name:"Endoscopy"`

	// 处方单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prescription *Prescription `json:"Prescription,omitnil" name:"Prescription"`

	// 疫苗接种凭证
	// 注意：此字段可能返回 null，表示取不到有效值。
	VaccineCertificate *VaccineCertificate `json:"VaccineCertificate,omitnil" name:"VaccineCertificate"`

	// OCR文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrText *string `json:"OcrText,omitnil" name:"OcrText"`

	// OCR拼接后文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrResult *string `json:"OcrResult,omitnil" name:"OcrResult"`

	// 报告类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTypeDesc *string `json:"ReportTypeDesc,omitnil" name:"ReportTypeDesc"`

	// 病理报告v2
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologyV2 *PathologyV2 `json:"PathologyV2,omitnil" name:"PathologyV2"`

	// 碳14尿素呼气试验
	// 注意：此字段可能返回 null，表示取不到有效值。
	C14 *Indicator `json:"C14,omitnil" name:"C14"`

	// 体检结论
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exame *Exame `json:"Exame,omitnil" name:"Exame"`

	// 出院报告v2，入院报告v2，门诊病历v2
	// 注意：此字段可能返回 null，表示取不到有效值。
	MedDocV2 *DischargeInfoBlock `json:"MedDocV2,omitnil" name:"MedDocV2"`

	// 检验报告v3
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndicatorV3 *IndicatorV3 `json:"IndicatorV3,omitnil" name:"IndicatorV3"`

	// 核酸报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Covid *CovidItemsInfo `json:"Covid,omitnil" name:"Covid"`

	// 孕产报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Maternity *Maternity `json:"Maternity,omitnil" name:"Maternity"`

	// 眼科报告
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eye *EyeItemsInfo `json:"Eye,omitnil" name:"Eye"`

	// 出生证明
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthCert *BirthCert `json:"BirthCert,omitnil" name:"BirthCert"`

	// 时间轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeline *TimelineInformation `json:"Timeline,omitnil" name:"Timeline"`
}

// Predefined struct for user
type TextToClassRequestParams struct {
	// 报告文本
	Text *string `json:"Text,omitnil" name:"Text"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil" name:"UserType"`
}

type TextToClassRequest struct {
	*tchttp.BaseRequest
	
	// 报告文本
	Text *string `json:"Text,omitnil" name:"Text"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil" name:"UserType"`
}

func (r *TextToClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToClassRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToClassRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToClassResponseParams struct {
	// 分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitnil" name:"TextTypeList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TextToClassResponse struct {
	*tchttp.BaseResponse
	Response *TextToClassResponseParams `json:"Response"`
}

func (r *TextToClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToClassResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToObjectRequestParams struct {
	// 报告文本
	Text *string `json:"Text,omitnil" name:"Text"`

	// 报告类型，目前支持12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为False，则Type字段不能为0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitnil" name:"IsUsedClassify"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病历记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitnil" name:"ReportTypeVersion"`
}

type TextToObjectRequest struct {
	*tchttp.BaseRequest
	
	// 报告文本
	Text *string `json:"Text,omitnil" name:"Text"`

	// 报告类型，目前支持12（检查报告），15（病理报告），28（出院报告），29（入院报告），210（门诊病历），212（手术记录），218（诊断证明），363（心电图），27（内窥镜检查），215（处方单），219（免疫接种证明），301（C14呼气试验）。如果不清楚报告类型，可以使用分类引擎，该字段传0（同时IsUsedClassify字段必须为True，否则无法输出结果）
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 是否使用分类引擎，当不确定报告类型时，可以使用收费的报告分类引擎服务。若该字段为False，则Type字段不能为0，否则无法输出结果。
	// 注意：当 IsUsedClassify 为True 时，表示使用收费的报告分类服务，将会产生额外的费用，具体收费标准参见 [购买指南的产品价格](https://cloud.tencent.com/document/product/1314/54264)。
	IsUsedClassify *bool `json:"IsUsedClassify,omitnil" name:"IsUsedClassify"`

	// 后付费的用户类型，新客户传1，老客户可不传或传 0。2022 年 12 月 15 新增了计费项，在此时间之前已经通过商务指定优惠价格的大客户，请不传这个字段或传 0，如果传 1 会导致以前获得的折扣价格失效。在 2022 年 12 月 15 日之后，通过商务指定优惠价格的大客户请传 1。
	UserType *uint64 `json:"UserType,omitnil" name:"UserType"`

	// 可选。用于指定不同报告使用的结构化引擎版本，不同版本返回的JSON 数据结果不兼容。若不指定版本号，就默认用旧的版本号。
	// （1）检验报告 11，默认使用 V2，最高支持 V3。
	// （2）病理报告 15，默认使用 V1，最高支持 V2。
	// （3）入院记录29、出院记录 28、病历记录 216、病程记录 217、门诊记录 210，默认使用 V1，最高支持 V2。
	ReportTypeVersion []*ReportTypeVersion `json:"ReportTypeVersion,omitnil" name:"ReportTypeVersion"`
}

func (r *TextToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Type")
	delete(f, "IsUsedClassify")
	delete(f, "UserType")
	delete(f, "ReportTypeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToObjectResponseParams struct {
	// 报告结构化结果
	Template *Template `json:"Template,omitnil" name:"Template"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TextToObjectResponse struct {
	*tchttp.BaseResponse
	Response *TextToObjectResponseParams `json:"Response"`
}

func (r *TextToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextType struct {
	// 类别Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 类别层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 类别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type TextTypeListBlock struct {
	// 文本类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTypeList []*TextType `json:"TextTypeList,omitnil" name:"TextTypeList"`
}

type Time struct {
	// 具体时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 时间值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type TimelineEvent struct {
	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 原文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 事件子类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// 事件发生时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil" name:"Time"`

	// 事件值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 位置坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rectangle *Rectangle `json:"Rectangle,omitnil" name:"Rectangle"`

	// 事件发生地点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Place *string `json:"Place,omitnil" name:"Place"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type TimelineInformation struct {
	// 时间轴
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeline []*TimelineEvent `json:"Timeline,omitnil" name:"Timeline"`
}

type TransfusionHistoryBlock struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *bool `json:"State,omitnil" name:"State"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type TreatmentRecord struct {
	// 入院
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DmissionCondition is deprecated.
	DmissionCondition *string `json:"DmissionCondition,omitnil" name:"DmissionCondition"`

	// 主诉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaint *string `json:"ChiefComplaint,omitnil" name:"ChiefComplaint"`

	// 现病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseasePresent *string `json:"DiseasePresent,omitnil" name:"DiseasePresent"`

	// 主要症状体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSigns *string `json:"SymptomsAndSigns,omitnil" name:"SymptomsAndSigns"`

	// 辅助检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExamination *string `json:"AuxiliaryExamination,omitnil" name:"AuxiliaryExamination"`

	// 体格检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyExamination *string `json:"BodyExamination,omitnil" name:"BodyExamination"`

	// 专科检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialistExamination *string `json:"SpecialistExamination,omitnil" name:"SpecialistExamination"`

	// 精神检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	MentalExamination *string `json:"MentalExamination,omitnil" name:"MentalExamination"`

	// 检查记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecord *string `json:"CheckRecord,omitnil" name:"CheckRecord"`

	// 化验结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectResult *string `json:"InspectResult,omitnil" name:"InspectResult"`

	// 切口愈合情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealing *string `json:"IncisionHealing,omitnil" name:"IncisionHealing"`

	// 处理意见
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentSuggestion *string `json:"TreatmentSuggestion,omitnil" name:"TreatmentSuggestion"`

	// 门诊随访要求
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowUpRequirements *string `json:"FollowUpRequirements,omitnil" name:"FollowUpRequirements"`

	// 诊疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcess *string `json:"CheckAndTreatmentProcess,omitnil" name:"CheckAndTreatmentProcess"`

	// 手术经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryCondition *string `json:"SurgeryCondition,omitnil" name:"SurgeryCondition"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionChanges *string `json:"ConditionChanges,omitnil" name:"ConditionChanges"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeCondition *string `json:"DischargeCondition,omitnil" name:"DischargeCondition"`

	// pTNM信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *string `json:"PTNM,omitnil" name:"PTNM"`

	// pTNMM信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMM *string `json:"PTNMM,omitnil" name:"PTNMM"`

	// pTNMN信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMN *string `json:"PTNMN,omitnil" name:"PTNMN"`

	// pTNMT信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNMT *string `json:"PTNMT,omitnil" name:"PTNMT"`

	// ECOG信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ECOG *string `json:"ECOG,omitnil" name:"ECOG"`

	// NRS信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NRS *string `json:"NRS,omitnil" name:"NRS"`

	// KPS信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KPS *string `json:"KPS,omitnil" name:"KPS"`

	// 死亡日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeathDate *string `json:"DeathDate,omitnil" name:"DeathDate"`

	// 复发日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelapseDate *string `json:"RelapseDate,omitnil" name:"RelapseDate"`

	// 观测天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObservationDays *string `json:"ObservationDays,omitnil" name:"ObservationDays"`

	// 入院
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionCondition *string `json:"AdmissionCondition,omitnil" name:"AdmissionCondition"`
}

type TreatmentRecordBlock struct {
	// 免疫组化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Immunohistochemistry *ImmunohistochemistryBlock `json:"Immunohistochemistry,omitnil" name:"Immunohistochemistry"`

	// 主诉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaint *ChiefComplaintBlock `json:"ChiefComplaint,omitnil" name:"ChiefComplaint"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionCondition *AdmissionConditionBlock `json:"AdmissionCondition,omitnil" name:"AdmissionCondition"`

	// 查体
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyExamination *BodyExaminationBlock `json:"BodyExamination,omitnil" name:"BodyExamination"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionDiagnosis,omitnil" name:"AdmissionDiagnosis"`

	// 入院中医诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionTraditionalDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionTraditionalDiagnosis,omitnil" name:"AdmissionTraditionalDiagnosis"`

	// 入院西医诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionModernDiagnosis *AdmissionDiagnosisBlock `json:"AdmissionModernDiagnosis,omitnil" name:"AdmissionModernDiagnosis"`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosis *PathologicalDiagnosisBlock `json:"PathologicalDiagnosis,omitnil" name:"PathologicalDiagnosis"`

	// 现病史
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiseasePresent *DiseasePresentBlock `json:"DiseasePresent,omitnil" name:"DiseasePresent"`

	// 体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSigns *DiseasePresentBlock `json:"SymptomsAndSigns,omitnil" name:"SymptomsAndSigns"`

	// 辅助检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExamination *DiseasePresentBlock `json:"AuxiliaryExamination,omitnil" name:"AuxiliaryExamination"`

	// 特殊检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialistExamination *DiseasePresentBlock `json:"SpecialistExamination,omitnil" name:"SpecialistExamination"`

	// 精神检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	MentalExamination *DiseasePresentBlock `json:"MentalExamination,omitnil" name:"MentalExamination"`

	// 检查记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecord *DiseasePresentBlock `json:"CheckRecord,omitnil" name:"CheckRecord"`

	// 检查结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectResult *DiseasePresentBlock `json:"InspectResult,omitnil" name:"InspectResult"`

	// 治疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcess *DiseasePresentBlock `json:"CheckAndTreatmentProcess,omitnil" name:"CheckAndTreatmentProcess"`

	// 手术经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryCondition *SurgeryConditionBlock `json:"SurgeryCondition,omitnil" name:"SurgeryCondition"`

	// 切口愈合
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealing *DiseasePresentBlock `json:"IncisionHealing,omitnil" name:"IncisionHealing"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosis *DischargeDiagnosisBlock `json:"DischargeDiagnosis,omitnil" name:"DischargeDiagnosis"`

	// 出院中医诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeTraditionalDiagnosis *DiseasePresentBlock `json:"DischargeTraditionalDiagnosis,omitnil" name:"DischargeTraditionalDiagnosis"`

	// 出院西医诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeModernDiagnosis *DischargeDiagnosisBlock `json:"DischargeModernDiagnosis,omitnil" name:"DischargeModernDiagnosis"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeCondition *DischargeConditionBlock `json:"DischargeCondition,omitnil" name:"DischargeCondition"`

	// 出院医嘱
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeInstructions *DiseasePresentBlock `json:"DischargeInstructions,omitnil" name:"DischargeInstructions"`

	// 治疗建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatmentSuggestion *DiseasePresentBlock `json:"TreatmentSuggestion,omitnil" name:"TreatmentSuggestion"`

	// 随访
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowUpRequirements *DiseasePresentBlock `json:"FollowUpRequirements,omitnil" name:"FollowUpRequirements"`

	// 治疗情况变化
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionChanges *DiseasePresentBlock `json:"ConditionChanges,omitnil" name:"ConditionChanges"`

	// 收缩压
	// 注意：此字段可能返回 null，表示取不到有效值。
	PulmonaryArterySystolicPressure *DiseasePresentBlock `json:"PulmonaryArterySystolicPressure,omitnil" name:"PulmonaryArterySystolicPressure"`

	// bclc分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BCLC *DiseasePresentBlock `json:"BCLC,omitnil" name:"BCLC"`

	// PTNM分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PTNM *PTNMBlock `json:"PTNM,omitnil" name:"PTNM"`

	// ECOG评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	ECOG *DiseasePresentBlock `json:"ECOG,omitnil" name:"ECOG"`

	// NRS评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	NRS *DiseasePresentBlock `json:"NRS,omitnil" name:"NRS"`

	// kps评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	KPS *DiseasePresentBlock `json:"KPS,omitnil" name:"KPS"`

	// 癌症分期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cancerstaging *ClinicalStaging `json:"Cancerstaging,omitnil" name:"Cancerstaging"`

	// 死亡时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeathDate *DeathDateBlock `json:"DeathDate,omitnil" name:"DeathDate"`

	// 复发日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelapseDate *RelapseDateBlock `json:"RelapseDate,omitnil" name:"RelapseDate"`

	// 观察日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObservationDays *DeathDateBlock `json:"ObservationDays,omitnil" name:"ObservationDays"`

	// 切口愈合情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncisionHealingText *string `json:"IncisionHealingText,omitnil" name:"IncisionHealingText"`

	// 辅助检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryExaminationText *string `json:"AuxiliaryExaminationText,omitnil" name:"AuxiliaryExaminationText"`

	// 特殊检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialExamText *string `json:"SpecialExamText,omitnil" name:"SpecialExamText"`

	// 门诊诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutpatientDiagnosisText *string `json:"OutpatientDiagnosisText,omitnil" name:"OutpatientDiagnosisText"`

	// 入院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionConditionText *string `json:"AdmissionConditionText,omitnil" name:"AdmissionConditionText"`

	// 诊疗经过
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckAndTreatmentProcessText *string `json:"CheckAndTreatmentProcessText,omitnil" name:"CheckAndTreatmentProcessText"`

	// 体征
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymptomsAndSignsText *string `json:"SymptomsAndSignsText,omitnil" name:"SymptomsAndSignsText"`

	// 出院医嘱
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeInstructionsText *string `json:"DischargeInstructionsText,omitnil" name:"DischargeInstructionsText"`

	// 入院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdmissionDiagnosisText *string `json:"AdmissionDiagnosisText,omitnil" name:"AdmissionDiagnosisText"`

	// 手术情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	SurgeryConditionText *string `json:"SurgeryConditionText,omitnil" name:"SurgeryConditionText"`

	// 病理诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathologicalDiagnosisText *string `json:"PathologicalDiagnosisText,omitnil" name:"PathologicalDiagnosisText"`

	// 出院情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeConditionText *string `json:"DischargeConditionText,omitnil" name:"DischargeConditionText"`

	// 检查记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckRecordText *string `json:"CheckRecordText,omitnil" name:"CheckRecordText"`

	// 主诉
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChiefComplaintText *string `json:"ChiefComplaintText,omitnil" name:"ChiefComplaintText"`

	// 出院诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	DischargeDiagnosisText *string `json:"DischargeDiagnosisText,omitnil" name:"DischargeDiagnosisText"`
}

type TuberInfo struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *BlockInfo `json:"Type,omitnil" name:"Type"`

	// 部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Part *Part `json:"Part,omitnil" name:"Part"`

	// 大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size []*Size `json:"Size,omitnil" name:"Size"`

	// 多发
	// 注意：此字段可能返回 null，表示取不到有效值。
	Multiple *Multiple `json:"Multiple,omitnil" name:"Multiple"`

	// 纵横比
	// 注意：此字段可能返回 null，表示取不到有效值。
	AspectRatio *AspectRatio `json:"AspectRatio,omitnil" name:"AspectRatio"`

	// 边缘
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edge *BlockInfo `json:"Edge,omitnil" name:"Edge"`

	// 内部回声
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEcho *BlockInfo `json:"InnerEcho,omitnil" name:"InnerEcho"`

	// 外部回声
	// 注意：此字段可能返回 null，表示取不到有效值。
	RearEcho *BlockInfo `json:"RearEcho,omitnil" name:"RearEcho"`

	// 弹性质地
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elastic *Elastic `json:"Elastic,omitnil" name:"Elastic"`

	// 形状
	// 注意：此字段可能返回 null，表示取不到有效值。
	Shape *BlockInfo `json:"Shape,omitnil" name:"Shape"`

	// 形态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShapeAttr *BlockInfo `json:"ShapeAttr,omitnil" name:"ShapeAttr"`

	// 皮髓质信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkinMedulla *BlockInfo `json:"SkinMedulla,omitnil" name:"SkinMedulla"`

	// 变化趋势
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trend *BlockInfo `json:"Trend,omitnil" name:"Trend"`

	// 钙化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Calcification *BlockInfo `json:"Calcification,omitnil" name:"Calcification"`

	// 包膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envelope *BlockInfo `json:"Envelope,omitnil" name:"Envelope"`

	// 强化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enhancement *BlockInfo `json:"Enhancement,omitnil" name:"Enhancement"`

	// 淋巴结
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphEnlargement *BlockInfo `json:"LymphEnlargement,omitnil" name:"LymphEnlargement"`

	// 淋巴门
	// 注意：此字段可能返回 null，表示取不到有效值。
	LymphDoor *BlockInfo `json:"LymphDoor,omitnil" name:"LymphDoor"`

	// 活动度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Activity *BlockInfo `json:"Activity,omitnil" name:"Activity"`

	// 手术情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *BlockInfo `json:"Operation,omitnil" name:"Operation"`

	// 血液cdfi
	// 注意：此字段可能返回 null，表示取不到有效值。
	CDFI *BlockInfo `json:"CDFI,omitnil" name:"CDFI"`

	// 原文位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index []*int64 `json:"Index,omitnil" name:"Index"`

	// 大小状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	SizeStatus *BlockInfo `json:"SizeStatus,omitnil" name:"SizeStatus"`

	// 内部回声分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoDistribution *BlockInfo `json:"InnerEchoDistribution,omitnil" name:"InnerEchoDistribution"`

	// 内部回声类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerEchoType []*BlockInfo `json:"InnerEchoType,omitnil" name:"InnerEchoType"`

	// 轮廓
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outline *BlockInfo `json:"Outline,omitnil" name:"Outline"`

	// 结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	Structure *BlockInfo `json:"Structure,omitnil" name:"Structure"`

	// 密度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Density *BlockInfo `json:"Density,omitnil" name:"Density"`

	// 血管
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vas *BlockInfo `json:"Vas,omitnil" name:"Vas"`

	// 囊壁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cysticwall *BlockInfo `json:"Cysticwall,omitnil" name:"Cysticwall"`

	// 被膜
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capsule *BlockInfo `json:"Capsule,omitnil" name:"Capsule"`

	// 峡部厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IsthmusThicknese is deprecated.
	IsthmusThicknese *Size `json:"IsthmusThicknese,omitnil" name:"IsthmusThicknese"`

	// 原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Src *string `json:"Src,omitnil" name:"Src"`

	// 透声度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transparent *BlockInfo `json:"Transparent,omitnil" name:"Transparent"`

	// MRI ADC
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriAdc *BlockInfo `json:"MriAdc,omitnil" name:"MriAdc"`

	// MRI DWI
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriDwi *BlockInfo `json:"MriDwi,omitnil" name:"MriDwi"`

	// MRI T1信号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriT1 *BlockInfo `json:"MriT1,omitnil" name:"MriT1"`

	// MRI T2信号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MriT2 *BlockInfo `json:"MriT2,omitnil" name:"MriT2"`

	// CT HU值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CtHu *BlockInfo `json:"CtHu,omitnil" name:"CtHu"`

	// SUmax值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suvmax *BlockInfo `json:"Suvmax,omitnil" name:"Suvmax"`

	// 代谢情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metabolism *BlockInfo `json:"Metabolism,omitnil" name:"Metabolism"`

	// 放射性摄取
	// 注意：此字段可能返回 null，表示取不到有效值。
	RadioactiveUptake *BlockInfo `json:"RadioactiveUptake,omitnil" name:"RadioactiveUptake"`

	// 病变
	// 注意：此字段可能返回 null，表示取不到有效值。
	SymDesc *BlockInfo `json:"SymDesc,omitnil" name:"SymDesc"`

	// 影像特征
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageFeature *BlockInfo `json:"ImageFeature,omitnil" name:"ImageFeature"`

	// 在报告图片中的坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coords []*Coord `json:"Coords,omitnil" name:"Coords"`

	// 峡部厚度
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsthmusThickness *Size `json:"IsthmusThickness,omitnil" name:"IsthmusThickness"`
}

// Predefined struct for user
type TurnPDFToObjectAsyncGetResultRequestParams struct {
	// 加密任务ID。在上一步通过TurnPDFToObjectAsync 接口返回的TaskID。
	// 1、建议在上一步调用TurnPDFToObjectAsync接口传入PDF之后，等5-10分钟再调用此接口获取 json 结果。如果任务还没完成，可以等待几分钟之后再重新调用此接口获取 json 结果。
	// 2、临时加密存储的 json 结果会 24 小时后定时自动删除，因此TaskID 仅 24 小时内有效。
	// 3、TaskID 与腾讯云的账号绑定，通过 TurnPDFToObjectAsync 传入 PDF 文件和通过 TurnPDFToObjectAsyncGetResult 获取 json 结果，必须是同一个腾讯云账号，否则无法获取到 json 结果。
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`
}

type TurnPDFToObjectAsyncGetResultRequest struct {
	*tchttp.BaseRequest
	
	// 加密任务ID。在上一步通过TurnPDFToObjectAsync 接口返回的TaskID。
	// 1、建议在上一步调用TurnPDFToObjectAsync接口传入PDF之后，等5-10分钟再调用此接口获取 json 结果。如果任务还没完成，可以等待几分钟之后再重新调用此接口获取 json 结果。
	// 2、临时加密存储的 json 结果会 24 小时后定时自动删除，因此TaskID 仅 24 小时内有效。
	// 3、TaskID 与腾讯云的账号绑定，通过 TurnPDFToObjectAsync 传入 PDF 文件和通过 TurnPDFToObjectAsyncGetResult 获取 json 结果，必须是同一个腾讯云账号，否则无法获取到 json 结果。
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`
}

func (r *TurnPDFToObjectAsyncGetResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectAsyncGetResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TurnPDFToObjectAsyncGetResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectAsyncGetResultResponseParams struct {
	// 报告结构化结果
	Template *Template `json:"Template,omitnil" name:"Template"`

	// 多级分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitnil" name:"TextTypeList"`

	// 报告结构化结果(体检报告PDF结构化接口返回的 json 内容非常多，建议通过本地代码调用)
	Block *Block `json:"Block,omitnil" name:"Block"`

	// 是否使用Block字段
	IsBlock *bool `json:"IsBlock,omitnil" name:"IsBlock"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TurnPDFToObjectAsyncGetResultResponse struct {
	*tchttp.BaseResponse
	Response *TurnPDFToObjectAsyncGetResultResponseParams `json:"Response"`
}

func (r *TurnPDFToObjectAsyncGetResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectAsyncGetResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectAsyncRequestParams struct {
	// 体检报告PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64)
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil" name:"PdfInfo"`
}

type TurnPDFToObjectAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 体检报告PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64)
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil" name:"PdfInfo"`
}

func (r *TurnPDFToObjectAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PdfInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TurnPDFToObjectAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectAsyncResponseParams struct {
	// 加密任务ID。 
	// 1、此 ID 是经过加密生成，是用于获取 PDF 返回 json 的凭证，需要由客户存储该 TaskID。
	// 2、建议在获取到TaskID 后，5-10分钟后再调用 TurnPDFToObjectAsyncGetResult 接口获取 json 结果。
	// 3、使用此接口，腾讯不会存储传入的 PDF 文件，但是会临时加密存储对应的 json 结果。如果不希望腾讯临时加密存储 json 结果，请使用 TurnPDFToObject 接口。
	// 4、加密存储的 json 结果会24小时后定时自动删除，因此TaskID 仅 24 小时内有效，请在24小时内调用接口 TurnPDFToObjectAsyncGetResult 获取对应 json 结果。
	// 5、TaskID 与腾讯云的账号绑定，通过 TurnPDFToObjectAsync 传入PDF文件和通过 TurnPDFToObjectAsyncGetResult 获取 json 结果，必须是同一个腾讯云账号。即其它人就算获取到 TaskID 也无法获取到 json 结果。
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TurnPDFToObjectAsyncResponse struct {
	*tchttp.BaseResponse
	Response *TurnPDFToObjectAsyncResponseParams `json:"Response"`
}

func (r *TurnPDFToObjectAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectRequestParams struct {
	// 体检报告PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64)
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil" name:"PdfInfo"`
}

type TurnPDFToObjectRequest struct {
	*tchttp.BaseRequest
	
	// 体检报告PDF文件信息, 目前只支持传PDF文件的Base64编码字符(PDF文件不能超过10MB，如果超过建议先压缩PDF，再转成base64)
	PdfInfo *PdfInfo `json:"PdfInfo,omitnil" name:"PdfInfo"`
}

func (r *TurnPDFToObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PdfInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TurnPDFToObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TurnPDFToObjectResponseParams struct {
	// 报告结构化结果
	Template *Template `json:"Template,omitnil" name:"Template"`

	// 多级分类结果
	TextTypeList []*TextType `json:"TextTypeList,omitnil" name:"TextTypeList"`

	// 报告结构化结果(体检报告PDF结构化接口返回的 json 内容非常多，建议通过本地代码调用)
	Block *Block `json:"Block,omitnil" name:"Block"`

	// 是否使用Block字段
	IsBlock *bool `json:"IsBlock,omitnil" name:"IsBlock"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TurnPDFToObjectResponse struct {
	*tchttp.BaseResponse
	Response *TurnPDFToObjectResponseParams `json:"Response"`
}

func (r *TurnPDFToObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TurnPDFToObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Vaccination struct {
	// 序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 疫苗名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vaccine *string `json:"Vaccine,omitnil" name:"Vaccine"`

	// 剂次
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dose *string `json:"Dose,omitnil" name:"Dose"`

	// 接种日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitnil" name:"Date"`

	// 疫苗批号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LotNumber *string `json:"LotNumber,omitnil" name:"LotNumber"`

	// 生产企业
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manufacturer *string `json:"Manufacturer,omitnil" name:"Manufacturer"`

	// 接种单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clinic *string `json:"Clinic,omitnil" name:"Clinic"`

	// 接种部位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Site *string `json:"Site,omitnil" name:"Site"`

	// 接种者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Provider *string `json:"Provider,omitnil" name:"Provider"`

	// 疫苗批号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lot *string `json:"Lot,omitnil" name:"Lot"`
}

type VaccineCertificate struct {
	// 免疫接种列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VaccineList []*Vaccination `json:"VaccineList,omitnil" name:"VaccineList"`
}

type Value struct {
	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitnil" name:"Grade"`

	// 百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent []*float64 `json:"Percent,omitnil" name:"Percent"`

	// 阳性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitnil" name:"Positive"`
}

type ValueBlock struct {
	// 等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitnil" name:"Grade"`

	// 百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent []*float64 `json:"Percent,omitnil" name:"Percent"`

	// 阳性阴性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Positive *string `json:"Positive,omitnil" name:"Positive"`
}

type ValueUnitItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 项目原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *PhysicalBaseItem `json:"Item,omitnil" name:"Item"`

	// 数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PhysicalBaseItem `json:"Result,omitnil" name:"Result"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *PhysicalBaseItem `json:"Unit,omitnil" name:"Unit"`
}
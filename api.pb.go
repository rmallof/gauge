// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GetProjectRootRequest
	GetProjectRootResponse
	GetInstallationRootRequest
	GetInstallationRootResponse
	GetAllStepsRequest
	GetAllStepsResponse
	GetAllSpecsRequest
	GetAllSpecsResponse
	GetAllConceptsRequest
	GetAllConceptsResponse
	ConceptInfo
	GetStepValueRequest
	GetStepValueResponse
	GetLanguagePluginLibPathRequest
	GetLanguagePluginLibPathResponse
	ErrorResponse
	APIMessage
*/
package main

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type APIMessage_APIMessageType int32

const (
	APIMessage_GetProjectRootRequest            APIMessage_APIMessageType = 1
	APIMessage_GetProjectRootResponse           APIMessage_APIMessageType = 2
	APIMessage_GetInstallationRootRequest       APIMessage_APIMessageType = 3
	APIMessage_GetInstallationRootResponse      APIMessage_APIMessageType = 4
	APIMessage_GetAllStepsRequest               APIMessage_APIMessageType = 5
	APIMessage_GetAllStepResponse               APIMessage_APIMessageType = 6
	APIMessage_GetAllSpecsRequest               APIMessage_APIMessageType = 7
	APIMessage_GetAllSpecsResponse              APIMessage_APIMessageType = 8
	APIMessage_GetStepValueRequest              APIMessage_APIMessageType = 9
	APIMessage_GetStepValueResponse             APIMessage_APIMessageType = 10
	APIMessage_GetAllConceptsRequest            APIMessage_APIMessageType = 11
	APIMessage_GetAllConceptsResponse           APIMessage_APIMessageType = 12
	APIMessage_GetLanguagePluginLibPathRequest  APIMessage_APIMessageType = 13
	APIMessage_GetLanguagePluginLibPathResponse APIMessage_APIMessageType = 14
	APIMessage_ErrorResponse                    APIMessage_APIMessageType = 15
)

var APIMessage_APIMessageType_name = map[int32]string{
	1:  "GetProjectRootRequest",
	2:  "GetProjectRootResponse",
	3:  "GetInstallationRootRequest",
	4:  "GetInstallationRootResponse",
	5:  "GetAllStepsRequest",
	6:  "GetAllStepResponse",
	7:  "GetAllSpecsRequest",
	8:  "GetAllSpecsResponse",
	9:  "GetStepValueRequest",
	10: "GetStepValueResponse",
	11: "GetAllConceptsRequest",
	12: "GetAllConceptsResponse",
	13: "GetLanguagePluginLibPathRequest",
	14: "GetLanguagePluginLibPathResponse",
	15: "ErrorResponse",
}
var APIMessage_APIMessageType_value = map[string]int32{
	"GetProjectRootRequest":            1,
	"GetProjectRootResponse":           2,
	"GetInstallationRootRequest":       3,
	"GetInstallationRootResponse":      4,
	"GetAllStepsRequest":               5,
	"GetAllStepResponse":               6,
	"GetAllSpecsRequest":               7,
	"GetAllSpecsResponse":              8,
	"GetStepValueRequest":              9,
	"GetStepValueResponse":             10,
	"GetAllConceptsRequest":            11,
	"GetAllConceptsResponse":           12,
	"GetLanguagePluginLibPathRequest":  13,
	"GetLanguagePluginLibPathResponse": 14,
	"ErrorResponse":                    15,
}

func (x APIMessage_APIMessageType) Enum() *APIMessage_APIMessageType {
	p := new(APIMessage_APIMessageType)
	*p = x
	return p
}
func (x APIMessage_APIMessageType) String() string {
	return proto.EnumName(APIMessage_APIMessageType_name, int32(x))
}
func (x *APIMessage_APIMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(APIMessage_APIMessageType_value, data, "APIMessage_APIMessageType")
	if err != nil {
		return err
	}
	*x = APIMessage_APIMessageType(value)
	return nil
}

type GetProjectRootRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetProjectRootRequest) Reset()         { *m = GetProjectRootRequest{} }
func (m *GetProjectRootRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootRequest) ProtoMessage()    {}

type GetProjectRootResponse struct {
	ProjectRoot      *string `protobuf:"bytes,1,req,name=projectRoot" json:"projectRoot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetProjectRootResponse) Reset()         { *m = GetProjectRootResponse{} }
func (m *GetProjectRootResponse) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootResponse) ProtoMessage()    {}

func (m *GetProjectRootResponse) GetProjectRoot() string {
	if m != nil && m.ProjectRoot != nil {
		return *m.ProjectRoot
	}
	return ""
}

type GetInstallationRootRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetInstallationRootRequest) Reset()         { *m = GetInstallationRootRequest{} }
func (m *GetInstallationRootRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstallationRootRequest) ProtoMessage()    {}

type GetInstallationRootResponse struct {
	InstallationRoot *string `protobuf:"bytes,1,req,name=installationRoot" json:"installationRoot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetInstallationRootResponse) Reset()         { *m = GetInstallationRootResponse{} }
func (m *GetInstallationRootResponse) String() string { return proto.CompactTextString(m) }
func (*GetInstallationRootResponse) ProtoMessage()    {}

func (m *GetInstallationRootResponse) GetInstallationRoot() string {
	if m != nil && m.InstallationRoot != nil {
		return *m.InstallationRoot
	}
	return ""
}

type GetAllStepsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllStepsRequest) Reset()         { *m = GetAllStepsRequest{} }
func (m *GetAllStepsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsRequest) ProtoMessage()    {}

type GetAllStepsResponse struct {
	AllSteps         []*ProtoStepValue `protobuf:"bytes,1,rep,name=allSteps" json:"allSteps,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *GetAllStepsResponse) Reset()         { *m = GetAllStepsResponse{} }
func (m *GetAllStepsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsResponse) ProtoMessage()    {}

func (m *GetAllStepsResponse) GetAllSteps() []*ProtoStepValue {
	if m != nil {
		return m.AllSteps
	}
	return nil
}

type GetAllSpecsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllSpecsRequest) Reset()         { *m = GetAllSpecsRequest{} }
func (m *GetAllSpecsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsRequest) ProtoMessage()    {}

type GetAllSpecsResponse struct {
	Specs            []*ProtoSpec `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GetAllSpecsResponse) Reset()         { *m = GetAllSpecsResponse{} }
func (m *GetAllSpecsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsResponse) ProtoMessage()    {}

func (m *GetAllSpecsResponse) GetSpecs() []*ProtoSpec {
	if m != nil {
		return m.Specs
	}
	return nil
}

type GetAllConceptsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllConceptsRequest) Reset()         { *m = GetAllConceptsRequest{} }
func (m *GetAllConceptsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllConceptsRequest) ProtoMessage()    {}

type GetAllConceptsResponse struct {
	Concepts         []*ConceptInfo `protobuf:"bytes,1,rep,name=concepts" json:"concepts,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *GetAllConceptsResponse) Reset()         { *m = GetAllConceptsResponse{} }
func (m *GetAllConceptsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllConceptsResponse) ProtoMessage()    {}

func (m *GetAllConceptsResponse) GetConcepts() []*ConceptInfo {
	if m != nil {
		return m.Concepts
	}
	return nil
}

type ConceptInfo struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Filepath         *string `protobuf:"bytes,2,req,name=filepath" json:"filepath,omitempty"`
	LineNumber       *int32  `protobuf:"varint,3,req,name=lineNumber" json:"lineNumber,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ConceptInfo) Reset()         { *m = ConceptInfo{} }
func (m *ConceptInfo) String() string { return proto.CompactTextString(m) }
func (*ConceptInfo) ProtoMessage()    {}

func (m *ConceptInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ConceptInfo) GetFilepath() string {
	if m != nil && m.Filepath != nil {
		return *m.Filepath
	}
	return ""
}

func (m *ConceptInfo) GetLineNumber() int32 {
	if m != nil && m.LineNumber != nil {
		return *m.LineNumber
	}
	return 0
}

type GetStepValueRequest struct {
	StepText         *string `protobuf:"bytes,1,req,name=stepText" json:"stepText,omitempty"`
	HasInlineTable   *bool   `protobuf:"varint,2,opt,name=hasInlineTable" json:"hasInlineTable,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetStepValueRequest) Reset()         { *m = GetStepValueRequest{} }
func (m *GetStepValueRequest) String() string { return proto.CompactTextString(m) }
func (*GetStepValueRequest) ProtoMessage()    {}

func (m *GetStepValueRequest) GetStepText() string {
	if m != nil && m.StepText != nil {
		return *m.StepText
	}
	return ""
}

func (m *GetStepValueRequest) GetHasInlineTable() bool {
	if m != nil && m.HasInlineTable != nil {
		return *m.HasInlineTable
	}
	return false
}

type GetStepValueResponse struct {
	StepValue        *ProtoStepValue `protobuf:"bytes,1,req,name=stepValue" json:"stepValue,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GetStepValueResponse) Reset()         { *m = GetStepValueResponse{} }
func (m *GetStepValueResponse) String() string { return proto.CompactTextString(m) }
func (*GetStepValueResponse) ProtoMessage()    {}

func (m *GetStepValueResponse) GetStepValue() *ProtoStepValue {
	if m != nil {
		return m.StepValue
	}
	return nil
}

type GetLanguagePluginLibPathRequest struct {
	Language         *string `protobuf:"bytes,1,req,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetLanguagePluginLibPathRequest) Reset()         { *m = GetLanguagePluginLibPathRequest{} }
func (m *GetLanguagePluginLibPathRequest) String() string { return proto.CompactTextString(m) }
func (*GetLanguagePluginLibPathRequest) ProtoMessage()    {}

func (m *GetLanguagePluginLibPathRequest) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

type GetLanguagePluginLibPathResponse struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetLanguagePluginLibPathResponse) Reset()         { *m = GetLanguagePluginLibPathResponse{} }
func (m *GetLanguagePluginLibPathResponse) String() string { return proto.CompactTextString(m) }
func (*GetLanguagePluginLibPathResponse) ProtoMessage()    {}

func (m *GetLanguagePluginLibPathResponse) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

type ErrorResponse struct {
	Error            *string `protobuf:"bytes,1,req,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}

func (m *ErrorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type APIMessage struct {
	MessageType              *APIMessage_APIMessageType        `protobuf:"varint,1,req,name=messageType,enum=main.APIMessage_APIMessageType" json:"messageType,omitempty"`
	MessageId                *int64                            `protobuf:"varint,2,req,name=messageId" json:"messageId,omitempty"`
	ProjectRootRequest       *GetProjectRootRequest            `protobuf:"bytes,3,opt,name=projectRootRequest" json:"projectRootRequest,omitempty"`
	ProjectRootResponse      *GetProjectRootResponse           `protobuf:"bytes,4,opt,name=projectRootResponse" json:"projectRootResponse,omitempty"`
	InstallationRootRequest  *GetInstallationRootRequest       `protobuf:"bytes,5,opt,name=installationRootRequest" json:"installationRootRequest,omitempty"`
	InstallationRootResponse *GetInstallationRootResponse      `protobuf:"bytes,6,opt,name=installationRootResponse" json:"installationRootResponse,omitempty"`
	AllStepsRequest          *GetAllStepsRequest               `protobuf:"bytes,7,opt,name=allStepsRequest" json:"allStepsRequest,omitempty"`
	AllStepsResponse         *GetAllStepsResponse              `protobuf:"bytes,8,opt,name=allStepsResponse" json:"allStepsResponse,omitempty"`
	AllSpecsRequest          *GetAllSpecsRequest               `protobuf:"bytes,9,opt,name=allSpecsRequest" json:"allSpecsRequest,omitempty"`
	AllSpecsResponse         *GetAllSpecsResponse              `protobuf:"bytes,10,opt,name=allSpecsResponse" json:"allSpecsResponse,omitempty"`
	AllConceptsRequest       *GetAllConceptsRequest            `protobuf:"bytes,11,opt,name=allConceptsRequest" json:"allConceptsRequest,omitempty"`
	AllConceptsResponse      *GetAllConceptsResponse           `protobuf:"bytes,12,opt,name=allConceptsResponse" json:"allConceptsResponse,omitempty"`
	StepValueRequest         *GetStepValueRequest              `protobuf:"bytes,13,opt,name=stepValueRequest" json:"stepValueRequest,omitempty"`
	StepValueResponse        *GetStepValueResponse             `protobuf:"bytes,14,opt,name=stepValueResponse" json:"stepValueResponse,omitempty"`
	LibPathRequest           *GetLanguagePluginLibPathRequest  `protobuf:"bytes,15,opt,name=libPathRequest" json:"libPathRequest,omitempty"`
	LibPathResponse          *GetLanguagePluginLibPathResponse `protobuf:"bytes,16,opt,name=libPathResponse" json:"libPathResponse,omitempty"`
	Error                    *ErrorResponse                    `protobuf:"bytes,17,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized         []byte                            `json:"-"`
}

func (m *APIMessage) Reset()         { *m = APIMessage{} }
func (m *APIMessage) String() string { return proto.CompactTextString(m) }
func (*APIMessage) ProtoMessage()    {}

func (m *APIMessage) GetMessageType() APIMessage_APIMessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return APIMessage_GetProjectRootRequest
}

func (m *APIMessage) GetMessageId() int64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *APIMessage) GetProjectRootRequest() *GetProjectRootRequest {
	if m != nil {
		return m.ProjectRootRequest
	}
	return nil
}

func (m *APIMessage) GetProjectRootResponse() *GetProjectRootResponse {
	if m != nil {
		return m.ProjectRootResponse
	}
	return nil
}

func (m *APIMessage) GetInstallationRootRequest() *GetInstallationRootRequest {
	if m != nil {
		return m.InstallationRootRequest
	}
	return nil
}

func (m *APIMessage) GetInstallationRootResponse() *GetInstallationRootResponse {
	if m != nil {
		return m.InstallationRootResponse
	}
	return nil
}

func (m *APIMessage) GetAllStepsRequest() *GetAllStepsRequest {
	if m != nil {
		return m.AllStepsRequest
	}
	return nil
}

func (m *APIMessage) GetAllStepsResponse() *GetAllStepsResponse {
	if m != nil {
		return m.AllStepsResponse
	}
	return nil
}

func (m *APIMessage) GetAllSpecsRequest() *GetAllSpecsRequest {
	if m != nil {
		return m.AllSpecsRequest
	}
	return nil
}

func (m *APIMessage) GetAllSpecsResponse() *GetAllSpecsResponse {
	if m != nil {
		return m.AllSpecsResponse
	}
	return nil
}

func (m *APIMessage) GetAllConceptsRequest() *GetAllConceptsRequest {
	if m != nil {
		return m.AllConceptsRequest
	}
	return nil
}

func (m *APIMessage) GetAllConceptsResponse() *GetAllConceptsResponse {
	if m != nil {
		return m.AllConceptsResponse
	}
	return nil
}

func (m *APIMessage) GetStepValueRequest() *GetStepValueRequest {
	if m != nil {
		return m.StepValueRequest
	}
	return nil
}

func (m *APIMessage) GetStepValueResponse() *GetStepValueResponse {
	if m != nil {
		return m.StepValueResponse
	}
	return nil
}

func (m *APIMessage) GetLibPathRequest() *GetLanguagePluginLibPathRequest {
	if m != nil {
		return m.LibPathRequest
	}
	return nil
}

func (m *APIMessage) GetLibPathResponse() *GetLanguagePluginLibPathResponse {
	if m != nil {
		return m.LibPathResponse
	}
	return nil
}

func (m *APIMessage) GetError() *ErrorResponse {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.APIMessage_APIMessageType", APIMessage_APIMessageType_name, APIMessage_APIMessageType_value)
}

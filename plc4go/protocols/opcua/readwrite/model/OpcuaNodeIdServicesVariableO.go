/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableO is an enum
type OpcuaNodeIdServicesVariableO int32

type IOpcuaNodeIdServicesVariableO interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceUri                              OpcuaNodeIdServicesVariableO = 15958
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceVersion                          OpcuaNodeIdServicesVariableO = 15959
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespacePublicationDate                  OpcuaNodeIdServicesVariableO = 15960
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_IsNamespaceSubset                         OpcuaNodeIdServicesVariableO = 15961
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNodeIdTypes                         OpcuaNodeIdServicesVariableO = 15962
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNumericNodeIdRange                  OpcuaNodeIdServicesVariableO = 15963
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticStringNodeIdPattern                 OpcuaNodeIdServicesVariableO = 15964
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Size                        OpcuaNodeIdServicesVariableO = 15966
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Writable                    OpcuaNodeIdServicesVariableO = 15967
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_UserWritable                OpcuaNodeIdServicesVariableO = 15968
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_OpenCount                   OpcuaNodeIdServicesVariableO = 15969
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MimeType                    OpcuaNodeIdServicesVariableO = 15970
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_InputArguments         OpcuaNodeIdServicesVariableO = 15972
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_OutputArguments        OpcuaNodeIdServicesVariableO = 15973
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Close_InputArguments        OpcuaNodeIdServicesVariableO = 15975
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_InputArguments         OpcuaNodeIdServicesVariableO = 15977
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_OutputArguments        OpcuaNodeIdServicesVariableO = 15978
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Write_InputArguments        OpcuaNodeIdServicesVariableO = 15980
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_InputArguments  OpcuaNodeIdServicesVariableO = 15982
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_OutputArguments OpcuaNodeIdServicesVariableO = 15983
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_SetPosition_InputArguments  OpcuaNodeIdServicesVariableO = 15985
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultRolePermissions                    OpcuaNodeIdServicesVariableO = 16134
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultUserRolePermissions                OpcuaNodeIdServicesVariableO = 16135
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultAccessRestrictions                 OpcuaNodeIdServicesVariableO = 16136
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MaxByteStringLength         OpcuaNodeIdServicesVariableO = 24243
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_LastModifiedTime            OpcuaNodeIdServicesVariableO = 25199
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ConfigurationVersion                      OpcuaNodeIdServicesVariableO = 25266
	OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ModelVersion                              OpcuaNodeIdServicesVariableO = 32408
)

var OpcuaNodeIdServicesVariableOValues []OpcuaNodeIdServicesVariableO

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableOValues = []OpcuaNodeIdServicesVariableO{
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceUri,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceVersion,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespacePublicationDate,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_IsNamespaceSubset,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNodeIdTypes,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNumericNodeIdRange,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticStringNodeIdPattern,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Size,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Writable,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_UserWritable,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_OpenCount,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MimeType,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_InputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_OutputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Close_InputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_InputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_OutputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Write_InputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_InputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_OutputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_SetPosition_InputArguments,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultRolePermissions,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultUserRolePermissions,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultAccessRestrictions,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MaxByteStringLength,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_LastModifiedTime,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ConfigurationVersion,
		OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ModelVersion,
	}
}

func OpcuaNodeIdServicesVariableOByValue(value int32) (enum OpcuaNodeIdServicesVariableO, ok bool) {
	switch value {
	case 15958:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceUri, true
	case 15959:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceVersion, true
	case 15960:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespacePublicationDate, true
	case 15961:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_IsNamespaceSubset, true
	case 15962:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNodeIdTypes, true
	case 15963:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNumericNodeIdRange, true
	case 15964:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticStringNodeIdPattern, true
	case 15966:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Size, true
	case 15967:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Writable, true
	case 15968:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_UserWritable, true
	case 15969:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_OpenCount, true
	case 15970:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MimeType, true
	case 15972:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_InputArguments, true
	case 15973:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_OutputArguments, true
	case 15975:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Close_InputArguments, true
	case 15977:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_InputArguments, true
	case 15978:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_OutputArguments, true
	case 15980:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Write_InputArguments, true
	case 15982:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_InputArguments, true
	case 15983:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_OutputArguments, true
	case 15985:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_SetPosition_InputArguments, true
	case 16134:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultRolePermissions, true
	case 16135:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultUserRolePermissions, true
	case 16136:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultAccessRestrictions, true
	case 24243:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MaxByteStringLength, true
	case 25199:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_LastModifiedTime, true
	case 25266:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ConfigurationVersion, true
	case 32408:
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ModelVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOByName(value string) (enum OpcuaNodeIdServicesVariableO, ok bool) {
	switch value {
	case "OPCUANamespaceMetadata_NamespaceUri":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceUri, true
	case "OPCUANamespaceMetadata_NamespaceVersion":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceVersion, true
	case "OPCUANamespaceMetadata_NamespacePublicationDate":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespacePublicationDate, true
	case "OPCUANamespaceMetadata_IsNamespaceSubset":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_IsNamespaceSubset, true
	case "OPCUANamespaceMetadata_StaticNodeIdTypes":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNodeIdTypes, true
	case "OPCUANamespaceMetadata_StaticNumericNodeIdRange":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNumericNodeIdRange, true
	case "OPCUANamespaceMetadata_StaticStringNodeIdPattern":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticStringNodeIdPattern, true
	case "OPCUANamespaceMetadata_NamespaceFile_Size":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Size, true
	case "OPCUANamespaceMetadata_NamespaceFile_Writable":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Writable, true
	case "OPCUANamespaceMetadata_NamespaceFile_UserWritable":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_UserWritable, true
	case "OPCUANamespaceMetadata_NamespaceFile_OpenCount":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_OpenCount, true
	case "OPCUANamespaceMetadata_NamespaceFile_MimeType":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MimeType, true
	case "OPCUANamespaceMetadata_NamespaceFile_Open_InputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_InputArguments, true
	case "OPCUANamespaceMetadata_NamespaceFile_Open_OutputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_OutputArguments, true
	case "OPCUANamespaceMetadata_NamespaceFile_Close_InputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Close_InputArguments, true
	case "OPCUANamespaceMetadata_NamespaceFile_Read_InputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_InputArguments, true
	case "OPCUANamespaceMetadata_NamespaceFile_Read_OutputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_OutputArguments, true
	case "OPCUANamespaceMetadata_NamespaceFile_Write_InputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Write_InputArguments, true
	case "OPCUANamespaceMetadata_NamespaceFile_GetPosition_InputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_InputArguments, true
	case "OPCUANamespaceMetadata_NamespaceFile_GetPosition_OutputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_OutputArguments, true
	case "OPCUANamespaceMetadata_NamespaceFile_SetPosition_InputArguments":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_SetPosition_InputArguments, true
	case "OPCUANamespaceMetadata_DefaultRolePermissions":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultRolePermissions, true
	case "OPCUANamespaceMetadata_DefaultUserRolePermissions":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultUserRolePermissions, true
	case "OPCUANamespaceMetadata_DefaultAccessRestrictions":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultAccessRestrictions, true
	case "OPCUANamespaceMetadata_NamespaceFile_MaxByteStringLength":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MaxByteStringLength, true
	case "OPCUANamespaceMetadata_NamespaceFile_LastModifiedTime":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_LastModifiedTime, true
	case "OPCUANamespaceMetadata_ConfigurationVersion":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ConfigurationVersion, true
	case "OPCUANamespaceMetadata_ModelVersion":
		return OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ModelVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableOValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableO(structType any) OpcuaNodeIdServicesVariableO {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableO {
		if sOpcuaNodeIdServicesVariableO, ok := typ.(OpcuaNodeIdServicesVariableO); ok {
			return sOpcuaNodeIdServicesVariableO
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableO) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableO) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableOParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableO, error) {
	return OpcuaNodeIdServicesVariableOParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableOParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableO, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableO", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableO")
	}
	if enum, ok := OpcuaNodeIdServicesVariableOByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableO")
		return OpcuaNodeIdServicesVariableO(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableO) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableO) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableO", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableO) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceUri:
		return "OPCUANamespaceMetadata_NamespaceUri"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceVersion:
		return "OPCUANamespaceMetadata_NamespaceVersion"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespacePublicationDate:
		return "OPCUANamespaceMetadata_NamespacePublicationDate"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_IsNamespaceSubset:
		return "OPCUANamespaceMetadata_IsNamespaceSubset"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNodeIdTypes:
		return "OPCUANamespaceMetadata_StaticNodeIdTypes"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticNumericNodeIdRange:
		return "OPCUANamespaceMetadata_StaticNumericNodeIdRange"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_StaticStringNodeIdPattern:
		return "OPCUANamespaceMetadata_StaticStringNodeIdPattern"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Size:
		return "OPCUANamespaceMetadata_NamespaceFile_Size"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Writable:
		return "OPCUANamespaceMetadata_NamespaceFile_Writable"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_UserWritable:
		return "OPCUANamespaceMetadata_NamespaceFile_UserWritable"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_OpenCount:
		return "OPCUANamespaceMetadata_NamespaceFile_OpenCount"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MimeType:
		return "OPCUANamespaceMetadata_NamespaceFile_MimeType"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_InputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_Open_InputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Open_OutputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_Open_OutputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Close_InputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_Close_InputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_InputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_Read_InputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Read_OutputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_Read_OutputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_Write_InputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_Write_InputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_InputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_GetPosition_InputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_GetPosition_OutputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_GetPosition_OutputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_SetPosition_InputArguments:
		return "OPCUANamespaceMetadata_NamespaceFile_SetPosition_InputArguments"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultRolePermissions:
		return "OPCUANamespaceMetadata_DefaultRolePermissions"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultUserRolePermissions:
		return "OPCUANamespaceMetadata_DefaultUserRolePermissions"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_DefaultAccessRestrictions:
		return "OPCUANamespaceMetadata_DefaultAccessRestrictions"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_MaxByteStringLength:
		return "OPCUANamespaceMetadata_NamespaceFile_MaxByteStringLength"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_NamespaceFile_LastModifiedTime:
		return "OPCUANamespaceMetadata_NamespaceFile_LastModifiedTime"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ConfigurationVersion:
		return "OPCUANamespaceMetadata_ConfigurationVersion"
	case OpcuaNodeIdServicesVariableO_OPCUANamespaceMetadata_ModelVersion:
		return "OPCUANamespaceMetadata_ModelVersion"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableO) String() string {
	return e.PLC4XEnumName()
}

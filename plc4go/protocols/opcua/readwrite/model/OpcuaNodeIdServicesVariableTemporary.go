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

// OpcuaNodeIdServicesVariableTemporary is an enum
type OpcuaNodeIdServicesVariableTemporary int32

type IOpcuaNodeIdServicesVariableTemporary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_ClientProcessingTimeout                                          OpcuaNodeIdServicesVariableTemporary = 15745
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_InputArguments                               OpcuaNodeIdServicesVariableTemporary = 15747
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_OutputArguments                              OpcuaNodeIdServicesVariableTemporary = 15748
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_OutputArguments                             OpcuaNodeIdServicesVariableTemporary = 15750
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_InputArguments                                    OpcuaNodeIdServicesVariableTemporary = 15752
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_OutputArguments                                   OpcuaNodeIdServicesVariableTemporary = 15753
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState                           OpcuaNodeIdServicesVariableTemporary = 15755
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Id                        OpcuaNodeIdServicesVariableTemporary = 15756
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Name                      OpcuaNodeIdServicesVariableTemporary = 15757
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Number                    OpcuaNodeIdServicesVariableTemporary = 15758
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_EffectiveDisplayName      OpcuaNodeIdServicesVariableTemporary = 15759
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition                         OpcuaNodeIdServicesVariableTemporary = 15760
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Id                      OpcuaNodeIdServicesVariableTemporary = 15761
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Name                    OpcuaNodeIdServicesVariableTemporary = 15762
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Number                  OpcuaNodeIdServicesVariableTemporary = 15763
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_TransitionTime          OpcuaNodeIdServicesVariableTemporary = 15764
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_EffectiveTransitionTime OpcuaNodeIdServicesVariableTemporary = 15765
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_InputArguments                              OpcuaNodeIdServicesVariableTemporary = 16359
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableStates                        OpcuaNodeIdServicesVariableTemporary = 17637
	OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableTransitions                   OpcuaNodeIdServicesVariableTemporary = 17638
)

var OpcuaNodeIdServicesVariableTemporaryValues []OpcuaNodeIdServicesVariableTemporary

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTemporaryValues = []OpcuaNodeIdServicesVariableTemporary{
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_ClientProcessingTimeout,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_InputArguments,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_OutputArguments,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_OutputArguments,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_InputArguments,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_OutputArguments,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Id,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Name,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Number,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Id,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Name,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Number,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_TransitionTime,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_EffectiveTransitionTime,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_InputArguments,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableStates,
		OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableTransitions,
	}
}

func OpcuaNodeIdServicesVariableTemporaryByValue(value int32) (enum OpcuaNodeIdServicesVariableTemporary, ok bool) {
	switch value {
	case 15745:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_ClientProcessingTimeout, true
	case 15747:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_InputArguments, true
	case 15748:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_OutputArguments, true
	case 15750:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_OutputArguments, true
	case 15752:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_InputArguments, true
	case 15753:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_OutputArguments, true
	case 15755:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState, true
	case 15756:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Id, true
	case 15757:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Name, true
	case 15758:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Number, true
	case 15759:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_EffectiveDisplayName, true
	case 15760:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition, true
	case 15761:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Id, true
	case 15762:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Name, true
	case 15763:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Number, true
	case 15764:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_TransitionTime, true
	case 15765:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_EffectiveTransitionTime, true
	case 16359:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_InputArguments, true
	case 17637:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableStates, true
	case 17638:
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableTransitions, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTemporaryByName(value string) (enum OpcuaNodeIdServicesVariableTemporary, ok bool) {
	switch value {
	case "TemporaryFileTransferType_ClientProcessingTimeout":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_ClientProcessingTimeout, true
	case "TemporaryFileTransferType_GenerateFileForRead_InputArguments":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_InputArguments, true
	case "TemporaryFileTransferType_GenerateFileForRead_OutputArguments":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_OutputArguments, true
	case "TemporaryFileTransferType_GenerateFileForWrite_OutputArguments":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_OutputArguments, true
	case "TemporaryFileTransferType_CloseAndCommit_InputArguments":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_InputArguments, true
	case "TemporaryFileTransferType_CloseAndCommit_OutputArguments":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_OutputArguments, true
	case "TemporaryFileTransferType_TransferState_Placeholder_CurrentState":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState, true
	case "TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Id":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Id, true
	case "TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Name":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Name, true
	case "TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Number":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Number, true
	case "TemporaryFileTransferType_TransferState_Placeholder_CurrentState_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_EffectiveDisplayName, true
	case "TemporaryFileTransferType_TransferState_Placeholder_LastTransition":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition, true
	case "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Id":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Id, true
	case "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Name":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Name, true
	case "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Number":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Number, true
	case "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_TransitionTime":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_TransitionTime, true
	case "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_EffectiveTransitionTime, true
	case "TemporaryFileTransferType_GenerateFileForWrite_InputArguments":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_InputArguments, true
	case "TemporaryFileTransferType_TransferState_Placeholder_AvailableStates":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableStates, true
	case "TemporaryFileTransferType_TransferState_Placeholder_AvailableTransitions":
		return OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableTransitions, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTemporaryKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTemporaryValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableTemporary(structType any) OpcuaNodeIdServicesVariableTemporary {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableTemporary {
		if sOpcuaNodeIdServicesVariableTemporary, ok := typ.(OpcuaNodeIdServicesVariableTemporary); ok {
			return sOpcuaNodeIdServicesVariableTemporary
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableTemporary) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableTemporary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTemporaryParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableTemporary, error) {
	return OpcuaNodeIdServicesVariableTemporaryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTemporaryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableTemporary, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableTemporary", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableTemporary")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTemporaryByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableTemporary")
		return OpcuaNodeIdServicesVariableTemporary(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableTemporary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableTemporary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableTemporary", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableTemporary) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_ClientProcessingTimeout:
		return "TemporaryFileTransferType_ClientProcessingTimeout"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_InputArguments:
		return "TemporaryFileTransferType_GenerateFileForRead_InputArguments"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForRead_OutputArguments:
		return "TemporaryFileTransferType_GenerateFileForRead_OutputArguments"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_OutputArguments:
		return "TemporaryFileTransferType_GenerateFileForWrite_OutputArguments"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_InputArguments:
		return "TemporaryFileTransferType_CloseAndCommit_InputArguments"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_CloseAndCommit_OutputArguments:
		return "TemporaryFileTransferType_CloseAndCommit_OutputArguments"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState:
		return "TemporaryFileTransferType_TransferState_Placeholder_CurrentState"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Id:
		return "TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Id"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Name:
		return "TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Name"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Number:
		return "TemporaryFileTransferType_TransferState_Placeholder_CurrentState_Number"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_CurrentState_EffectiveDisplayName:
		return "TemporaryFileTransferType_TransferState_Placeholder_CurrentState_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition:
		return "TemporaryFileTransferType_TransferState_Placeholder_LastTransition"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Id:
		return "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Id"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Name:
		return "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Name"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Number:
		return "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_Number"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_TransitionTime:
		return "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_TransitionTime"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_LastTransition_EffectiveTransitionTime:
		return "TemporaryFileTransferType_TransferState_Placeholder_LastTransition_EffectiveTransitionTime"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_GenerateFileForWrite_InputArguments:
		return "TemporaryFileTransferType_GenerateFileForWrite_InputArguments"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableStates:
		return "TemporaryFileTransferType_TransferState_Placeholder_AvailableStates"
	case OpcuaNodeIdServicesVariableTemporary_TemporaryFileTransferType_TransferState_Placeholder_AvailableTransitions:
		return "TemporaryFileTransferType_TransferState_Placeholder_AvailableTransitions"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableTemporary) String() string {
	return e.PLC4XEnumName()
}

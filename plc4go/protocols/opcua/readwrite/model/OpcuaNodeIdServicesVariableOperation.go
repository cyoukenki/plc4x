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

// OpcuaNodeIdServicesVariableOperation is an enum
type OpcuaNodeIdServicesVariableOperation int32

type IOpcuaNodeIdServicesVariableOperation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRead                          OpcuaNodeIdServicesVariableOperation = 11565
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerWrite                         OpcuaNodeIdServicesVariableOperation = 11567
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerMethodCall                    OpcuaNodeIdServicesVariableOperation = 11569
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerBrowse                        OpcuaNodeIdServicesVariableOperation = 11570
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRegisterNodes                 OpcuaNodeIdServicesVariableOperation = 11571
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerTranslateBrowsePathsToNodeIds OpcuaNodeIdServicesVariableOperation = 11572
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerNodeManagement                OpcuaNodeIdServicesVariableOperation = 11573
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxMonitoredItemsPerCall                 OpcuaNodeIdServicesVariableOperation = 11574
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadData               OpcuaNodeIdServicesVariableOperation = 12161
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadEvents             OpcuaNodeIdServicesVariableOperation = 12162
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateData             OpcuaNodeIdServicesVariableOperation = 12163
	OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateEvents           OpcuaNodeIdServicesVariableOperation = 12164
)

var OpcuaNodeIdServicesVariableOperationValues []OpcuaNodeIdServicesVariableOperation

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableOperationValues = []OpcuaNodeIdServicesVariableOperation{
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRead,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerWrite,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerMethodCall,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerBrowse,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRegisterNodes,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerTranslateBrowsePathsToNodeIds,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerNodeManagement,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxMonitoredItemsPerCall,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadData,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadEvents,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateData,
		OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateEvents,
	}
}

func OpcuaNodeIdServicesVariableOperationByValue(value int32) (enum OpcuaNodeIdServicesVariableOperation, ok bool) {
	switch value {
	case 11565:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRead, true
	case 11567:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerWrite, true
	case 11569:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerMethodCall, true
	case 11570:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerBrowse, true
	case 11571:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRegisterNodes, true
	case 11572:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerTranslateBrowsePathsToNodeIds, true
	case 11573:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerNodeManagement, true
	case 11574:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxMonitoredItemsPerCall, true
	case 12161:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadData, true
	case 12162:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadEvents, true
	case 12163:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateData, true
	case 12164:
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateEvents, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOperationByName(value string) (enum OpcuaNodeIdServicesVariableOperation, ok bool) {
	switch value {
	case "OperationLimitsType_MaxNodesPerRead":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRead, true
	case "OperationLimitsType_MaxNodesPerWrite":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerWrite, true
	case "OperationLimitsType_MaxNodesPerMethodCall":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerMethodCall, true
	case "OperationLimitsType_MaxNodesPerBrowse":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerBrowse, true
	case "OperationLimitsType_MaxNodesPerRegisterNodes":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRegisterNodes, true
	case "OperationLimitsType_MaxNodesPerTranslateBrowsePathsToNodeIds":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerTranslateBrowsePathsToNodeIds, true
	case "OperationLimitsType_MaxNodesPerNodeManagement":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerNodeManagement, true
	case "OperationLimitsType_MaxMonitoredItemsPerCall":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxMonitoredItemsPerCall, true
	case "OperationLimitsType_MaxNodesPerHistoryReadData":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadData, true
	case "OperationLimitsType_MaxNodesPerHistoryReadEvents":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadEvents, true
	case "OperationLimitsType_MaxNodesPerHistoryUpdateData":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateData, true
	case "OperationLimitsType_MaxNodesPerHistoryUpdateEvents":
		return OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateEvents, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOperationKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableOperationValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableOperation(structType any) OpcuaNodeIdServicesVariableOperation {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableOperation {
		if sOpcuaNodeIdServicesVariableOperation, ok := typ.(OpcuaNodeIdServicesVariableOperation); ok {
			return sOpcuaNodeIdServicesVariableOperation
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableOperation) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableOperation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableOperationParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableOperation, error) {
	return OpcuaNodeIdServicesVariableOperationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableOperationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableOperation, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableOperation", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableOperation")
	}
	if enum, ok := OpcuaNodeIdServicesVariableOperationByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableOperation")
		return OpcuaNodeIdServicesVariableOperation(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableOperation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableOperation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableOperation", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableOperation) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRead:
		return "OperationLimitsType_MaxNodesPerRead"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerWrite:
		return "OperationLimitsType_MaxNodesPerWrite"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerMethodCall:
		return "OperationLimitsType_MaxNodesPerMethodCall"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerBrowse:
		return "OperationLimitsType_MaxNodesPerBrowse"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerRegisterNodes:
		return "OperationLimitsType_MaxNodesPerRegisterNodes"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerTranslateBrowsePathsToNodeIds:
		return "OperationLimitsType_MaxNodesPerTranslateBrowsePathsToNodeIds"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerNodeManagement:
		return "OperationLimitsType_MaxNodesPerNodeManagement"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxMonitoredItemsPerCall:
		return "OperationLimitsType_MaxMonitoredItemsPerCall"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadData:
		return "OperationLimitsType_MaxNodesPerHistoryReadData"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryReadEvents:
		return "OperationLimitsType_MaxNodesPerHistoryReadEvents"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateData:
		return "OperationLimitsType_MaxNodesPerHistoryUpdateData"
	case OpcuaNodeIdServicesVariableOperation_OperationLimitsType_MaxNodesPerHistoryUpdateEvents:
		return "OperationLimitsType_MaxNodesPerHistoryUpdateEvents"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableOperation) String() string {
	return e.PLC4XEnumName()
}
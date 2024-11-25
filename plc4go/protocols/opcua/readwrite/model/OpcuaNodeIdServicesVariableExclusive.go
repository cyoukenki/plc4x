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

// OpcuaNodeIdServicesVariableExclusive is an enum
type OpcuaNodeIdServicesVariableExclusive int32

type IOpcuaNodeIdServicesVariableExclusive interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHigh_StateNumber                       OpcuaNodeIdServicesVariableExclusive = 9330
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_High_StateNumber                           OpcuaNodeIdServicesVariableExclusive = 9332
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_Low_StateNumber                            OpcuaNodeIdServicesVariableExclusive = 9334
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLow_StateNumber                         OpcuaNodeIdServicesVariableExclusive = 9336
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState                                       OpcuaNodeIdServicesVariableExclusive = 9398
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Id                                    OpcuaNodeIdServicesVariableExclusive = 9399
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Name                                  OpcuaNodeIdServicesVariableExclusive = 9400
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Number                                OpcuaNodeIdServicesVariableExclusive = 9401
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveDisplayName                  OpcuaNodeIdServicesVariableExclusive = 9402
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TransitionTime                        OpcuaNodeIdServicesVariableExclusive = 9403
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveTransitionTime               OpcuaNodeIdServicesVariableExclusive = 9404
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TrueState                             OpcuaNodeIdServicesVariableExclusive = 9405
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_FalseState                            OpcuaNodeIdServicesVariableExclusive = 9406
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState                           OpcuaNodeIdServicesVariableExclusive = 9456
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Id                        OpcuaNodeIdServicesVariableExclusive = 9457
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Name                      OpcuaNodeIdServicesVariableExclusive = 9458
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Number                    OpcuaNodeIdServicesVariableExclusive = 9459
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_EffectiveDisplayName      OpcuaNodeIdServicesVariableExclusive = 9460
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition                         OpcuaNodeIdServicesVariableExclusive = 9461
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Id                      OpcuaNodeIdServicesVariableExclusive = 9462
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Name                    OpcuaNodeIdServicesVariableExclusive = 9463
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Number                  OpcuaNodeIdServicesVariableExclusive = 9464
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_TransitionTime          OpcuaNodeIdServicesVariableExclusive = 9465
	OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_SetpointNode                                  OpcuaNodeIdServicesVariableExclusive = 9905
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLowToLow_TransitionNumber               OpcuaNodeIdServicesVariableExclusive = 11340
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowToLowLow_TransitionNumber               OpcuaNodeIdServicesVariableExclusive = 11341
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHighToHigh_TransitionNumber            OpcuaNodeIdServicesVariableExclusive = 11342
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighToHighHigh_TransitionNumber            OpcuaNodeIdServicesVariableExclusive = 11343
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_EffectiveTransitionTime OpcuaNodeIdServicesVariableExclusive = 11470
	OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_BaseSetpointNode                              OpcuaNodeIdServicesVariableExclusive = 16817
	OpcuaNodeIdServicesVariableExclusive_ExclusiveRateOfChangeAlarmType_EngineeringUnits                           OpcuaNodeIdServicesVariableExclusive = 16899
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableStates                        OpcuaNodeIdServicesVariableExclusive = 17670
	OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableTransitions                   OpcuaNodeIdServicesVariableExclusive = 17671
)

var OpcuaNodeIdServicesVariableExclusiveValues []OpcuaNodeIdServicesVariableExclusive

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableExclusiveValues = []OpcuaNodeIdServicesVariableExclusive{
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHigh_StateNumber,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_High_StateNumber,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_Low_StateNumber,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLow_StateNumber,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Id,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Name,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Number,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TransitionTime,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveTransitionTime,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TrueState,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_FalseState,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Id,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Name,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Number,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Id,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Name,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Number,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_TransitionTime,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_SetpointNode,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLowToLow_TransitionNumber,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowToLowLow_TransitionNumber,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHighToHigh_TransitionNumber,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighToHighHigh_TransitionNumber,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_EffectiveTransitionTime,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_BaseSetpointNode,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveRateOfChangeAlarmType_EngineeringUnits,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableStates,
		OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableTransitions,
	}
}

func OpcuaNodeIdServicesVariableExclusiveByValue(value int32) (enum OpcuaNodeIdServicesVariableExclusive, ok bool) {
	switch value {
	case 11340:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLowToLow_TransitionNumber, true
	case 11341:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowToLowLow_TransitionNumber, true
	case 11342:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHighToHigh_TransitionNumber, true
	case 11343:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighToHighHigh_TransitionNumber, true
	case 11470:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_EffectiveTransitionTime, true
	case 16817:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_BaseSetpointNode, true
	case 16899:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveRateOfChangeAlarmType_EngineeringUnits, true
	case 17670:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableStates, true
	case 17671:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableTransitions, true
	case 9330:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHigh_StateNumber, true
	case 9332:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_High_StateNumber, true
	case 9334:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_Low_StateNumber, true
	case 9336:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLow_StateNumber, true
	case 9398:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState, true
	case 9399:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Id, true
	case 9400:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Name, true
	case 9401:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Number, true
	case 9402:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveDisplayName, true
	case 9403:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TransitionTime, true
	case 9404:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveTransitionTime, true
	case 9405:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TrueState, true
	case 9406:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_FalseState, true
	case 9456:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState, true
	case 9457:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Id, true
	case 9458:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Name, true
	case 9459:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Number, true
	case 9460:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_EffectiveDisplayName, true
	case 9461:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition, true
	case 9462:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Id, true
	case 9463:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Name, true
	case 9464:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Number, true
	case 9465:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_TransitionTime, true
	case 9905:
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_SetpointNode, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableExclusiveByName(value string) (enum OpcuaNodeIdServicesVariableExclusive, ok bool) {
	switch value {
	case "ExclusiveLimitStateMachineType_LowLowToLow_TransitionNumber":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLowToLow_TransitionNumber, true
	case "ExclusiveLimitStateMachineType_LowToLowLow_TransitionNumber":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowToLowLow_TransitionNumber, true
	case "ExclusiveLimitStateMachineType_HighHighToHigh_TransitionNumber":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHighToHigh_TransitionNumber, true
	case "ExclusiveLimitStateMachineType_HighToHighHigh_TransitionNumber":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighToHighHigh_TransitionNumber, true
	case "ExclusiveLimitAlarmType_LimitState_LastTransition_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_EffectiveTransitionTime, true
	case "ExclusiveDeviationAlarmType_BaseSetpointNode":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_BaseSetpointNode, true
	case "ExclusiveRateOfChangeAlarmType_EngineeringUnits":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveRateOfChangeAlarmType_EngineeringUnits, true
	case "ExclusiveLimitAlarmType_LimitState_AvailableStates":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableStates, true
	case "ExclusiveLimitAlarmType_LimitState_AvailableTransitions":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableTransitions, true
	case "ExclusiveLimitStateMachineType_HighHigh_StateNumber":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHigh_StateNumber, true
	case "ExclusiveLimitStateMachineType_High_StateNumber":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_High_StateNumber, true
	case "ExclusiveLimitStateMachineType_Low_StateNumber":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_Low_StateNumber, true
	case "ExclusiveLimitStateMachineType_LowLow_StateNumber":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLow_StateNumber, true
	case "ExclusiveLimitAlarmType_ActiveState":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState, true
	case "ExclusiveLimitAlarmType_ActiveState_Id":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Id, true
	case "ExclusiveLimitAlarmType_ActiveState_Name":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Name, true
	case "ExclusiveLimitAlarmType_ActiveState_Number":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Number, true
	case "ExclusiveLimitAlarmType_ActiveState_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveDisplayName, true
	case "ExclusiveLimitAlarmType_ActiveState_TransitionTime":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TransitionTime, true
	case "ExclusiveLimitAlarmType_ActiveState_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveTransitionTime, true
	case "ExclusiveLimitAlarmType_ActiveState_TrueState":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TrueState, true
	case "ExclusiveLimitAlarmType_ActiveState_FalseState":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_FalseState, true
	case "ExclusiveLimitAlarmType_LimitState_CurrentState":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState, true
	case "ExclusiveLimitAlarmType_LimitState_CurrentState_Id":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Id, true
	case "ExclusiveLimitAlarmType_LimitState_CurrentState_Name":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Name, true
	case "ExclusiveLimitAlarmType_LimitState_CurrentState_Number":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Number, true
	case "ExclusiveLimitAlarmType_LimitState_CurrentState_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_EffectiveDisplayName, true
	case "ExclusiveLimitAlarmType_LimitState_LastTransition":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition, true
	case "ExclusiveLimitAlarmType_LimitState_LastTransition_Id":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Id, true
	case "ExclusiveLimitAlarmType_LimitState_LastTransition_Name":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Name, true
	case "ExclusiveLimitAlarmType_LimitState_LastTransition_Number":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Number, true
	case "ExclusiveLimitAlarmType_LimitState_LastTransition_TransitionTime":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_TransitionTime, true
	case "ExclusiveDeviationAlarmType_SetpointNode":
		return OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_SetpointNode, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableExclusiveKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableExclusiveValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableExclusive(structType any) OpcuaNodeIdServicesVariableExclusive {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableExclusive {
		if sOpcuaNodeIdServicesVariableExclusive, ok := typ.(OpcuaNodeIdServicesVariableExclusive); ok {
			return sOpcuaNodeIdServicesVariableExclusive
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableExclusive) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableExclusive) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableExclusiveParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableExclusive, error) {
	return OpcuaNodeIdServicesVariableExclusiveParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableExclusiveParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableExclusive, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableExclusive", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableExclusive")
	}
	if enum, ok := OpcuaNodeIdServicesVariableExclusiveByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableExclusive")
		return OpcuaNodeIdServicesVariableExclusive(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableExclusive) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableExclusive) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableExclusive", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableExclusive) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLowToLow_TransitionNumber:
		return "ExclusiveLimitStateMachineType_LowLowToLow_TransitionNumber"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowToLowLow_TransitionNumber:
		return "ExclusiveLimitStateMachineType_LowToLowLow_TransitionNumber"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHighToHigh_TransitionNumber:
		return "ExclusiveLimitStateMachineType_HighHighToHigh_TransitionNumber"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighToHighHigh_TransitionNumber:
		return "ExclusiveLimitStateMachineType_HighToHighHigh_TransitionNumber"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_EffectiveTransitionTime:
		return "ExclusiveLimitAlarmType_LimitState_LastTransition_EffectiveTransitionTime"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_BaseSetpointNode:
		return "ExclusiveDeviationAlarmType_BaseSetpointNode"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveRateOfChangeAlarmType_EngineeringUnits:
		return "ExclusiveRateOfChangeAlarmType_EngineeringUnits"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableStates:
		return "ExclusiveLimitAlarmType_LimitState_AvailableStates"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_AvailableTransitions:
		return "ExclusiveLimitAlarmType_LimitState_AvailableTransitions"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_HighHigh_StateNumber:
		return "ExclusiveLimitStateMachineType_HighHigh_StateNumber"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_High_StateNumber:
		return "ExclusiveLimitStateMachineType_High_StateNumber"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_Low_StateNumber:
		return "ExclusiveLimitStateMachineType_Low_StateNumber"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitStateMachineType_LowLow_StateNumber:
		return "ExclusiveLimitStateMachineType_LowLow_StateNumber"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState:
		return "ExclusiveLimitAlarmType_ActiveState"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Id:
		return "ExclusiveLimitAlarmType_ActiveState_Id"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Name:
		return "ExclusiveLimitAlarmType_ActiveState_Name"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_Number:
		return "ExclusiveLimitAlarmType_ActiveState_Number"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveDisplayName:
		return "ExclusiveLimitAlarmType_ActiveState_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TransitionTime:
		return "ExclusiveLimitAlarmType_ActiveState_TransitionTime"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_EffectiveTransitionTime:
		return "ExclusiveLimitAlarmType_ActiveState_EffectiveTransitionTime"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_TrueState:
		return "ExclusiveLimitAlarmType_ActiveState_TrueState"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_ActiveState_FalseState:
		return "ExclusiveLimitAlarmType_ActiveState_FalseState"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState:
		return "ExclusiveLimitAlarmType_LimitState_CurrentState"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Id:
		return "ExclusiveLimitAlarmType_LimitState_CurrentState_Id"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Name:
		return "ExclusiveLimitAlarmType_LimitState_CurrentState_Name"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_Number:
		return "ExclusiveLimitAlarmType_LimitState_CurrentState_Number"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_CurrentState_EffectiveDisplayName:
		return "ExclusiveLimitAlarmType_LimitState_CurrentState_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition:
		return "ExclusiveLimitAlarmType_LimitState_LastTransition"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Id:
		return "ExclusiveLimitAlarmType_LimitState_LastTransition_Id"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Name:
		return "ExclusiveLimitAlarmType_LimitState_LastTransition_Name"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_Number:
		return "ExclusiveLimitAlarmType_LimitState_LastTransition_Number"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveLimitAlarmType_LimitState_LastTransition_TransitionTime:
		return "ExclusiveLimitAlarmType_LimitState_LastTransition_TransitionTime"
	case OpcuaNodeIdServicesVariableExclusive_ExclusiveDeviationAlarmType_SetpointNode:
		return "ExclusiveDeviationAlarmType_SetpointNode"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableExclusive) String() string {
	return e.PLC4XEnumName()
}

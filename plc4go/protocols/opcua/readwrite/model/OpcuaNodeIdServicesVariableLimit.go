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

// OpcuaNodeIdServicesVariableLimit is an enum
type OpcuaNodeIdServicesVariableLimit int32

type IOpcuaNodeIdServicesVariableLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighLimit     OpcuaNodeIdServicesVariableLimit = 11124
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighLimit         OpcuaNodeIdServicesVariableLimit = 11125
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLimit          OpcuaNodeIdServicesVariableLimit = 11126
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowLimit       OpcuaNodeIdServicesVariableLimit = 11127
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighHighLimit OpcuaNodeIdServicesVariableLimit = 16572
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighLimit     OpcuaNodeIdServicesVariableLimit = 16573
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLimit      OpcuaNodeIdServicesVariableLimit = 16574
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLowLimit   OpcuaNodeIdServicesVariableLimit = 16575
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHighHigh  OpcuaNodeIdServicesVariableLimit = 24770
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHigh      OpcuaNodeIdServicesVariableLimit = 24771
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLow       OpcuaNodeIdServicesVariableLimit = 24772
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLowLow    OpcuaNodeIdServicesVariableLimit = 24773
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighDeadband  OpcuaNodeIdServicesVariableLimit = 24774
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighDeadband      OpcuaNodeIdServicesVariableLimit = 24775
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowDeadband       OpcuaNodeIdServicesVariableLimit = 24776
	OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowDeadband    OpcuaNodeIdServicesVariableLimit = 24777
)

var OpcuaNodeIdServicesVariableLimitValues []OpcuaNodeIdServicesVariableLimit

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableLimitValues = []OpcuaNodeIdServicesVariableLimit{
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighLimit,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighLimit,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLimit,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowLimit,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighHighLimit,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighLimit,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLimit,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLowLimit,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHighHigh,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHigh,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLow,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLowLow,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighDeadband,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighDeadband,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowDeadband,
		OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowDeadband,
	}
}

func OpcuaNodeIdServicesVariableLimitByValue(value int32) (enum OpcuaNodeIdServicesVariableLimit, ok bool) {
	switch value {
	case 11124:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighLimit, true
	case 11125:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighLimit, true
	case 11126:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLimit, true
	case 11127:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowLimit, true
	case 16572:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighHighLimit, true
	case 16573:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighLimit, true
	case 16574:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLimit, true
	case 16575:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLowLimit, true
	case 24770:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHighHigh, true
	case 24771:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHigh, true
	case 24772:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLow, true
	case 24773:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLowLow, true
	case 24774:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighDeadband, true
	case 24775:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighDeadband, true
	case 24776:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowDeadband, true
	case 24777:
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowDeadband, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableLimitByName(value string) (enum OpcuaNodeIdServicesVariableLimit, ok bool) {
	switch value {
	case "LimitAlarmType_HighHighLimit":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighLimit, true
	case "LimitAlarmType_HighLimit":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighLimit, true
	case "LimitAlarmType_LowLimit":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLimit, true
	case "LimitAlarmType_LowLowLimit":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowLimit, true
	case "LimitAlarmType_BaseHighHighLimit":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighHighLimit, true
	case "LimitAlarmType_BaseHighLimit":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighLimit, true
	case "LimitAlarmType_BaseLowLimit":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLimit, true
	case "LimitAlarmType_BaseLowLowLimit":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLowLimit, true
	case "LimitAlarmType_SeverityHighHigh":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHighHigh, true
	case "LimitAlarmType_SeverityHigh":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHigh, true
	case "LimitAlarmType_SeverityLow":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLow, true
	case "LimitAlarmType_SeverityLowLow":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLowLow, true
	case "LimitAlarmType_HighHighDeadband":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighDeadband, true
	case "LimitAlarmType_HighDeadband":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighDeadband, true
	case "LimitAlarmType_LowDeadband":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowDeadband, true
	case "LimitAlarmType_LowLowDeadband":
		return OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowDeadband, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableLimitKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableLimitValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableLimit(structType any) OpcuaNodeIdServicesVariableLimit {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableLimit {
		if sOpcuaNodeIdServicesVariableLimit, ok := typ.(OpcuaNodeIdServicesVariableLimit); ok {
			return sOpcuaNodeIdServicesVariableLimit
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableLimit) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableLimitParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableLimit, error) {
	return OpcuaNodeIdServicesVariableLimitParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableLimit, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableLimit", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableLimit")
	}
	if enum, ok := OpcuaNodeIdServicesVariableLimitByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableLimit")
		return OpcuaNodeIdServicesVariableLimit(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableLimit", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableLimit) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighLimit:
		return "LimitAlarmType_HighHighLimit"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighLimit:
		return "LimitAlarmType_HighLimit"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLimit:
		return "LimitAlarmType_LowLimit"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowLimit:
		return "LimitAlarmType_LowLowLimit"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighHighLimit:
		return "LimitAlarmType_BaseHighHighLimit"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseHighLimit:
		return "LimitAlarmType_BaseHighLimit"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLimit:
		return "LimitAlarmType_BaseLowLimit"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_BaseLowLowLimit:
		return "LimitAlarmType_BaseLowLowLimit"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHighHigh:
		return "LimitAlarmType_SeverityHighHigh"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityHigh:
		return "LimitAlarmType_SeverityHigh"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLow:
		return "LimitAlarmType_SeverityLow"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_SeverityLowLow:
		return "LimitAlarmType_SeverityLowLow"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighHighDeadband:
		return "LimitAlarmType_HighHighDeadband"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_HighDeadband:
		return "LimitAlarmType_HighDeadband"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowDeadband:
		return "LimitAlarmType_LowDeadband"
	case OpcuaNodeIdServicesVariableLimit_LimitAlarmType_LowLowDeadband:
		return "LimitAlarmType_LowLowDeadband"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableLimit) String() string {
	return e.PLC4XEnumName()
}
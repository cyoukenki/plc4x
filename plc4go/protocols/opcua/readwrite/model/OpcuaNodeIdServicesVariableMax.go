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

// OpcuaNodeIdServicesVariableMax is an enum
type OpcuaNodeIdServicesVariableMax int32

type IOpcuaNodeIdServicesVariableMax interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableMax_MaxStringLength     OpcuaNodeIdServicesVariableMax = 11498
	OpcuaNodeIdServicesVariableMax_MaxArrayLength      OpcuaNodeIdServicesVariableMax = 11512
	OpcuaNodeIdServicesVariableMax_MaxByteStringLength OpcuaNodeIdServicesVariableMax = 12908
	OpcuaNodeIdServicesVariableMax_MaxCharacters       OpcuaNodeIdServicesVariableMax = 15002
)

var OpcuaNodeIdServicesVariableMaxValues []OpcuaNodeIdServicesVariableMax

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableMaxValues = []OpcuaNodeIdServicesVariableMax{
		OpcuaNodeIdServicesVariableMax_MaxStringLength,
		OpcuaNodeIdServicesVariableMax_MaxArrayLength,
		OpcuaNodeIdServicesVariableMax_MaxByteStringLength,
		OpcuaNodeIdServicesVariableMax_MaxCharacters,
	}
}

func OpcuaNodeIdServicesVariableMaxByValue(value int32) (enum OpcuaNodeIdServicesVariableMax, ok bool) {
	switch value {
	case 11498:
		return OpcuaNodeIdServicesVariableMax_MaxStringLength, true
	case 11512:
		return OpcuaNodeIdServicesVariableMax_MaxArrayLength, true
	case 12908:
		return OpcuaNodeIdServicesVariableMax_MaxByteStringLength, true
	case 15002:
		return OpcuaNodeIdServicesVariableMax_MaxCharacters, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableMaxByName(value string) (enum OpcuaNodeIdServicesVariableMax, ok bool) {
	switch value {
	case "MaxStringLength":
		return OpcuaNodeIdServicesVariableMax_MaxStringLength, true
	case "MaxArrayLength":
		return OpcuaNodeIdServicesVariableMax_MaxArrayLength, true
	case "MaxByteStringLength":
		return OpcuaNodeIdServicesVariableMax_MaxByteStringLength, true
	case "MaxCharacters":
		return OpcuaNodeIdServicesVariableMax_MaxCharacters, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableMaxKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableMaxValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableMax(structType any) OpcuaNodeIdServicesVariableMax {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableMax {
		if sOpcuaNodeIdServicesVariableMax, ok := typ.(OpcuaNodeIdServicesVariableMax); ok {
			return sOpcuaNodeIdServicesVariableMax
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableMax) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableMax) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableMaxParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableMax, error) {
	return OpcuaNodeIdServicesVariableMaxParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableMaxParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableMax, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableMax", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableMax")
	}
	if enum, ok := OpcuaNodeIdServicesVariableMaxByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableMax")
		return OpcuaNodeIdServicesVariableMax(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableMax) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableMax) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableMax", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableMax) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableMax_MaxStringLength:
		return "MaxStringLength"
	case OpcuaNodeIdServicesVariableMax_MaxArrayLength:
		return "MaxArrayLength"
	case OpcuaNodeIdServicesVariableMax_MaxByteStringLength:
		return "MaxByteStringLength"
	case OpcuaNodeIdServicesVariableMax_MaxCharacters:
		return "MaxCharacters"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableMax) String() string {
	return e.PLC4XEnumName()
}
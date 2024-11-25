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

// OpcuaNodeIdServicesVariableFind is an enum
type OpcuaNodeIdServicesVariableFind int32

type IOpcuaNodeIdServicesVariableFind interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableFind_FindAliasMethodType_InputArguments  OpcuaNodeIdServicesVariableFind = 23466
	OpcuaNodeIdServicesVariableFind_FindAliasMethodType_OutputArguments OpcuaNodeIdServicesVariableFind = 23467
)

var OpcuaNodeIdServicesVariableFindValues []OpcuaNodeIdServicesVariableFind

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableFindValues = []OpcuaNodeIdServicesVariableFind{
		OpcuaNodeIdServicesVariableFind_FindAliasMethodType_InputArguments,
		OpcuaNodeIdServicesVariableFind_FindAliasMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableFindByValue(value int32) (enum OpcuaNodeIdServicesVariableFind, ok bool) {
	switch value {
	case 23466:
		return OpcuaNodeIdServicesVariableFind_FindAliasMethodType_InputArguments, true
	case 23467:
		return OpcuaNodeIdServicesVariableFind_FindAliasMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableFindByName(value string) (enum OpcuaNodeIdServicesVariableFind, ok bool) {
	switch value {
	case "FindAliasMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableFind_FindAliasMethodType_InputArguments, true
	case "FindAliasMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableFind_FindAliasMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableFindKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableFindValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableFind(structType any) OpcuaNodeIdServicesVariableFind {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableFind {
		if sOpcuaNodeIdServicesVariableFind, ok := typ.(OpcuaNodeIdServicesVariableFind); ok {
			return sOpcuaNodeIdServicesVariableFind
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableFind) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableFind) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableFindParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableFind, error) {
	return OpcuaNodeIdServicesVariableFindParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableFindParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableFind, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableFind", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableFind")
	}
	if enum, ok := OpcuaNodeIdServicesVariableFindByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableFind")
		return OpcuaNodeIdServicesVariableFind(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableFind) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableFind) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableFind", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableFind) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableFind_FindAliasMethodType_InputArguments:
		return "FindAliasMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableFind_FindAliasMethodType_OutputArguments:
		return "FindAliasMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableFind) String() string {
	return e.PLC4XEnumName()
}
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

// OpcuaNodeIdServicesVariableUpdate is an enum
type OpcuaNodeIdServicesVariableUpdate int32

type IOpcuaNodeIdServicesVariableUpdate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_InputArguments  OpcuaNodeIdServicesVariableUpdate = 12579
	OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_OutputArguments OpcuaNodeIdServicesVariableUpdate = 12580
)

var OpcuaNodeIdServicesVariableUpdateValues []OpcuaNodeIdServicesVariableUpdate

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableUpdateValues = []OpcuaNodeIdServicesVariableUpdate{
		OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_InputArguments,
		OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableUpdateByValue(value int32) (enum OpcuaNodeIdServicesVariableUpdate, ok bool) {
	switch value {
	case 12579:
		return OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_InputArguments, true
	case 12580:
		return OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableUpdateByName(value string) (enum OpcuaNodeIdServicesVariableUpdate, ok bool) {
	switch value {
	case "UpdateCertificateMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_InputArguments, true
	case "UpdateCertificateMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableUpdateKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableUpdateValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableUpdate(structType any) OpcuaNodeIdServicesVariableUpdate {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableUpdate {
		if sOpcuaNodeIdServicesVariableUpdate, ok := typ.(OpcuaNodeIdServicesVariableUpdate); ok {
			return sOpcuaNodeIdServicesVariableUpdate
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableUpdate) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableUpdate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableUpdateParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableUpdate, error) {
	return OpcuaNodeIdServicesVariableUpdateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableUpdateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableUpdate, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableUpdate", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableUpdate")
	}
	if enum, ok := OpcuaNodeIdServicesVariableUpdateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableUpdate")
		return OpcuaNodeIdServicesVariableUpdate(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableUpdate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableUpdate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableUpdate", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableUpdate) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_InputArguments:
		return "UpdateCertificateMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableUpdate_UpdateCertificateMethodType_OutputArguments:
		return "UpdateCertificateMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableUpdate) String() string {
	return e.PLC4XEnumName()
}

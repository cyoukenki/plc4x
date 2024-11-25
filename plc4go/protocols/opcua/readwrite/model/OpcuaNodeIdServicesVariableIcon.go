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

// OpcuaNodeIdServicesVariableIcon is an enum
type OpcuaNodeIdServicesVariableIcon int32

type IOpcuaNodeIdServicesVariableIcon interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableIcon_Icon OpcuaNodeIdServicesVariableIcon = 3067
)

var OpcuaNodeIdServicesVariableIconValues []OpcuaNodeIdServicesVariableIcon

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableIconValues = []OpcuaNodeIdServicesVariableIcon{
		OpcuaNodeIdServicesVariableIcon_Icon,
	}
}

func OpcuaNodeIdServicesVariableIconByValue(value int32) (enum OpcuaNodeIdServicesVariableIcon, ok bool) {
	switch value {
	case 3067:
		return OpcuaNodeIdServicesVariableIcon_Icon, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableIconByName(value string) (enum OpcuaNodeIdServicesVariableIcon, ok bool) {
	switch value {
	case "Icon":
		return OpcuaNodeIdServicesVariableIcon_Icon, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableIconKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableIconValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableIcon(structType any) OpcuaNodeIdServicesVariableIcon {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableIcon {
		if sOpcuaNodeIdServicesVariableIcon, ok := typ.(OpcuaNodeIdServicesVariableIcon); ok {
			return sOpcuaNodeIdServicesVariableIcon
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableIcon) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableIcon) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableIconParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableIcon, error) {
	return OpcuaNodeIdServicesVariableIconParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableIconParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableIcon, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableIcon", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableIcon")
	}
	if enum, ok := OpcuaNodeIdServicesVariableIconByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableIcon")
		return OpcuaNodeIdServicesVariableIcon(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableIcon) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableIcon) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableIcon", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableIcon) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableIcon_Icon:
		return "Icon"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableIcon) String() string {
	return e.PLC4XEnumName()
}
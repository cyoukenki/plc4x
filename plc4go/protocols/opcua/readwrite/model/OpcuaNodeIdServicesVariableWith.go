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

// OpcuaNodeIdServicesVariableWith is an enum
type OpcuaNodeIdServicesVariableWith int32

type IOpcuaNodeIdServicesVariableWith interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableWith_WithCommentMethodType_InputArguments OpcuaNodeIdServicesVariableWith = 24327
)

var OpcuaNodeIdServicesVariableWithValues []OpcuaNodeIdServicesVariableWith

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableWithValues = []OpcuaNodeIdServicesVariableWith{
		OpcuaNodeIdServicesVariableWith_WithCommentMethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableWithByValue(value int32) (enum OpcuaNodeIdServicesVariableWith, ok bool) {
	switch value {
	case 24327:
		return OpcuaNodeIdServicesVariableWith_WithCommentMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableWithByName(value string) (enum OpcuaNodeIdServicesVariableWith, ok bool) {
	switch value {
	case "WithCommentMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableWith_WithCommentMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableWithKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableWithValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableWith(structType any) OpcuaNodeIdServicesVariableWith {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableWith {
		if sOpcuaNodeIdServicesVariableWith, ok := typ.(OpcuaNodeIdServicesVariableWith); ok {
			return sOpcuaNodeIdServicesVariableWith
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableWith) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableWith) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableWithParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableWith, error) {
	return OpcuaNodeIdServicesVariableWithParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableWithParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableWith, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableWith", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableWith")
	}
	if enum, ok := OpcuaNodeIdServicesVariableWithByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableWith")
		return OpcuaNodeIdServicesVariableWith(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableWith) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableWith) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableWith", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableWith) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableWith_WithCommentMethodType_InputArguments:
		return "WithCommentMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableWith) String() string {
	return e.PLC4XEnumName()
}
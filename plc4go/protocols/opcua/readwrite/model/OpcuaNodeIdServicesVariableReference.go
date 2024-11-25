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

// OpcuaNodeIdServicesVariableReference is an enum
type OpcuaNodeIdServicesVariableReference int32

type IOpcuaNodeIdServicesVariableReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableReference_ReferenceDescriptionVariableType_ReferenceRefinement OpcuaNodeIdServicesVariableReference = 32658
)

var OpcuaNodeIdServicesVariableReferenceValues []OpcuaNodeIdServicesVariableReference

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableReferenceValues = []OpcuaNodeIdServicesVariableReference{
		OpcuaNodeIdServicesVariableReference_ReferenceDescriptionVariableType_ReferenceRefinement,
	}
}

func OpcuaNodeIdServicesVariableReferenceByValue(value int32) (enum OpcuaNodeIdServicesVariableReference, ok bool) {
	switch value {
	case 32658:
		return OpcuaNodeIdServicesVariableReference_ReferenceDescriptionVariableType_ReferenceRefinement, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableReferenceByName(value string) (enum OpcuaNodeIdServicesVariableReference, ok bool) {
	switch value {
	case "ReferenceDescriptionVariableType_ReferenceRefinement":
		return OpcuaNodeIdServicesVariableReference_ReferenceDescriptionVariableType_ReferenceRefinement, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableReferenceKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableReferenceValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableReference(structType any) OpcuaNodeIdServicesVariableReference {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableReference {
		if sOpcuaNodeIdServicesVariableReference, ok := typ.(OpcuaNodeIdServicesVariableReference); ok {
			return sOpcuaNodeIdServicesVariableReference
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableReference) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableReferenceParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableReference, error) {
	return OpcuaNodeIdServicesVariableReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableReference, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableReference", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableReference")
	}
	if enum, ok := OpcuaNodeIdServicesVariableReferenceByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableReference")
		return OpcuaNodeIdServicesVariableReference(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableReference", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableReference) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableReference_ReferenceDescriptionVariableType_ReferenceRefinement:
		return "ReferenceDescriptionVariableType_ReferenceRefinement"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableReference) String() string {
	return e.PLC4XEnumName()
}
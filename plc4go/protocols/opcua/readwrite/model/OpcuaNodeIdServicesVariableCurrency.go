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

// OpcuaNodeIdServicesVariableCurrency is an enum
type OpcuaNodeIdServicesVariableCurrency int32

type IOpcuaNodeIdServicesVariableCurrency interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableCurrency_CurrencyUnit OpcuaNodeIdServicesVariableCurrency = 23501
)

var OpcuaNodeIdServicesVariableCurrencyValues []OpcuaNodeIdServicesVariableCurrency

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableCurrencyValues = []OpcuaNodeIdServicesVariableCurrency{
		OpcuaNodeIdServicesVariableCurrency_CurrencyUnit,
	}
}

func OpcuaNodeIdServicesVariableCurrencyByValue(value int32) (enum OpcuaNodeIdServicesVariableCurrency, ok bool) {
	switch value {
	case 23501:
		return OpcuaNodeIdServicesVariableCurrency_CurrencyUnit, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableCurrencyByName(value string) (enum OpcuaNodeIdServicesVariableCurrency, ok bool) {
	switch value {
	case "CurrencyUnit":
		return OpcuaNodeIdServicesVariableCurrency_CurrencyUnit, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableCurrencyKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableCurrencyValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableCurrency(structType any) OpcuaNodeIdServicesVariableCurrency {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableCurrency {
		if sOpcuaNodeIdServicesVariableCurrency, ok := typ.(OpcuaNodeIdServicesVariableCurrency); ok {
			return sOpcuaNodeIdServicesVariableCurrency
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableCurrency) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableCurrency) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableCurrencyParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableCurrency, error) {
	return OpcuaNodeIdServicesVariableCurrencyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableCurrencyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableCurrency, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableCurrency", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableCurrency")
	}
	if enum, ok := OpcuaNodeIdServicesVariableCurrencyByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableCurrency")
		return OpcuaNodeIdServicesVariableCurrency(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableCurrency) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableCurrency) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableCurrency", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableCurrency) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableCurrency_CurrencyUnit:
		return "CurrencyUnit"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableCurrency) String() string {
	return e.PLC4XEnumName()
}

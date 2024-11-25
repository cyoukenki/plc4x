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

// HVACHumidityModeAndFlagsMode is an enum
type HVACHumidityModeAndFlagsMode uint8

type IHVACHumidityModeAndFlagsMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	HVACHumidityModeAndFlagsMode_OFF              HVACHumidityModeAndFlagsMode = 0x0
	HVACHumidityModeAndFlagsMode_HUMIDIFY_ONLY    HVACHumidityModeAndFlagsMode = 0x1
	HVACHumidityModeAndFlagsMode_DEHUMIDIFY_ONLY  HVACHumidityModeAndFlagsMode = 0x2
	HVACHumidityModeAndFlagsMode_HUMIDITY_CONTROL HVACHumidityModeAndFlagsMode = 0x3
)

var HVACHumidityModeAndFlagsModeValues []HVACHumidityModeAndFlagsMode

func init() {
	_ = errors.New
	HVACHumidityModeAndFlagsModeValues = []HVACHumidityModeAndFlagsMode{
		HVACHumidityModeAndFlagsMode_OFF,
		HVACHumidityModeAndFlagsMode_HUMIDIFY_ONLY,
		HVACHumidityModeAndFlagsMode_DEHUMIDIFY_ONLY,
		HVACHumidityModeAndFlagsMode_HUMIDITY_CONTROL,
	}
}

func HVACHumidityModeAndFlagsModeByValue(value uint8) (enum HVACHumidityModeAndFlagsMode, ok bool) {
	switch value {
	case 0x0:
		return HVACHumidityModeAndFlagsMode_OFF, true
	case 0x1:
		return HVACHumidityModeAndFlagsMode_HUMIDIFY_ONLY, true
	case 0x2:
		return HVACHumidityModeAndFlagsMode_DEHUMIDIFY_ONLY, true
	case 0x3:
		return HVACHumidityModeAndFlagsMode_HUMIDITY_CONTROL, true
	}
	return 0, false
}

func HVACHumidityModeAndFlagsModeByName(value string) (enum HVACHumidityModeAndFlagsMode, ok bool) {
	switch value {
	case "OFF":
		return HVACHumidityModeAndFlagsMode_OFF, true
	case "HUMIDIFY_ONLY":
		return HVACHumidityModeAndFlagsMode_HUMIDIFY_ONLY, true
	case "DEHUMIDIFY_ONLY":
		return HVACHumidityModeAndFlagsMode_DEHUMIDIFY_ONLY, true
	case "HUMIDITY_CONTROL":
		return HVACHumidityModeAndFlagsMode_HUMIDITY_CONTROL, true
	}
	return 0, false
}

func HVACHumidityModeAndFlagsModeKnows(value uint8) bool {
	for _, typeValue := range HVACHumidityModeAndFlagsModeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastHVACHumidityModeAndFlagsMode(structType any) HVACHumidityModeAndFlagsMode {
	castFunc := func(typ any) HVACHumidityModeAndFlagsMode {
		if sHVACHumidityModeAndFlagsMode, ok := typ.(HVACHumidityModeAndFlagsMode); ok {
			return sHVACHumidityModeAndFlagsMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m HVACHumidityModeAndFlagsMode) GetLengthInBits(ctx context.Context) uint16 {
	return 3
}

func (m HVACHumidityModeAndFlagsMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACHumidityModeAndFlagsModeParse(ctx context.Context, theBytes []byte) (HVACHumidityModeAndFlagsMode, error) {
	return HVACHumidityModeAndFlagsModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACHumidityModeAndFlagsModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidityModeAndFlagsMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("HVACHumidityModeAndFlagsMode", 3)
	if err != nil {
		return 0, errors.Wrap(err, "error reading HVACHumidityModeAndFlagsMode")
	}
	if enum, ok := HVACHumidityModeAndFlagsModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for HVACHumidityModeAndFlagsMode")
		return HVACHumidityModeAndFlagsMode(val), nil
	} else {
		return enum, nil
	}
}

func (e HVACHumidityModeAndFlagsMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e HVACHumidityModeAndFlagsMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("HVACHumidityModeAndFlagsMode", 3, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e HVACHumidityModeAndFlagsMode) PLC4XEnumName() string {
	switch e {
	case HVACHumidityModeAndFlagsMode_OFF:
		return "OFF"
	case HVACHumidityModeAndFlagsMode_HUMIDIFY_ONLY:
		return "HUMIDIFY_ONLY"
	case HVACHumidityModeAndFlagsMode_DEHUMIDIFY_ONLY:
		return "DEHUMIDIFY_ONLY"
	case HVACHumidityModeAndFlagsMode_HUMIDITY_CONTROL:
		return "HUMIDITY_CONTROL"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e HVACHumidityModeAndFlagsMode) String() string {
	return e.PLC4XEnumName()
}
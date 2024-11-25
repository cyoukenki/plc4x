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

// ModeTransitionType is an enum
type ModeTransitionType uint8

type IModeTransitionType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ModeTransitionType_STOP         ModeTransitionType = 0x00
	ModeTransitionType_WARM_RESTART ModeTransitionType = 0x01
	ModeTransitionType_RUN          ModeTransitionType = 0x02
	ModeTransitionType_HOT_RESTART  ModeTransitionType = 0x03
	ModeTransitionType_HOLD         ModeTransitionType = 0x04
	ModeTransitionType_COLD_RESTART ModeTransitionType = 0x06
	ModeTransitionType_RUN_R        ModeTransitionType = 0x09
	ModeTransitionType_LINK_UP      ModeTransitionType = 0x11
	ModeTransitionType_UPDATE       ModeTransitionType = 0x12
)

var ModeTransitionTypeValues []ModeTransitionType

func init() {
	_ = errors.New
	ModeTransitionTypeValues = []ModeTransitionType{
		ModeTransitionType_STOP,
		ModeTransitionType_WARM_RESTART,
		ModeTransitionType_RUN,
		ModeTransitionType_HOT_RESTART,
		ModeTransitionType_HOLD,
		ModeTransitionType_COLD_RESTART,
		ModeTransitionType_RUN_R,
		ModeTransitionType_LINK_UP,
		ModeTransitionType_UPDATE,
	}
}

func ModeTransitionTypeByValue(value uint8) (enum ModeTransitionType, ok bool) {
	switch value {
	case 0x00:
		return ModeTransitionType_STOP, true
	case 0x01:
		return ModeTransitionType_WARM_RESTART, true
	case 0x02:
		return ModeTransitionType_RUN, true
	case 0x03:
		return ModeTransitionType_HOT_RESTART, true
	case 0x04:
		return ModeTransitionType_HOLD, true
	case 0x06:
		return ModeTransitionType_COLD_RESTART, true
	case 0x09:
		return ModeTransitionType_RUN_R, true
	case 0x11:
		return ModeTransitionType_LINK_UP, true
	case 0x12:
		return ModeTransitionType_UPDATE, true
	}
	return 0, false
}

func ModeTransitionTypeByName(value string) (enum ModeTransitionType, ok bool) {
	switch value {
	case "STOP":
		return ModeTransitionType_STOP, true
	case "WARM_RESTART":
		return ModeTransitionType_WARM_RESTART, true
	case "RUN":
		return ModeTransitionType_RUN, true
	case "HOT_RESTART":
		return ModeTransitionType_HOT_RESTART, true
	case "HOLD":
		return ModeTransitionType_HOLD, true
	case "COLD_RESTART":
		return ModeTransitionType_COLD_RESTART, true
	case "RUN_R":
		return ModeTransitionType_RUN_R, true
	case "LINK_UP":
		return ModeTransitionType_LINK_UP, true
	case "UPDATE":
		return ModeTransitionType_UPDATE, true
	}
	return 0, false
}

func ModeTransitionTypeKnows(value uint8) bool {
	for _, typeValue := range ModeTransitionTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastModeTransitionType(structType any) ModeTransitionType {
	castFunc := func(typ any) ModeTransitionType {
		if sModeTransitionType, ok := typ.(ModeTransitionType); ok {
			return sModeTransitionType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ModeTransitionType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m ModeTransitionType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModeTransitionTypeParse(ctx context.Context, theBytes []byte) (ModeTransitionType, error) {
	return ModeTransitionTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModeTransitionTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModeTransitionType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("ModeTransitionType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ModeTransitionType")
	}
	if enum, ok := ModeTransitionTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ModeTransitionType")
		return ModeTransitionType(val), nil
	} else {
		return enum, nil
	}
}

func (e ModeTransitionType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ModeTransitionType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("ModeTransitionType", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ModeTransitionType) PLC4XEnumName() string {
	switch e {
	case ModeTransitionType_STOP:
		return "STOP"
	case ModeTransitionType_WARM_RESTART:
		return "WARM_RESTART"
	case ModeTransitionType_RUN:
		return "RUN"
	case ModeTransitionType_HOT_RESTART:
		return "HOT_RESTART"
	case ModeTransitionType_HOLD:
		return "HOLD"
	case ModeTransitionType_COLD_RESTART:
		return "COLD_RESTART"
	case ModeTransitionType_RUN_R:
		return "RUN_R"
	case ModeTransitionType_LINK_UP:
		return "LINK_UP"
	case ModeTransitionType_UPDATE:
		return "UPDATE"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e ModeTransitionType) String() string {
	return e.PLC4XEnumName()
}
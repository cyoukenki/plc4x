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

// TriggerControlLabelFlavour is an enum
type TriggerControlLabelFlavour uint8

type ITriggerControlLabelFlavour interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	TriggerControlLabelFlavour_FLAVOUR_0 TriggerControlLabelFlavour = 0
	TriggerControlLabelFlavour_FLAVOUR_1 TriggerControlLabelFlavour = 1
	TriggerControlLabelFlavour_FLAVOUR_2 TriggerControlLabelFlavour = 2
	TriggerControlLabelFlavour_FLAVOUR_3 TriggerControlLabelFlavour = 3
)

var TriggerControlLabelFlavourValues []TriggerControlLabelFlavour

func init() {
	_ = errors.New
	TriggerControlLabelFlavourValues = []TriggerControlLabelFlavour{
		TriggerControlLabelFlavour_FLAVOUR_0,
		TriggerControlLabelFlavour_FLAVOUR_1,
		TriggerControlLabelFlavour_FLAVOUR_2,
		TriggerControlLabelFlavour_FLAVOUR_3,
	}
}

func TriggerControlLabelFlavourByValue(value uint8) (enum TriggerControlLabelFlavour, ok bool) {
	switch value {
	case 0:
		return TriggerControlLabelFlavour_FLAVOUR_0, true
	case 1:
		return TriggerControlLabelFlavour_FLAVOUR_1, true
	case 2:
		return TriggerControlLabelFlavour_FLAVOUR_2, true
	case 3:
		return TriggerControlLabelFlavour_FLAVOUR_3, true
	}
	return 0, false
}

func TriggerControlLabelFlavourByName(value string) (enum TriggerControlLabelFlavour, ok bool) {
	switch value {
	case "FLAVOUR_0":
		return TriggerControlLabelFlavour_FLAVOUR_0, true
	case "FLAVOUR_1":
		return TriggerControlLabelFlavour_FLAVOUR_1, true
	case "FLAVOUR_2":
		return TriggerControlLabelFlavour_FLAVOUR_2, true
	case "FLAVOUR_3":
		return TriggerControlLabelFlavour_FLAVOUR_3, true
	}
	return 0, false
}

func TriggerControlLabelFlavourKnows(value uint8) bool {
	for _, typeValue := range TriggerControlLabelFlavourValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTriggerControlLabelFlavour(structType any) TriggerControlLabelFlavour {
	castFunc := func(typ any) TriggerControlLabelFlavour {
		if sTriggerControlLabelFlavour, ok := typ.(TriggerControlLabelFlavour); ok {
			return sTriggerControlLabelFlavour
		}
		return 0
	}
	return castFunc(structType)
}

func (m TriggerControlLabelFlavour) GetLengthInBits(ctx context.Context) uint16 {
	return 2
}

func (m TriggerControlLabelFlavour) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TriggerControlLabelFlavourParse(ctx context.Context, theBytes []byte) (TriggerControlLabelFlavour, error) {
	return TriggerControlLabelFlavourParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TriggerControlLabelFlavourParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlLabelFlavour, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("TriggerControlLabelFlavour", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TriggerControlLabelFlavour")
	}
	if enum, ok := TriggerControlLabelFlavourByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TriggerControlLabelFlavour")
		return TriggerControlLabelFlavour(val), nil
	} else {
		return enum, nil
	}
}

func (e TriggerControlLabelFlavour) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TriggerControlLabelFlavour) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("TriggerControlLabelFlavour", 2, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TriggerControlLabelFlavour) PLC4XEnumName() string {
	switch e {
	case TriggerControlLabelFlavour_FLAVOUR_0:
		return "FLAVOUR_0"
	case TriggerControlLabelFlavour_FLAVOUR_1:
		return "FLAVOUR_1"
	case TriggerControlLabelFlavour_FLAVOUR_2:
		return "FLAVOUR_2"
	case TriggerControlLabelFlavour_FLAVOUR_3:
		return "FLAVOUR_3"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e TriggerControlLabelFlavour) String() string {
	return e.PLC4XEnumName()
}

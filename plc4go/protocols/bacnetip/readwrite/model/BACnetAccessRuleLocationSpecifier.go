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

// BACnetAccessRuleLocationSpecifier is an enum
type BACnetAccessRuleLocationSpecifier uint8

type IBACnetAccessRuleLocationSpecifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetAccessRuleLocationSpecifier_SPECIFIED BACnetAccessRuleLocationSpecifier = 0
	BACnetAccessRuleLocationSpecifier_ALL       BACnetAccessRuleLocationSpecifier = 1
)

var BACnetAccessRuleLocationSpecifierValues []BACnetAccessRuleLocationSpecifier

func init() {
	_ = errors.New
	BACnetAccessRuleLocationSpecifierValues = []BACnetAccessRuleLocationSpecifier{
		BACnetAccessRuleLocationSpecifier_SPECIFIED,
		BACnetAccessRuleLocationSpecifier_ALL,
	}
}

func BACnetAccessRuleLocationSpecifierByValue(value uint8) (enum BACnetAccessRuleLocationSpecifier, ok bool) {
	switch value {
	case 0:
		return BACnetAccessRuleLocationSpecifier_SPECIFIED, true
	case 1:
		return BACnetAccessRuleLocationSpecifier_ALL, true
	}
	return 0, false
}

func BACnetAccessRuleLocationSpecifierByName(value string) (enum BACnetAccessRuleLocationSpecifier, ok bool) {
	switch value {
	case "SPECIFIED":
		return BACnetAccessRuleLocationSpecifier_SPECIFIED, true
	case "ALL":
		return BACnetAccessRuleLocationSpecifier_ALL, true
	}
	return 0, false
}

func BACnetAccessRuleLocationSpecifierKnows(value uint8) bool {
	for _, typeValue := range BACnetAccessRuleLocationSpecifierValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessRuleLocationSpecifier(structType any) BACnetAccessRuleLocationSpecifier {
	castFunc := func(typ any) BACnetAccessRuleLocationSpecifier {
		if sBACnetAccessRuleLocationSpecifier, ok := typ.(BACnetAccessRuleLocationSpecifier); ok {
			return sBACnetAccessRuleLocationSpecifier
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessRuleLocationSpecifier) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetAccessRuleLocationSpecifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessRuleLocationSpecifierParse(ctx context.Context, theBytes []byte) (BACnetAccessRuleLocationSpecifier, error) {
	return BACnetAccessRuleLocationSpecifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessRuleLocationSpecifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessRuleLocationSpecifier, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetAccessRuleLocationSpecifier", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccessRuleLocationSpecifier")
	}
	if enum, ok := BACnetAccessRuleLocationSpecifierByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetAccessRuleLocationSpecifier")
		return BACnetAccessRuleLocationSpecifier(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccessRuleLocationSpecifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAccessRuleLocationSpecifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetAccessRuleLocationSpecifier", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessRuleLocationSpecifier) PLC4XEnumName() string {
	switch e {
	case BACnetAccessRuleLocationSpecifier_SPECIFIED:
		return "SPECIFIED"
	case BACnetAccessRuleLocationSpecifier_ALL:
		return "ALL"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetAccessRuleLocationSpecifier) String() string {
	return e.PLC4XEnumName()
}
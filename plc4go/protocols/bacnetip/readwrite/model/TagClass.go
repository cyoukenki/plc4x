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

// TagClass is an enum
type TagClass uint8

type ITagClass interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	TagClass_APPLICATION_TAGS      TagClass = 0x0
	TagClass_CONTEXT_SPECIFIC_TAGS TagClass = 0x1
)

var TagClassValues []TagClass

func init() {
	_ = errors.New
	TagClassValues = []TagClass{
		TagClass_APPLICATION_TAGS,
		TagClass_CONTEXT_SPECIFIC_TAGS,
	}
}

func TagClassByValue(value uint8) (enum TagClass, ok bool) {
	switch value {
	case 0x0:
		return TagClass_APPLICATION_TAGS, true
	case 0x1:
		return TagClass_CONTEXT_SPECIFIC_TAGS, true
	}
	return 0, false
}

func TagClassByName(value string) (enum TagClass, ok bool) {
	switch value {
	case "APPLICATION_TAGS":
		return TagClass_APPLICATION_TAGS, true
	case "CONTEXT_SPECIFIC_TAGS":
		return TagClass_CONTEXT_SPECIFIC_TAGS, true
	}
	return 0, false
}

func TagClassKnows(value uint8) bool {
	for _, typeValue := range TagClassValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTagClass(structType any) TagClass {
	castFunc := func(typ any) TagClass {
		if sTagClass, ok := typ.(TagClass); ok {
			return sTagClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m TagClass) GetLengthInBits(ctx context.Context) uint16 {
	return 1
}

func (m TagClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TagClassParse(ctx context.Context, theBytes []byte) (TagClass, error) {
	return TagClassParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TagClassParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TagClass, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("TagClass", 1)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TagClass")
	}
	if enum, ok := TagClassByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TagClass")
		return TagClass(val), nil
	} else {
		return enum, nil
	}
}

func (e TagClass) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TagClass) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("TagClass", 1, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TagClass) PLC4XEnumName() string {
	switch e {
	case TagClass_APPLICATION_TAGS:
		return "APPLICATION_TAGS"
	case TagClass_CONTEXT_SPECIFIC_TAGS:
		return "CONTEXT_SPECIFIC_TAGS"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e TagClass) String() string {
	return e.PLC4XEnumName()
}
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

// CIPAttributes is the corresponding interface of CIPAttributes
type CIPAttributes interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetClassId returns ClassId (property field)
	GetClassId() []uint16
	// GetNumberAvailable returns NumberAvailable (property field)
	GetNumberAvailable() *uint16
	// GetNumberActive returns NumberActive (property field)
	GetNumberActive() *uint16
	// GetData returns Data (property field)
	GetData() []byte
}

// CIPAttributesExactly can be used when we want exactly this type and not a type which fulfills CIPAttributes.
// This is useful for switch cases.
type CIPAttributesExactly interface {
	CIPAttributes
	isCIPAttributes() bool
}

// _CIPAttributes is the data-structure of this message
type _CIPAttributes struct {
	ClassId         []uint16
	NumberAvailable *uint16
	NumberActive    *uint16
	Data            []byte

	// Arguments.
	PacketLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPAttributes) GetClassId() []uint16 {
	return m.ClassId
}

func (m *_CIPAttributes) GetNumberAvailable() *uint16 {
	return m.NumberAvailable
}

func (m *_CIPAttributes) GetNumberActive() *uint16 {
	return m.NumberActive
}

func (m *_CIPAttributes) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPAttributes factory function for _CIPAttributes
func NewCIPAttributes(classId []uint16, numberAvailable *uint16, numberActive *uint16, data []byte, packetLength uint16) *_CIPAttributes {
	return &_CIPAttributes{ClassId: classId, NumberAvailable: numberAvailable, NumberActive: numberActive, Data: data, PacketLength: packetLength}
}

// Deprecated: use the interface for direct cast
func CastCIPAttributes(structType any) CIPAttributes {
	if casted, ok := structType.(CIPAttributes); ok {
		return casted
	}
	if casted, ok := structType.(*CIPAttributes); ok {
		return *casted
	}
	return nil
}

func (m *_CIPAttributes) GetTypeName() string {
	return "CIPAttributes"
}

func (m *_CIPAttributes) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (numberOfClasses)
	lengthInBits += 16

	// Array field
	if len(m.ClassId) > 0 {
		lengthInBits += 16 * uint16(len(m.ClassId))
	}

	// Optional Field (numberAvailable)
	if m.NumberAvailable != nil {
		lengthInBits += 16
	}

	// Optional Field (numberActive)
	if m.NumberActive != nil {
		lengthInBits += 16
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CIPAttributes) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPAttributesParse(ctx context.Context, theBytes []byte, packetLength uint16) (CIPAttributes, error) {
	return CIPAttributesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), packetLength)
}

func CIPAttributesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, packetLength uint16) (CIPAttributes, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CIPAttributes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPAttributes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (numberOfClasses) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numberOfClasses, _numberOfClassesErr := readBuffer.ReadUint16("numberOfClasses", 16)
	_ = numberOfClasses
	if _numberOfClassesErr != nil {
		return nil, errors.Wrap(_numberOfClassesErr, "Error parsing 'numberOfClasses' field of CIPAttributes")
	}

	// Array field (classId)
	if pullErr := readBuffer.PullContext("classId", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for classId")
	}
	// Count array
	classId := make([]uint16, utils.Max(numberOfClasses, 0))
	// This happens when the size is set conditional to 0
	if len(classId) == 0 {
		classId = nil
	}
	{
		_numItems := uint16(utils.Max(numberOfClasses, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'classId' field of CIPAttributes")
			}
			classId[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("classId", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for classId")
	}

	// Optional Field (numberAvailable) (Can be skipped, if a given expression evaluates to false)
	var numberAvailable *uint16 = nil
	if bool((packetLength) >= (((numberOfClasses) * (2)) + (4))) {
		_val, _err := readBuffer.ReadUint16("numberAvailable", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'numberAvailable' field of CIPAttributes")
		}
		numberAvailable = &_val
	}

	// Optional Field (numberActive) (Can be skipped, if a given expression evaluates to false)
	var numberActive *uint16 = nil
	if bool((packetLength) >= (((numberOfClasses) * (2)) + (6))) {
		_val, _err := readBuffer.ReadUint16("numberActive", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'numberActive' field of CIPAttributes")
		}
		numberActive = &_val
	}
	// Byte Array field (data)
	numberOfBytesdata := int(utils.InlineIf((bool((packetLength) > (((numberOfClasses) * (2)) + (6)))), func() any {
		return uint16(uint16(packetLength) - uint16((uint16((uint16(numberOfClasses) * uint16(uint16(2)))) + uint16(uint16(6)))))
	}, func() any { return uint16(uint16(0)) }).(uint16))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of CIPAttributes")
	}

	if closeErr := readBuffer.CloseContext("CIPAttributes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPAttributes")
	}

	// Create the instance
	return &_CIPAttributes{
		PacketLength:    packetLength,
		ClassId:         classId,
		NumberAvailable: numberAvailable,
		NumberActive:    numberActive,
		Data:            data,
	}, nil
}

func (m *_CIPAttributes) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPAttributes) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CIPAttributes"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CIPAttributes")
	}

	// Implicit Field (numberOfClasses) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numberOfClasses := uint16(uint16(len(m.GetClassId())))
	_numberOfClassesErr := writeBuffer.WriteUint16("numberOfClasses", 16, uint16((numberOfClasses)))
	if _numberOfClassesErr != nil {
		return errors.Wrap(_numberOfClassesErr, "Error serializing 'numberOfClasses' field")
	}

	// Array Field (classId)
	if pushErr := writeBuffer.PushContext("classId", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for classId")
	}
	for _curItem, _element := range m.GetClassId() {
		_ = _curItem
		_elementErr := writeBuffer.WriteUint16("", 16, uint16(_element))
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'classId' field")
		}
	}
	if popErr := writeBuffer.PopContext("classId", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for classId")
	}

	// Optional Field (numberAvailable) (Can be skipped, if the value is null)
	var numberAvailable *uint16 = nil
	if m.GetNumberAvailable() != nil {
		numberAvailable = m.GetNumberAvailable()
		_numberAvailableErr := writeBuffer.WriteUint16("numberAvailable", 16, uint16(*(numberAvailable)))
		if _numberAvailableErr != nil {
			return errors.Wrap(_numberAvailableErr, "Error serializing 'numberAvailable' field")
		}
	}

	// Optional Field (numberActive) (Can be skipped, if the value is null)
	var numberActive *uint16 = nil
	if m.GetNumberActive() != nil {
		numberActive = m.GetNumberActive()
		_numberActiveErr := writeBuffer.WriteUint16("numberActive", 16, uint16(*(numberActive)))
		if _numberActiveErr != nil {
			return errors.Wrap(_numberActiveErr, "Error serializing 'numberActive' field")
		}
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("CIPAttributes"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CIPAttributes")
	}
	return nil
}

////
// Arguments Getter

func (m *_CIPAttributes) GetPacketLength() uint16 {
	return m.PacketLength
}

//
////

func (m *_CIPAttributes) isCIPAttributes() bool {
	return true
}

func (m *_CIPAttributes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
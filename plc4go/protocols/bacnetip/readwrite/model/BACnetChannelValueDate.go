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

// BACnetChannelValueDate is the corresponding interface of BACnetChannelValueDate
type BACnetChannelValueDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
}

// BACnetChannelValueDateExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueDate.
// This is useful for switch cases.
type BACnetChannelValueDateExactly interface {
	BACnetChannelValueDate
	isBACnetChannelValueDate() bool
}

// _BACnetChannelValueDate is the data-structure of this message
type _BACnetChannelValueDate struct {
	*_BACnetChannelValue
	DateValue BACnetApplicationTagDate
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueDate) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueDate) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueDate) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueDate factory function for _BACnetChannelValueDate
func NewBACnetChannelValueDate(dateValue BACnetApplicationTagDate, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueDate {
	_result := &_BACnetChannelValueDate{
		DateValue:           dateValue,
		_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueDate(structType any) BACnetChannelValueDate {
	if casted, ok := structType.(BACnetChannelValueDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueDate) GetTypeName() string {
	return "BACnetChannelValueDate"
}

func (m *_BACnetChannelValueDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetChannelValueDateParse(ctx context.Context, theBytes []byte) (BACnetChannelValueDate, error) {
	return BACnetChannelValueDateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetChannelValueDateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetChannelValueDate, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetChannelValueDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateValue)
	if pullErr := readBuffer.PullContext("dateValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateValue")
	}
	_dateValue, _dateValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _dateValueErr != nil {
		return nil, errors.Wrap(_dateValueErr, "Error parsing 'dateValue' field of BACnetChannelValueDate")
	}
	dateValue := _dateValue.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("dateValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueDate{
		_BACnetChannelValue: &_BACnetChannelValue{},
		DateValue:           dateValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueDate")
		}

		// Simple Field (dateValue)
		if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateValue")
		}
		_dateValueErr := writeBuffer.WriteSerializable(ctx, m.GetDateValue())
		if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateValue")
		}
		if _dateValueErr != nil {
			return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueDate")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueDate) isBACnetChannelValueDate() bool {
	return true
}

func (m *_BACnetChannelValueDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
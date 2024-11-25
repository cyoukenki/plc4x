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

// BACnetDailySchedule is the corresponding interface of BACnetDailySchedule
type BACnetDailySchedule interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetDaySchedule returns DaySchedule (property field)
	GetDaySchedule() []BACnetTimeValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetDailyScheduleExactly can be used when we want exactly this type and not a type which fulfills BACnetDailySchedule.
// This is useful for switch cases.
type BACnetDailyScheduleExactly interface {
	BACnetDailySchedule
	isBACnetDailySchedule() bool
}

// _BACnetDailySchedule is the data-structure of this message
type _BACnetDailySchedule struct {
	OpeningTag  BACnetOpeningTag
	DaySchedule []BACnetTimeValue
	ClosingTag  BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDailySchedule) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDailySchedule) GetDaySchedule() []BACnetTimeValue {
	return m.DaySchedule
}

func (m *_BACnetDailySchedule) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDailySchedule factory function for _BACnetDailySchedule
func NewBACnetDailySchedule(openingTag BACnetOpeningTag, daySchedule []BACnetTimeValue, closingTag BACnetClosingTag) *_BACnetDailySchedule {
	return &_BACnetDailySchedule{OpeningTag: openingTag, DaySchedule: daySchedule, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastBACnetDailySchedule(structType any) BACnetDailySchedule {
	if casted, ok := structType.(BACnetDailySchedule); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDailySchedule); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDailySchedule) GetTypeName() string {
	return "BACnetDailySchedule"
}

func (m *_BACnetDailySchedule) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.DaySchedule) > 0 {
		for _, element := range m.DaySchedule {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDailySchedule) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDailyScheduleParse(ctx context.Context, theBytes []byte) (BACnetDailySchedule, error) {
	return BACnetDailyScheduleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDailyScheduleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDailySchedule, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetDailySchedule"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDailySchedule")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetDailySchedule")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (daySchedule)
	if pullErr := readBuffer.PullContext("daySchedule", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for daySchedule")
	}
	// Terminated array
	var daySchedule []BACnetTimeValue
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 0)) {
			_item, _err := BACnetTimeValueParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'daySchedule' field of BACnetDailySchedule")
			}
			daySchedule = append(daySchedule, _item.(BACnetTimeValue))
		}
	}
	if closeErr := readBuffer.CloseContext("daySchedule", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for daySchedule")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetDailySchedule")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetDailySchedule"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDailySchedule")
	}

	// Create the instance
	return &_BACnetDailySchedule{
		OpeningTag:  openingTag,
		DaySchedule: daySchedule,
		ClosingTag:  closingTag,
	}, nil
}

func (m *_BACnetDailySchedule) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDailySchedule) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDailySchedule"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDailySchedule")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (daySchedule)
	if pushErr := writeBuffer.PushContext("daySchedule", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for daySchedule")
	}
	for _curItem, _element := range m.GetDaySchedule() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDaySchedule()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'daySchedule' field")
		}
	}
	if popErr := writeBuffer.PopContext("daySchedule", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for daySchedule")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDailySchedule"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDailySchedule")
	}
	return nil
}

func (m *_BACnetDailySchedule) isBACnetDailySchedule() bool {
	return true
}

func (m *_BACnetDailySchedule) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
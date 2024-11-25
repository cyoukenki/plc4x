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

// BACnetEventParameterChangeOfTimer is the corresponding interface of BACnetEventParameterChangeOfTimer
type BACnetEventParameterChangeOfTimer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() BACnetEventParameterChangeOfTimerAlarmValue
	// GetUpdateTimeReference returns UpdateTimeReference (property field)
	GetUpdateTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfTimerExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfTimer.
// This is useful for switch cases.
type BACnetEventParameterChangeOfTimerExactly interface {
	BACnetEventParameterChangeOfTimer
	isBACnetEventParameterChangeOfTimer() bool
}

// _BACnetEventParameterChangeOfTimer is the data-structure of this message
type _BACnetEventParameterChangeOfTimer struct {
	*_BACnetEventParameter
	OpeningTag          BACnetOpeningTag
	TimeDelay           BACnetContextTagUnsignedInteger
	AlarmValues         BACnetEventParameterChangeOfTimerAlarmValue
	UpdateTimeReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag          BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfTimer) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterChangeOfTimer) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfTimer) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfTimer) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfTimer) GetAlarmValues() BACnetEventParameterChangeOfTimerAlarmValue {
	return m.AlarmValues
}

func (m *_BACnetEventParameterChangeOfTimer) GetUpdateTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.UpdateTimeReference
}

func (m *_BACnetEventParameterChangeOfTimer) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfTimer factory function for _BACnetEventParameterChangeOfTimer
func NewBACnetEventParameterChangeOfTimer(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, alarmValues BACnetEventParameterChangeOfTimerAlarmValue, updateTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterChangeOfTimer {
	_result := &_BACnetEventParameterChangeOfTimer{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		AlarmValues:           alarmValues,
		UpdateTimeReference:   updateTimeReference,
		ClosingTag:            closingTag,
		_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfTimer(structType any) BACnetEventParameterChangeOfTimer {
	if casted, ok := structType.(BACnetEventParameterChangeOfTimer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfTimer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfTimer) GetTypeName() string {
	return "BACnetEventParameterChangeOfTimer"
}

func (m *_BACnetEventParameterChangeOfTimer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (alarmValues)
	lengthInBits += m.AlarmValues.GetLengthInBits(ctx)

	// Simple field (updateTimeReference)
	lengthInBits += m.UpdateTimeReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfTimer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfTimerParse(ctx context.Context, theBytes []byte) (BACnetEventParameterChangeOfTimer, error) {
	return BACnetEventParameterChangeOfTimerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterChangeOfTimerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfTimer, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfTimer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfTimer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(uint8(22)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfTimer")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field of BACnetEventParameterChangeOfTimer")
	}
	timeDelay := _timeDelay.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (alarmValues)
	if pullErr := readBuffer.PullContext("alarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmValues")
	}
	_alarmValues, _alarmValuesErr := BACnetEventParameterChangeOfTimerAlarmValueParseWithBuffer(ctx, readBuffer, uint8(uint8(1)))
	if _alarmValuesErr != nil {
		return nil, errors.Wrap(_alarmValuesErr, "Error parsing 'alarmValues' field of BACnetEventParameterChangeOfTimer")
	}
	alarmValues := _alarmValues.(BACnetEventParameterChangeOfTimerAlarmValue)
	if closeErr := readBuffer.CloseContext("alarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmValues")
	}

	// Simple Field (updateTimeReference)
	if pullErr := readBuffer.PullContext("updateTimeReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for updateTimeReference")
	}
	_updateTimeReference, _updateTimeReferenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _updateTimeReferenceErr != nil {
		return nil, errors.Wrap(_updateTimeReferenceErr, "Error parsing 'updateTimeReference' field of BACnetEventParameterChangeOfTimer")
	}
	updateTimeReference := _updateTimeReference.(BACnetDeviceObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("updateTimeReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for updateTimeReference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(uint8(22)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfTimer")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfTimer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfTimer")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfTimer{
		_BACnetEventParameter: &_BACnetEventParameter{},
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		AlarmValues:           alarmValues,
		UpdateTimeReference:   updateTimeReference,
		ClosingTag:            closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfTimer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfTimer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfTimer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfTimer")
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

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeDelay")
		}
		_timeDelayErr := writeBuffer.WriteSerializable(ctx, m.GetTimeDelay())
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeDelay")
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (alarmValues)
		if pushErr := writeBuffer.PushContext("alarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmValues")
		}
		_alarmValuesErr := writeBuffer.WriteSerializable(ctx, m.GetAlarmValues())
		if popErr := writeBuffer.PopContext("alarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmValues")
		}
		if _alarmValuesErr != nil {
			return errors.Wrap(_alarmValuesErr, "Error serializing 'alarmValues' field")
		}

		// Simple Field (updateTimeReference)
		if pushErr := writeBuffer.PushContext("updateTimeReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for updateTimeReference")
		}
		_updateTimeReferenceErr := writeBuffer.WriteSerializable(ctx, m.GetUpdateTimeReference())
		if popErr := writeBuffer.PopContext("updateTimeReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for updateTimeReference")
		}
		if _updateTimeReferenceErr != nil {
			return errors.Wrap(_updateTimeReferenceErr, "Error serializing 'updateTimeReference' field")
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

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfTimer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfTimer")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfTimer) isBACnetEventParameterChangeOfTimer() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfTimer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

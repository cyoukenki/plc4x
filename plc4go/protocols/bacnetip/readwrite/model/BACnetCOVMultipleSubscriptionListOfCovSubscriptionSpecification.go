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

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfCovSubscriptionSpecificationEntry returns ListOfCovSubscriptionSpecificationEntry (property field)
	GetListOfCovSubscriptionSpecificationEntry() []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationExactly can be used when we want exactly this type and not a type which fulfills BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification.
// This is useful for switch cases.
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationExactly interface {
	BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
	isBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification() bool
}

// _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is the data-structure of this message
type _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification struct {
	OpeningTag                              BACnetOpeningTag
	ListOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	ClosingTag                              BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetListOfCovSubscriptionSpecificationEntry() []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	return m.ListOfCovSubscriptionSpecificationEntry
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification factory function for _BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(openingTag BACnetOpeningTag, listOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification{OpeningTag: openingTag, ListOfCovSubscriptionSpecificationEntry: listOfCovSubscriptionSpecificationEntry, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(structType any) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfCovSubscriptionSpecificationEntry) > 0 {
		for _, element := range m.ListOfCovSubscriptionSpecificationEntry {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	return BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfCovSubscriptionSpecificationEntry)
	if pullErr := readBuffer.PullContext("listOfCovSubscriptionSpecificationEntry", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfCovSubscriptionSpecificationEntry")
	}
	// Terminated array
	var listOfCovSubscriptionSpecificationEntry []BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfCovSubscriptionSpecificationEntry' field of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
			}
			listOfCovSubscriptionSpecificationEntry = append(listOfCovSubscriptionSpecificationEntry, _item.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry))
		}
	}
	if closeErr := readBuffer.CloseContext("listOfCovSubscriptionSpecificationEntry", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfCovSubscriptionSpecificationEntry")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}

	// Create the instance
	return &_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification{
		TagNumber:                               tagNumber,
		OpeningTag:                              openingTag,
		ListOfCovSubscriptionSpecificationEntry: listOfCovSubscriptionSpecificationEntry,
		ClosingTag:                              closingTag,
	}, nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
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

	// Array Field (listOfCovSubscriptionSpecificationEntry)
	if pushErr := writeBuffer.PushContext("listOfCovSubscriptionSpecificationEntry", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfCovSubscriptionSpecificationEntry")
	}
	for _curItem, _element := range m.GetListOfCovSubscriptionSpecificationEntry() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetListOfCovSubscriptionSpecificationEntry()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfCovSubscriptionSpecificationEntry' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfCovSubscriptionSpecificationEntry", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfCovSubscriptionSpecificationEntry")
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

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) isBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification() bool {
	return true
}

func (m *_BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
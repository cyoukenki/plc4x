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

// BACnetNameValueCollection is the corresponding interface of BACnetNameValueCollection
type BACnetNameValueCollection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMembers returns Members (property field)
	GetMembers() []BACnetNameValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetNameValueCollectionExactly can be used when we want exactly this type and not a type which fulfills BACnetNameValueCollection.
// This is useful for switch cases.
type BACnetNameValueCollectionExactly interface {
	BACnetNameValueCollection
	isBACnetNameValueCollection() bool
}

// _BACnetNameValueCollection is the data-structure of this message
type _BACnetNameValueCollection struct {
	OpeningTag BACnetOpeningTag
	Members    []BACnetNameValue
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNameValueCollection) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNameValueCollection) GetMembers() []BACnetNameValue {
	return m.Members
}

func (m *_BACnetNameValueCollection) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNameValueCollection factory function for _BACnetNameValueCollection
func NewBACnetNameValueCollection(openingTag BACnetOpeningTag, members []BACnetNameValue, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNameValueCollection {
	return &_BACnetNameValueCollection{OpeningTag: openingTag, Members: members, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetNameValueCollection(structType any) BACnetNameValueCollection {
	if casted, ok := structType.(BACnetNameValueCollection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNameValueCollection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNameValueCollection) GetTypeName() string {
	return "BACnetNameValueCollection"
}

func (m *_BACnetNameValueCollection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.Members) > 0 {
		for _, element := range m.Members {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNameValueCollection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNameValueCollectionParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetNameValueCollection, error) {
	return BACnetNameValueCollectionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNameValueCollectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNameValueCollection, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetNameValueCollection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNameValueCollection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetNameValueCollection")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (members)
	if pullErr := readBuffer.PullContext("members", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for members")
	}
	// Terminated array
	var members []BACnetNameValue
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetNameValueParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'members' field of BACnetNameValueCollection")
			}
			members = append(members, _item.(BACnetNameValue))
		}
	}
	if closeErr := readBuffer.CloseContext("members", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for members")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetNameValueCollection")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNameValueCollection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNameValueCollection")
	}

	// Create the instance
	return &_BACnetNameValueCollection{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		Members:    members,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetNameValueCollection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNameValueCollection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNameValueCollection"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNameValueCollection")
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

	// Array Field (members)
	if pushErr := writeBuffer.PushContext("members", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for members")
	}
	for _curItem, _element := range m.GetMembers() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetMembers()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'members' field")
		}
	}
	if popErr := writeBuffer.PopContext("members", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for members")
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

	if popErr := writeBuffer.PopContext("BACnetNameValueCollection"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNameValueCollection")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNameValueCollection) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetNameValueCollection) isBACnetNameValueCollection() bool {
	return true
}

func (m *_BACnetNameValueCollection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
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

// BACnetLogStatusTagged is the corresponding interface of BACnetLogStatusTagged
type BACnetLogStatusTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetLogDisabled returns LogDisabled (virtual field)
	GetLogDisabled() bool
	// GetBufferPurged returns BufferPurged (virtual field)
	GetBufferPurged() bool
	// GetLogInterrupted returns LogInterrupted (virtual field)
	GetLogInterrupted() bool
}

// BACnetLogStatusTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLogStatusTagged.
// This is useful for switch cases.
type BACnetLogStatusTaggedExactly interface {
	BACnetLogStatusTagged
	isBACnetLogStatusTagged() bool
}

// _BACnetLogStatusTagged is the data-structure of this message
type _BACnetLogStatusTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogStatusTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLogStatusTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLogStatusTagged) GetLogDisabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() any { return bool(m.GetPayload().GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetLogStatusTagged) GetBufferPurged() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() any { return bool(m.GetPayload().GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetLogStatusTagged) GetLogInterrupted() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() any { return bool(m.GetPayload().GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogStatusTagged factory function for _BACnetLogStatusTagged
func NewBACnetLogStatusTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetLogStatusTagged {
	return &_BACnetLogStatusTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetLogStatusTagged(structType any) BACnetLogStatusTagged {
	if casted, ok := structType.(BACnetLogStatusTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogStatusTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogStatusTagged) GetTypeName() string {
	return "BACnetLogStatusTagged"
}

func (m *_BACnetLogStatusTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetLogStatusTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogStatusTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLogStatusTagged, error) {
	return BACnetLogStatusTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLogStatusTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLogStatusTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetLogStatusTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogStatusTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetLogStatusTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParseWithBuffer(ctx, readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetLogStatusTagged")
	}
	payload := _payload.(BACnetTagPayloadBitString)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_logDisabled := utils.InlineIf((bool((len(payload.GetData())) > (0))), func() any { return bool(payload.GetData()[0]) }, func() any { return bool(bool(false)) }).(bool)
	logDisabled := bool(_logDisabled)
	_ = logDisabled

	// Virtual field
	_bufferPurged := utils.InlineIf((bool((len(payload.GetData())) > (1))), func() any { return bool(payload.GetData()[1]) }, func() any { return bool(bool(false)) }).(bool)
	bufferPurged := bool(_bufferPurged)
	_ = bufferPurged

	// Virtual field
	_logInterrupted := utils.InlineIf((bool((len(payload.GetData())) > (2))), func() any { return bool(payload.GetData()[2]) }, func() any { return bool(bool(false)) }).(bool)
	logInterrupted := bool(_logInterrupted)
	_ = logInterrupted

	if closeErr := readBuffer.CloseContext("BACnetLogStatusTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogStatusTagged")
	}

	// Create the instance
	return &_BACnetLogStatusTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Payload:   payload,
	}, nil
}

func (m *_BACnetLogStatusTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogStatusTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLogStatusTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogStatusTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Simple Field (payload)
	if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for payload")
	}
	_payloadErr := writeBuffer.WriteSerializable(ctx, m.GetPayload())
	if popErr := writeBuffer.PopContext("payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for payload")
	}
	if _payloadErr != nil {
		return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
	}
	// Virtual field
	logDisabled := m.GetLogDisabled()
	_ = logDisabled
	if _logDisabledErr := writeBuffer.WriteVirtual(ctx, "logDisabled", m.GetLogDisabled()); _logDisabledErr != nil {
		return errors.Wrap(_logDisabledErr, "Error serializing 'logDisabled' field")
	}
	// Virtual field
	bufferPurged := m.GetBufferPurged()
	_ = bufferPurged
	if _bufferPurgedErr := writeBuffer.WriteVirtual(ctx, "bufferPurged", m.GetBufferPurged()); _bufferPurgedErr != nil {
		return errors.Wrap(_bufferPurgedErr, "Error serializing 'bufferPurged' field")
	}
	// Virtual field
	logInterrupted := m.GetLogInterrupted()
	_ = logInterrupted
	if _logInterruptedErr := writeBuffer.WriteVirtual(ctx, "logInterrupted", m.GetLogInterrupted()); _logInterruptedErr != nil {
		return errors.Wrap(_logInterruptedErr, "Error serializing 'logInterrupted' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogStatusTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogStatusTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLogStatusTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLogStatusTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLogStatusTagged) isBACnetLogStatusTagged() bool {
	return true
}

func (m *_BACnetLogStatusTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
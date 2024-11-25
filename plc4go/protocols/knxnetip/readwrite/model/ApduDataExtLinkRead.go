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

// ApduDataExtLinkRead is the corresponding interface of ApduDataExtLinkRead
type ApduDataExtLinkRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtLinkReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtLinkRead.
// This is useful for switch cases.
type ApduDataExtLinkReadExactly interface {
	ApduDataExtLinkRead
	isApduDataExtLinkRead() bool
}

// _ApduDataExtLinkRead is the data-structure of this message
type _ApduDataExtLinkRead struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtLinkRead) GetExtApciType() uint8 {
	return 0x25
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtLinkRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtLinkRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtLinkRead factory function for _ApduDataExtLinkRead
func NewApduDataExtLinkRead(length uint8) *_ApduDataExtLinkRead {
	_result := &_ApduDataExtLinkRead{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtLinkRead(structType any) ApduDataExtLinkRead {
	if casted, ok := structType.(ApduDataExtLinkRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtLinkRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtLinkRead) GetTypeName() string {
	return "ApduDataExtLinkRead"
}

func (m *_ApduDataExtLinkRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtLinkRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtLinkReadParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtLinkRead, error) {
	return ApduDataExtLinkReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtLinkReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtLinkRead, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataExtLinkRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtLinkRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtLinkRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtLinkRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtLinkRead{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtLinkRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtLinkRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtLinkRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtLinkRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtLinkRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtLinkRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtLinkRead) isApduDataExtLinkRead() bool {
	return true
}

func (m *_ApduDataExtLinkRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
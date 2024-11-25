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

// AnsiExtendedSymbolSegment is the corresponding interface of AnsiExtendedSymbolSegment
type AnsiExtendedSymbolSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DataSegmentType
	// GetSymbol returns Symbol (property field)
	GetSymbol() string
	// GetPad returns Pad (property field)
	GetPad() *uint8
}

// AnsiExtendedSymbolSegmentExactly can be used when we want exactly this type and not a type which fulfills AnsiExtendedSymbolSegment.
// This is useful for switch cases.
type AnsiExtendedSymbolSegmentExactly interface {
	AnsiExtendedSymbolSegment
	isAnsiExtendedSymbolSegment() bool
}

// _AnsiExtendedSymbolSegment is the data-structure of this message
type _AnsiExtendedSymbolSegment struct {
	*_DataSegmentType
	Symbol string
	Pad    *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AnsiExtendedSymbolSegment) GetDataSegmentType() uint8 {
	return 0x11
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AnsiExtendedSymbolSegment) InitializeParent(parent DataSegmentType) {}

func (m *_AnsiExtendedSymbolSegment) GetParent() DataSegmentType {
	return m._DataSegmentType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AnsiExtendedSymbolSegment) GetSymbol() string {
	return m.Symbol
}

func (m *_AnsiExtendedSymbolSegment) GetPad() *uint8 {
	return m.Pad
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAnsiExtendedSymbolSegment factory function for _AnsiExtendedSymbolSegment
func NewAnsiExtendedSymbolSegment(symbol string, pad *uint8) *_AnsiExtendedSymbolSegment {
	_result := &_AnsiExtendedSymbolSegment{
		Symbol:           symbol,
		Pad:              pad,
		_DataSegmentType: NewDataSegmentType(),
	}
	_result._DataSegmentType._DataSegmentTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAnsiExtendedSymbolSegment(structType any) AnsiExtendedSymbolSegment {
	if casted, ok := structType.(AnsiExtendedSymbolSegment); ok {
		return casted
	}
	if casted, ok := structType.(*AnsiExtendedSymbolSegment); ok {
		return *casted
	}
	return nil
}

func (m *_AnsiExtendedSymbolSegment) GetTypeName() string {
	return "AnsiExtendedSymbolSegment"
}

func (m *_AnsiExtendedSymbolSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (dataSize)
	lengthInBits += 8

	// Simple field (symbol)
	lengthInBits += uint16(int32(uint8(len(m.GetSymbol()))) * int32(int32(8)))

	// Optional Field (pad)
	if m.Pad != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_AnsiExtendedSymbolSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AnsiExtendedSymbolSegmentParse(ctx context.Context, theBytes []byte) (AnsiExtendedSymbolSegment, error) {
	return AnsiExtendedSymbolSegmentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AnsiExtendedSymbolSegmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AnsiExtendedSymbolSegment, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AnsiExtendedSymbolSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AnsiExtendedSymbolSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (dataSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	dataSize, _dataSizeErr := readBuffer.ReadUint8("dataSize", 8)
	_ = dataSize
	if _dataSizeErr != nil {
		return nil, errors.Wrap(_dataSizeErr, "Error parsing 'dataSize' field of AnsiExtendedSymbolSegment")
	}

	// Simple Field (symbol)
	_symbol, _symbolErr := readBuffer.ReadString("symbol", uint32((dataSize)*(8)), "UTF-8")
	if _symbolErr != nil {
		return nil, errors.Wrap(_symbolErr, "Error parsing 'symbol' field of AnsiExtendedSymbolSegment")
	}
	symbol := _symbol

	// Optional Field (pad) (Can be skipped, if a given expression evaluates to false)
	var pad *uint8 = nil
	if bool(((len(symbol)) % (2)) != (0)) {
		_val, _err := readBuffer.ReadUint8("pad", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'pad' field of AnsiExtendedSymbolSegment")
		}
		pad = &_val
	}

	if closeErr := readBuffer.CloseContext("AnsiExtendedSymbolSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AnsiExtendedSymbolSegment")
	}

	// Create a partially initialized instance
	_child := &_AnsiExtendedSymbolSegment{
		_DataSegmentType: &_DataSegmentType{},
		Symbol:           symbol,
		Pad:              pad,
	}
	_child._DataSegmentType._DataSegmentTypeChildRequirements = _child
	return _child, nil
}

func (m *_AnsiExtendedSymbolSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AnsiExtendedSymbolSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AnsiExtendedSymbolSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AnsiExtendedSymbolSegment")
		}

		// Implicit Field (dataSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		dataSize := uint8(uint8(len(m.GetSymbol())))
		_dataSizeErr := writeBuffer.WriteUint8("dataSize", 8, uint8((dataSize)))
		if _dataSizeErr != nil {
			return errors.Wrap(_dataSizeErr, "Error serializing 'dataSize' field")
		}

		// Simple Field (symbol)
		symbol := string(m.GetSymbol())
		_symbolErr := writeBuffer.WriteString("symbol", uint32((uint8(len(m.GetSymbol())))*(8)), "UTF-8", (symbol))
		if _symbolErr != nil {
			return errors.Wrap(_symbolErr, "Error serializing 'symbol' field")
		}

		// Optional Field (pad) (Can be skipped, if the value is null)
		var pad *uint8 = nil
		if m.GetPad() != nil {
			pad = m.GetPad()
			_padErr := writeBuffer.WriteUint8("pad", 8, uint8(*(pad)))
			if _padErr != nil {
				return errors.Wrap(_padErr, "Error serializing 'pad' field")
			}
		}

		if popErr := writeBuffer.PopContext("AnsiExtendedSymbolSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AnsiExtendedSymbolSegment")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AnsiExtendedSymbolSegment) isAnsiExtendedSymbolSegment() bool {
	return true
}

func (m *_AnsiExtendedSymbolSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

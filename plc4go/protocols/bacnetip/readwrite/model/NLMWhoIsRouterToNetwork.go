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

// NLMWhoIsRouterToNetwork is the corresponding interface of NLMWhoIsRouterToNetwork
type NLMWhoIsRouterToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress (property field)
	GetDestinationNetworkAddress() *uint16
}

// NLMWhoIsRouterToNetworkExactly can be used when we want exactly this type and not a type which fulfills NLMWhoIsRouterToNetwork.
// This is useful for switch cases.
type NLMWhoIsRouterToNetworkExactly interface {
	NLMWhoIsRouterToNetwork
	isNLMWhoIsRouterToNetwork() bool
}

// _NLMWhoIsRouterToNetwork is the data-structure of this message
type _NLMWhoIsRouterToNetwork struct {
	*_NLM
	DestinationNetworkAddress *uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMWhoIsRouterToNetwork) GetMessageType() uint8 {
	return 0x00
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMWhoIsRouterToNetwork) InitializeParent(parent NLM) {}

func (m *_NLMWhoIsRouterToNetwork) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMWhoIsRouterToNetwork) GetDestinationNetworkAddress() *uint16 {
	return m.DestinationNetworkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMWhoIsRouterToNetwork factory function for _NLMWhoIsRouterToNetwork
func NewNLMWhoIsRouterToNetwork(destinationNetworkAddress *uint16, apduLength uint16) *_NLMWhoIsRouterToNetwork {
	_result := &_NLMWhoIsRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		_NLM:                      NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMWhoIsRouterToNetwork(structType any) NLMWhoIsRouterToNetwork {
	if casted, ok := structType.(NLMWhoIsRouterToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMWhoIsRouterToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMWhoIsRouterToNetwork) GetTypeName() string {
	return "NLMWhoIsRouterToNetwork"
}

func (m *_NLMWhoIsRouterToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (destinationNetworkAddress)
	if m.DestinationNetworkAddress != nil {
		lengthInBits += 16
	}

	return lengthInBits
}

func (m *_NLMWhoIsRouterToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMWhoIsRouterToNetworkParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMWhoIsRouterToNetwork, error) {
	return NLMWhoIsRouterToNetworkParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMWhoIsRouterToNetworkParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMWhoIsRouterToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMWhoIsRouterToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMWhoIsRouterToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (destinationNetworkAddress) (Can be skipped, if a given expression evaluates to false)
	var destinationNetworkAddress *uint16 = nil
	{
		_val, _err := readBuffer.ReadUint16("destinationNetworkAddress", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'destinationNetworkAddress' field of NLMWhoIsRouterToNetwork")
		}
		destinationNetworkAddress = &_val
	}

	if closeErr := readBuffer.CloseContext("NLMWhoIsRouterToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMWhoIsRouterToNetwork")
	}

	// Create a partially initialized instance
	_child := &_NLMWhoIsRouterToNetwork{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		DestinationNetworkAddress: destinationNetworkAddress,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMWhoIsRouterToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMWhoIsRouterToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMWhoIsRouterToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMWhoIsRouterToNetwork")
		}

		// Optional Field (destinationNetworkAddress) (Can be skipped, if the value is null)
		var destinationNetworkAddress *uint16 = nil
		if m.GetDestinationNetworkAddress() != nil {
			destinationNetworkAddress = m.GetDestinationNetworkAddress()
			_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, uint16(*(destinationNetworkAddress)))
			if _destinationNetworkAddressErr != nil {
				return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
			}
		}

		if popErr := writeBuffer.PopContext("NLMWhoIsRouterToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMWhoIsRouterToNetwork")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMWhoIsRouterToNetwork) isNLMWhoIsRouterToNetwork() bool {
	return true
}

func (m *_NLMWhoIsRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
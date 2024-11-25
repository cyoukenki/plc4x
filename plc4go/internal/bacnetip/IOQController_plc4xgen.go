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

// Code generated by "plc4xgenerator -type=IOQController"; DO NOT EDIT.

package bacnetip

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *IOQController) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *IOQController) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("IOQController"); err != nil {
		return err
	}
	if err := d.IOController.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.state)

		if err := writeBuffer.WriteString("state", uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}
	{
		_value := fmt.Sprintf("%v", d.activeIOCB)

		if err := writeBuffer.WriteString("activeIOCB", uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}
	{
		_value := fmt.Sprintf("%v", d.ioQueue)

		if err := writeBuffer.WriteString("ioQueue", uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}

	if err := writeBuffer.WriteString("waitTime", uint32(len(d.waitTime.String())*8), "UTF-8", d.waitTime.String()); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.rootStruct)

		if err := writeBuffer.WriteString("rootStruct", uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("IOQController"); err != nil {
		return err
	}
	return nil
}

func (d *IOQController) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

#ifndef PLC4C_S7_READ_WRITE_COTP_PROTOCOL_CLASS_H_
#define PLC4C_S7_READ_WRITE_COTP_PROTOCOL_CLASS_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/spi/context.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>

// Code generated by code-generation. DO NOT EDIT.


enum plc4c_s7_read_write_cotp_protocol_class {
  plc4c_s7_read_write_cotp_protocol_class_CLASS_0 = 0x00,
  plc4c_s7_read_write_cotp_protocol_class_CLASS_1 = 0x10,
  plc4c_s7_read_write_cotp_protocol_class_CLASS_2 = 0x20,
  plc4c_s7_read_write_cotp_protocol_class_CLASS_3 = 0x30,
  plc4c_s7_read_write_cotp_protocol_class_CLASS_4 = 0x40
};
typedef enum plc4c_s7_read_write_cotp_protocol_class plc4c_s7_read_write_cotp_protocol_class;

// Get an empty NULL-struct
plc4c_s7_read_write_cotp_protocol_class plc4c_s7_read_write_cotp_protocol_class_null();

plc4c_return_code plc4c_s7_read_write_cotp_protocol_class_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_cotp_protocol_class* message);

plc4c_return_code plc4c_s7_read_write_cotp_protocol_class_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_cotp_protocol_class* message);

plc4c_s7_read_write_cotp_protocol_class plc4c_s7_read_write_cotp_protocol_class_for_value(uint8_t value);

plc4c_s7_read_write_cotp_protocol_class plc4c_s7_read_write_cotp_protocol_class_value_of(char* value_string);

int plc4c_s7_read_write_cotp_protocol_class_num_values();

plc4c_s7_read_write_cotp_protocol_class plc4c_s7_read_write_cotp_protocol_class_value_for_index(int index);

uint16_t plc4c_s7_read_write_cotp_protocol_class_length_in_bytes(plc4x_spi_context ctx, plc4c_s7_read_write_cotp_protocol_class* message);

uint16_t plc4c_s7_read_write_cotp_protocol_class_length_in_bits(plc4x_spi_context ctx, plc4c_s7_read_write_cotp_protocol_class* message);


#endif  // PLC4C_S7_READ_WRITE_COTP_PROTOCOL_CLASS_H_

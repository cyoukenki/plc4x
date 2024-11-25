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

#ifndef PLC4C_S7_READ_WRITE_ASSOCIATED_VALUE_TYPE_H_
#define PLC4C_S7_READ_WRITE_ASSOCIATED_VALUE_TYPE_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/spi/context.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/utils/list.h>
#include "data_transport_error_code.h"
#include "data_transport_size.h"

// Code generated by code-generation. DO NOT EDIT.


struct plc4c_s7_read_write_associated_value_type {
  /* Properties */
  plc4c_s7_read_write_data_transport_error_code return_code;
  plc4c_s7_read_write_data_transport_size transport_size;
  uint16_t value_length;
  plc4c_list* data;
};
typedef struct plc4c_s7_read_write_associated_value_type plc4c_s7_read_write_associated_value_type;

// Create an empty NULL-struct
plc4c_s7_read_write_associated_value_type plc4c_s7_read_write_associated_value_type_null();

plc4c_return_code plc4c_s7_read_write_associated_value_type_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_associated_value_type** message);

plc4c_return_code plc4c_s7_read_write_associated_value_type_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_associated_value_type* message);

uint16_t plc4c_s7_read_write_associated_value_type_length_in_bytes(plc4x_spi_context ctx, plc4c_s7_read_write_associated_value_type* message);

uint16_t plc4c_s7_read_write_associated_value_type_length_in_bits(plc4x_spi_context ctx, plc4c_s7_read_write_associated_value_type* message);

#endif  // PLC4C_S7_READ_WRITE_ASSOCIATED_VALUE_TYPE_H_

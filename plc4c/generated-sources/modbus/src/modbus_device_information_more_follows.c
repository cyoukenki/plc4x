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

#include <string.h>
#include <plc4c/driver_modbus_static.h>

#include "modbus_device_information_more_follows.h"

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_modbus_read_write_modbus_device_information_more_follows plc4c_modbus_read_write_modbus_device_information_more_follows_null_const;

plc4c_modbus_read_write_modbus_device_information_more_follows plc4c_modbus_read_write_modbus_device_information_more_follows_null() {
  return plc4c_modbus_read_write_modbus_device_information_more_follows_null_const;
}

// Parse function.
plc4c_return_code plc4c_modbus_read_write_modbus_device_information_more_follows_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_modbus_read_write_modbus_device_information_more_follows* _message) {
    plc4c_return_code _res = OK;

    uint8_t value;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
    *_message = plc4c_modbus_read_write_modbus_device_information_more_follows_for_value(value);

    return _res;
}

plc4c_return_code plc4c_modbus_read_write_modbus_device_information_more_follows_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_modbus_read_write_modbus_device_information_more_follows* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, *_message);

    return _res;
}

plc4c_modbus_read_write_modbus_device_information_more_follows plc4c_modbus_read_write_modbus_device_information_more_follows_for_value(uint8_t value) {
    for(int i = 0; i < plc4c_modbus_read_write_modbus_device_information_more_follows_num_values(); i++) {
        if(plc4c_modbus_read_write_modbus_device_information_more_follows_value_for_index(i) == value) {
            return plc4c_modbus_read_write_modbus_device_information_more_follows_value_for_index(i);
        }
    }
    return -1;
}

plc4c_modbus_read_write_modbus_device_information_more_follows plc4c_modbus_read_write_modbus_device_information_more_follows_value_of(char* value_string) {
    if(strcmp(value_string, "NO_MORE_OBJECTS_AVAILABLE") == 0) {
        return plc4c_modbus_read_write_modbus_device_information_more_follows_NO_MORE_OBJECTS_AVAILABLE;
    }
    if(strcmp(value_string, "MORE_OBJECTS_AVAILABLE") == 0) {
        return plc4c_modbus_read_write_modbus_device_information_more_follows_MORE_OBJECTS_AVAILABLE;
    }
    return -1;
}

int plc4c_modbus_read_write_modbus_device_information_more_follows_num_values() {
  return 2;
}

plc4c_modbus_read_write_modbus_device_information_more_follows plc4c_modbus_read_write_modbus_device_information_more_follows_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_modbus_read_write_modbus_device_information_more_follows_NO_MORE_OBJECTS_AVAILABLE;
      }
      case 1: {
        return plc4c_modbus_read_write_modbus_device_information_more_follows_MORE_OBJECTS_AVAILABLE;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_modbus_read_write_modbus_device_information_more_follows_length_in_bytes(plc4x_spi_context ctx, plc4c_modbus_read_write_modbus_device_information_more_follows* _message) {
    return plc4c_modbus_read_write_modbus_device_information_more_follows_length_in_bits(ctx, _message) / 8;
}

uint16_t plc4c_modbus_read_write_modbus_device_information_more_follows_length_in_bits(plc4x_spi_context ctx, plc4c_modbus_read_write_modbus_device_information_more_follows* _message) {
    return 8;
}

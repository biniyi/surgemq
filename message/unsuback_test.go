// Copyright (c) 2014 The SurgeMQ Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnSubAckMessageFields(t *testing.T) {
	msg := NewUnSubAckMessage()

	msg.SetPacketID(100)

	require.Equal(t, 100, int(msg.PacketID()))
}

func TestUnSubAckMessageDecode(t *testing.T) {
	msgBytes := []byte{
		byte(UNSUBACK << 4),
		2,
		0, // packet ID MSB (0)
		7, // packet ID LSB (7)
	}

	m, n, err := Decode(msgBytes)
	msg, ok := m.(*UnSubAckMessage)
	require.Equal(t, true, ok, "Invalid message type")

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n, "Error decoding message.")
	require.Equal(t, UNSUBACK, msg.Type(), "Error decoding message.")
	require.Equal(t, 7, int(msg.PacketID()), "Error decoding message.")
}

// test insufficient bytes
func TestUnSubAckMessageDecode2(t *testing.T) {
	msgBytes := []byte{
		byte(UNSUBACK << 4),
		2,
		7, // packet ID LSB (7)
	}

	_, _, err := Decode(msgBytes)

	require.Error(t, err)
}

func TestUnSubAckMessageEncode(t *testing.T) {
	msgBytes := []byte{
		byte(UNSUBACK << 4),
		2,
		0, // packet ID MSB (0)
		7, // packet ID LSB (7)
	}

	msg := NewUnSubAckMessage()
	msg.SetPacketID(7)

	dst := make([]byte, 10)
	n, err := msg.Encode(dst)

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n, "Error decoding message.")
	require.Equal(t, msgBytes, dst[:n], "Error decoding message.")
}

// test to ensure encoding and decoding are the same
// decode, encode, and decode again
func TestUnSubAckDecodeEncodeEquiv(t *testing.T) {
	msgBytes := []byte{
		byte(UNSUBACK << 4),
		2,
		0, // packet ID MSB (0)
		7, // packet ID LSB (7)
	}

	m, n, err := Decode(msgBytes)
	msg, ok := m.(*UnSubAckMessage)
	require.Equal(t, true, ok, "Invalid message type")

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n, "Error decoding message.")

	dst := make([]byte, 100)
	n2, err := msg.Encode(dst)

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n2, "Error decoding message.")
	require.Equal(t, msgBytes, dst[:n2], "Error decoding message.")

	_, n3, err := Decode(dst)

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n3, "Error decoding message.")
}

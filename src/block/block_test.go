package block

import (
	"fmt"
	"testing"
	"time"
)

func TestBlock_setHash(t *testing.T) {
	date := time.Date(2018, time.February, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		block    *Block
		expected []byte
	}{
		{
			name:     "GenesisBlock",
			block:    &Block{date.Unix(), []byte("Genesis Block"), []byte{}, []byte{}},
			expected: []byte("8ed4a7ec93295dd573a58a84fb0436eb4a989e0044fb2ab2183722ae54bfec2b"),
		},
		{
			name:     "Block 2",
			block:    &Block{date.Add(time.Second * 10).Unix(), []byte("Block 2"), []byte("8ed4a7ec93295dd573a58a84fb0436eb4a989e0044fb2ab2183722ae54bfec2b"), []byte{}},
			expected: []byte("d5cda0fb57c75d4829883c1fe94674770ebfac417ff4de951d5019a6d5cc0209"),
		},
		{
			name:     "Block 3",
			block:    &Block{date.Add(time.Second * 10).Unix(), []byte("Block 3"), []byte("d5cda0fb57c75d4829883c1fe94674770ebfac417ff4de951d5019a6d5cc0209"), []byte{}},
			expected: []byte("8f55a8bbcaa64e1797abcd8156f9785f37c0810d8b533fb9541bed8b455365a4"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.block.SetHash()
			if fmt.Sprintf("%x", tt.block.Hash) != string(tt.expected) {
				t.Fatalf("expected %s, got: %x", tt.expected, tt.block.Hash)
			}
		})
	}
}

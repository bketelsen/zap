// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package backlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONLoggerCheck(t *testing.T) {
	withJSONLogger(t, opts(InfoLevel), func(jl *jsonLogger, output func() []string) {
		assert.False(t, jl.Check(DebugLevel, "Debug.").OK(), "Expected CheckedMessage to be not OK at disabled levels.")

		cm := jl.Check(InfoLevel, "Info.")
		require.True(t, cm.OK(), "Expected CheckedMessage to be OK at enabled levels.")
		cm.Write(Int("magic", 42))
		assert.Equal(
			t,
			`{"msg":"Info.","level":"info","ts":0,"fields":{"magic":42}}`,
			output()[0],
			"Unexpected output after writing a CheckedMessage.",
		)
	})
}

func TestCheckedMessageIsSingleUse(t *testing.T) {
	expected := []string{
		`{"msg":"Single-use.","level":"info","ts":0,"fields":{}}`,
		`{"msg":"Shouldn't re-use a CheckedMessage.","level":"error","ts":0,"fields":{"original":"Single-use."}}`,
	}
	withJSONLogger(t, nil, func(jl *jsonLogger, output func() []string) {
		cm := jl.Check(InfoLevel, "Single-use.")
		cm.Write() // ok
		cm.Write() // first re-use logs error
		cm.Write() // second re-use is silently ignored
		assert.Equal(t, expected, output(), "Expected re-using a CheckedMessage to log an error.")
	})
}

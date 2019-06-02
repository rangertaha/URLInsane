// The MIT License (MIT)
//
// Copyright © 2018 rangertaha rangertaha@gmail.com
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

package typo

import (
	"fmt"
	"testing"
)

// TestLevenshtein ...
func TestLevenshtein(t *testing.T) {
	var str1 = []rune("Asheville")
	var str2 = []rune("Arizona")
	fmt.Println("Distance between Asheville and Arizona:", Levenshtein(str1, str2))

	str1 = []rune("Python")
	str2 = []rune("Peithen")
	fmt.Println("Distance between Python and Peithen:", Levenshtein(str1, str2))

	str1 = []rune("Orange")
	str2 = []rune("Apple")
	fmt.Println("Distance between Orange and Apple:", Levenshtein(str1, str2))
}

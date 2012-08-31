// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
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

package oglematchers_test

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"testing"
)

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type PointeeTest struct {}
func init()                     { RegisterTestSuite(&PointeeTest{}) }

func TestPointee(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *PointeeTest) Description() {
	wrapped := &fakeMatcher{nil, "taco"}
	matcher := Pointee(wrapped)

	ExpectEq("pointee(taco)", matcher.Description())
}

func (t *PointeeTest) CandidateIsNotAPointer() {
	matcher := Pointee(HasSubstr(""))
	err := matcher.Matches([]byte{})

	ExpectThat(err, Error(Equals("which is not a pointer")))
	ExpectTrue(isFatal(err))
}

func (t *PointeeTest) CandidateIsANilLiteral() {
	matcher := Pointee(HasSubstr(""))
	err := matcher.Matches(nil)

	ExpectThat(err, Error(Equals("which is not a pointer")))
	ExpectTrue(isFatal(err))
}

func (t *PointeeTest) CandidateIsANilPointer() {
	matcher := Pointee(HasSubstr(""))
	err := matcher.Matches((*int)(nil))

	ExpectThat(err, Error(Equals("")))
	ExpectTrue(isFatal(err))
}

func (t *PointeeTest) CallsWrapped() {
	ExpectEq("TODO", "")
}

func (t *PointeeTest) WrappedReturnsTrue() {
	ExpectEq("TODO", "")
}

func (t *PointeeTest) WrappedReturnsNonFatalError() {
	ExpectEq("TODO", "")
}

func (t *PointeeTest) WrappedReturnsFatalError() {
	ExpectEq("TODO", "")
}

// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
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
//
// SPDX-License-Identifier: Apache-2.0

package analyzer_test

import (
	"testing"

	"github.com/specterops/bloodhound/cypher/analyzer"
	"github.com/specterops/bloodhound/cypher/frontend"
	"github.com/specterops/bloodhound/cypher/test"
	"github.com/stretchr/testify/require"
)

func TestQueryComplexity(t *testing.T) {
	// Walk through all positive test cases to ensure that the walker can handle all supported types
	for _, caseLoad := range []string{test.PositiveTestCases, test.MutationTestCases} {
		for _, testCase := range test.LoadFixture(t, caseLoad).RunnableCases() {
			t.Run(testCase.Name, func(t *testing.T) {
				// Only bother with the string match tests
				if testCase.Type == test.TypeStringMatch {
					var (
						details              = test.UnmarshallTestCaseDetails[test.StringMatchTest](t, testCase)
						parseContext         = frontend.NewContext()
						queryModel, parseErr = frontend.ParseCypher(parseContext, details.Query)
					)

					if parseErr != nil {
						t.Fatalf("Parser errors: %s", parseErr.Error())
					}

					if details.Complexity != nil {
						complexity, analyzerErr := analyzer.QueryComplexity(queryModel)
						require.Nil(t, analyzerErr)

						require.Equal(t, *details.Complexity, complexity.Weight)
					}
				}
			})
		}
	}
}

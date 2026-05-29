// Copyright 2026 Google LLC
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

package skilltoolset

import (
	"strings"
	"testing"
)

// Guards against regressions of issue #912: the default system instruction
// must tell the model to call load_skill with the parameter name that matches
// LoadSkillArgs json tag ("name"), not "skill_name".
func TestDefaultInstructionUsesCorrectLoadSkillParam(t *testing.T) {
	if !strings.Contains(defaultSkillSystemInstruction, "`load_skill` tool with `name=\"") {
		t.Errorf("default instruction must instruct load_skill with name=, got:\n%s", defaultSkillSystemInstruction)
	}
	if strings.Contains(defaultSkillSystemInstruction, "`load_skill` tool with `skill_name=\"") {
		t.Errorf("default instruction still uses skill_name= for load_skill")
	}
}

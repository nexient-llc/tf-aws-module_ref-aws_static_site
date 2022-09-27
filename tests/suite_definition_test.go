// Copyright 2022 Nexient LLC
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

package test

import (
	"path"
	"testing"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type TerraTestSuite struct {
	suite.Suite
	TerraformOptions *terraform.Options
	PlanStruct       *terraform.PlanStruct
	ApplyIo          string
}

// Convenience method to avoid constantly dropping the suite scope in log calls
// You can use
// suite.Log(...)
// instead of
// logger.Log(suite.T(), ...)
func (suite *TerraTestSuite) Log(args ...interface{}) {
	logger.Log(suite.T(), args...)
}

// This is run before _any_ test is run
// It defines our working session, plans, and applies before executing tests
// Failure here must defer a suite.TearDown() then register suite.T().FailNow()
func (suite *TerraTestSuite) SetupSuite() {
	// We need to copy to a tmp folder to avoid touching local files
	tempTestFolder := test_structure.CopyTerraformFolderToTemp(suite.T(), "..", ".")
	// This is required for asdf users
	_ = files.CopyFile(path.Join("..", ".tool-versions"), path.Join(tempTestFolder, ".tool-versions"))
	// Define the options for running Terraform
	suite.TerraformOptions = terraform.WithDefaultRetryableErrors(suite.T(), &terraform.Options{
		// Use the temp directory to run the TF
		TerraformDir: tempTestFolder,
		// Save the plan file so it can be reused
		PlanFilePath: "terraform.tfplan",
		// Vars for the module
		Vars: map[string]interface{}{
			"application":  "test",
			"region":       "us-west-2",
			"env":          "dev",
			"env_instance": "000",
		},
	})

	// Handle plan/apply timeouts if we go over
	timedOut := true
	defer (func() {
		if timedOut {
			suite.Log("Apply timed out")
			defer suite.TearDownSuite()
			suite.T().FailNow()
		}
	})()

	// Using the ...E forms allows us to handle errors during the suite setup
	// See https://github.com/stretchr/testify/issues/1123
	// See https://github.com/stretchr/testify/issues/849

	// Plan and store its output in a Go struct for testing
	var planErr error
	suite.PlanStruct, planErr = terraform.InitAndPlanAndShowWithStructE(suite.T(), suite.TerraformOptions)
	if nil != planErr {
		suite.Log(planErr)
		defer suite.TearDownSuite()
		suite.T().FailNow()
	}

	// Apply twice to ensure idempotency
	// Store the stdout/stderr for testing
	var applyErr error
	suite.ApplyIo, applyErr = terraform.ApplyAndIdempotentE(suite.T(), suite.TerraformOptions)
	if nil != applyErr {
		suite.Log(applyErr)
		defer suite.TearDownSuite()
		suite.T().FailNow()
	}

	// We have finished applying everything within the allotted time
	timedOut = false
}

// This is run after _all_ tests are run (or if a failure is called)
func (suite *TerraTestSuite) TearDownSuite() {
	terraform.Destroy(suite.T(), suite.TerraformOptions)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TerraTestSuite))
}

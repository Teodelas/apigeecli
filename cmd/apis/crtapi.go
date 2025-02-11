// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apis

import (
	"github.com/spf13/cobra"
)

// CreateCmd to create api
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an API proxy in an Apigee Org",
	Long:  "Creates an API proxy in an Apigee Org",
}

var (
	targetURLRef                     string
	importProxy, skipPolicy, addCORS bool
)

func init() {
	CreateCmd.AddCommand(OasCreateCmd)
	CreateCmd.AddCommand(GhCreateCmd)
	CreateCmd.AddCommand(BundleCreateCmd)
	CreateCmd.AddCommand(GqlCreateCmd)
	CreateCmd.AddCommand(IntegrationCmd)
	CreateCmd.AddCommand(SwaggerCreateCmd)
}

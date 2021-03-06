// Copyright 2016-2018, Pulumi Corporation.
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

package boundary

import (
	"fmt"
	"github.com/matdehaast/terraform-provider-boundary/boundary"
	"path/filepath"

	"github.com/matdehaast/pulumi-boundary/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "boundary"
	// modules:
	mainMod = "index" // the boundary module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(boundary.NewProvider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "boundary",
		DisplayName:       "",
		Publisher:         "Pulumi",
		LogoURL:           "",
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing boundary cloud resources.",
		Keywords:          []string{"pulumi", "boundary", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://www.pulumi.com",
		Repository:        "https://github.com/matdehaast/pulumi-boundary",
		GitHubOrg:         "",
		Config:            map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"boundary_account_oidc":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AccountOidc")},
			"boundary_account_password":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AccountPassword")},
			"boundary_auth_method":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AuthMethod")},
			"boundary_auth_method_oidc":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AuthMethodOidc")},
			"boundary_auth_method_password":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "AuthMethodPassword")},
			"boundary_credential_library_vault": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CredentialLibraryVault")},
			"boundary_group":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Group")},
			"boundary_host_catalog_plugin":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HostCatalogPlugin")},
			"boundary_host_catalog_static":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HostCatalogStatic")},
			"boundary_host_set_plugin":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HostSetPlugin")},
			"boundary_host_set_static":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HostSetStatic")},
			"boundary_host_static":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "HostStatic")},
			"boundary_managed_group":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ManagedGroup")},
			"boundary_role":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Role")},
			"boundary_scope":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Scope")},
			"boundary_target":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Target")},
			"boundary_user":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
		},
		//JavaScript: &tfbridge.JavaScriptInfo{
		//	// List any npm dependencies and their versions
		//	Dependencies: map[string]string{
		//		"@pulumi/pulumi": "^3.0.0",
		//	},
		//	DevDependencies: map[string]string{
		//		"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
		//		"@types/mime": "^2.0.0",
		//	},
		//	// See the documentation for tfbridge.OverlayInfo for how to lay out this
		//	// section, or refer to the AWS provider. Delete this section if there are
		//	// no overlay files.
		//	//Overlay: &tfbridge.OverlayInfo{},
		//},
		//Python: &tfbridge.PythonInfo{
		//	// List any Python dependencies and their version ranges
		//	Requires: map[string]string{
		//		"pulumi": ">=3.0.0,<4.0.0",
		//	},
		//},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/matdehaast/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		//CSharp: &tfbridge.CSharpInfo{
		//	PackageReferences: map[string]string{
		//		"Pulumi": "3.*",
		//	},
		//},
	}

	prov.SetAutonaming(255, "-")

	return prov
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "boundary:index/accountOidc:AccountOidc":
		r = &AccountOidc{}
	case "boundary:index/accountPassword:AccountPassword":
		r = &AccountPassword{}
	case "boundary:index/authMethod:AuthMethod":
		r = &AuthMethod{}
	case "boundary:index/authMethodOidc:AuthMethodOidc":
		r = &AuthMethodOidc{}
	case "boundary:index/authMethodPassword:AuthMethodPassword":
		r = &AuthMethodPassword{}
	case "boundary:index/credentialLibraryVault:CredentialLibraryVault":
		r = &CredentialLibraryVault{}
	case "boundary:index/group:Group":
		r = &Group{}
	case "boundary:index/hostCatalogPlugin:HostCatalogPlugin":
		r = &HostCatalogPlugin{}
	case "boundary:index/hostCatalogStatic:HostCatalogStatic":
		r = &HostCatalogStatic{}
	case "boundary:index/hostSetPlugin:HostSetPlugin":
		r = &HostSetPlugin{}
	case "boundary:index/hostSetStatic:HostSetStatic":
		r = &HostSetStatic{}
	case "boundary:index/hostStatic:HostStatic":
		r = &HostStatic{}
	case "boundary:index/managedGroup:ManagedGroup":
		r = &ManagedGroup{}
	case "boundary:index/role:Role":
		r = &Role{}
	case "boundary:index/scope:Scope":
		r = &Scope{}
	case "boundary:index/target:Target":
		r = &Target{}
	case "boundary:index/user:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:boundary" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"boundary",
		"index/accountOidc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/accountPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/authMethod",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/authMethodOidc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/authMethodPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/credentialLibraryVault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/hostCatalogPlugin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/hostCatalogStatic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/hostSetPlugin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/hostSetStatic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/hostStatic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/managedGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/scope",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/target",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"boundary",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"boundary",
		&pkg{version},
	)
}

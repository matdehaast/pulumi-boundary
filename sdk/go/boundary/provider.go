// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the boundary package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
	// var.
	Addr pulumi.StringOutput `pulumi:"addr"`
	// The auth method ID e.g. ampw_1234567890
	AuthMethodId pulumi.StringPtrOutput `pulumi:"authMethodId"`
	// The auth method login name for password-style auth methods
	PasswordAuthMethodLoginName pulumi.StringPtrOutput `pulumi:"passwordAuthMethodLoginName"`
	// The auth method password for password-style auth methods
	PasswordAuthMethodPassword pulumi.StringPtrOutput `pulumi:"passwordAuthMethodPassword"`
	// Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
	// mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
	// used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
	RecoveryKmsHcl pulumi.StringPtrOutput `pulumi:"recoveryKmsHcl"`
	// The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
	// used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
	// will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
	Token pulumi.StringPtrOutput `pulumi:"token"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Addr == nil {
		return nil, errors.New("invalid value for required argument 'Addr'")
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:boundary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
	// var.
	Addr string `pulumi:"addr"`
	// The auth method ID e.g. ampw_1234567890
	AuthMethodId *string `pulumi:"authMethodId"`
	// The auth method login name for password-style auth methods
	PasswordAuthMethodLoginName *string `pulumi:"passwordAuthMethodLoginName"`
	// The auth method password for password-style auth methods
	PasswordAuthMethodPassword *string `pulumi:"passwordAuthMethodPassword"`
	// Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
	// mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
	// used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
	RecoveryKmsHcl *string `pulumi:"recoveryKmsHcl"`
	// When set to true, does not validate the Boundary API endpoint certificate
	TlsInsecure *bool `pulumi:"tlsInsecure"`
	// The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
	// used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
	// will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
	// var.
	Addr pulumi.StringInput
	// The auth method ID e.g. ampw_1234567890
	AuthMethodId pulumi.StringPtrInput
	// The auth method login name for password-style auth methods
	PasswordAuthMethodLoginName pulumi.StringPtrInput
	// The auth method password for password-style auth methods
	PasswordAuthMethodPassword pulumi.StringPtrInput
	// Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
	// mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
	// used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
	RecoveryKmsHcl pulumi.StringPtrInput
	// When set to true, does not validate the Boundary API endpoint certificate
	TlsInsecure pulumi.BoolPtrInput
	// The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
	// used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
	// will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
	Token pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env
// var.
func (o ProviderOutput) Addr() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Addr }).(pulumi.StringOutput)
}

// The auth method ID e.g. ampw_1234567890
func (o ProviderOutput) AuthMethodId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AuthMethodId }).(pulumi.StringPtrOutput)
}

// The auth method login name for password-style auth methods
func (o ProviderOutput) PasswordAuthMethodLoginName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.PasswordAuthMethodLoginName }).(pulumi.StringPtrOutput)
}

// The auth method password for password-style auth methods
func (o ProviderOutput) PasswordAuthMethodPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.PasswordAuthMethodPassword }).(pulumi.StringPtrOutput)
}

// Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS
// mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be
// used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
func (o ProviderOutput) RecoveryKmsHcl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.RecoveryKmsHcl }).(pulumi.StringPtrOutput)
}

// The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be
// used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism
// will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.
func (o ProviderOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}

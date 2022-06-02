// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccountPassword struct {
	pulumi.CustomResourceState

	// The resource ID for the auth method.
	AuthMethodId pulumi.StringOutput `pulumi:"authMethodId"`
	// The account description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The login name for this account.
	LoginName pulumi.StringPtrOutput `pulumi:"loginName"`
	// The account name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The account password. Only set on create, changes will not be reflected when updating account.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccountPassword registers a new resource with the given unique name, arguments, and options.
func NewAccountPassword(ctx *pulumi.Context,
	name string, args *AccountPasswordArgs, opts ...pulumi.ResourceOption) (*AccountPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthMethodId == nil {
		return nil, errors.New("invalid value for required argument 'AuthMethodId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource AccountPassword
	err := ctx.RegisterResource("boundary:index/accountPassword:AccountPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountPassword gets an existing AccountPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountPasswordState, opts ...pulumi.ResourceOption) (*AccountPassword, error) {
	var resource AccountPassword
	err := ctx.ReadResource("boundary:index/accountPassword:AccountPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountPassword resources.
type accountPasswordState struct {
	// The resource ID for the auth method.
	AuthMethodId *string `pulumi:"authMethodId"`
	// The account description.
	Description *string `pulumi:"description"`
	// The login name for this account.
	LoginName *string `pulumi:"loginName"`
	// The account name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The account password. Only set on create, changes will not be reflected when updating account.
	Password *string `pulumi:"password"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type AccountPasswordState struct {
	// The resource ID for the auth method.
	AuthMethodId pulumi.StringPtrInput
	// The account description.
	Description pulumi.StringPtrInput
	// The login name for this account.
	LoginName pulumi.StringPtrInput
	// The account name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The account password. Only set on create, changes will not be reflected when updating account.
	Password pulumi.StringPtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (AccountPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPasswordState)(nil)).Elem()
}

type accountPasswordArgs struct {
	// The resource ID for the auth method.
	AuthMethodId string `pulumi:"authMethodId"`
	// The account description.
	Description *string `pulumi:"description"`
	// The login name for this account.
	LoginName *string `pulumi:"loginName"`
	// The account name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The account password. Only set on create, changes will not be reflected when updating account.
	Password *string `pulumi:"password"`
	// The resource type.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a AccountPassword resource.
type AccountPasswordArgs struct {
	// The resource ID for the auth method.
	AuthMethodId pulumi.StringInput
	// The account description.
	Description pulumi.StringPtrInput
	// The login name for this account.
	LoginName pulumi.StringPtrInput
	// The account name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The account password. Only set on create, changes will not be reflected when updating account.
	Password pulumi.StringPtrInput
	// The resource type.
	Type pulumi.StringInput
}

func (AccountPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountPasswordArgs)(nil)).Elem()
}

type AccountPasswordInput interface {
	pulumi.Input

	ToAccountPasswordOutput() AccountPasswordOutput
	ToAccountPasswordOutputWithContext(ctx context.Context) AccountPasswordOutput
}

func (*AccountPassword) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPassword)(nil)).Elem()
}

func (i *AccountPassword) ToAccountPasswordOutput() AccountPasswordOutput {
	return i.ToAccountPasswordOutputWithContext(context.Background())
}

func (i *AccountPassword) ToAccountPasswordOutputWithContext(ctx context.Context) AccountPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordOutput)
}

// AccountPasswordArrayInput is an input type that accepts AccountPasswordArray and AccountPasswordArrayOutput values.
// You can construct a concrete instance of `AccountPasswordArrayInput` via:
//
//          AccountPasswordArray{ AccountPasswordArgs{...} }
type AccountPasswordArrayInput interface {
	pulumi.Input

	ToAccountPasswordArrayOutput() AccountPasswordArrayOutput
	ToAccountPasswordArrayOutputWithContext(context.Context) AccountPasswordArrayOutput
}

type AccountPasswordArray []AccountPasswordInput

func (AccountPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPassword)(nil)).Elem()
}

func (i AccountPasswordArray) ToAccountPasswordArrayOutput() AccountPasswordArrayOutput {
	return i.ToAccountPasswordArrayOutputWithContext(context.Background())
}

func (i AccountPasswordArray) ToAccountPasswordArrayOutputWithContext(ctx context.Context) AccountPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordArrayOutput)
}

// AccountPasswordMapInput is an input type that accepts AccountPasswordMap and AccountPasswordMapOutput values.
// You can construct a concrete instance of `AccountPasswordMapInput` via:
//
//          AccountPasswordMap{ "key": AccountPasswordArgs{...} }
type AccountPasswordMapInput interface {
	pulumi.Input

	ToAccountPasswordMapOutput() AccountPasswordMapOutput
	ToAccountPasswordMapOutputWithContext(context.Context) AccountPasswordMapOutput
}

type AccountPasswordMap map[string]AccountPasswordInput

func (AccountPasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPassword)(nil)).Elem()
}

func (i AccountPasswordMap) ToAccountPasswordMapOutput() AccountPasswordMapOutput {
	return i.ToAccountPasswordMapOutputWithContext(context.Background())
}

func (i AccountPasswordMap) ToAccountPasswordMapOutputWithContext(ctx context.Context) AccountPasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountPasswordMapOutput)
}

type AccountPasswordOutput struct{ *pulumi.OutputState }

func (AccountPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountPassword)(nil)).Elem()
}

func (o AccountPasswordOutput) ToAccountPasswordOutput() AccountPasswordOutput {
	return o
}

func (o AccountPasswordOutput) ToAccountPasswordOutputWithContext(ctx context.Context) AccountPasswordOutput {
	return o
}

// The resource ID for the auth method.
func (o AccountPasswordOutput) AuthMethodId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountPassword) pulumi.StringOutput { return v.AuthMethodId }).(pulumi.StringOutput)
}

// The account description.
func (o AccountPasswordOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPassword) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The login name for this account.
func (o AccountPasswordOutput) LoginName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPassword) pulumi.StringPtrOutput { return v.LoginName }).(pulumi.StringPtrOutput)
}

// The account name. Defaults to the resource name.
func (o AccountPasswordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountPassword) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The account password. Only set on create, changes will not be reflected when updating account.
func (o AccountPasswordOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountPassword) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The resource type.
func (o AccountPasswordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountPassword) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type AccountPasswordArrayOutput struct{ *pulumi.OutputState }

func (AccountPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountPassword)(nil)).Elem()
}

func (o AccountPasswordArrayOutput) ToAccountPasswordArrayOutput() AccountPasswordArrayOutput {
	return o
}

func (o AccountPasswordArrayOutput) ToAccountPasswordArrayOutputWithContext(ctx context.Context) AccountPasswordArrayOutput {
	return o
}

func (o AccountPasswordArrayOutput) Index(i pulumi.IntInput) AccountPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountPassword {
		return vs[0].([]*AccountPassword)[vs[1].(int)]
	}).(AccountPasswordOutput)
}

type AccountPasswordMapOutput struct{ *pulumi.OutputState }

func (AccountPasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountPassword)(nil)).Elem()
}

func (o AccountPasswordMapOutput) ToAccountPasswordMapOutput() AccountPasswordMapOutput {
	return o
}

func (o AccountPasswordMapOutput) ToAccountPasswordMapOutputWithContext(ctx context.Context) AccountPasswordMapOutput {
	return o
}

func (o AccountPasswordMapOutput) MapIndex(k pulumi.StringInput) AccountPasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountPassword {
		return vs[0].(map[string]*AccountPassword)[vs[1].(string)]
	}).(AccountPasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordInput)(nil)).Elem(), &AccountPassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordArrayInput)(nil)).Elem(), AccountPasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountPasswordMapInput)(nil)).Elem(), AccountPasswordMap{})
	pulumi.RegisterOutputType(AccountPasswordOutput{})
	pulumi.RegisterOutputType(AccountPasswordArrayOutput{})
	pulumi.RegisterOutputType(AccountPasswordMapOutput{})
}
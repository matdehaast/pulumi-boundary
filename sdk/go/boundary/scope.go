// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Scope struct {
	pulumi.CustomResourceState

	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
	// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
	// HCL but results in role resources that are unmanaged by Terraform.
	AutoCreateAdminRole pulumi.BoolPtrOutput `pulumi:"autoCreateAdminRole"`
	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
	// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
	// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
	// that are unmanaged by Terraform.
	AutoCreateDefaultRole pulumi.BoolPtrOutput `pulumi:"autoCreateDefaultRole"`
	// The scope description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it
	// to be imported and managed.
	GlobalScope pulumi.BoolPtrOutput `pulumi:"globalScope"`
	// The scope name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scope ID containing the sub scope resource.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
}

// NewScope registers a new resource with the given unique name, arguments, and options.
func NewScope(ctx *pulumi.Context,
	name string, args *ScopeArgs, opts ...pulumi.ResourceOption) (*Scope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	var resource Scope
	err := ctx.RegisterResource("boundary:index/scope:Scope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScope gets an existing Scope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScopeState, opts ...pulumi.ResourceOption) (*Scope, error) {
	var resource Scope
	err := ctx.ReadResource("boundary:index/scope:Scope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Scope resources.
type scopeState struct {
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
	// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
	// HCL but results in role resources that are unmanaged by Terraform.
	AutoCreateAdminRole *bool `pulumi:"autoCreateAdminRole"`
	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
	// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
	// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
	// that are unmanaged by Terraform.
	AutoCreateDefaultRole *bool `pulumi:"autoCreateDefaultRole"`
	// The scope description.
	Description *string `pulumi:"description"`
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it
	// to be imported and managed.
	GlobalScope *bool `pulumi:"globalScope"`
	// The scope name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID containing the sub scope resource.
	ScopeId *string `pulumi:"scopeId"`
}

type ScopeState struct {
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
	// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
	// HCL but results in role resources that are unmanaged by Terraform.
	AutoCreateAdminRole pulumi.BoolPtrInput
	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
	// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
	// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
	// that are unmanaged by Terraform.
	AutoCreateDefaultRole pulumi.BoolPtrInput
	// The scope description.
	Description pulumi.StringPtrInput
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it
	// to be imported and managed.
	GlobalScope pulumi.BoolPtrInput
	// The scope name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID containing the sub scope resource.
	ScopeId pulumi.StringPtrInput
}

func (ScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeState)(nil)).Elem()
}

type scopeArgs struct {
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
	// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
	// HCL but results in role resources that are unmanaged by Terraform.
	AutoCreateAdminRole *bool `pulumi:"autoCreateAdminRole"`
	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
	// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
	// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
	// that are unmanaged by Terraform.
	AutoCreateDefaultRole *bool `pulumi:"autoCreateDefaultRole"`
	// The scope description.
	Description *string `pulumi:"description"`
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it
	// to be imported and managed.
	GlobalScope *bool `pulumi:"globalScope"`
	// The scope name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID containing the sub scope resource.
	ScopeId string `pulumi:"scopeId"`
}

// The set of arguments for constructing a Scope resource.
type ScopeArgs struct {
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
	// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
	// HCL but results in role resources that are unmanaged by Terraform.
	AutoCreateAdminRole pulumi.BoolPtrInput
	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
	// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
	// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
	// that are unmanaged by Terraform.
	AutoCreateDefaultRole pulumi.BoolPtrInput
	// The scope description.
	Description pulumi.StringPtrInput
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it
	// to be imported and managed.
	GlobalScope pulumi.BoolPtrInput
	// The scope name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID containing the sub scope resource.
	ScopeId pulumi.StringInput
}

func (ScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeArgs)(nil)).Elem()
}

type ScopeInput interface {
	pulumi.Input

	ToScopeOutput() ScopeOutput
	ToScopeOutputWithContext(ctx context.Context) ScopeOutput
}

func (*Scope) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (i *Scope) ToScopeOutput() ScopeOutput {
	return i.ToScopeOutputWithContext(context.Background())
}

func (i *Scope) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput)
}

// ScopeArrayInput is an input type that accepts ScopeArray and ScopeArrayOutput values.
// You can construct a concrete instance of `ScopeArrayInput` via:
//
//          ScopeArray{ ScopeArgs{...} }
type ScopeArrayInput interface {
	pulumi.Input

	ToScopeArrayOutput() ScopeArrayOutput
	ToScopeArrayOutputWithContext(context.Context) ScopeArrayOutput
}

type ScopeArray []ScopeInput

func (ScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Scope)(nil)).Elem()
}

func (i ScopeArray) ToScopeArrayOutput() ScopeArrayOutput {
	return i.ToScopeArrayOutputWithContext(context.Background())
}

func (i ScopeArray) ToScopeArrayOutputWithContext(ctx context.Context) ScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeArrayOutput)
}

// ScopeMapInput is an input type that accepts ScopeMap and ScopeMapOutput values.
// You can construct a concrete instance of `ScopeMapInput` via:
//
//          ScopeMap{ "key": ScopeArgs{...} }
type ScopeMapInput interface {
	pulumi.Input

	ToScopeMapOutput() ScopeMapOutput
	ToScopeMapOutputWithContext(context.Context) ScopeMapOutput
}

type ScopeMap map[string]ScopeInput

func (ScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Scope)(nil)).Elem()
}

func (i ScopeMap) ToScopeMapOutput() ScopeMapOutput {
	return i.ToScopeMapOutputWithContext(context.Background())
}

func (i ScopeMap) ToScopeMapOutputWithContext(ctx context.Context) ScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeMapOutput)
}

type ScopeOutput struct{ *pulumi.OutputState }

func (ScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (o ScopeOutput) ToScopeOutput() ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return o
}

// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role
// in the new scope and gives permissions to manage the scope to the provider's user. Marking this true makes for simpler
// HCL but results in role resources that are unmanaged by Terraform.
func (o ScopeOutput) AutoCreateAdminRole() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Scope) pulumi.BoolPtrOutput { return v.AutoCreateAdminRole }).(pulumi.BoolPtrOutput)
}

// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the
// functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the
// ability to authenticate to the anonymous user. Marking this true makes for simpler HCL but results in role resources
// that are unmanaged by Terraform.
func (o ScopeOutput) AutoCreateDefaultRole() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Scope) pulumi.BoolPtrOutput { return v.AutoCreateDefaultRole }).(pulumi.BoolPtrOutput)
}

// The scope description.
func (o ScopeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Scope) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it
// to be imported and managed.
func (o ScopeOutput) GlobalScope() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Scope) pulumi.BoolPtrOutput { return v.GlobalScope }).(pulumi.BoolPtrOutput)
}

// The scope name. Defaults to the resource name.
func (o ScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Scope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The scope ID containing the sub scope resource.
func (o ScopeOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Scope) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

type ScopeArrayOutput struct{ *pulumi.OutputState }

func (ScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Scope)(nil)).Elem()
}

func (o ScopeArrayOutput) ToScopeArrayOutput() ScopeArrayOutput {
	return o
}

func (o ScopeArrayOutput) ToScopeArrayOutputWithContext(ctx context.Context) ScopeArrayOutput {
	return o
}

func (o ScopeArrayOutput) Index(i pulumi.IntInput) ScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Scope {
		return vs[0].([]*Scope)[vs[1].(int)]
	}).(ScopeOutput)
}

type ScopeMapOutput struct{ *pulumi.OutputState }

func (ScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Scope)(nil)).Elem()
}

func (o ScopeMapOutput) ToScopeMapOutput() ScopeMapOutput {
	return o
}

func (o ScopeMapOutput) ToScopeMapOutputWithContext(ctx context.Context) ScopeMapOutput {
	return o
}

func (o ScopeMapOutput) MapIndex(k pulumi.StringInput) ScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Scope {
		return vs[0].(map[string]*Scope)[vs[1].(string)]
	}).(ScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScopeInput)(nil)).Elem(), &Scope{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScopeArrayInput)(nil)).Elem(), ScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScopeMapInput)(nil)).Elem(), ScopeMap{})
	pulumi.RegisterOutputType(ScopeOutput{})
	pulumi.RegisterOutputType(ScopeArrayOutput{})
	pulumi.RegisterOutputType(ScopeMapOutput{})
}

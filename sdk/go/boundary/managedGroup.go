// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedGroup struct {
	pulumi.CustomResourceState

	// The resource ID for the auth method.
	AuthMethodId pulumi.StringOutput `pulumi:"authMethodId"`
	// The managed group description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Boolean expression to filter the workers for this managed group.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// The managed group name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewManagedGroup registers a new resource with the given unique name, arguments, and options.
func NewManagedGroup(ctx *pulumi.Context,
	name string, args *ManagedGroupArgs, opts ...pulumi.ResourceOption) (*ManagedGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthMethodId == nil {
		return nil, errors.New("invalid value for required argument 'AuthMethodId'")
	}
	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	var resource ManagedGroup
	err := ctx.RegisterResource("boundary:index/managedGroup:ManagedGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedGroup gets an existing ManagedGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedGroupState, opts ...pulumi.ResourceOption) (*ManagedGroup, error) {
	var resource ManagedGroup
	err := ctx.ReadResource("boundary:index/managedGroup:ManagedGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedGroup resources.
type managedGroupState struct {
	// The resource ID for the auth method.
	AuthMethodId *string `pulumi:"authMethodId"`
	// The managed group description.
	Description *string `pulumi:"description"`
	// Boolean expression to filter the workers for this managed group.
	Filter *string `pulumi:"filter"`
	// The managed group name. Defaults to the resource name.
	Name *string `pulumi:"name"`
}

type ManagedGroupState struct {
	// The resource ID for the auth method.
	AuthMethodId pulumi.StringPtrInput
	// The managed group description.
	Description pulumi.StringPtrInput
	// Boolean expression to filter the workers for this managed group.
	Filter pulumi.StringPtrInput
	// The managed group name. Defaults to the resource name.
	Name pulumi.StringPtrInput
}

func (ManagedGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedGroupState)(nil)).Elem()
}

type managedGroupArgs struct {
	// The resource ID for the auth method.
	AuthMethodId string `pulumi:"authMethodId"`
	// The managed group description.
	Description *string `pulumi:"description"`
	// Boolean expression to filter the workers for this managed group.
	Filter string `pulumi:"filter"`
	// The managed group name. Defaults to the resource name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ManagedGroup resource.
type ManagedGroupArgs struct {
	// The resource ID for the auth method.
	AuthMethodId pulumi.StringInput
	// The managed group description.
	Description pulumi.StringPtrInput
	// Boolean expression to filter the workers for this managed group.
	Filter pulumi.StringInput
	// The managed group name. Defaults to the resource name.
	Name pulumi.StringPtrInput
}

func (ManagedGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedGroupArgs)(nil)).Elem()
}

type ManagedGroupInput interface {
	pulumi.Input

	ToManagedGroupOutput() ManagedGroupOutput
	ToManagedGroupOutputWithContext(ctx context.Context) ManagedGroupOutput
}

func (*ManagedGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedGroup)(nil)).Elem()
}

func (i *ManagedGroup) ToManagedGroupOutput() ManagedGroupOutput {
	return i.ToManagedGroupOutputWithContext(context.Background())
}

func (i *ManagedGroup) ToManagedGroupOutputWithContext(ctx context.Context) ManagedGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedGroupOutput)
}

// ManagedGroupArrayInput is an input type that accepts ManagedGroupArray and ManagedGroupArrayOutput values.
// You can construct a concrete instance of `ManagedGroupArrayInput` via:
//
//          ManagedGroupArray{ ManagedGroupArgs{...} }
type ManagedGroupArrayInput interface {
	pulumi.Input

	ToManagedGroupArrayOutput() ManagedGroupArrayOutput
	ToManagedGroupArrayOutputWithContext(context.Context) ManagedGroupArrayOutput
}

type ManagedGroupArray []ManagedGroupInput

func (ManagedGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedGroup)(nil)).Elem()
}

func (i ManagedGroupArray) ToManagedGroupArrayOutput() ManagedGroupArrayOutput {
	return i.ToManagedGroupArrayOutputWithContext(context.Background())
}

func (i ManagedGroupArray) ToManagedGroupArrayOutputWithContext(ctx context.Context) ManagedGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedGroupArrayOutput)
}

// ManagedGroupMapInput is an input type that accepts ManagedGroupMap and ManagedGroupMapOutput values.
// You can construct a concrete instance of `ManagedGroupMapInput` via:
//
//          ManagedGroupMap{ "key": ManagedGroupArgs{...} }
type ManagedGroupMapInput interface {
	pulumi.Input

	ToManagedGroupMapOutput() ManagedGroupMapOutput
	ToManagedGroupMapOutputWithContext(context.Context) ManagedGroupMapOutput
}

type ManagedGroupMap map[string]ManagedGroupInput

func (ManagedGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedGroup)(nil)).Elem()
}

func (i ManagedGroupMap) ToManagedGroupMapOutput() ManagedGroupMapOutput {
	return i.ToManagedGroupMapOutputWithContext(context.Background())
}

func (i ManagedGroupMap) ToManagedGroupMapOutputWithContext(ctx context.Context) ManagedGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedGroupMapOutput)
}

type ManagedGroupOutput struct{ *pulumi.OutputState }

func (ManagedGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedGroup)(nil)).Elem()
}

func (o ManagedGroupOutput) ToManagedGroupOutput() ManagedGroupOutput {
	return o
}

func (o ManagedGroupOutput) ToManagedGroupOutputWithContext(ctx context.Context) ManagedGroupOutput {
	return o
}

// The resource ID for the auth method.
func (o ManagedGroupOutput) AuthMethodId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedGroup) pulumi.StringOutput { return v.AuthMethodId }).(pulumi.StringOutput)
}

// The managed group description.
func (o ManagedGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Boolean expression to filter the workers for this managed group.
func (o ManagedGroupOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedGroup) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// The managed group name. Defaults to the resource name.
func (o ManagedGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type ManagedGroupArrayOutput struct{ *pulumi.OutputState }

func (ManagedGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedGroup)(nil)).Elem()
}

func (o ManagedGroupArrayOutput) ToManagedGroupArrayOutput() ManagedGroupArrayOutput {
	return o
}

func (o ManagedGroupArrayOutput) ToManagedGroupArrayOutputWithContext(ctx context.Context) ManagedGroupArrayOutput {
	return o
}

func (o ManagedGroupArrayOutput) Index(i pulumi.IntInput) ManagedGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedGroup {
		return vs[0].([]*ManagedGroup)[vs[1].(int)]
	}).(ManagedGroupOutput)
}

type ManagedGroupMapOutput struct{ *pulumi.OutputState }

func (ManagedGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedGroup)(nil)).Elem()
}

func (o ManagedGroupMapOutput) ToManagedGroupMapOutput() ManagedGroupMapOutput {
	return o
}

func (o ManagedGroupMapOutput) ToManagedGroupMapOutputWithContext(ctx context.Context) ManagedGroupMapOutput {
	return o
}

func (o ManagedGroupMapOutput) MapIndex(k pulumi.StringInput) ManagedGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedGroup {
		return vs[0].(map[string]*ManagedGroup)[vs[1].(string)]
	}).(ManagedGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedGroupInput)(nil)).Elem(), &ManagedGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedGroupArrayInput)(nil)).Elem(), ManagedGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedGroupMapInput)(nil)).Elem(), ManagedGroupMap{})
	pulumi.RegisterOutputType(ManagedGroupOutput{})
	pulumi.RegisterOutputType(ManagedGroupArrayOutput{})
	pulumi.RegisterOutputType(ManagedGroupMapOutput{})
}

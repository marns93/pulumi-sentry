// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # SentryTeam Resource
//
// Sentry Team resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-sentry/sdk/go/sentry"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sentry.NewSentryTeam(ctx, "default", &sentry.SentryTeamArgs{
// 			Organization: pulumi.String("my-organization"),
// 			Slug:         pulumi.String("my-team"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// This resource can be imported using an ID made up of the organization slug and project slugbash
//
// ```sh
//  $ pulumi import sentry:index/sentryTeam:SentryTeam default org-slug/team-slug
// ```
type SentryTeam struct {
	pulumi.CustomResourceState

	HasAccess pulumi.BoolOutput `pulumi:"hasAccess"`
	IsMember  pulumi.BoolOutput `pulumi:"isMember"`
	IsPending pulumi.BoolOutput `pulumi:"isPending"`
	// The human readable name for the team.
	Name pulumi.StringOutput `pulumi:"name"`
	// The slug of the organization the team should be created for.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The unique URL slug for this team. If this is not provided a slug is automatically generated based on the name.
	Slug   pulumi.StringOutput `pulumi:"slug"`
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewSentryTeam registers a new resource with the given unique name, arguments, and options.
func NewSentryTeam(ctx *pulumi.Context,
	name string, args *SentryTeamArgs, opts ...pulumi.ResourceOption) (*SentryTeam, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SentryTeam
	err := ctx.RegisterResource("sentry:index/sentryTeam:SentryTeam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSentryTeam gets an existing SentryTeam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSentryTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SentryTeamState, opts ...pulumi.ResourceOption) (*SentryTeam, error) {
	var resource SentryTeam
	err := ctx.ReadResource("sentry:index/sentryTeam:SentryTeam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SentryTeam resources.
type sentryTeamState struct {
	HasAccess *bool `pulumi:"hasAccess"`
	IsMember  *bool `pulumi:"isMember"`
	IsPending *bool `pulumi:"isPending"`
	// The human readable name for the team.
	Name *string `pulumi:"name"`
	// The slug of the organization the team should be created for.
	Organization *string `pulumi:"organization"`
	// The unique URL slug for this team. If this is not provided a slug is automatically generated based on the name.
	Slug   *string `pulumi:"slug"`
	TeamId *string `pulumi:"teamId"`
}

type SentryTeamState struct {
	HasAccess pulumi.BoolPtrInput
	IsMember  pulumi.BoolPtrInput
	IsPending pulumi.BoolPtrInput
	// The human readable name for the team.
	Name pulumi.StringPtrInput
	// The slug of the organization the team should be created for.
	Organization pulumi.StringPtrInput
	// The unique URL slug for this team. If this is not provided a slug is automatically generated based on the name.
	Slug   pulumi.StringPtrInput
	TeamId pulumi.StringPtrInput
}

func (SentryTeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*sentryTeamState)(nil)).Elem()
}

type sentryTeamArgs struct {
	// The human readable name for the team.
	Name *string `pulumi:"name"`
	// The slug of the organization the team should be created for.
	Organization string `pulumi:"organization"`
	// The unique URL slug for this team. If this is not provided a slug is automatically generated based on the name.
	Slug *string `pulumi:"slug"`
}

// The set of arguments for constructing a SentryTeam resource.
type SentryTeamArgs struct {
	// The human readable name for the team.
	Name pulumi.StringPtrInput
	// The slug of the organization the team should be created for.
	Organization pulumi.StringInput
	// The unique URL slug for this team. If this is not provided a slug is automatically generated based on the name.
	Slug pulumi.StringPtrInput
}

func (SentryTeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sentryTeamArgs)(nil)).Elem()
}

type SentryTeamInput interface {
	pulumi.Input

	ToSentryTeamOutput() SentryTeamOutput
	ToSentryTeamOutputWithContext(ctx context.Context) SentryTeamOutput
}

func (*SentryTeam) ElementType() reflect.Type {
	return reflect.TypeOf((**SentryTeam)(nil)).Elem()
}

func (i *SentryTeam) ToSentryTeamOutput() SentryTeamOutput {
	return i.ToSentryTeamOutputWithContext(context.Background())
}

func (i *SentryTeam) ToSentryTeamOutputWithContext(ctx context.Context) SentryTeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryTeamOutput)
}

// SentryTeamArrayInput is an input type that accepts SentryTeamArray and SentryTeamArrayOutput values.
// You can construct a concrete instance of `SentryTeamArrayInput` via:
//
//          SentryTeamArray{ SentryTeamArgs{...} }
type SentryTeamArrayInput interface {
	pulumi.Input

	ToSentryTeamArrayOutput() SentryTeamArrayOutput
	ToSentryTeamArrayOutputWithContext(context.Context) SentryTeamArrayOutput
}

type SentryTeamArray []SentryTeamInput

func (SentryTeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SentryTeam)(nil)).Elem()
}

func (i SentryTeamArray) ToSentryTeamArrayOutput() SentryTeamArrayOutput {
	return i.ToSentryTeamArrayOutputWithContext(context.Background())
}

func (i SentryTeamArray) ToSentryTeamArrayOutputWithContext(ctx context.Context) SentryTeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryTeamArrayOutput)
}

// SentryTeamMapInput is an input type that accepts SentryTeamMap and SentryTeamMapOutput values.
// You can construct a concrete instance of `SentryTeamMapInput` via:
//
//          SentryTeamMap{ "key": SentryTeamArgs{...} }
type SentryTeamMapInput interface {
	pulumi.Input

	ToSentryTeamMapOutput() SentryTeamMapOutput
	ToSentryTeamMapOutputWithContext(context.Context) SentryTeamMapOutput
}

type SentryTeamMap map[string]SentryTeamInput

func (SentryTeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SentryTeam)(nil)).Elem()
}

func (i SentryTeamMap) ToSentryTeamMapOutput() SentryTeamMapOutput {
	return i.ToSentryTeamMapOutputWithContext(context.Background())
}

func (i SentryTeamMap) ToSentryTeamMapOutputWithContext(ctx context.Context) SentryTeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SentryTeamMapOutput)
}

type SentryTeamOutput struct{ *pulumi.OutputState }

func (SentryTeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SentryTeam)(nil)).Elem()
}

func (o SentryTeamOutput) ToSentryTeamOutput() SentryTeamOutput {
	return o
}

func (o SentryTeamOutput) ToSentryTeamOutputWithContext(ctx context.Context) SentryTeamOutput {
	return o
}

type SentryTeamArrayOutput struct{ *pulumi.OutputState }

func (SentryTeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SentryTeam)(nil)).Elem()
}

func (o SentryTeamArrayOutput) ToSentryTeamArrayOutput() SentryTeamArrayOutput {
	return o
}

func (o SentryTeamArrayOutput) ToSentryTeamArrayOutputWithContext(ctx context.Context) SentryTeamArrayOutput {
	return o
}

func (o SentryTeamArrayOutput) Index(i pulumi.IntInput) SentryTeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SentryTeam {
		return vs[0].([]*SentryTeam)[vs[1].(int)]
	}).(SentryTeamOutput)
}

type SentryTeamMapOutput struct{ *pulumi.OutputState }

func (SentryTeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SentryTeam)(nil)).Elem()
}

func (o SentryTeamMapOutput) ToSentryTeamMapOutput() SentryTeamMapOutput {
	return o
}

func (o SentryTeamMapOutput) ToSentryTeamMapOutputWithContext(ctx context.Context) SentryTeamMapOutput {
	return o
}

func (o SentryTeamMapOutput) MapIndex(k pulumi.StringInput) SentryTeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SentryTeam {
		return vs[0].(map[string]*SentryTeam)[vs[1].(string)]
	}).(SentryTeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SentryTeamInput)(nil)).Elem(), &SentryTeam{})
	pulumi.RegisterInputType(reflect.TypeOf((*SentryTeamArrayInput)(nil)).Elem(), SentryTeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SentryTeamMapInput)(nil)).Elem(), SentryTeamMap{})
	pulumi.RegisterOutputType(SentryTeamOutput{})
	pulumi.RegisterOutputType(SentryTeamArrayOutput{})
	pulumi.RegisterOutputType(SentryTeamMapOutput{})
}

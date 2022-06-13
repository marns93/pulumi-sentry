// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sentry
{
    /// <summary>
    /// ## # sentry.SentryOrganization Resource
    /// 
    /// Sentry Organization resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Sentry = Pulumi.Sentry;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create an organization
    ///         var @default = new Sentry.SentryOrganization("default", new Sentry.SentryOrganizationArgs
    ///         {
    ///             AgreeTerms = true,
    ///             Slug = "my-organization",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an ID made up of the organization slugbash
    /// 
    /// ```sh
    ///  $ pulumi import sentry:index/sentryOrganization:SentryOrganization default org-slug
    /// ```
    /// </summary>
    [SentryResourceType("sentry:index/sentryOrganization:SentryOrganization")]
    public partial class SentryOrganization : Pulumi.CustomResource
    {
        /// <summary>
        /// You agree to the applicable terms of service and privacy policy.
        /// </summary>
        [Output("agreeTerms")]
        public Output<bool> AgreeTerms { get; private set; } = null!;

        /// <summary>
        /// The human readable name for the organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The unique URL slug for this organization. If this is not provided a slug is automatically generated based on the name.
        /// </summary>
        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;


        /// <summary>
        /// Create a SentryOrganization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SentryOrganization(string name, SentryOrganizationArgs args, CustomResourceOptions? options = null)
            : base("sentry:index/sentryOrganization:SentryOrganization", name, args ?? new SentryOrganizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SentryOrganization(string name, Input<string> id, SentryOrganizationState? state = null, CustomResourceOptions? options = null)
            : base("sentry:index/sentryOrganization:SentryOrganization", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/pulumiverse/pulumi-sentry/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SentryOrganization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SentryOrganization Get(string name, Input<string> id, SentryOrganizationState? state = null, CustomResourceOptions? options = null)
        {
            return new SentryOrganization(name, id, state, options);
        }
    }

    public sealed class SentryOrganizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// You agree to the applicable terms of service and privacy policy.
        /// </summary>
        [Input("agreeTerms", required: true)]
        public Input<bool> AgreeTerms { get; set; } = null!;

        /// <summary>
        /// The human readable name for the organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique URL slug for this organization. If this is not provided a slug is automatically generated based on the name.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public SentryOrganizationArgs()
        {
        }
    }

    public sealed class SentryOrganizationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// You agree to the applicable terms of service and privacy policy.
        /// </summary>
        [Input("agreeTerms")]
        public Input<bool>? AgreeTerms { get; set; }

        /// <summary>
        /// The human readable name for the organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique URL slug for this organization. If this is not provided a slug is automatically generated based on the name.
        /// </summary>
        [Input("slug")]
        public Input<string>? Slug { get; set; }

        public SentryOrganizationState()
        {
        }
    }
}

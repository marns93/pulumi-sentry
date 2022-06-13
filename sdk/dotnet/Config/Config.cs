// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Sentry
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("sentry");

        private static readonly __Value<string?> _baseUrl = new __Value<string?>(() => __config.Get("baseUrl"));
        /// <summary>
        /// The Sentry Base API URL
        /// </summary>
        public static string? BaseUrl
        {
            get => _baseUrl.Get();
            set => _baseUrl.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        /// <summary>
        /// The authentication token used to connect to Sentry
        /// </summary>
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

    }
}

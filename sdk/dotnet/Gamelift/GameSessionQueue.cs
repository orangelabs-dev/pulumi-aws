// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GameLift
{
    /// <summary>
    /// Provides an Gamelift Game Session Queue resource.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/gamelift_game_session_queue.html.markdown.
    /// </summary>
    public partial class GameSessionQueue : Pulumi.CustomResource
    {
        /// <summary>
        /// Game Session Queue ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// List of fleet/alias ARNs used by session queue for placing game sessions.
        /// </summary>
        [Output("destinations")]
        public Output<ImmutableArray<string>> Destinations { get; private set; } = null!;

        /// <summary>
        /// Name of the session queue.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more policies used to choose fleet based on player latency. See below.
        /// </summary>
        [Output("playerLatencyPolicies")]
        public Output<ImmutableArray<Outputs.GameSessionQueuePlayerLatencyPolicies>> PlayerLatencyPolicies { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Maximum time a game session request can remain in the queue.
        /// </summary>
        [Output("timeoutInSeconds")]
        public Output<int?> TimeoutInSeconds { get; private set; } = null!;


        /// <summary>
        /// Create a GameSessionQueue resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GameSessionQueue(string name, GameSessionQueueArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:gamelift/gameSessionQueue:GameSessionQueue", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GameSessionQueue(string name, Input<string> id, GameSessionQueueState? state = null, CustomResourceOptions? options = null)
            : base("aws:gamelift/gameSessionQueue:GameSessionQueue", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GameSessionQueue resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GameSessionQueue Get(string name, Input<string> id, GameSessionQueueState? state = null, CustomResourceOptions? options = null)
        {
            return new GameSessionQueue(name, id, state, options);
        }
    }

    public sealed class GameSessionQueueArgs : Pulumi.ResourceArgs
    {
        [Input("destinations")]
        private InputList<string>? _destinations;

        /// <summary>
        /// List of fleet/alias ARNs used by session queue for placing game sessions.
        /// </summary>
        public InputList<string> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<string>());
            set => _destinations = value;
        }

        /// <summary>
        /// Name of the session queue.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("playerLatencyPolicies")]
        private InputList<Inputs.GameSessionQueuePlayerLatencyPoliciesArgs>? _playerLatencyPolicies;

        /// <summary>
        /// One or more policies used to choose fleet based on player latency. See below.
        /// </summary>
        public InputList<Inputs.GameSessionQueuePlayerLatencyPoliciesArgs> PlayerLatencyPolicies
        {
            get => _playerLatencyPolicies ?? (_playerLatencyPolicies = new InputList<Inputs.GameSessionQueuePlayerLatencyPoliciesArgs>());
            set => _playerLatencyPolicies = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Maximum time a game session request can remain in the queue.
        /// </summary>
        [Input("timeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        public GameSessionQueueArgs()
        {
        }
    }

    public sealed class GameSessionQueueState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Game Session Queue ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("destinations")]
        private InputList<string>? _destinations;

        /// <summary>
        /// List of fleet/alias ARNs used by session queue for placing game sessions.
        /// </summary>
        public InputList<string> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<string>());
            set => _destinations = value;
        }

        /// <summary>
        /// Name of the session queue.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("playerLatencyPolicies")]
        private InputList<Inputs.GameSessionQueuePlayerLatencyPoliciesGetArgs>? _playerLatencyPolicies;

        /// <summary>
        /// One or more policies used to choose fleet based on player latency. See below.
        /// </summary>
        public InputList<Inputs.GameSessionQueuePlayerLatencyPoliciesGetArgs> PlayerLatencyPolicies
        {
            get => _playerLatencyPolicies ?? (_playerLatencyPolicies = new InputList<Inputs.GameSessionQueuePlayerLatencyPoliciesGetArgs>());
            set => _playerLatencyPolicies = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Maximum time a game session request can remain in the queue.
        /// </summary>
        [Input("timeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        public GameSessionQueueState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class GameSessionQueuePlayerLatencyPoliciesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum latency value that is allowed for any player.
        /// </summary>
        [Input("maximumIndividualPlayerLatencyMilliseconds", required: true)]
        public Input<int> MaximumIndividualPlayerLatencyMilliseconds { get; set; } = null!;

        /// <summary>
        /// Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
        /// </summary>
        [Input("policyDurationSeconds")]
        public Input<int>? PolicyDurationSeconds { get; set; }

        public GameSessionQueuePlayerLatencyPoliciesArgs()
        {
        }
    }

    public sealed class GameSessionQueuePlayerLatencyPoliciesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum latency value that is allowed for any player.
        /// </summary>
        [Input("maximumIndividualPlayerLatencyMilliseconds", required: true)]
        public Input<int> MaximumIndividualPlayerLatencyMilliseconds { get; set; } = null!;

        /// <summary>
        /// Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
        /// </summary>
        [Input("policyDurationSeconds")]
        public Input<int>? PolicyDurationSeconds { get; set; }

        public GameSessionQueuePlayerLatencyPoliciesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GameSessionQueuePlayerLatencyPolicies
    {
        /// <summary>
        /// Maximum latency value that is allowed for any player.
        /// </summary>
        public readonly int MaximumIndividualPlayerLatencyMilliseconds;
        /// <summary>
        /// Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
        /// </summary>
        public readonly int? PolicyDurationSeconds;

        [OutputConstructor]
        private GameSessionQueuePlayerLatencyPolicies(
            int maximumIndividualPlayerLatencyMilliseconds,
            int? policyDurationSeconds)
        {
            MaximumIndividualPlayerLatencyMilliseconds = maximumIndividualPlayerLatencyMilliseconds;
            PolicyDurationSeconds = policyDurationSeconds;
        }
    }
    }
}

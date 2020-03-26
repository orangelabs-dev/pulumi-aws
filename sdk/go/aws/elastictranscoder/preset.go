// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elastictranscoder

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic Transcoder preset resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elastictranscoder_preset.html.markdown.
type Preset struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// Audio parameters object (documented below).
	Audio PresetAudioPtrOutput `pulumi:"audio"`
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions PresetAudioCodecOptionsPtrOutput `pulumi:"audioCodecOptions"`
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringOutput `pulumi:"container"`
	// A description of the preset (maximum 255 characters)
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringOutput `pulumi:"name"`
	// Thumbnail parameters object (documented below)
	Thumbnails PresetThumbnailsPtrOutput `pulumi:"thumbnails"`
	Type       pulumi.StringOutput       `pulumi:"type"`
	// Video parameters object (documented below)
	Video             PresetVideoPtrOutput `pulumi:"video"`
	VideoCodecOptions pulumi.MapOutput     `pulumi:"videoCodecOptions"`
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks PresetVideoWatermarkArrayOutput `pulumi:"videoWatermarks"`
}

// NewPreset registers a new resource with the given unique name, arguments, and options.
func NewPreset(ctx *pulumi.Context,
	name string, args *PresetArgs, opts ...pulumi.ResourceOption) (*Preset, error) {
	if args == nil || args.Container == nil {
		return nil, errors.New("missing required argument 'Container'")
	}
	if args == nil {
		args = &PresetArgs{}
	}
	var resource Preset
	err := ctx.RegisterResource("aws:elastictranscoder/preset:Preset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPreset gets an existing Preset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PresetState, opts ...pulumi.ResourceOption) (*Preset, error) {
	var resource Preset
	err := ctx.ReadResource("aws:elastictranscoder/preset:Preset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Preset resources.
type presetState struct {
	Arn *string `pulumi:"arn"`
	// Audio parameters object (documented below).
	Audio *PresetAudio `pulumi:"audio"`
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions *PresetAudioCodecOptions `pulumi:"audioCodecOptions"`
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container *string `pulumi:"container"`
	// A description of the preset (maximum 255 characters)
	Description *string `pulumi:"description"`
	// The name of the preset. (maximum 40 characters)
	Name *string `pulumi:"name"`
	// Thumbnail parameters object (documented below)
	Thumbnails *PresetThumbnails `pulumi:"thumbnails"`
	Type       *string           `pulumi:"type"`
	// Video parameters object (documented below)
	Video             *PresetVideo           `pulumi:"video"`
	VideoCodecOptions map[string]interface{} `pulumi:"videoCodecOptions"`
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks []PresetVideoWatermark `pulumi:"videoWatermarks"`
}

type PresetState struct {
	Arn pulumi.StringPtrInput
	// Audio parameters object (documented below).
	Audio PresetAudioPtrInput
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions PresetAudioCodecOptionsPtrInput
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringPtrInput
	// A description of the preset (maximum 255 characters)
	Description pulumi.StringPtrInput
	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringPtrInput
	// Thumbnail parameters object (documented below)
	Thumbnails PresetThumbnailsPtrInput
	Type       pulumi.StringPtrInput
	// Video parameters object (documented below)
	Video             PresetVideoPtrInput
	VideoCodecOptions pulumi.MapInput
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks PresetVideoWatermarkArrayInput
}

func (PresetState) ElementType() reflect.Type {
	return reflect.TypeOf((*presetState)(nil)).Elem()
}

type presetArgs struct {
	// Audio parameters object (documented below).
	Audio *PresetAudio `pulumi:"audio"`
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions *PresetAudioCodecOptions `pulumi:"audioCodecOptions"`
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container string `pulumi:"container"`
	// A description of the preset (maximum 255 characters)
	Description *string `pulumi:"description"`
	// The name of the preset. (maximum 40 characters)
	Name *string `pulumi:"name"`
	// Thumbnail parameters object (documented below)
	Thumbnails *PresetThumbnails `pulumi:"thumbnails"`
	Type       *string           `pulumi:"type"`
	// Video parameters object (documented below)
	Video             *PresetVideo           `pulumi:"video"`
	VideoCodecOptions map[string]interface{} `pulumi:"videoCodecOptions"`
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks []PresetVideoWatermark `pulumi:"videoWatermarks"`
}

// The set of arguments for constructing a Preset resource.
type PresetArgs struct {
	// Audio parameters object (documented below).
	Audio PresetAudioPtrInput
	// Codec options for the audio parameters (documented below)
	AudioCodecOptions PresetAudioCodecOptionsPtrInput
	// The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
	Container pulumi.StringInput
	// A description of the preset (maximum 255 characters)
	Description pulumi.StringPtrInput
	// The name of the preset. (maximum 40 characters)
	Name pulumi.StringPtrInput
	// Thumbnail parameters object (documented below)
	Thumbnails PresetThumbnailsPtrInput
	Type       pulumi.StringPtrInput
	// Video parameters object (documented below)
	Video             PresetVideoPtrInput
	VideoCodecOptions pulumi.MapInput
	// Watermark parameters for the video parameters (documented below)
	// * `videoCodecOptions` (Optional, Forces new resource) Codec options for the video parameters
	VideoWatermarks PresetVideoWatermarkArrayInput
}

func (PresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*presetArgs)(nil)).Elem()
}

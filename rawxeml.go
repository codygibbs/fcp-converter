package converter

import (
	"encoding/xml"
	"strconv"

	"github.com/google/uuid"
)

const (
	compositeNormal     = "normal"
	compositeAdd        = "add"
	compositeSubtract   = "subtract"
	compositeDifference = "difference"
	compositeMultiply   = "multiply"
	compositeScreen     = "screen"
	compositeTexturize  = "texturize"
	compositeHardLight  = "hardlight"
	compositeSoftLight  = "softlight"
	compositeDarken     = "darken"
	compositeLighten    = "lighten"
	compositeMask       = "mask"
	compositeLumaMask   = "lumamask"
)

// Major Elements

// RawXEML describes raw XEML data (before validation, inheritance, etc takes place)
type RawXEML struct {
	XMLName xml.Name `xml:"xmeml"`
	Version int      `xml:"version,attr"`
	// ImportOptions
	// Project Project
	// Bin     Bin
	Clip     *Clip     `xml:"clip,omitempty"`
	Sequence *Sequence `xml:"sequence,omitempty"`
}

// type Project struct {
// 	Name name
// 	Children
// }

// type Bin struct {
// 	Name name
// 	Children
// 	Labels
// 	Comments
// }

// type children struct {
// 	clip     *clip
// 	sequence *sequence
// 	bin      *bin
// }

// Sequence describes a collection of clips and generators sequenced in relation to each other by time, layer, and position.
type Sequence struct {
	Name             name             `xml:"name"`
	Duration         duration         `xml:"duration"`
	Rate             *Rate            `xml:"rate"`
	In               in               `xml:"in,omitempty"`
	Out              out              `xml:"out,omitempty"`
	TimeCode         *TimeCode        `xml:"timecode,omitempty"`
	Media            *Media           `xml:"media,omitempty"`
	Marker           *Marker          `xml:"marker,omitempty"`
	Sequence         *Sequence        `xml:"sequence,omitempty"`
	Labels           *Labels          `xml:"labels,omitempty"`
	Comment          comment          `xml:"comment,omitempty"`
	MasterClipID     masterClipID     `xml:"masterclipid,omitempty"`
	IsMasterClip     isMasterClip     `xml:"ismasterclip,omitempty"`
	LoggingInfo      *LoggingInfo     `xml:"descriptionlogginginfo,omitempty"`
	FilmData         *FilmData        `xml:"filmdata,omitempty"`
	File             *File            `xml:"file,omitempty"`
	PixelAspectRatio pixelAspectRatio `xml:"pixelAspectRatio,omitempty"`
	UUID             *uuid.UUID       `xml:"uuid,omitempty"`
	UpdateBehavior   updateBehavior   `xml:"updatebehavior,omitempty"`
}

// Media describes specific media tracks for a clip or a sequence.
type Media struct {
	Audio *Audio `xml:"audio,omitempty"`
	Video *Video `xml:"video,omitempty"`
}

// Track describes data specific to one or more video or audio elements for a track.
type Track struct {
	ClipItem *ClipItem `xml:"clipitem,omitempty"`
	Enabled  enabled   `xml:"enabled,omitempty"`
	Locked   locked    `xml:"locked,omitempty"`
}

type locked bool

type outputChannelIndex int

// Link describes a link between different clips in a sequence.
type Link struct {
	LinkClipPref linkClipPref `xml:"linkclippref,omitempty"`
	MediaType    mediaType    `xml:"mediatype,omitempty"`
	TrackIndex   trackIndex   `xml:"trackindex,omitempty"`
	ClipIndex    clipIndex    `xml:"clipindex,omitempty"`
	GroupIndex   groupIndex   `xml:"groupindex,omitempty"`
}

type linkClipPref string

type clipIndex int

type groupIndex int

type updateBehavior string // enum

// Section: Clips

// Clip describes an encoded clip in the Browser.
type Clip struct {
	Name         name         `xml:"name"`
	Duration     duration     `xml:"duration"`
	Rate         *Rate        `xml:"rate"`
	In           in           `xml:"in,omitempty"`
	Out          out          `xml:"out,omitempty"`
	MasterClipID masterClipID `xml:"masterclipid,omitempty"`
	IsMasterClip isMasterClip `xml:"ismasterclip,omitempty"`
	Enabled      enabled      `xml:"enabled,omitempty"`
	// media
	// marker
	Anamorphic   anamorphic   `xml:"anamorphic,omitempty"`
	AlphaType    alphaType    `xml:"alphatype,omitempty"`
	AlphaReverse alphaReverse `xml:"alphareverse,omitempty"`
	// labels
	// comments
	// sourceTrack
	CompositeMode compositeMode `xml:"compositemode,omitempty"`
	// subClipInfo
	// filter
	StillFrame       stillFrame       `xml:"stillframe,omitempty"`
	StillFrameOffset stillFrameOffset `xml:"stillframeoffset,omitempty"`
	StartOffset      startOffset      `xml:"startoffset,omitempty"`
	EndOffset        endOffset        `xml:"endoffset,omitempty"`
	File             *File            `xml:"file,omitempty"`
	// loggingInfo
	// timeCode
}

// ClipItem describes a clip in a track.
type ClipItem struct {
	Name             name             `xml:"name"`
	Duration         duration         `xml:"duration"`
	Rate             *Rate            `xml:"rate"`
	In               in               `xml:"in,omitempty"`
	Out              out              `xml:"out,omitempty"`
	MasterClipID     masterClipID     `xml:"masterclipid,omitempty"`
	IsMasterClip     isMasterClip     `xml:"ismasterclip,omitempty"`
	Enabled          enabled          `xml:"enabled,omitempty"`
	Start            start            `xml:"start"`
	End              end              `xml:"end"`
	Link             *Link            `xml:"link,omitempty"`
	SyncOffset       syncOffset       `xml:"syncoffset,omitempty"`
	LoggingInfo      *LoggingInfo     `xml:"logginginfo,omitempty"`
	File             *File            `xml:"file,omitempty"`
	TimeCode         *TimeCode        `xml:"timecode,omitempty"`
	Marker           *Marker          `xml:"marker,omitempty"`
	Anamorphic       anamorphic       `xml:"anamorphic,omitempty"`
	AlphaType        alphaType        `xml:"alphatype,omitempty"`
	AlphaReverse     alphaReverse     `xml:"alphareverse,omitempty"`
	Labels           *Labels          `xml:"labels,omitempty"`
	Comments         *Comments        `xml:"comments,omitempty"`
	SourceTrack      *SourceTrack     `xml:"sourcetrack,omitempty"`
	CompositeMode    compositeMode    `xml:"compositemode,omitempty"`
	SubClipInfo      *SubClipInfo     `xml:"sublipinfo,omitempty"`
	Filter           *Filter          `xml:"filter,omitempty"`
	StillFrame       stillFrame       `xml:"stillframe,omitempty"`
	StillFrameOffset stillFrameOffset `xml:"stillframeoffset,omitempty"`
	Sequence         *Sequence        `xml:"sequence,omitempty"`
	StartOffset      startOffset      `xml:"startoffset,omitempty"`
	EndOffset        endOffset        `xml:"endoffset,omitempty"`
}

type anamorphic bool

type alphaType string // enum alphatype

type alphaReverse bool

type compositeMode string // enum compositemode

type masterClipID string

type isMasterClip bool

// LoggingInfo describes logging information for a clip.
type LoggingInfo struct {
	Description description `xml:"description,omitempty"`
	Scene       scene       `xml:"scene,omitempty"`
	ShotTake    shotTake    `xml:"shottake,omitempty"`
	LogNote     logNote     `xml:"lognote,omitempty"`
	Good        good        `xml:"good,omitempty"`
}

type description string

type scene string

type shotTake string

type logNote string

type good bool

// Labels describes Label and Label 2 information for a clip.
type Labels struct {
	Label  label `xml:"label,omitempty"`
	Label2 label `xml:"label2,omitempty"`
}

type label string

// Comments describes comment information for a clip.
type Comments struct {
	MasterComment1 comment `xml:"mastercomment1,omitempty"`
	MasterComment2 comment `xml:"mastercomment2,omitempty"`
	MasterComment3 comment `xml:"mastercomment3,omitempty"`
	MasterComment4 comment `xml:"mastercomment4,omitempty"`
	ClipCommentA   comment `xml:"clipcommenta,omitempty"`
	ClipCommentB   comment `xml:"clipcommentb,omitempty"`
}

// SourceTrack describes details of the media connected with a clip.
type SourceTrack struct {
	MediaType  mediaType  `xml:"mediatype,omitempty"`
	TrackIndex trackIndex `xml:"trackindex,omitempty"`
}

type start int

type end int

// SubClipInfo describes offset information for a subclip.
type SubClipInfo struct {
	StartOffset startOffset `xml:"startoffset,omitempty"`
	EndOffset   endOffset   `xml:"endoffset,omitempty"`
}

type startOffset int

type endOffset int

type stillFrame bool

type stillFrameOffset int

type syncOffset int

// Section: Video and Audio

// Video describes data specific to video media.
type Video struct {
	Track                 *Track                 `xml:"track,omitempty"`
	Duration              duration               `xml:"duration,omitempty"`
	Format                *Format                `xml:"forma,omitemptyt"`
	SampleCharacteristics *SampleCharacteristics `xml:"samplecharacteristics,omitempty"`
	In                    in                     `xml:"in,omitempty"`
	Out                   out                    `xml:"out,omitempty"`
}

// Audio describes data specific to audio media.
type Audio struct {
	Track                 *Track                 `xml:"track,omitempty"`
	Format                *Format                `xml:"format,omitempty"`
	Outputs               *Outputs               `xml:"outputs,omitempty"`
	In                    in                     `xml:"in,omitempty"`
	Out                   out                    `xml:"out,omitempty"`
	ChannelCount          channelCount           `xml:"channelcount,omitempty"`
	SampleCharacteristics *SampleCharacteristics `xml:"samplecharacteristics,omitempty"`
	TrackCount            trackCount             `xml:"trackcount,omitempty"`
	Rate                  *Rate                  `xml:"rate,omitempty"`
	Duration              duration               `xml:"duration,omitempty"`
}

type channelCount int

type trackCount int

type channelDescription string // enum

// Section: Common Elements

type name string

type duration int

type enabled bool

// File describes an encoded media file used by a Clip.
type File struct {
	ID       string    `xml:"id,attr"`
	Duration duration  `xml:"duration"`
	Rate     *Rate     `xml:"rate"`
	Name     name      `xml:"name,omitempty"`
	PathURL  pathURL   `xml:"pathurl,omitempty"`
	TimeCode *TimeCode `xml:"timecode,omitempty"`
	Media    *Media    `xml:"media,omitempty"`
}

type pathURL string

// Marker describes a named time or range of time in a clip or sequence.
type Marker struct {
	Name    name    `xml:"name"`
	In      in      `xml:"in"`
	Out     out     `xml:"out"`
	Comment comment `xml:"comment,omitempty"`
}

type comment string

type in int

type out int

type mediaType string // enum mediatype

type trackIndex int

type layerIndex int

// Section: Rate and Timecode

// Rate describes an encoded time scale to interpret time values for a higher component.
type Rate struct {
	TimeBase int  `xml:"timebase,omitempty"`
	NTSC     bool `xml:"ntsc,omitempty"`
}

type timebase int

type ntsc bool

// TimeCode describes an encoded value for a clip, sequence, or file.
type TimeCode struct {
	TimeCodeString timeCodeString `xml:"string,omitempty"`
	Frame          frame          `xml:"frame,omitempty"`
	DisplayFormat  displayFormat  `xml:"displayformat,omitempty"`
	Rate           *Rate          `xml:"rate"`
}

type timeCodeString string

type frame int

type displayFormat string // enum displayformat

type field int

// type reel struct {
// 	name
// }

type source string

// Section: Effects

// type generatorItem struct {
// 	name
// 	duration
// 	rate
// 	in
// 	out
// 	start
// 	end
// 	enabled
// 	anamorphic
// 	alphaType
// 	effect
// 	sourceTrack
// }

// type transitionItem struct {
// 	rate
// 	start
// 	end
// 	alignment
// 	effect
// 	name
// 	duration
// }

type alignment string // enum alignment

// Filter describes a filter effect.
type Filter struct {
	Enabled enabled `xml:"enabled,omitempty"`
	Start   start   `xml:"start,omitempty"`
	End     end     `xml:"end,omitempty"`
	Effect  *Effect `xml:"effect,omitempty"`
}

// Effect describes an effect or processing operation.
type Effect struct {
	Name           name           `xml:"name"`
	EffectID       effectID       `xml:"effectid"`
	EffectType     effectType     `xml:"effecttype"`
	MediaType      mediaType      `xml:"mediatype"`
	EffectCategory effectCategory `xml:"effectcategory,omitempty"`
	Parameter      *Parameter     `xml:"parameter,omitempty"`
}

type effectID string

type effectCategory string

type effectType string // enum effecttype

type wipeCode int

type startRatio float32

type endRatio float32

type reverse bool

// Parameter describes a parameter for an effect.
type Parameter struct {
	ParameterID     string           `xml:"parameterid,omitempty"`
	Name            name             `xml:"name,omitempty"`
	Value           *Value           `xml:"value,omitempty"`
	KeyFrame        *KeyFrame        `xml:"keyframe,omitempty"`
	ValueMin        valueMin         `xml:"valuemin,omitempty"`
	ValueMax        valueMin         `xml:"valuemax,omitempty"`
	ValueList       *ValueList       `xml:"valuelist,omitempty"`
	Interpolation   *Interpolation   `xml:"interpolation,omitempty"`
	AppSpecificData *AppSpecificData `xml:"appspecificdata,omitempty"`
}

type parameterID string

type valueMin int

type valueMax int

// ValueList describes information about a pop-up list in a parameter.
type ValueList struct {
	ValueEntry `xml:"valueentry,omitempty"`
}

// ValueEntry describes information about the choice in a pop-up list in a parameter.
type ValueEntry struct {
	Name  name  `xml:"name,omitempty"`
	Value Value `xml:"value,omitempty"`
}

// Value describes a fixed value for an effect parameter or a keyframe.
type Value struct {
	Data  string `xml:",chardata"`
	Red   int    `xml:"red,omitempty"`
	Blue  int    `xml:"blue,omitempty"`
	Green int    `xml:"green,omitempty"`
	Alpha int    `xml:"alpha,omitempty"`
	Horiz horiz  `xml:"horiz,omitempty"`
	Vert  vert   `xml:"vert,omitempty"`
}

// ColorValue describes color information that can be pulled from a value.
type ColorValue struct {
	Red   int
	Blue  int
	Green int
	Alpha int
}

// PositionValue describes position information that can be pulled from a value.
type PositionValue struct {
	Horiz horiz
	Vert  vert
}

// GetBool pulls collected boolean data from a value.
func (v Value) GetBool() bool {
	b, err := strconv.ParseBool(v.Data)
	if err != nil {
		return false
	}

	return b
}

// GetColor pulls collected color data from a value.
func (v Value) GetColor() ColorValue {
	return ColorValue{
		Red:   v.Red,
		Blue:  v.Blue,
		Green: v.Green,
		Alpha: v.Alpha,
	}
}

// GetNumber pulls collected number data from a value.
func (v Value) GetNumber() int {
	i, err := strconv.Atoi(v.Data)
	if err != nil {
		return 0
	}

	return i
}

// GetPosition pulls collected position data from a value.
func (v Value) GetPosition() PositionValue {
	return PositionValue{
		Horiz: v.Horiz,
		Vert:  v.Vert,
	}
}

// KeyFrame describes a keyframe for an effect.
type KeyFrame struct {
	When          when           `xml:"when"`
	Value         *Value         `xml:"value"`
	Interpolation *Interpolation `xml:"interpolation,omitempty"`
	InScale       inScale        `xml:"inscale,omitempty"`
	OutScale      outScale       `xml:"outscale,omitempty"`
	InBEZ         *InBEZ         `xml:"inbez,omitempty"`
	OutBEZ        *OutBEZ        `xml:"outbez,omitempty"`
}

type when int

type inScale int

// InBEZ describes the incoming handle value for a keyframe.
type InBEZ struct {
	Horiz horiz `xml:"horiz,omitempty"`
	Vert  vert  `xml:"vert,omitempty"`
}

type outScale int

// OutBEZ describes the outgoing handle value for a keyframe.
type OutBEZ struct {
	Horiz horiz `xml:"horiz,omitempty"`
	Vert  vert  `xml:"vert,omitempty"`
}

type horiz int

type vert int

// Interpolation describes the type of curve interpretation and data to use in the parent element.
type Interpolation struct {
	Name name `xml:"name"`
}

// Section: Sequence Settings

// Format describes format information for video or audio media in a track.
type Format struct {
	SampleCharacteristics *SampleCharacteristics `xml:"samplecharacteristics,omitempty"`
	AppSpecificData       *AppSpecificData       `xml:"appspecificdata,omitempty"`
}

// SampleCharacteristics describes characteristics of video or audio media.
type SampleCharacteristics struct {
	Width            width            `xml:"width,omitempty"`
	Height           height           `xml:"height,omitempty"`
	Anamorphic       anamorphic       `xml:"anamorphic,omitempty"`
	PixelAspectRatio pixelAspectRatio `xml:"pixelaspectratio,omitempty"`
	FieldDominance   fieldDominance   `xml:"fielddominance,omitempty"`
	Rate             *Rate            `xml:"rate,omitempty"`
	ColorDepth       colorDepth       `xml:"colordepth,omitempty"`
	Codec            *Codec           `xml:"codec,omitempty"`
}

type width int

type height int

type pixelAspectRatio string // enum pixelaspectratio

type fieldDominance string

type colorDepth int // enum colordepth

// Codec describes details about a codec.
type Codec struct {
	Name            name             `xml:"name,omitempty"`
	AppSpecificData *AppSpecificData `xml:"appspecificdata,omitempty"`
}

type depth int // enum 8|16

type sampleRate int // enum 32000|44100|48000

// Outputs describes information about audio outputs.
type Outputs struct {
	Group *Group `xml:"group,omitempty"`
}

// Group describes information about a group of audio output channels.
type Group struct {
	Index       index        `xml:"index,omitempty"`
	NumChannels channelCount `xml:"numchannels,omitempty"`
	DownMix     downMix      `xml:"downmix,omitempty"`
	Channel     *Channel     `xml:"channel,omitempty"`
}

type index int

type downMix int // enum downmix

// Channel describes the output device index of a channel in a group.
type Channel struct {
	Index index `xml:"index,omitempty"`
}

// Section: Application Specific Data

// AppSpecificData describes application-specific data.
type AppSpecificData struct {
	AppName         appName         `xml:"appname,omitempty"`
	AppManufacturer appManufacturer `xml:"appmanufacturer,omitempty"`
	AppVersion      appVersion      `xml:"appversion,omitempty"`
	// Data Data
}

type appName string

type appManufacturer string

type appVersion string

// type data struct {
// 	fcpImageProcessing
// 	qtCodec
// }

// type fcpImageProcessing struct {
// 	useYUV
// 	useSuperWhite
// 	renderMode
// }

type useYUV bool

type useSuperWhite bool

type renderMode string // enum rendermode

// type qtCodec struct {
// 	codecName
// 	codecTypeName
// 	codecTypeCode
// 	codeVendorCode
// 	spatialQuality
// 	temporalQuality
// 	keyFrameRate
// 	dataRate
// }

type codecName string

type codecTypeName string

type codecTypeCode string

type codeVendorCode string

type spatialQuality int // range 0 - 1023

type temporalQuality int // range 0 - 1023

type keyFrameRate int

type dataRate int

// Section: Film Data

// FilmData describes metadata imported from Cinema Tools.
type FilmData struct{}

// Section: Import Options

// type importOptions struct {
// 	createNewProject
// 	targetProjectName
// 	defSequencePresetName
// 	filterReconnectMediaFiles
// 	filterIncludeMarkers
// 	filterIncludeEffects
// 	filterIncludeSequenceSettings
// }

type createNewProject bool

type targetProjectName string

type defSequencePresetName string

type filterReconnectMediaFiles bool // default: true

type filterIncludeMarkers bool // default: true

type filterIncludeEffects bool // default: true

type filterIncludeSequenceSettings bool // default: true

// ImportRawXEML imports XML into a raw XEML data tree
func ImportRawXEML(s []byte) RawXEML {
	var xs RawXEML
	if err := xml.Unmarshal(s, &xs); err != nil {
		panic(err)
	}

	return RawXEML(xs)
}

package converter

import (
	"encoding/xml"

	"github.com/google/uuid"
)

// For XMEML v1 details, see http://mirror.informatimago.com/next/developer.apple.com/documentation/AppleApplications/Conceptual/FinalCutPro_XML/FinalCutPro_XML.pdf
// For XMEML v2-5 changes, see https://developer.apple.com/library/archive/documentation/AppleApplications/Reference/FinalCutPro_XML/VersionsoftheInterchangeFormat/VersionsoftheInterchangeFormat.html

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
	Clip     Clip     `xml:"clip"`
	Sequence Sequence `xml:"sequence"`
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
	Rate             Rate             `xml:"rate"`
	In               in               `xml:"in"`
	Out              out              `xml:"out"`
	TimeCode         TimeCode         `xml:"timecode"`
	Media            Media            `xml:"media"`
	Marker           Marker           `xml:"marker"`
	Sequence         *Sequence        `xml:"sequence"`
	Labels           Labels           `xml:"labels"`
	Comment          comment          `xml:"comment"`
	MasterClipID     masterClipID     `xml:"masterclipid"`
	IsMasterClip     isMasterClip     `xml:"ismasterclip"`
	LoggingInfo      LoggingInfo      `xml:"descriptionlogginginfo"`
	FilmData         FilmData         `xml:"filmdata"`
	File             *File            `xml:"file"`
	PixelAspectRatio pixelAspectRatio `xml:"pixelAspectRatio"`
	UUID             uuid.UUID        `xml:"uuid"`
	UpdateBehavior   updateBehavior   `xml:"updatebehavior"`
}

// Media describes specific media tracks for a clip or a sequence.
type Media struct {
	Audio Audio `xml:"audio"`
	Video Video `xml:"video"`
}

// Video describes data specific to video media.
type Video struct {
	Duration              duration              `xml:"duration"`
	SampleCharacteristics SampleCharacteristics `xml:"samplecharacteristics"`
}

// Audio describes data specific to audio media.
type Audio struct {
	Track                 *Track                `xml:"track"`
	Format                Format                `xml:"format"`
	Outputs               Outputs               `xml:"outputs"`
	In                    in                    `xml:"in"`
	Out                   out                   `xml:"out"`
	ChannelCount          numChannels           `xml:"channelcount"`
	SampleCharacteristics SampleCharacteristics `xml:"samplecharacteristics"`
	TrackCount            trackCount            `xml:"trackcount"`
	Rate                  Rate                  `xml:"rate"`
	Duration              duration              `xml:"duration"`
}

type trackCount int

type channelDescription string

// Track describes data specific to one or more video or audio elements for a track.
type Track struct {
	ClipItem ClipItem `xml:"clipitem"`
	Enabled  enabled  `xml:"enabled"`
	Locked   locked   `xml:"locked"`
}

type locked bool

type outputChannelIndex int

// Link describes a link between different clips in a sequence.
type Link struct {
	LinkClipPref linkClipPref
	MediaType    mediaType
	TrackIndex   trackIndex
	ClipIndex    clipIndex
	GroupIndex   groupIndex
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
	Rate         Rate         `xml:"rate"`
	In           in           `xml:"in"`
	Out          out          `xml:"out"`
	MasterClipID masterClipID `xml:"masterclipid"`
	IsMasterClip isMasterClip `xml:"ismasterclip"`
	Enabled      enabled      `xml:"enabled"`
	// media
	// marker
	Anamorphic   anamorphic   `xml:"anamorphic"`
	AlphaType    alphaType    `xml:"alphatype"`
	AlphaReverse alphaReverse `xml:"alphareverse"`
	// labels
	// comments
	// sourceTrack
	CompositeMode compositeMode `xml:"compositemode"`
	// subClipInfo
	// filter
	StillFrame       stillFrame       `xml:"stillframe"`
	StillFrameOffset stillFrameOffset `xml:"stillframeoffset"`
	StartOffset      startOffset      `xml:"startoffset"`
	EndOffset        endOffset        `xml:"endoffset"`
	File             File             `xml:"file"`
	// loggingInfo
	// timeCode
}

// ClipItem describes a clip in a track.
type ClipItem struct {
	Name         name         `xml:"name"`
	Duration     duration     `xml:"duration"`
	Rate         Rate         `xml:"rate"`
	In           in           `xml:"in"`
	Out          out          `xml:"out"`
	MasterClipID masterClipID `xml:"masterclipid"`
	IsMasterClip isMasterClip `xml:"ismasterclip"`
	Enabled      enabled      `xml:"enabled"`
	Start        start
	End          end
	Link         Link
	SyncOffset   syncOffset
	LoggingInfo  LoggingInfo
	// File             File
	TimeCode         TimeCode
	Marker           Marker
	Anamorphic       anamorphic
	AlphaType        alphaType
	AlphaReverse     alphaReverse
	Labels           Labels
	Comments         Comments
	SourceTrack      SourceTrack
	CompositeMode    compositeMode
	SubClipInfo      SubClipInfo
	Filter           Filter
	StillFrame       stillFrame
	StillFrameOffset stillFrameOffset
	Sequence         Sequence
	StartOffset      startOffset
	EndOffset        endOffset
}

type anamorphic bool

type alphaType string // enum alphatype

type alphaReverse bool

type compositeMode string // enum compositemode

type masterClipID string

type isMasterClip bool

// LoggingInfo describes logging information for a clip.
type LoggingInfo struct {
	Description description `xml:"description"`
	Scene       scene       `xml:"scene"`
	ShotTake    shotTake    `xml:"shottake"`
	LogNote     logNote     `xml:"lognote"`
	Good        good        `xml:"good"`
}

type description string

type scene string

type shotTake string

type logNote string

type good bool

// Labels describes Label and Label 2 information for a clip.
type Labels struct {
	Label  label `xml:"label"`
	Label2 label `xml:"label2"`
}

type label string

// Comments describes comment information for a clip.
type Comments struct {
	MasterComment1 comment
	MasterComment2 comment
	MasterComment3 comment
	MasterComment4 comment
	ClipCommentA   comment
	ClipCommentB   comment
}

// SourceTrack describes details of the media connected with a clip.
type SourceTrack struct {
	MediaType  mediaType  `xml:"mediatype"`
	TrackIndex trackIndex `xml:"trackindex"`
}

type start int

type end int

// SubClipInfo describes offset information for a subclip.
type SubClipInfo struct {
	StartOffset startOffset `xml:"startoffset"`
	EndOffset   endOffset   `xml:"endoffset"`
}

type startOffset int

type endOffset int

type stillFrame bool

type stillFrameOffset int

type syncOffset int

// Section: Common Elements

type name string

type duration int

type enabled bool

// File describes an encoded media file used by a Clip.
type File struct {
	ID       string   `xml:"id,attr"`
	Duration duration `xml:"duration"`
	Rate     Rate     `xml:"rate"`
	Name     name     `xml:"name"`
	PathURL  pathURL  `xml:"pathurl"`
	TimeCode TimeCode `xml:"timecode"`
	Media    Media    `xml:"media"`
}

type pathURL string

// Marker describes a named time or range of time in a clip or sequence.
type Marker struct {
	Name    name    `xml:"name"`
	In      in      `xml:"in"`
	Out     out     `xml:"out"`
	Comment comment `xml:"comment"`
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
	TimeBase int  `xml:"timebase"`
	NTSC     bool `xml:"ntsc"`
}

type timebase int

type ntsc bool

// TimeCode describes an encoded value for a clip, sequence, or file.
type TimeCode struct {
	TimeCodeString timeCodeString `xml:"string"`
	Frame          frame          `xml:"frame"`
	DisplayFormat  displayFormat  `xml:"displayformat"`
	Rate           Rate           `xml:"rate"`
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
	Enabled enabled `xml:"enabled"`
	Start   start   `xml:"start"`
	End     end     `xml:"end"`
	// effect
}

// type effect struct {
// 	name
// 	effectID
// 	effectType
// 	mediaType
// 	effectCategory
// 	parameter
// }

type effectID string

type effectCategory string

type effectType string // enum effecttype

type wipeCode int

type startRatio float32

type endRatio float32

type reverse bool

// type parameter struct {
// 	parameterID
// 	name
// 	value
// 	keyFrame
// 	valueMin
// 	valueMax
// 	valueList
// 	interpolation
// 	appSpecificData
// }

type parameterID string

type valueMin int

type valueMax int

// type valueList struct {
// 	valueEntry
// }

// type valueEntry struct {
// 	name
// 	value
// }

// type value struct {
// 	number  int
// 	boolean bool
// 	red     int
// 	blue    int
// 	green   int
// 	alpha   int
// 	horiz
// 	vert
// }

// type keyFrame struct {
// 	when
// 	value
// 	interpolation
// 	inScale
// 	outScale
// 	inBEZ
// 	outBEZ
// }

type when int

type inScale int

// type inBEZ struct {
// 	horiz
// 	vert
// }

type outScale int

// type outBEZ struct {
// 	horiz
// 	vert
// }

type horiz int

type vert int

// type interpolation struct {
// 	name
// }

// Section: Sequence Settings

// Format describes format information for video or audio media in a track.
type Format struct {
	SampleCharacteristics SampleCharacteristics `xml:"samplecharacteristics"`
	AppSpecificData       AppSpecificData       `xml:"appspecificdata"`
}

// SampleCharacteristics describes characteristics of video or audio media.
type SampleCharacteristics struct {
	Width            width
	Height           height
	Anamorphic       anamorphic
	PixelAspectRatio pixelAspectRatio
	FieldDominance   fieldDominance
	Rate             Rate `xml:"rate"`
	ColorDepth       colorDepth
	Codec            Codec
}

type width int

type height int

type pixelAspectRatio string // enum pixelaspectratio

type fieldDominance string

type colorDepth int // enum colordepth

// Codec describes details about a codec.
type Codec struct {
	Name            name            `xml:"name"`
	AppSpecificData AppSpecificData `xml:"appspecificdata"`
}

type depth int // enum 8|16

type sampleRate int // enum 32000|44100|48000

// Outputs describes information about audio outputs.
type Outputs struct {
	Group Group `xml:"group"`
}

// Group describes information about a group of audio output channels.
type Group struct {
	Index       index       `xml:"index"`
	NumChannels numChannels `xml:"numchannels"`
	DownMix     downMix     `xml:"downmix"`
	Channel     Channel     `xml:"channel"`
}

type index int

type numChannels int

type downMix int // enum downmix

// Channel describes the output device index of a channel in a group.
type Channel struct {
	Index index `xml:"index"`
}

// Section: Application Specific Data

// AppSpecificData describes application-specific data.
type AppSpecificData struct {
	AppName         appName         `xml:"appname"`
	AppManufacturer appManufacturer `xml:"appmanufacturer"`
	AppVersion      appVersion      `xml:"appversion"`
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

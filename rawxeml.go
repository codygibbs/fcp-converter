package converter

import (
	"encoding/xml"
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
	Name     name     `xml:"name"`
	Duration duration `xml:"duration"`
	Rate     Rate     `xml:"rate"`
	In       in       `xml:"in"`
	Out      out      `xml:"out"`
	TimeCode TimeCode `xml:"timecode"`
}

// type media struct {
// 	audio
// 	video
// }

// type video struct {
// 	duration
// 	sample sampleCharacteristics
// }

// type audio struct {
// 	track
// 	format
// 	outputs
// 	in
// 	out
// 	channelCount int
// 	sampleCharacteristics
// 	trackCount int
// 	rate
// 	duration
// }

type trackCount int

type channelDescription string

// type track struct {
// 	clipItem
// 	enabled bool
// 	locked
// }

type locked bool

type outputChannelIndex int

// type link struct {
// 	linkClipPref
// 	mediaType
// 	trackIndex
// 	clipIndex
// 	groupIndex
// }

type linkClipPref string

type clipIndex int

type groupIndex int

// Section: Clips

// Clip describes an encoded clip in the Browser
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

// type clipItem struct {
// 	name
// 	duration
// 	rate
// 	start
// 	end
// 	link
// 	syncOffset
// 	enabled
// 	in
// 	out
// 	masterClipID
// 	isMasterClip
// 	loggingInfo
// 	file
// 	timeCode
// 	marker
// 	anamorphic
// 	alphaType
// 	alphaReverse
// 	labels
// 	comments
// 	sourceTrack
// 	compositeMode
// 	subClipInfo
// 	filter
// 	stillFrame
// 	stillFrameOffset
// 	sequence
// 	startOffset
// 	endOffset
// }

type anamorphic bool

type alphaType string // enum alphatype

type alphaReverse bool

type compositeMode string // enum compositemode

type masterClipID string

type isMasterClip bool

// type loggingInfo struct {
// 	description
// 	scene
// 	shotTake
// 	logNote
// 	good
// }

type description string

type scene string

type shotTake string

type logNote string

type good bool

// type labels struct {
// 	label
// 	label2
// }

type label string

type label2 string

// type comments struct {
// 	masterComment1 string
// 	masterComment2 string
// 	masterComment3 string
// 	masterComment4 string
// 	clipCommentA   string
// 	clipCommentB   string
// }

// type sourceTrack struct {
// 	mediaType
// 	trackIndex
// }

type start int

type end int

// type subClipInfo struct {
// 	startOffset
// 	endOffset
// }

type startOffset int

type endOffset int

type stillFrame bool

type stillFrameOffset int

type syncOffset int

// Section: Common Elements

type name string

type duration int

type enabled bool

// File describes an encoded media file used by a Clip
type File struct {
	ID string `xml:"id,attr"`
	// duration
	Rate Rate `xml:"rate"`
	// name
	PathURL pathURL `xml:"pathurl"`
	// timeCode
	// media
}

type pathURL string

// type marker struct {
// 	name
// 	in
// 	out
// 	comment
// }

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

// type filter struct {
// 	enabled
// 	start
// 	end
// 	effect
// }

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

// type format struct {
// 	sampleCharacteristics
// }

// type sampleCharacteristics struct {
// 	width
// 	height
// 	anamorphic
// 	pixelAspectRatio
// 	fieldDominance
// 	rate
// 	colorDepth
// 	codec
// }

type width int

type height int

type pixelAspectRatio string // enum pixelaspectratio

type fieldDominance string

type colorDepth int // enum colordepth

// type codec struct {
// 	name
// 	appSpecificData
// }

type depth int // enum 8|16

type sampleRate int // enum 32000|44100|48000

// type outputs struct {
// 	group
// }

// type group struct {
// 	index
// 	numChannels
// 	downMix
// 	channel
// }

type index int

type numChannels int

type downMix int // enum downmix

// type channel struct {
// 	index
// }

// Section: Application Specific Data

// type appSpecificData struct {
// 	appName
// 	appManufacturer
// 	appVersion
// 	data
// }

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

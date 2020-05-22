package converter

// See http://mirror.informatimago.com/next/developer.apple.com/documentation/AppleApplications/Conceptual/FinalCutPro_XML/FinalCutPro_XML.pdf

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

type clipItem struct {
	name          string
	duration      int
	rate          rate
	start         int
	end           int
	enabled       bool
	in            int
	out           int
	compositeMode string
}

type effect struct {
	name           string
	effectID       string
	effectType     string // enum effecttype
	mediaType      string
	effectCategory string
	parameter      parameter
}

type file struct {
	id       string
	name     string
	duration int
	rate     rate
	pathURL  string
	timeCode timeCode
	media    media
}

type filter struct {
	enabled bool
	start   int
	end     int
	effect  effect
}

type media struct {
	audio mediaAudio
	video mediaVideo
}

type mediaAudio struct {
	in           int
	out          int
	channelCount int
	sample       sampleCharacteristics
	trackCount   int
}

type mediaVideo struct {
	duration int
	sample   sampleCharacteristics
}

type rate struct {
	timebase int
	ntsc     bool
}

type parameter struct {
	id       string
	name     string
	value    int // complex type
	valueMin int
	valueMax int
}

type sampleCharacteristics struct {
	height int
	width  int
}

type sequence struct {
	name     string
	duration int
	rate     rate
	in       int
	out      int
	timeCode timeCode
}

type timeCode struct {
	timeCodeString string
	frame          int
	displayFormat  string
	rate           rate
}

type track struct {
	clipItem clipItem
	enabled  bool
	locked   bool
}

type value struct {
	number  int
	boolean bool
	color   valueColor
	shift   valueShift
}

type valueColor struct {
	red   int
	blue  int
	green int
	alpha int
}

type valueShift struct {
	horiz int
	vert  int
}

// FCP describes the Final-Cut Pro XML Interchange Format
type FCP struct {
	sequences []sequence
}

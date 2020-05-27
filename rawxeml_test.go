package converter

import (
	"encoding/xml"
	"testing"
)

func TestImportingAClip(t *testing.T) {
	xc := `<?xml version="1.0" encoding="UTF-8"?>
		<!DOCTYPE xmeml>
		<xmeml version="1">
			<clip>
				<name>Jeremy Solo</name>
				<duration>188</duration>
				<in>5</in>
				<out>10</out>
				<masterclipid>master-foo</masterclipid>
				<ismasterclip>TRUE</ismasterclip>
				<enabled>TRUE</enabled>
				<anamorphic>FALSE</anamorphic>
				<alphatype>straight</alphatype>
				<alphareverse>TRUE</alphareverse>
				<compositemode>hardlight</compositemode>
				<stillframe>TRUE</stillframe>
				<stillframeoffset>10</stillframeoffset>
				<startoffset>20</startoffset>
				<endoffset>30</endoffset>
				<rate>
					<timebase>30</timebase>
				</rate>
			</clip>
		</xmeml>
	`

	x := ImportRawXEML([]byte(xc))

	if x.Clip.Name != "Jeremy Solo" {
		t.Error("Clip name not imported")
	}

	if x.Clip.Duration != 188 {
		t.Error("Clip duration not imported")
	}

	if x.Clip.In != 5 {
		t.Error("Clip start time not imported")
	}

	if x.Clip.Out != 10 {
		t.Error("Clip end time not imported")
	}

	if x.Clip.MasterClipID != "master-foo" {
		t.Error("Clip master ID not imported")
	}

	if x.Clip.IsMasterClip != true {
		t.Error("Clip master clip setting not imported")
	}

	if x.Clip.Enabled != true {
		t.Error("Clip enabled setting not imported")
	}

	if x.Clip.Anamorphic != false {
		t.Error("Clip anamorphic setting not imported")
	}

	if x.Clip.AlphaType != "straight" {
		t.Error("Clip alpha type setting not imported")
	}

	if x.Clip.AlphaReverse != true {
		t.Error("Clip alpha reverse setting not imported")
	}

	if x.Clip.CompositeMode != "hardlight" {
		t.Error("Clip composite mode setting not imported")
	}

	if x.Clip.StillFrame != true {
		t.Error("Clip still frame setting not imported")
	}

	if x.Clip.StillFrameOffset != 10 {
		t.Error("Clip still frame offset setting not imported")
	}

	if x.Clip.StartOffset != 20 {
		t.Error("Clip start offset setting not imported")
	}

	if x.Clip.EndOffset != 30 {
		t.Error("Clip end offset setting not imported")
	}

	if x.Clip.Rate.TimeBase != 30 {
		t.Error("Rate time base does not match expectations")
	}
}

func TestImportingASequence(t *testing.T) {
	xc := `<?xml version="1.0" encoding="UTF-8"?>
		<!DOCTYPE xmeml>
		<xmeml version="1">
			<sequence>
				<name>Example Sequence</name>
				<duration>3839</duration>
				<in>-1</in>
				<out>-1</out>
				<rate>
					<timebase>24</timebase>
				</rate>
				<timecode>
						<string>00:00:00:00</string>
						<frame>1</frame>
						<displayformat>NDF</displayformat>
						<rate>
								<timebase>24</timebase>
						</rate>
				</timecode>
			</sequence>
		</xmeml>
	`

	x := ImportRawXEML([]byte(xc))

	if x.Sequence.Name != "Example Sequence" {
		t.Error("Sequence name not imported")
	}

	if x.Sequence.Duration != 3839 {
		t.Error("Sequence duration not imported")
	}

	if x.Sequence.In != -1 {
		t.Error("Sequence start time not imported")
	}

	if x.Sequence.Out != -1 {
		t.Error("Sequence end time not imported")
	}

	if x.Sequence.Rate.TimeBase != 24 {
		t.Error("Sequence rate time base not imported")
	}

	if x.Sequence.TimeCode.TimeCodeString != "00:00:00:00" {
		t.Error("Sequence time code string not imported")
	}

	if x.Sequence.TimeCode.Frame != 1 {
		t.Error("Sequence time code frame not imported")
	}

	if x.Sequence.TimeCode.DisplayFormat != "NDF" {
		t.Error("Sequence time code format not imported")
	}

	if x.Sequence.TimeCode.Rate.TimeBase != 24 {
		t.Error("Sequence time code rate time base not imported")
	}
}

func TestXEMLVersionImport(t *testing.T) {
	xc := `<xmeml version="3"></xmeml>`

	xmeml := ImportRawXEML([]byte(xc))

	if xmeml.Version != 3 {
		t.Error("XEML version was not imported correctly")
	}
}

func TestIntValueImport(t *testing.T) {
	s := `
		<parameter>
			<value>100</value>
		</parameter>
	`
	var xs Parameter

	err := xml.Unmarshal([]byte(s), &xs)
	if err != nil {
		t.Error("parameter source could not be imported: " + err.Error())
	}

	if xs.Value.GetNumber() != 100 {
		t.Error("value number not imported")
	}
}

func TestBoolValueImport(t *testing.T) {
	s := `
		<parameter>
			<value>TRUE</value>
		</parameter>
	`
	var xs Parameter

	err := xml.Unmarshal([]byte(s), &xs)
	if err != nil {
		t.Error("parameter source could not be imported: " + err.Error())
	}

	if xs.Value.GetBool() != true {
		t.Error("value boolean not imported")
	}
}

func TestColorValueImport(t *testing.T) {
	s := `
		<parameter>
			<value>
				<red>3</red>
				<blue>5</blue>
				<green>7</green>
				<alpha>9</alpha>
			</value>
		</parameter>
	`
	var xs Parameter

	err := xml.Unmarshal([]byte(s), &xs)
	if err != nil {
		t.Error("parameter source could not be imported: " + err.Error())
	}

	color := xs.Value.GetColor()

	if color.Red != 3 {
		t.Error("red color value not imported")
	}

	if color.Blue != 5 {
		t.Error("blue color value not imported")
	}

	if color.Green != 7 {
		t.Error("green color value not imported")
	}

	if color.Alpha != 9 {
		t.Error("alpha color value not imported")
	}
}

func TestPositionValueImport(t *testing.T) {
	s := `
		<parameter>
			<value>
				<horiz>3</horiz>
				<vert>5</vert>
			</value>
		</parameter>
	`
	var xs Parameter

	err := xml.Unmarshal([]byte(s), &xs)
	if err != nil {
		t.Error("parameter source could not be imported: " + err.Error())
	}

	position := xs.Value.GetPosition()

	if position.Horiz != 3 {
		t.Error("horizontal position value not imported")
	}

	if position.Vert != 5 {
		t.Error("vertical position value not imported")
	}
}

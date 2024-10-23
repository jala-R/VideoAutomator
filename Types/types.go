package types

import (
	"fmt"
	"reflect"
)

type Xmeml struct {
	XMLName  string   `xml:"xmeml"`
	Version  string   `xml:"version,attr"`
	Sequence Sequence `xml:"sequence"`
}

var defautXmemlOptions = Xmeml{
	Version: "4",
}

func NewXmeml(opt *Xmeml) *Xmeml {
	result := &Xmeml{}

	if opt.Version == "" {
		result.Version = defautXmemlOptions.Version
	}

	return result
}

var defautSequenceOptions = Sequence{
	Id:                                      "sequence",
	TL_SQAudioVisibleBase:                   "0",
	TL_SQVideoVisibleBase:                   "0",
	TL_SQVisibleBaseTime:                    "19964966382570",
	TL_SQAVDividerPosition:                  "0.5",
	TL_SQHideShyTracks:                      "0",
	TL_SQHeaderWidth:                        "236",
	TL_SQDataTrackViewControlState:          "1",
	Monitor_ProgramZoomOut:                  "190897452345600",
	Monitor_ProgramZoomIn:                   "0",
	TL_SQTimePerPixel:                       "0.12206059932398784",
	MZ_EditLine:                             "34326452160000",
	MZ_Sequence_PreviewFrameSizeHeight:      "1080",
	MZ_Sequence_PreviewFrameSizeWidth:       "1920",
	MZ_Sequence_AudioTimeDisplayFormat:      "200",
	MZ_Sequence_PreviewRenderingClassID:     "1061109567",
	MZ_Sequence_PreviewRenderingPresetCodec: "1634755443",
	MZ_Sequence_PreviewRenderingPresetPath:  `EncoderPresets\SequencePreview\9678af98-a7b7-4bdb-b477-7ac9c8df4a4e\QuickTime.epr`,
	MZ_Sequence_PreviewUseMaxRenderQuality:  "false",
	MZ_Sequence_PreviewUseMaxBitDepth:       "false",
	MZ_Sequence_EditingModeGUID:             "9678af98-a7b7-4bdb-b477-7ac9c8df4a4e",
	MZ_Sequence_VideoTimeDisplayFormat:      "102",
	MZ_WorkOutPoint:                         "219435023808000",
	MZ_WorkInPoint:                          "212061193344000",
	ExplodedTracks:                          "true",
}

type Sequence struct {
	Id                                      string      `xml:"id,attr"`
	TL_SQAudioVisibleBase                   string      `xml:"TL.SQAudioVisibleBase,attr"`
	TL_SQVideoVisibleBase                   string      `xml:"TL.SQVideoVisibleBase,attr"`
	TL_SQVisibleBaseTime                    string      `xml:"TL.SQVisibleBaseTime,attr"`
	TL_SQAVDividerPosition                  string      `xml:"TL.SQAVDividerPosition,attr"`
	TL_SQHideShyTracks                      string      `xml:"TL.SQHideShyTracks,attr"`
	TL_SQHeaderWidth                        string      `xml:"TL.SQHeaderWidth,attr"`
	TL_SQDataTrackViewControlState          string      `xml:"TL.SQDataTrackViewControlState,attr"`
	Monitor_ProgramZoomOut                  string      `xml:"Monitor.ProgramZoomOut,attr"`
	Monitor_ProgramZoomIn                   string      `xml:"Monitor.ProgramZoomIn,attr"`
	TL_SQTimePerPixel                       string      `xml:"TL.SQTimePerPixel,attr"`
	MZ_EditLine                             string      `xml:"MZ.EditLine,attr"`
	MZ_Sequence_PreviewFrameSizeHeight      string      `xml:"MZ.Sequence.PreviewFrameSizeHeight,attr"`
	MZ_Sequence_PreviewFrameSizeWidth       string      `xml:"MZ.Sequence.PreviewFrameSizeWidth,attr"`
	MZ_Sequence_AudioTimeDisplayFormat      string      `xml:"MZ.Sequence.AudioTimeDisplayFormat,attr"`
	MZ_Sequence_PreviewRenderingClassID     string      `xml:"MZ.Sequence.PreviewRenderingClassID,attr"`
	MZ_Sequence_PreviewRenderingPresetCodec string      `xml:"MZ.Sequence.PreviewRenderingPresetCodec,attr"`
	MZ_Sequence_PreviewRenderingPresetPath  string      `xml:"MZ.Sequence.PreviewRenderingPresetPath,attr"`
	MZ_Sequence_PreviewUseMaxRenderQuality  string      `xml:"MZ.Sequence.PreviewUseMaxRenderQuality,attr"`
	MZ_Sequence_PreviewUseMaxBitDepth       string      `xml:"MZ.Sequence.PreviewUseMaxBitDepth,attr"`
	MZ_Sequence_EditingModeGUID             string      `xml:"MZ_Sequence_EditingModeGUID,attr"`
	MZ_Sequence_VideoTimeDisplayFormat      string      `xml:"MZ.Sequence.VideoTimeDisplayFormat,attr"`
	MZ_WorkOutPoint                         string      `xml:"MZ.WorkOutPoint,attr"`
	MZ_WorkInPoint                          string      `xml:"MZ.WorkInPoint,attr"`
	ExplodedTracks                          string      `xml:"explodedTracks,attr"`
	Uuid                                    string      `xml:"uuid"`
	Duration                                int64       `xml:"duration"`
	Rate                                    Rate        `xml:"rate"`
	Name                                    string      `xml:"name"`
	Media                                   Media       `xml:"media"`
	Timecode                                Timecode    `xml:"timecode"`
	Label                                   Label       `xml:"labels"`
	LoggingInfo                             LoggingInfo `xml:"logginginfo"`
}

type Rate struct {
	Timebase int64  `xml:"timebase"`
	Ntsc     string `xml:"ntsc"`
}

var defaultRate = Rate{
	30,
	"TRUE",
}

type Timecode struct {
	Rate          Rate   `xml:"rate"`
	String        string `xml:"string"`
	Frame         int64  `xml:"frame"`
	DispalyFormat string `xml:"displayformat"`
}

type Label struct {
	Label2 string `xml:"label2,omitempty"`
}

var defaultLabel = Label{
	"Forest",
}

type LoggingInfo struct {
	Description           string `xml:"description"`
	Scene                 string `xml:"scene"`
	Shottake              string `xml:"shottake"`
	Lognote               string `xml:"lognote"`
	Good                  string `xml:"good"`
	Originalvideofilename string `xml:"originalvideofilename"`
	Originalaudiofilename string `xml:"originalaudiofilename"`
}

var defautLoggingInfo = LoggingInfo{}

type Media struct {
	Video Video `xml:"video,omitempty"`
	Audio Audio `xml:"audio,omitempty"`
}

type Video struct {
	Format                Format                `xml:"format,omitempty"`
	Track                 []Track               `xml:"track,omitempty"`
	Samplecharacteristics Samplecharacteristics `xml:"samplecharacteristics,omitempty"`
}

type Audio struct {
	NumOutputChannels int     `xml:"numOutputChannels"`
	Format            Format  `xml:"format"`
	Outputs           Outputs `xml:"outputs"`
	Track             []Track `xml:"track"`
}

type Format struct {
	Samplecharacteristics Samplecharacteristics `xml:"samplecharacteristics"`
}

type Outputs struct {
	Group []Group `xml:"group"`
}
type Group struct {
	Index         int `xml:"index"`
	Numchannel    int `xml:"numchannels"`
	Downmix       int `xml:"downmix"`
	IChannelIndex int `xml:"channel>index"`
}

type Track struct {
	TL_SQTrackShy            string     `xml:"TL.SQTrackShy,attr"`
	TL_SQTrackExpandedHeight string     `xml:"TL.SQTrackExpandedHeight,attr"`
	TL_SQTrackExpanded       string     `xml:"TL.SQTrackExpanded,attr"`
	MZ_TrackTargeted         string     `xml:"MZ.TrackTargeted,attr"`
	Clipitems                []Clipitem `xml:"clipitem"`
	Enabled                  string     `xml:"enabled"`
	Locked                   string     `xml:"locked"`
}

type Clipitem struct {
	Id               string      `xml:"id,attr"`
	MasterClipId     string      `xml:"masterclipid"`
	Name             string      `xml:"name"`
	Enabled          string      `xml:"enabled"`
	Duration         int64       `xml:"duration"`
	Rate             Rate        `xml:"rate"`
	Start            int         `xml:"start"`
	End              int         `xml:"end"`
	In               int64       `xml:"in"`
	Out              int64       `xml:"out"`
	PproTicksIn      string      `xml:"pproTicksIn"`
	PproTicksOut     string      `xml:"pproTicksOut"`
	Alphatype        string      `xml:"alphatype"`
	Pixelaspectratio string      `xml:"pixelaspectratio"`
	Anamorphic       string      `xml:"anamorphic"`
	File             File        `xml:"file"`
	Filter           Filter      `xml:"filter"`
	LoggingInfo      LoggingInfo `xml:"logginginfo"`
	ColourInfo       ColourInfo  `xml:"colorinfo"`
	Label            Label       `xml:"labels"`
}

type ColourInfo struct {
	Lut    string `xml:"lut"`
	Lut1   string `xml:"lut1"`
	AscSop string `xml:"asc_sop"`
	AscSat string `xml:"asc_sat"`
	Lut2   string `xml:"lut2"`
}

type File struct {
	Id       string   `xml:"id,attr"`
	Name     string   `xml:"name"`
	PathUrl  string   `xml:"pathurl"`
	Rate     Rate     `xml:"rate"`
	Timecode Timecode `xml:"timecode"`
	Media    Media    `xml:"media"`
}

type Filter struct {
	Effect Effect `xml:"effect"`
}

type Effect struct {
	Name           string      `xml:"name"`
	EffectId       string      `xml:"effectid"`
	Effectcategory string      `xml:"effectcategory"`
	Effecttype     string      `xml:"effecttype"`
	Mediatype      string      `xml:"mediatype"`
	PproBypass     bool        `xml:"pproBypass"`
	Parameter      []Parameter `xml:"parameter"`
}

type Parameter struct {
	AuthoringApp string  `xml:"authoringApp,attr"`
	Parameterid  string  `xml:"parameterid"`
	Name         string  `xml:"name"`
	Valuemin     int     `xml:"valuemin"`
	Valuemax     int     `xml:"valuemax"`
	Value        float64 `xml:"value"`
}

type Codec struct {
	Name            string `xml:"name"`
	AppName         string `xml:"appspecificdata>appname"`
	Appmanufacturer string `xml:"appspecificdata>appmanufacturer"`
	Appversion      string `xml:"appspecificdata>appversion"`
	Codecname       string `xml:"appspecificdata>data>qtcodec>codecname"`
	Codectypename   string `xml:"appspecificdata>data>qtcodec>codectypename"`
	Codectypecode   string `xml:"appspecificdata>data>qtcodec>codectypecode"`
	Codecvendorcode string `xml:"appspecificdata>data>qtcodec>codecvendorcode"`
	Spatialquality  string `xml:"appspecificdata>data>qtcodec>spatialquality"`
	Temporalquality string `xml:"appspecificdata>data>qtcodec>temporalquality"`
	Keyframerate    string `xml:"appspecificdata>data>qtcodec>keyframerate"`
	Datarate        string `xml:"appspecificdata>data>qtcodec>datarate"`
}

type Samplecharacteristics struct {
	Rate             Rate   `xml:"rate,omitempty"`
	Codec            Codec  `xml:"codec,omitempty"`
	Width            int    `xml:"width,omitempty"`
	Height           int    `xml:"height,omitempty"`
	Anamorphic       string `xml:"anamorphic,omitempty"`
	Pixelaspectratio string `xml:"pixelaspectratio,omitempty"`
	Fielddominance   string `xml:"fielddominance,omitempty"`
	Colordepth       string `xml:"colordepth,omitempty"`
	Depth            int    `xml:"depth,omitempty"`
	SampleRate       int    `xml:"samplerate,omitempty"`
}

func Constructor(ip Samplecharacteristics) Samplecharacteristics {
	var op Samplecharacteristics

	val := reflect.ValueOf(ip)

	for i := 0; i < val.NumField(); i++ {
		feildVal := val.Field(i)
		fmt.Println(feildVal)
	}

	return op
}

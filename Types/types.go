package types

import (
	"fmt"
	"reflect"
)

var deafultValues = map[reflect.Type]interface{}{
	reflect.TypeOf(Xmeml{}): Xmeml{
		Version: "4",
	},
	reflect.TypeOf(Sequence{}): Sequence{
		Id:                                      "sequence-1",
		TL_SQAudioVisibleBase:                   "0",
		TL_SQVideoVisibleBase:                   "0",
		TL_SQVisibleBaseTime:                    "0",
		TL_SQAVDividerPosition:                  "0.5",
		TL_SQHideShyTracks:                      "0",
		TL_SQHeaderWidth:                        "236",
		TL_SQDataTrackViewControlState:          "1",
		Monitor_ProgramZoomOut:                  "5373573004800",
		Monitor_ProgramZoomIn:                   "0",
		TL_SQTimePerPixel:                       "0.01196542318491731",
		MZ_EditLine:                             "2542700160000",
		MZ_Sequence_PreviewFrameSizeHeight:      "1168",
		MZ_Sequence_PreviewFrameSizeWidth:       "876",
		MZ_Sequence_AudioTimeDisplayFormat:      "200",
		MZ_Sequence_PreviewRenderingClassID:     "1061109567",
		MZ_Sequence_PreviewRenderingPresetCodec: "1634755443",
		MZ_Sequence_PreviewRenderingPresetPath:  `EncoderPresets\SequencePreview\9678af98-a7b7-4bdb-b477-7ac9c8df4a4e\QuickTime.epr`,
		MZ_Sequence_PreviewUseMaxRenderQuality:  "false",
		MZ_Sequence_PreviewUseMaxBitDepth:       "false",
		MZ_Sequence_EditingModeGUID:             "9678af98-a7b7-4bdb-b477-7ac9c8df4a4e",
		MZ_Sequence_VideoTimeDisplayFormat:      "102",
		MZ_WorkOutPoint:                         "1262874240000",
		MZ_WorkInPoint:                          "212061193344000",
		ExplodedTracks:                          "true",
	},
	reflect.TypeOf(Rate{}): Rate{
		30,
		"TRUE",
	},
	reflect.TypeOf(Timecode{}): Timecode{
		String:        "00:00:00:00",
		Frame:         0,
		DispalyFormat: "NDF",
	},
	reflect.TypeOf(Label{}): Label{
		Label2: "Caribbean",
	},
	reflect.TypeOf(LoggingInfo{}): LoggingInfo{},
	reflect.TypeOf(Media{}):       Media{},
	reflect.TypeOf(Video{}):       Video{},
	reflect.TypeOf(Audio{}):       Audio{},
	reflect.TypeOf(Format{}):      Format{},
	reflect.TypeOf(Outputs{}):     Outputs{},
	reflect.TypeOf(Group{}):       Group{},
	reflect.TypeOf(Track{}):       Track{},
	reflect.TypeOf(Clipitem{}):    Clipitem{},
	reflect.TypeOf(ColourInfo{}):  ColourInfo{},
	reflect.TypeOf(File{}):        File{},
	reflect.TypeOf(Filter{}):      Filter{},
	reflect.TypeOf(Effect{}):      Effect{},
	reflect.TypeOf(Parameter{}):   Parameter{},
	reflect.TypeOf(Codec{}): Codec{
		Name:            "Apple ProRes 422",
		AppName:         "Final Cut Pro",
		Appmanufacturer: "Apple Inc.",
		Appversion:      "7.0",
		Codecname:       "Apple ProRes 422",
		Codectypename:   "Apple ProRes 422",
		Codectypecode:   "apcn",
		Codecvendorcode: "appl",
		Spatialquality:  "1024",
		Temporalquality: "0",
		Keyframerate:    "0",
		Datarate:        "0",
	},
	reflect.TypeOf(Samplecharacteristics{}): Samplecharacteristics{
		Width:            1920,
		Height:           1080,
		Anamorphic:       "FALSE",
		Pixelaspectratio: "square",
		Fielddominance:   "none",
		Colordepth:       "24",
	},
}

type Xmeml struct {
	XMLName  string   `xml:"xmeml"`
	Version  string   `xml:"version,attr"`
	Sequence Sequence `xml:"sequence"`
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

type Timecode struct {
	Rate          Rate   `xml:"rate"`
	String        string `xml:"string"`
	Frame         int64  `xml:"frame"`
	DispalyFormat string `xml:"displayformat"`
}

type Label struct {
	Label2 string `xml:"label2,omitempty"`
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

type Media struct {
	Video Video `xml:"video,omitempty"`
	Audio Audio `xml:"audio,omitempty"`
}

type Video struct {
	Format                *Format                `xml:"format,omitempty"`
	Track                 []Track                `xml:"track,omitempty"`
	Samplecharacteristics *Samplecharacteristics `xml:"samplecharacteristics,omitempty"`
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
	Codec            *Codec `xml:"codec,omitempty"`
	Width            int    `xml:"width,omitempty"`
	Height           int    `xml:"height,omitempty"`
	Anamorphic       string `xml:"anamorphic,omitempty"`
	Pixelaspectratio string `xml:"pixelaspectratio,omitempty"`
	Fielddominance   string `xml:"fielddominance,omitempty"`
	Colordepth       string `xml:"colordepth,omitempty"`
	Depth            int    `xml:"depth,omitempty"`
	SampleRate       int    `xml:"samplerate,omitempty"`
}

func Constructor(ip ...any) {

	if len(ip) > 2 || len(ip) == 0 {
		panic(fmt.Sprintf("expecpted 1 or 2 got %d", len(ip)))
	}
	if _, ok := ip[0].(reflect.Value); !ok && reflect.ValueOf(ip[0]).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Got %s, Need %s", reflect.ValueOf(ip[0]).Kind(), reflect.Ptr))
	}

	var ogValue reflect.Value
	var ogType reflect.Type

	if temp, ok := ip[0].(reflect.Value); ok {
		ogValue = temp.Elem()
		ogType = ogValue.Type()
	} else {
		ogValue = reflect.ValueOf(ip[0]).Elem()
		ogType = ogValue.Type()

	}

	var defaultValue reflect.Value

	if len(ip) == 2 {
		defaultValue = reflect.ValueOf((ip[1]))
	} else {
		temp := deafultValues[ogType]
		defaultValue = reflect.ValueOf(temp)
	}

	// var typ reflect.Type
	// var val reflect.Value

	// if temp, ok := ip.(reflect.Value); ok {
	// 	val = temp.Elem()
	// 	typ = val.Type()
	// } else {
	// 	val = reflect.ValueOf(ip).Elem()
	// 	typ = val.Type()

	// }

	// defaultObj := deafultValues[typ]
	// defaultObjVal := reflect.ValueOf(defaultObj)

	deepCopyDefaults(ogValue, defaultValue)

	// if typ.String() != defaultObjVal.Type().String() {
	// 	panic("Type in default args are in compatiable")
	// }

	// // fmt.Println(val)

	// // fmt.Println(defaultObjVal)

	// for i := 0; i < typ.NumField(); i++ {
	// 	structFeild := val.Field(i)
	// 	defaultStrutObj := defaultObjVal.Field(i)

	// 	if structFeild.IsZero() {
	// 		if structFeild.Kind() == reflect.Struct {
	// 			Constructor(structFeild.Addr())
	// 		} else {
	// 			structFeild.Set(defaultStrutObj)
	// 		}
	// 	}
	// }

}

func deepCopyDefaults(obj, defaultObj reflect.Value) {
	if obj.Type() != defaultObj.Type() {
		panic("Type in default args are in compatiable")
	}

	// fmt.Println(obj.Type().Name())

	// fmt.Println(val)

	// fmt.Println(defaultObjVal)

	for i := 0; i < obj.Type().NumField(); i++ {
		structFeild := obj.Field(i)
		defaultStrutObj := defaultObj.Field(i)

		if (structFeild.Kind() == reflect.Ptr && structFeild.IsNil()) || structFeild.IsZero() || structFeild.Kind() == reflect.Struct {
			if structFeild.Kind() == reflect.Struct {
				Constructor(structFeild.Addr())
			} else if structFeild.Kind() == reflect.Ptr {
				typ := (structFeild.Type().Elem())
				addr := reflect.New(typ).Elem().Addr()
				Constructor(addr)
				structFeild.Set(addr)
			} else {
				structFeild.Set(defaultStrutObj)
			}
		}
	}
}

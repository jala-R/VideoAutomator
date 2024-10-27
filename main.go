package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	types "jalar.me/xml/Types"
)

const XMLHEADER = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE xmeml>
`

type Temp struct {
}

func main() {

	xeml := types.Xmeml{}
	types.Constructor(&xeml)
	xeml.Sequence.Uuid = uuid.NewString()
	frameRate := 30

	xeml.Sequence.Label.Label2 = "Forest"
	xeml.Sequence.Timecode.DispalyFormat = "DF"

	clipOneDuration := 4.96 * float64(frameRate)

	xeml.Sequence.Duration = int64(clipOneDuration)
	xeml.Sequence.Name = "Test1 project"

	tickConversion := 8475667200

	clip1 := types.Clipitem{}
	clip1.Id = "clipitem-1"
	clip1.MasterClipId = "masterclip-1"
	clip1.Name = "Test1 clippp"
	clip1.Enabled = "TRUE"
	clip1.Duration = 1294705
	clip1.Start = 0
	clip1.End = int(clipOneDuration)
	clip1.In = 107892
	clip1.Out = clip1.In + (int64(clip1.End) - int64(clip1.Start))
	clip1.PproTicksIn = strconv.Itoa(int(clip1.In) * int(tickConversion))
	clip1.PproTicksOut = strconv.Itoa(int(clip1.Out) * int(tickConversion))
	// fmt.Println(strconv.FormatUint())
	clip1.Alphatype = "none"
	clip1.Pixelaspectratio = "square"
	clip1.Anamorphic = "FALSE"

	ClipFile := types.File{
		Id:      "file-1",
		Name:    "File 1 name",
		PathUrl: `file:///D:/Etheral%20SEEkers/Pictures/new%20oilpainted%20pics/1.jpg`,
	}

	ClipFile.Media.Video.Samplecharacteristics = &types.Samplecharacteristics{
		Width:  876,
		Height: 1168,
	}

	// types.Constructor(ClipFile.Media.Video.Samplecharacteristics)

	types.Constructor(&ClipFile)
	types.Constructor(&clip1)
	clip1.File = ClipFile

	// xeml.Sequence.Media.Video.Format.Samplecharacteristics.Colordepth = ""
	xeml.Sequence.Media.Video.Track = []types.Track{
		{
			TL_SQTrackShy:            "0",
			TL_SQTrackExpandedHeight: "25",
			TL_SQTrackExpanded:       "0",
			MZ_TrackTargeted:         "1",
			Enabled:                  "TRUE",
			Locked:                   "FALSE",
			Clipitems:                []types.Clipitem{clip1},
		},
	}

	xmlData, _ := xml.MarshalIndent(xeml, "", "	")

	file, err := os.Create("./output1.xml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file.Write([]byte(XMLHEADER))
	file.Write(xmlData)
	file.Close()
}

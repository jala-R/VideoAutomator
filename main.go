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

func main() {

	a := types.Xmeml{}
	types.Constructor(&a)

	a.Sequence.Uuid = uuid.NewString()
	a.Sequence.Duration = 434
	a.Sequence.Name = "Test Trail 1"
	a.Sequence.Label.Label2 = "Forest"
	a.Sequence.Media.Video.Samplecharacteristics = nil

	//set video videoTracks
	var videoTracks []types.Track

	var v1 types.Track = types.Track{}
	var v2 types.Track = types.Track{}
	var v3 types.Track = types.Track{}
	types.Constructor(&v1)
	types.Constructor(&v2)
	types.Constructor(&v3)
	v1.MZ_TrackTargeted = "1"
	v1.Clipitems = getVideoClipItems()

	videoTracks = append(videoTracks, v1, v2, v3)
	a.Sequence.Media.Video.Track = videoTracks

	//set audio tracks
	// var audioTracks []types.Track

	a.Sequence.Media.Audio.NumOutputChannels = 2
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Rate = nil
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Codec = nil
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Width = 0
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Height = 0
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Anamorphic = ""
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Pixelaspectratio = ""
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Fielddominance = ""
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Colordepth = ""
	a.Sequence.Media.Audio.Format.Samplecharacteristics.Depth = 16
	a.Sequence.Media.Audio.Format.Samplecharacteristics.SampleRate = 44100

	a.Sequence.Media.Audio.Outputs = &types.Outputs{
		Group: []types.Group{
			GetGroups(1),
			GetGroups(2),
		},
	}

	//Add Audio Tracks

	for i := 1; i <= 1; i++ {
		temp := GetAudioTrack(i)
		types.Constructor(&temp)
		a.Sequence.Media.Audio.Track = append(a.Sequence.Media.Audio.Track, temp)
	}
	a.Sequence.Media.Audio.Samplecharacteristics = nil
	a.Sequence.Media.Audio.Audiochannel = nil

	xmlData, _ := xml.MarshalIndent(a, "", "	")
	xmlData = append(xmlData, byte('\n'))

	file, err := os.Create("./output1.xml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file.Write([]byte(XMLHEADER))
	file.Write(xmlData)
	file.Close()
}

// func GetVideoTrack() types.Track{
// 	var track = types.Track{}
// 	// track.TL_SQVisibleBaseTime
// }

func GetAudioTrack(ind int) types.Track {
	var track = types.Track{}
	types.Constructor(&track)
	track.TL_SQTrackAudioKeyframeStyle = "0"
	track.TL_SQTrackShy = "0"
	track.TL_SQTrackExpandedHeight = "25"
	track.TL_SQTrackAudioKeyframeStyle = "0"
	track.MZ_TrackTargeted = "1"
	track.PannerCurrentValue = "0.5"
	track.PannerStartKeyframe = "-91445760000000000,0.5,0,0,0,0,0,0"
	track.PannerName = "Balance"
	track.TotalExplodedTrackCount = "1"
	track.TL_SQTrackExpanded = "0"
	track.CurrentExplodedTrackIndex = fmt.Sprintf("%d", ind-1)
	track.PremiereTrackType = "Stereo"
	track.Outputchannelindex = ind
	track.Clipitems = GetAudioClipItems()
	return track
}

func GetAudioClipItems() []types.Clipitem {
	var clips []types.Clipitem
	clips = append(clips, getAudioClip(4, "fin.wav", 0, 91))
	clips = append(clips, getAudioClip(5, "fin.wav", 91, 240))
	clips = append(clips, getAudioClip(6, "fin.wav", 269, 434))
	return clips
}

func getAudioClip(ind int, fileName string, start, end int) types.Clipitem {
	var clip = types.Clipitem{}
	types.Constructor(&clip)
	var tickMeasure = 8475667200
	clip.Id = fmt.Sprintf("clipitem-%d", ind)
	clip.PremiereChannelType = "mono"
	clip.MasterClipId = fmt.Sprintf("masterclip-%d", ind)
	clip.Name = fileName
	clip.Duration = int64(49616)
	clip.Start = start
	clip.End = end
	clip.In = int64(start)
	clip.Out = int64(end)
	clip.PproTicksIn = strconv.FormatUint(uint64(clip.Start)*uint64(tickMeasure), 10)
	clip.PproTicksOut = strconv.FormatUint(uint64(clip.End)*uint64(tickMeasure), 10)
	clip.Alphatype = ""
	clip.Pixelaspectratio = ""
	clip.Anamorphic = ""
	clip.File = getAudioFile(ind, fileName)
	clip.Filter = nil
	clip.Sourcetrack.Mediatype = "audio"
	clip.Sourcetrack.Trackindex = 1
	clip.Label = nil
	return clip
}

func getAudioFile(ind int, fileName string) types.File {
	var file = types.File{}
	types.Constructor(&file)

	file.Id = fmt.Sprintf("file-%d", ind)
	file.Name = fileName
	file.PathUrl = fmt.Sprintf("file://localhost/E%%3a/Spirituality/Ethereal%%20seekers/61)Why%%20chosen%%20one%%20Become%%20Monsters/%s", fileName)
	file.Duration = 49616
	file.Media.Video = nil
	file.Media.Audio.Samplecharacteristics.Depth = 16
	file.Media.Audio.Samplecharacteristics.SampleRate = 44100
	file.Media.Audio.Samplecharacteristics.Rate = nil
	file.Media.Audio.Samplecharacteristics.Codec = nil
	file.Media.Audio.Samplecharacteristics.Width = 0
	file.Media.Audio.Samplecharacteristics.Height = 0
	file.Media.Audio.Samplecharacteristics.Anamorphic = ""
	file.Media.Audio.Samplecharacteristics.Pixelaspectratio = ""
	file.Media.Audio.Samplecharacteristics.Fielddominance = ""
	file.Media.Audio.Samplecharacteristics.Colordepth = ""
	file.Media.Audio.Channelcount = 1
	file.Media.Audio.Audiochannel.Sourcechannel = 1
	file.Media.Audio.Format = nil
	file.Media.Audio.Outputs = nil
	return file
}

func getVideoClipItems() []types.Clipitem {

	var clips []types.Clipitem
	var clip1 = GetVideoClipItem(1, "Clip1", 0, 91)
	var clip2 = GetVideoClipItem(2, "Clip2", 91, 240)
	var clip3 = GetVideoClipItem(3, "Clip3", 269, 434)

	clips = append(clips, clip1, clip2, clip3)
	return clips
}

func GetVideoClipItem(clipID int, name string, start int, end int) types.Clipitem {

	var clipIn int64 = 107892
	var tickMeasure = 8475667200
	var clip = types.Clipitem{}
	types.Constructor(&clip)
	clip.Id = fmt.Sprintf("clipitem-%d", clipID)
	clip.MasterClipId = fmt.Sprintf("masterclip-%d", clipID)
	clip.Name = name
	clip.Enabled = "TRUE"
	clip.Duration = 1294705
	clip.Start = start
	clip.End = end
	clip.In = clipIn
	clip.Out = clip.In + int64(clip.End) - int64(clip.Start)
	clip.PproTicksIn = strconv.FormatUint(uint64(tickMeasure)*uint64(clip.In), 10)
	clip.PproTicksOut = strconv.FormatUint(uint64(tickMeasure)*uint64(clip.Out), 10)
	clip.Label.Label2 = "Lavender"
	clip.File = getVideoFile(clipID)
	clip.Sourcetrack = nil
	clip.Filter = nil

	return clip
}

func getVideoFile(fileId int) types.File {
	var file types.File = types.File{}
	types.Constructor(&file)

	file.Id = fmt.Sprintf("file-%d", fileId)
	file.Name = fmt.Sprintf("%d.png", fileId)
	file.PathUrl = fmt.Sprintf("file://localhost/D%%3a/Etheral%%20SEEkers/Pictures/new%%20oilpainted%%20pics/%d.png", fileId)
	file.Media = types.Media{}
	types.Constructor(&file.Media)
	file.Media.Audio = nil
	file.Media.Video.Format = nil
	file.Media.Video.Samplecharacteristics.Codec = nil
	file.Media.Video.Samplecharacteristics.Colordepth = ""

	FileTimeCode(&file.Timecode)
	return file
}

func FileTimeCode(timecode *types.Timecode) {
	timecode.String = "00:00:00:00"
	timecode.DispalyFormat = "NDF"
}

func GetGroups(ind int) types.Group {
	var group = types.Group{}
	group.Index = ind
	group.Numchannel = 1
	group.Downmix = 0
	group.IChannelIndex = ind
	return group
}

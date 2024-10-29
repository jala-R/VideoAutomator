package main

import (
	"encoding/xml"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/hajimehoshi/go-mp3"
	types "jalar.me/xml/Types"
)

const XMLHEADER = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE xmeml>
`

func convertSecToSecFrames(seconds float64) float64 {
	totalFramesFullFrames := math.Floor(seconds * 29.97)
	sec := math.Floor(totalFramesFullFrames / 29.97)
	frame := math.Floor(totalFramesFullFrames - sec*29.97)

	return sec + frame/100
}

func GetDurations() []float64 {
	fileDirs, err := os.ReadDir("./audio")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	sort.Slice(fileDirs, func(i, j int) bool {
		f1, _ := strconv.ParseInt(strings.Split(fileDirs[i].Name(), ".")[0], 10, 64)
		f2, _ := strconv.ParseInt(strings.Split(fileDirs[j].Name(), ".")[0], 10, 64)

		return f1 < f2
	})

	var clipsDuration = make([]float64, 0, len(fileDirs))

	var cnt = 5

	for _, name := range fileDirs {
		if cnt == 0 {
			break
		}
		cnt--
		fileName := fmt.Sprintf("./audio/%s", name.Name())

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err.Error(), fileName)
			continue
		}
		defer file.Close()

		decoder, err := mp3.NewDecoder(file)
		if err != nil {
			fmt.Println(err.Error(), fileName)
			continue
		}

		length := decoder.Length()

		samplePerSec := decoder.SampleRate()

		// length := decoder.Length()

		totalSample := float64(length) / 4

		time := (totalSample) / float64(samplePerSec)

		clipsDuration = append(clipsDuration, convertSecToSecFrames(time))
		fmt.Println(time, convertSecToSecFrames(time), fileName)
	}

	return clipsDuration

}

func main() {

	timeClips := GetDurations()
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
	v1.Clipitems = getVideoClipItems(timeClips)

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
		temp := GetAudioTrack(i, timeClips)
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

func GetAudioTrack(ind int, timings []float64) types.Track {
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
	track.Clipitems = GetAudioClipItems(timings)
	return track
}

func GetAudioClipItems(timings []float64) []types.Clipitem {
	var clips []types.Clipitem

	var prev int

	for i := range timings {
		duration := convertDuration(timings[i])
		left := prev
		right := left + int(duration)

		prev = right
		clips = append(clips, getAudioClip(len(timings)+i+1, fmt.Sprintf("%d.wav", i+1+4), left, right))
	}
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
	clip.In = 0
	clip.Out = int64(end - start)
	clip.PproTicksIn = strconv.FormatUint(uint64(clip.In)*uint64(tickMeasure), 10)
	clip.PproTicksOut = strconv.FormatUint(uint64(clip.Out)*uint64(tickMeasure), 10)
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

func convertDuration(seconds float64) int64 {
	temp := int(seconds * 100)
	return int64(seconds)*30 + int64(temp%100)
}

func getVideoClipItems(timings []float64) []types.Clipitem {

	var clips []types.Clipitem

	var prev int64

	for i := range timings {
		seconds := timings[i]
		durations := convertDuration(seconds)

		left := prev
		right := prev + durations

		prev = right

		clips = append(clips, GetVideoClipItem(i+1, fmt.Sprintf("Clip%d", i+1), int(left), int(right)))
	}
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

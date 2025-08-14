package lib

import (
	"encoding/xml"
	"errors"
	"reflect"
	"strings"
)

var (
	ErrSonos = struct{ ErrUnexpectedResponse, ErrInvalidIPAdress, ErrNoZonePlayerFound, ErrInvalidEndpoint, ErrTagNotFound, ErrInvalidContentType, ErrInvalidPlayMode error }{
		ErrUnexpectedResponse: errors.New("unexpected response"),
		ErrInvalidIPAdress:    errors.New("unable to discover zone player"),
		ErrNoZonePlayerFound:  errors.New("unable to find zone player"),
		ErrTagNotFound:        errors.New("tag not found in response"),
		ErrInvalidEndpoint:    errors.New("invalid endpoint"),
		ErrInvalidPlayMode:    errors.New("invalid play mode"),
	}
	ContentTypes = struct{ MusicLibrarys, Artist, AlbumArtist, Album, Genre, Composer, Tracks, Playlists, Share, SonosPlaylists, SonosFavorites, Radios, RadioStations, RadioShows, Queues, QueueMain, QueueSecond string }{
		MusicLibrarys:  "A:",
		Artist:         "A:ARTIST",
		AlbumArtist:    "A:ALBUMARTIST",
		Album:          "A:ALBUM",
		Genre:          "A:GENRE",
		Composer:       "A:COMPOSER",
		Tracks:         "A:TRACKS",
		Playlists:      "A:PLAYLISTS",
		Share:          "S:",
		SonosPlaylists: "SQ:",
		SonosFavorites: "FV:2",
		Radios:         "R:0",
		RadioStations:  "R:0/0",
		RadioShows:     "R:0/1",
		Queues:         "Q:",
		QueueMain:      "Q:0",
		QueueSecond:    "Q:1",
	}
	SeekModes = struct{ Track, Relative, Delta string }{
		Track:    "TRACK_NR",
		Relative: "REL_TIME",
		Delta:    "TIME_DELTA",
	}
	TransportStates = struct{ Stopped, Playing, PausedPlayback, Transitioning string }{
		Stopped:        "STOPPED",
		Playing:        "PLAYING",
		PausedPlayback: "PAUSED_PLAYBACK",
		Transitioning:  "TRANSITIONING",
	}
	PlayModes = struct{ Normal, RepeatAll, RepeatOne, ShuffleNorepeat, Shuffle, ShuffleRepeaOne string }{
		Normal:          "NORMAL",
		RepeatAll:       "REPEAT_ALL",
		RepeatOne:       "REPEAT_ONE",
		ShuffleNorepeat: "SHUFFLE_NOREPEAT",
		Shuffle:         "SHUFFLE",
		ShuffleRepeaOne: "SHUFFLE_REPEAT_ONE",
	}
	PlayModeMap = map[string][3]bool{
		// "PlayMode": [2]bool{shuffle, repeat, repeatOne}
		PlayModes.Normal:          {false, false, false},
		PlayModes.RepeatAll:       {false, true, false},
		PlayModes.RepeatOne:       {false, false, true},
		PlayModes.ShuffleNorepeat: {true, false, false},
		PlayModes.Shuffle:         {true, true, false},
		PlayModes.ShuffleRepeaOne: {true, false, true},
	}
	PlayModeMapReversed = func() map[[3]bool]string {
		m := map[[3]bool]string{}
		for k, v := range PlayModeMap {
			m[v] = k
		}
		return m
	}()
	RecurrenceModes = struct{ Once, Weekdays, Weekends, Daily string }{
		Once:     "ONCE",
		Weekdays: "WEEKDAYS",
		Weekends: "WEEKENDS",
		Daily:    "DAILY",
	}
	AlbumArtistDisplayOptionModes = struct{ WMP, ITunes, None string }{
		WMP:    "WMP",
		ITunes: "ITUNES",
		None:   "NONE",
	}
	UpdateTypes = struct{ All, Software string }{
		All:      "All ",
		Software: "Software",
	}
)

func UnmarshalMetaData[T any](data string, v T) error {
	data = strings.ReplaceAll(data, "&apos;", "'")
	data = strings.ReplaceAll(data, "&quot;", "\"")
	data = strings.ReplaceAll(data, "&gt;", ">")
	data = strings.ReplaceAll(data, "&lt;", "<")
	if reflect.TypeOf(v).Elem().Kind() == reflect.Slice {
		vTmp := struct {
			XMLName xml.Name `xml:"DIDL-Lite"`
			Items   T        `xml:"item"`
		}{Items: v}
		if err := xml.Unmarshal([]byte(data), &vTmp); err != nil {
			return err
		}
		return nil
	}

	data, err := ExtractTag(data, "DIDL-Lite")
	if err != nil {
		return err
	}
	if err := xml.Unmarshal([]byte(data), v); err != nil {
		return err
	}
	return nil
}

func ExtractTag(data, tag string) (string, error) {
	if start, end := strings.Index(data, "<"+tag), strings.LastIndex(data, "</"+tag+">"); start != -1 && end != -1 {
		data = data[start+len(tag) : end]
		if mid := strings.Index(data, ">"); mid != -1 {
			return data[mid+1:], nil
		}
	}
	return data, ErrSonos.ErrTagNotFound
}

func BoolTo10(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func BoolToOnOff(b bool) string {
	if b {
		return "On"
	}
	return "Off"
}

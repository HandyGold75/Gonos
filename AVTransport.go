package Gonos

import (
	"strconv"
	"time"

	"github.com/HandyGold75/Gonos/lib"
)

type (
	TrackInfo struct {
		QuePosition int
		Duration    string
		URI         string
		Progress    string
		AlbumArtURI string
		Title       string
		Class       string
		Creator     string
		Album       string
	}
)

// Get simplified info about currently playing track.
func (zp *ZonePlayer) GetTrackInfo() (TrackInfo, error) {
	info, err := zp.AVTransport.GetPositionInfo()
	if err != nil {
		return TrackInfo{}, err
	}
	return TrackInfo{
		QuePosition: info.Track,
		Duration:    info.TrackDuration,
		URI:         info.TrackURI,
		Progress:    info.RelTime,
		AlbumArtURI: info.TrackMetaDataParsed.AlbumArtUri,
		Title:       info.TrackMetaDataParsed.Title,
		Class:       info.TrackMetaDataParsed.Class,
		Creator:     info.TrackMetaDataParsed.Creator,
		Album:       info.TrackMetaDataParsed.Album,
	}, nil
}

// Return may be one of `Gonos.TransportStates.*`.
func (zp *ZonePlayer) GetCurrentTransportState() (string, error) {
	res, err := zp.AVTransport.GetTransportInfo()
	return res.CurrentTransportState, err
}

// Short for `zp.AVTransport.Stop`
func (zp *ZonePlayer) Stop() error {
	return zp.AVTransport.Stop()
}

// Short for `zp.GetCurrentTransportState() == Gonos.TransportStates.Stopped`.
func (zp *ZonePlayer) GetStop() (bool, error) {
	state, err := zp.GetCurrentTransportState()
	return state == lib.TransportStates.Stopped, err
}

// Short for `zp.AVTransport.Play`
func (zp *ZonePlayer) Play() error {
	return zp.AVTransport.Play()
}

// Short for `zp.GetCurrentTransportState() == Gonos.TransportStates.Playing`.
func (zp *ZonePlayer) GetPlay() (bool, error) {
	state, err := zp.GetCurrentTransportState()
	return state == lib.TransportStates.Playing, err
}

// Short for `zp.AVTransport.Pause`
func (zp *ZonePlayer) Pause() error {
	return zp.AVTransport.Pause()
}

// Short for `zp.GetCurrentTransportState() == Gonos.TransportStates.PausedPlayback`.
func (zp *ZonePlayer) GetPause() (bool, error) {
	state, err := zp.GetCurrentTransportState()
	return state == lib.TransportStates.PausedPlayback, err
}

// Short for `zp.GetCurrentTransportState() == Gonos.TransportStates.Transitioning`.
func (zp *ZonePlayer) GetTransitioning() (bool, error) {
	state, err := zp.GetCurrentTransportState()
	return state == lib.TransportStates.Transitioning, err
}

// Short for `zp.AVTransport.Next`.
func (zp *ZonePlayer) Next() error {
	return zp.AVTransport.Next()
}

// Short for `zp.AVTransport.`.
func (zp *ZonePlayer) Previous() error {
	return zp.AVTransport.Previous()
}

// Get current transport status.
func (zp *ZonePlayer) GetCurrentTransportStatus() (string, error) {
	res, err := zp.AVTransport.GetTransportInfo()
	return res.CurrentTransportStatus, err
}

// Get current speed.
func (zp *ZonePlayer) GetCurrentSpeed() (string, error) {
	res, err := zp.AVTransport.GetTransportInfo()
	return res.CurrentSpeed, err
}

// Will always return false for all if a third party application is controling playback.
func (zp *ZonePlayer) GetPlayMode() (shuffle bool, repeat bool, repeatOne bool, err error) {
	res, err := zp.AVTransport.GetTransportSettings()
	if err != nil {
		return false, false, false, err
	}
	modeBools, ok := lib.PlayModeMap[res.PlayMode]
	if !ok {
		return false, false, false, lib.ErrSonos.ErrUnexpectedResponse
	}
	return modeBools[0], modeBools[1], modeBools[2], nil
}

// Will always return false if a third party application is controling playback.
func (zp *ZonePlayer) GetShuffle() (bool, error) {
	shuffle, _, _, err := zp.GetPlayMode()
	return shuffle, err
}

// Will always disable other playmodes if a third party application is controling playback, as we can not determine the actual state.
func (zp *ZonePlayer) SetShuffle(state bool) error {
	_, repeat, repeatOne, err := zp.GetPlayMode()
	if err != nil {
		return err
	}
	return zp.AVTransport.SetPlayMode(state, repeat, repeatOne)
}

// Will always return false if a third party application is controling playback.
func (zp *ZonePlayer) GetRepeat() (bool, error) {
	_, repeat, _, err := zp.GetPlayMode()
	return repeat, err
}

// If enabled then repeat one will be disabled.
//
// Will always disable other playmodes if a third party application is controling playback, as we can not determine the actual state.
func (zp *ZonePlayer) SetRepeat(state bool) error {
	shuffle, _, repeatOne, err := zp.GetPlayMode()
	if err != nil {
		return err
	}
	return zp.AVTransport.SetPlayMode(shuffle, state, repeatOne && !state)
}

// Will always return false if a third party application is controling playback.
func (zp *ZonePlayer) GetRepeatOne() (bool, error) {
	_, _, repeatOne, err := zp.GetPlayMode()
	return repeatOne, err
}

// If enabled then repeat will be disabled.
//
// Will always disable other playmodes if a third party application is controling playback, as we can not determine the actual state.
func (zp *ZonePlayer) SetRepeatOne(state bool) error {
	shuffle, repeat, _, err := zp.GetPlayMode()
	if err != nil {
		return err
	}
	return zp.AVTransport.SetPlayMode(shuffle, repeat && !state, state)
}

// Returns `NOT_IMPLEMENTED`.
func (zp *ZonePlayer) GetRecQualityMode() (string, error) {
	res, err := zp.AVTransport.GetTransportSettings()
	return res.RecQualityMode, err
}

// Go to track by index (index starts at 1).
//
// Will always fail if a third party application is controling playback.
func (zp *ZonePlayer) SeekTrack(track int) error {
	return zp.AVTransport.Seek(lib.SeekModes.Track, strconv.Itoa(max(1, track)))
}

// Go to track time (Absolute).
func (zp *ZonePlayer) SeekTime(seconds int) error {
	return zp.AVTransport.Seek(lib.SeekModes.Relative, time.Time.Add(time.Time{}, time.Second*time.Duration(max(0, seconds))).Format("15:04:05"))
}

// Go to track time (Delta).
func (zp *ZonePlayer) SeekTimeDelta(seconds int) error {
	prefix := "+"
	if seconds < 0 {
		seconds = -seconds
		prefix = "-"
	}
	return zp.AVTransport.Seek(lib.SeekModes.Delta, prefix+time.Time.Add(time.Time{}, time.Second*time.Duration(seconds)).Format("15:04:05"))
}

// Short for `zp.AVTransport.RemoveTrackFromQueue`.
func (zp *ZonePlayer) QueRemove(track int) error {
	return zp.AVTransport.RemoveTrackFromQueue(lib.ContentTypes.QueueMain, track)
}

// Short for `zp.AVTransport.RemoveAllTracksFromQueue`.
func (zp *ZonePlayer) QueClear() error {
	return zp.AVTransport.RemoveAllTracksFromQueue()
}

// Short for `zp.AVTransport.AddURIToQueue`.
func (zp *ZonePlayer) QueAdd(uri string, index int, next bool) error {
	_, err := zp.AVTransport.AddURIToQueue(uri, "", index, next)
	return err
}

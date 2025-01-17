# Gonos

This module aims to implement an easy and simple way to control Sonos devices while keeping advanced control possible.

The documentation from [svrooij](https://github.com/svrooij/sonos-api-docs) can be referenced for the base implementation.
Many thanks for the documentation as without it this module would not exist.

Next to the base implementation some helper function are present for easier use.
For all helper functions please refer to the helper files.

Note that most of this project is untested.
Some functions might not work as expected.

## Usage

Creating a ZonePlayer for controlling a Sonos device can be done by any of these methods:

```go
zp, err := Gonos.NewZonePlayer("127.0.0.1") // Use the IpAddress of the Sonos device.
zp, err := Gonos.DiscoverZonePlayer() // Discover a Sonos device using SSDP.
zp, err := Gonos.ScanZonePlayer("127.0.0.0/8") // Scan a network for Sonos devices.
```

After a ZonePlayer is successfully created the associated Sonos device can be controlled.
This can be done using either the Sonos services (base implementation) or the helpers

The available Sonos services are:

- `zp.AlarmClock`
- `zp.AudioIn`
- `zp.AVTransport`
- `zp.ConnectionManager`
- `zp.ContentDirectory`
- `zp.DeviceProperties`
- `zp.GroupManagement`
- `zp.GroupRenderingControl` (Not yet implemented)
- `zp.HTControl` (Not yet implemented)
- `zp.MusicServices` (Not yet implemented)
- `zp.QPlay` (Not yet implemented)
- `zp.Queue` (Not yet implemented)
- `zp.RenderingControl`
- `zp.SystemProperties` (Not yet implemented)
- `zp.VirtualLineIn` (Not yet implemented)
- `zp.ZoneGroupTopology` (Not yet implemented)

## Examples

Some examples for controlling a Sonos device using the Sonos services:

```go
err := zp.AlarmClock.GetTimeServer() // Get the current time server.
err := zp.AudioIn.SetLineInLevel(10, 10) // Set left and right line in level to 10.
err := zp.AVTransport.Play() // Play current track.
err := zp.ConnectionManager.GetCurrentConnectionIDs() // Get ids of current connections.
queInfo, err := zp.ContentDirectory.Browse(Gonos.lib.ContentTypes.QueueMain, "BrowseDirectChildren", "dc:title,res,dc:creator,upnp:artist,upnp:album,upnp:albumArtURI", 0, 0, "") // Get info of the current main que.
zoneAttributes, err := zp.DeviceProperties.GetZoneAttributes() // Get attributes of current zone.
err := zp.GroupManagement.RemoveMember("id") // Remove a group member.
err := zp.GroupRenderingControl.GetGroupVolume() // Get the current group volume.
err := zp.HTControl.SetLEDFeedbackState(true) // Set the LED feedback state.
err := zp.MusicServices
err := zp.QPlay
err := zp.Queue
err := zp.RenderingControl.SetVolume(10) // Set volume to 10.
err := zp.SystemProperties
err := zp.VirtualLineIn
err := zp.ZoneGroupTopology
```

Some examples for controlling a Sonos device using the helpers:

```go
err := zp.H.Play() // Play current track.
isPlaying, err := zp.H.GetPlay() // Check if current track is playing.

err := zp.H.Pause() // Pause current track.
isPaues, err := zp.H.GetPause() // Check if current track is paused.

err := zp.H.Stop() // Stop current track.
isPaues, err := zp.H.GetStop() // Check if current track is stopped.

isTransitioning, err := zp.H.GetTransitioning() // Check if current track is transitioning.

err := zp.H.Next() // Go to next track.
err := zp.H.Previous() // Go to previous track.

err := zp.H.SetShuffle(true) // Enable shuffle.
isShuffle, err := zp.H.GetSuffle() // Check if shuffle is enabled.

err := zp.H.SetRepeat(true) // Enable repeat (Disables reapeat one).
isRepeat, err := zp.H.GetRepeat() // Check if repeat is enabled.

err := zp.H.SetRepeatOne(true) // Enable reapeat one (Disables repeat).
isRepeatOne, err := zp.H.GetRepeatOne() // Check if repeat one is enabled.

err := zp.H.SeekTrack(10) // Go to 10th track in the que (Count start at 1).
err := zp.H.SeekTime(69) // Go to the 69th second in the track.
err := zp.H.SeekTimeDelta(-15) // Go back 15 seconds in the track.

queInfo, err := zp.GetQue() // Get the current que.
```

## Structure

This project is structured as follows:

- ZonePlayer (Gonos.go; Entry point and glues everything together)
  - lib (lib.go; Contains functions and variables that are used throughout the project)
  - Sonos Services (Ex: AVTransport.go; Implements base as documented in [svrooij](https://github.com/svrooij/sonos-api-docs))
  - Helper (Helper.go; Glues helper functions together)
    - Sonos service helpers (Ex: AVTransport.go; Build upon the base implementation for easier use)

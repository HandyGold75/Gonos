# Gonos

This module aims to implement an easy and simple way to control Sonos devices while keeping advanced control possible.

The documentation from [svrooij](https://github.com/svrooij/sonos-api-docs) can be referenced for the base implementation.  
Many thanks for the documentation as without it, this module would not exist.

Next to the base implementation some helper function are present for easier use.  
For all helper functions please refer to the files in the top level directory of this project.

Note that most of this project is untested, as such not all functions might not work as expected.  
Do you're own testing and see if the functionality you need actually works.

## Usage

Creating a ZonePlayer for controlling a Sonos device can be done by any of these methods:

```go
zp, err := Gonos.NewZonePlayer("127.0.0.1") // Use the IpAddress of the Sonos device.
zp, err := Gonos.DiscoverZonePlayer(time.Second) // Discover a Sonos device using SSDP.
zp, err := Gonos.ScanZonePlayer("127.0.0.0/8", time.Second) // Scan a network for Sonos devices.
```

After a ZonePlayer is successfully created the associated Sonos device can be controlled.  
This can be done using either the Sonos services (base implementation) or the helpers.

The available Sonos services are:

- [zp.AlarmClock](/AlarmClock/AlarmClock.go)
- [zp.AudioIn](/AudioIn/AudioIn.go)
- [zp.AVTransport](/AVTransport/AVTransport.go)
- [zp.ConnectionManager](/ConnectionManager/ConnectionManager.go)
- [zp.ContentDirectory](/ContentDirectory/ContentDirectory.go)
- [zp.DeviceProperties](/DeviceProperties/DeviceProperties.go)
- [zp.GroupManagement](/GroupManagement/GroupManagement.go)
- [zp.GroupRenderingControl](/GroupRenderingControl/GroupRenderingControl.go)
- [zp.HTControl](/HTControl/HTControl.go)
- [zp.MusicServices](/MusicServices/MusicServices.go)
- [zp.QPlay](/QPlay/QPlay.go)
- [zp.Queue](/Queue/Queue.go)
- [zp.RenderingControl](/RenderingControl/RenderingControl.go)
- [zp.SystemProperties](/SystemProperties/SystemProperties.go)
- [zp.VirtualLineIn](/VirtualLineIn/VirtualLineIn.go)
- [zp.ZoneGroupTopology](/ZoneGroupTopology/ZoneGroupTopology.go)

## Examples

Some examples for controlling a Sonos device using the Sonos services:

```go
timeServer, err := zp.AlarmClock.GetTimeServer() // Get the current time server.
err := zp.AudioIn.SetLineInLevel(10, 10) // Set left and right line in level to 10.
err := zp.AVTransport.Play() // Play current track.
connectionIDs, err := zp.ConnectionManager.GetCurrentConnectionIDs() // Get ids of current connections.
queInfo, err := zp.ContentDirectory.Browse(Gonos.lib.ContentTypes.QueueMain, "BrowseDirectChildren", "dc:title,res,dc:creator,upnp:artist,upnp:album,upnp:albumArtURI", 0, 0, "") // Get info of the current main que.
zoneAttributes, err := zp.DeviceProperties.GetZoneAttributes() // Get attributes of current zone.
err := zp.GroupManagement.RemoveMember("id") // Remove a group member.
volume, err := zp.GroupRenderingControl.GetGroupVolume() // Get the current group volume.
err := zp.HTControl.SetLEDFeedbackState(true) // Set the LED feedback state.
availableServices, err := zp.MusicServices.ListAvailableServices() // Get available music services.
qPlayAuth, err := zp.QPlay.QPlayAuth("seed") // Does something, probably, idk.
id, err := zp.Queue.RemoveTrackRange(10, 6) // Starting at track 10 remove 6 tracks from que.
err := zp.RenderingControl.SetVolume(10) // Set volume to 10.
err := zp.SystemProperties.EnableRDM(1) // Enable RDM.
err := zp.VirtualLineIn.Play() // Play virtual line in.
zoneGroupAttributes, err := zp.ZoneGroupTopology.GetZoneGroupAttributes() // Get attributes of current zone group.
```

Some examples for controlling a Sonos device using the helpers:

```go
err := zp.Play() // Play current track.
isPlaying, err := zp.GetPlay() // Check if current track is playing.

err := zp.Pause() // Pause current track.
isPaused, err := zp.GetPause() // Check if current track is paused.

err := zp.Stop() // Stop current track.
isStopped, err := zp.GetStop() // Check if current track is stopped.

isTransitioning, err := zp.GetTransitioning() // Check if current track is transitioning.

err := zp.Next() // Go to next track.
err := zp.Previous() // Go to previous track.

err := zp.SetShuffle(true) // Enable shuffle.
isShuffle, err := zp.GetSuffle() // Check if shuffle is enabled.

err := zp.SetRepeat(true) // Enable repeat (Disables reapeat one).
isRepeat, err := zp.GetRepeat() // Check if repeat is enabled.

err := zp.SetRepeatOne(true) // Enable reapeat one (Disables repeat).
isRepeatOne, err := zp.GetRepeatOne() // Check if repeat one is enabled.

err := zp.SeekTrack(10) // Go to 10th track in the que (Count start at 1).
err := zp.SeekTime(69) // Go to the 69th second in the track.
err := zp.SeekTimeDelta(-15) // Go back 15 seconds in the track.

queInfo, err := zp.GetQue() // Get the current que.
```

## Structure

This project is structured as follows:

- Gonos ([/Gonos.go](/Gonos.go); Entrypoint and main functions to get started)
- lib ([/lib/lib.go](/lib/lib.go); Contains functions and variables that are used throughout the project)
- Sonos Services ([/\*/\*.go](/AVTransport/AVTransport.go); Implements base as documented in [svrooij](https://github.com/svrooij/sonos-api-docs))
- Sonos Service Helpers (Ex: [/\*.go](/AVTransport.go); Build upon the base implementation for easier use)

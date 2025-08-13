package Gonos

import (
	"io"

	"github.com/HandyGold75/Gonos/lib"
)

type (
	browseResponseMetaDataLibrary struct {
		Title       string `xml:"title"`
		Class       string `xml:"class"`
		Ordinal     string `xml:"ordinal"`
		Res         string `xml:"res"`
		AlbumArtUri string `xml:"albumArtURI"`
		Type        string `xml:"type"`
		Description string `xml:"description"`
		ResMD       string `xml:"resMD"`
	}
	LibraryInfo struct {
		Count      int
		TotalCount int
		Librarys   []LibraryInfoItem
	}
	LibraryInfoItem struct {
		AlbumArtURI string
		Title       string
		Description string
		Ordinal     string
		Class       string
		Type        string
	}

	browseResponseMetaDataQuePlaylist struct {
		Title       string `xml:"title"`
		Class       string `xml:"class"`
		Ordinal     string `xml:"ordinal"`
		Res         string `xml:"res"`
		AlbumArtUri string `xml:"albumArtURI"`
		Type        string `xml:"type"`
		Description string `xml:"description"`
		ResMD       string `xml:"resMD"`
	}
	PLaylistInfo struct {
		Count      int
		TotalCount int
		Playlists  []PlaylistInfoItem
	}
	PlaylistInfoItem struct {
		AlbumArtURI string
		Title       string
		Description string
		Ordinal     string
		Class       string
		Type        string
	}

	browseResponseMetaDataQue struct {
		Res         string `xml:"res"`
		AlbumArtUri string `xml:"albumArtURI"`
		Title       string `xml:"title"`
		Class       string `xml:"class"`
		Creator     string `xml:"creator"`
		Album       string `xml:"album"`
	}
	QueInfo struct {
		Count      int
		TotalCount int
		Tracks     []QueInfoItem
	}
	QueInfoItem struct {
		AlbumArtURI string
		Title       string
		Class       string
		Creator     string
		Album       string
	}
)

// Prefer methods `zp.LibraryArtist`, `zp.LibraryAlbumArtist`, `zp.LibraryAlbum`, `zp.LibraryGenre`, `zp.LibraryComposer`, `zp.LibraryTracks`, `zp.LibraryPlaylists`.
//
// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (zp *ZonePlayer) BrowseMusicLibrary(objectID string) (LibraryInfo, error) {
	info, err := zp.ContentDirectory.Browse(objectID, "BrowseDirectChildren", "dc:title,res,dc:creator,upnp:artist,upnp:album,upnp:albumArtURI", 0, 0, "")
	if err != nil {
		return LibraryInfo{}, err
	}
	metadata := []browseResponseMetaDataLibrary{}
	err = lib.UnmarshalMetaData(info.Result, &metadata)
	if err == io.EOF {
		return LibraryInfo{}, nil
	} else if err != nil {
		return LibraryInfo{}, err
	}
	librarys := []LibraryInfoItem{}
	for _, library := range metadata {
		librarys = append(librarys, LibraryInfoItem{
			AlbumArtURI: library.AlbumArtUri,
			Title:       library.Title,
			Description: library.Description,
			Class:       library.Class,
			Type:        library.Type,
		})
	}
	return LibraryInfo{Count: info.NumberReturned, TotalCount: info.TotalMatches, Librarys: librarys}, nil
}

// Short for `zp.BrowseMusicLibrary(lib.ContentTypes.Artist)`.
func (zp *ZonePlayer) GetLibraryArtist() (LibraryInfo, error) {
	return zp.BrowseMusicLibrary(lib.ContentTypes.Artist)
}

// Short for `zp.BrowseMusicLibrary(lib.ContentTypes.AlbumArtist)`.
func (zp *ZonePlayer) GetLibraryAlbumArtist() (LibraryInfo, error) {
	return zp.BrowseMusicLibrary(lib.ContentTypes.AlbumArtist)
}

// Short for `zp.BrowseMusicLibrary(lib.ContentTypes.Album)`.
func (zp *ZonePlayer) GetLibraryAlbum() (LibraryInfo, error) {
	return zp.BrowseMusicLibrary(lib.ContentTypes.Album)
}

// Short for `zp.BrowseMusicLibrary(lib.ContentTypes.Genre)`.
func (zp *ZonePlayer) GetLibraryGenre() (LibraryInfo, error) {
	return zp.BrowseMusicLibrary(lib.ContentTypes.Genre)
}

// Short for `zp.BrowseMusicLibrary(lib.ContentTypes.Composer)`.
func (zp *ZonePlayer) GetLibraryComposer() (LibraryInfo, error) {
	return zp.BrowseMusicLibrary(lib.ContentTypes.Composer)
}

// Short for `zp.BrowseMusicLibrary(lib.ContentTypes.Share)`.
func (zp *ZonePlayer) GetLibraryTracks() (LibraryInfo, error) {
	return zp.BrowseMusicLibrary(lib.ContentTypes.Tracks)
}

// Short for `zp.BrowseMusicLibrary(lib.ContentTypes.Playlists)`.
func (zp *ZonePlayer) GetLibraryPlaylists() (LibraryInfo, error) {
	return zp.BrowseMusicLibrary(lib.ContentTypes.Playlists)
}

// Prefer methods `zp.GetShare`, `zp.GetSonosPlaylists`, `zp.GetSonosFavorites`, `zp.GetRadioStations` or `zp.GetRadioShows`.
//
// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (zp *ZonePlayer) BrowsePlaylist(objectID string) (PLaylistInfo, error) {
	info, err := zp.ContentDirectory.Browse(objectID, "BrowseDirectChildren", "dc:title,res,dc:creator,upnp:artist,upnp:album,upnp:albumArtURI", 0, 0, "")
	if err != nil {
		return PLaylistInfo{}, err
	}
	metadata := []browseResponseMetaDataQuePlaylist{}
	err = lib.UnmarshalMetaData(info.Result, &metadata)
	if err == io.EOF {
		return PLaylistInfo{}, nil
	} else if err != nil {
		return PLaylistInfo{}, err
	}
	playlists := []PlaylistInfoItem{}
	for _, playlist := range metadata {
		playlists = append(playlists, PlaylistInfoItem{
			AlbumArtURI: playlist.AlbumArtUri,
			Title:       playlist.Title,
			Description: playlist.Description,
			Class:       playlist.Class,
			Type:        playlist.Type,
		})
	}
	return PLaylistInfo{Count: info.NumberReturned, TotalCount: info.TotalMatches, Playlists: playlists}, nil
}

// Short for `zp.BrowsePlaylist(lib.ContentTypes.Share)`.
func (zp *ZonePlayer) GetShare() (PLaylistInfo, error) {
	return zp.BrowsePlaylist(lib.ContentTypes.Share)
}

// Short for `zp.BrowsePlaylist(lib.ContentTypes.SonosPlaylists)`.
func (zp *ZonePlayer) GetSonosPlaylists() (PLaylistInfo, error) {
	return zp.BrowsePlaylist(lib.ContentTypes.SonosPlaylists)
}

// Short for `zp.BrowsePlaylist(lib.ContentTypes.SonosFavorites)`.
func (zp *ZonePlayer) GetSonosFavorites() (PLaylistInfo, error) {
	return zp.BrowsePlaylist(lib.ContentTypes.SonosFavorites)
}

// Short for `zp.BrowsePlaylist(lib.ContentTypes.RadioStations)`.
func (zp *ZonePlayer) GetRadioStations() (PLaylistInfo, error) {
	return zp.BrowsePlaylist(lib.ContentTypes.RadioStations)
}

// Short for `zp.BrowsePlaylist(lib.ContentTypes.RadioShows)`.
func (zp *ZonePlayer) GetRadioShows() (PLaylistInfo, error) {
	return zp.BrowsePlaylist(lib.ContentTypes.RadioShows)
}

// Prefer methods `zp.GetQue` or `zp.GetQueSecond`.
func (zp *ZonePlayer) BrowseQue(objectID string) (QueInfo, error) {
	info, err := zp.ContentDirectory.Browse(objectID, "BrowseDirectChildren", "dc:title,res,dc:creator,upnp:artist,upnp:album,upnp:albumArtURI", 0, 0, "")
	if err != nil {
		return QueInfo{}, err
	}
	metadata := []browseResponseMetaDataQue{}
	err = lib.UnmarshalMetaData(info.Result, &metadata)
	if err == io.EOF {
		return QueInfo{}, nil
	} else if err != nil {
		return QueInfo{}, err
	}
	tracks := []QueInfoItem{}
	for _, track := range metadata {
		tracks = append(tracks, QueInfoItem{
			AlbumArtURI: track.AlbumArtUri,
			Title:       track.Title,
			Class:       track.Class,
			Creator:     track.Creator,
			Album:       track.Album,
		})
	}
	return QueInfo{Count: info.NumberReturned, TotalCount: info.TotalMatches, Tracks: tracks}, nil
}

// Get que, in case no que is active a empty que will be returned.
//
// Will return incorrect data if a third party application is controling playback.
func (zp *ZonePlayer) GetQue() (QueInfo, error) {
	return zp.BrowseQue(lib.ContentTypes.QueueMain)
}

// Get secondairy que, in case no que is active a empty que will be returned.
//
// Will return incorrect data if a third party application is controling playback.
func (zp *ZonePlayer) GetQueSecond() (QueInfo, error) {
	return zp.BrowseQue(lib.ContentTypes.QueueSecond)
}

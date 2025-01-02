package Helper

import (
	"Gonos/lib"
	"fmt"
	"io"
	"strings"
)

type (
	// To implement, couldn't get a example.
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
	// To implement, couldn't get a example.
	LibraryInfo struct {
		Count      int
		TotalCount int
		Librarys   []LibraryInfoItem
	}
	// To implement, couldn't get a example.
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

// Prefer methods `h.LibraryArtist`, `h.LibraryAlbumArtist`, `h.LibraryAlbum`, `h.LibraryGenre`, `h.LibraryComposer`, `h.LibraryTracks`, `h.LibraryPlaylists`.
//
// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (h *Helper) BrowseMusicLibrary(objectID string) (LibraryInfo, error) {
	info, err := h.contentDirectory.Browse(objectID, "BrowseDirectChildren", "dc:title,res,dc:creator,upnp:artist,upnp:album,upnp:albumArtURI", 0, 0, "")
	if err != nil {
		return LibraryInfo{}, err
	}
	metadata := []browseResponseMetaDataLibrary{}
	fmt.Println(strings.ReplaceAll(info.Result, "id=", "\r\nid="))
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

func (h *Helper) GetLibraryArtist() (LibraryInfo, error) {
	return h.BrowseMusicLibrary(lib.ContentTypes.Artist)
}

func (h *Helper) GetLibraryAlbumArtist() (LibraryInfo, error) {
	return h.BrowseMusicLibrary(lib.ContentTypes.AlbumArtist)
}

func (h *Helper) GetLibraryAlbum() (LibraryInfo, error) {
	return h.BrowseMusicLibrary(lib.ContentTypes.Album)
}

func (h *Helper) GetLibraryGenre() (LibraryInfo, error) {
	return h.BrowseMusicLibrary(lib.ContentTypes.Genre)
}

func (h *Helper) GetLibraryComposer() (LibraryInfo, error) {
	return h.BrowseMusicLibrary(lib.ContentTypes.Composer)
}

func (h *Helper) GetLibraryTracks() (LibraryInfo, error) {
	return h.BrowseMusicLibrary(lib.ContentTypes.Tracks)
}

func (h *Helper) GetLibraryPlaylists() (LibraryInfo, error) {
	return h.BrowseMusicLibrary(lib.ContentTypes.Playlists)
}

// Prefer methods `h.GetShare`, `h.GetSonosPlaylists`, `h.GetSonosFavorites`, `h.GetRadioStations` or `h.GetRadioShows`.
//
// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (h *Helper) BrowsePlaylist(objectID string) (PLaylistInfo, error) {
	info, err := h.contentDirectory.Browse(objectID, "BrowseDirectChildren", "dc:title,res,dc:creator,upnp:artist,upnp:album,upnp:albumArtURI", 0, 0, "")
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

func (h *Helper) GetShare() (PLaylistInfo, error) {
	return h.BrowsePlaylist(lib.ContentTypes.Share)
}

func (h *Helper) GetSonosPlaylists() (PLaylistInfo, error) {
	return h.BrowsePlaylist(lib.ContentTypes.SonosPlaylists)
}

// Get Sonos playlists, in case no sonos playlists are present a empty playlist will be returned
func (h *Helper) GetSonosFavorites() (PLaylistInfo, error) {
	return h.BrowsePlaylist(lib.ContentTypes.SonosFavorites)
}

func (h *Helper) GetRadioStations() (PLaylistInfo, error) {
	return h.BrowsePlaylist(lib.ContentTypes.RadioStations)
}

func (h *Helper) GetRadioShows() (PLaylistInfo, error) {
	return h.BrowsePlaylist(lib.ContentTypes.RadioShows)
}

// Prefer methods `h.GetQue` or `h.GetQueSecond`.
func (h *Helper) BrowseQue(objectID string) (QueInfo, error) {
	info, err := h.contentDirectory.Browse(objectID, "BrowseDirectChildren", "dc:title,res,dc:creator,upnp:artist,upnp:album,upnp:albumArtURI", 0, 0, "")
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
	fmt.Println(len(metadata))
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
func (h *Helper) GetQue() (QueInfo, error) {
	return h.BrowseQue(lib.ContentTypes.QueueMain)
}

// Get secondairy que, in case no que is active a empty que will be returned.
//
// Will return incorrect data if a third party application is controling playback.
func (h *Helper) GetQueSecond() (QueInfo, error) {
	return h.BrowseQue(lib.ContentTypes.QueueSecond)
}

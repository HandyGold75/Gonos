package ContentDirectory

import (
	"Gonos/lib"
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

type (
	ContentDirectory struct {
		Send func(action, body, targetTag string) (string, error)
	}

	browseResponse struct {
		XMLName xml.Name `xml:"BrowseResponse"`
		// Encoded DIDL-Lite XML.
		//
		// Should be unmarshaled into type of `browseResponseMetaData*`
		Result         string
		NumberReturned int
		TotalMatches   int
		UpdateID       int
	}
	createObjectResponse struct {
		XMLName  xml.Name `xml:"CreateObjectResponse"`
		ObjectID string
		Result   string
	}
	findPrefixResponse struct {
		XMLName       xml.Name `xml:"FindPrefixResponse"`
		StartingIndex int
		UpdateID      int
	}
	getAllPrefixLocationsResponse struct {
		XMLName           xml.Name `xml:"GetAllPrefixLocationsResponse"`
		TotalPrefixes     int
		PrefixAndIndexCSV string
		UpdateID          int
	}
)

func New(send func(action, body, targetTag string) (string, error)) ContentDirectory {
	return ContentDirectory{Send: send}
}

// Prefer methods `zp.BrowseMusicLibrary`, `zp.BrowsePlaylist`, `zp.BrowseQue`.
//
// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (s *ContentDirectory) Browse(objectID string, browseFlag string, filter string, startingIndex int, requestedCount int, sortCriteria string) (browseResponse, error) {
	res, err := s.Send("Browse", "<ObjectID>"+objectID+"</ObjectID><BrowseFlag>"+browseFlag+"</BrowseFlag><Filter>"+filter+"</Filter><StartingIndex>"+strconv.Itoa(startingIndex)+"</StartingIndex><RequestedCount>"+strconv.Itoa(requestedCount)+"</RequestedCount><SortCriteria>"+sortCriteria+"</SortCriteria>", "s:Body")
	if err != nil {
		return browseResponse{}, err
	}
	data := browseResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *ContentDirectory) CreateObject(containerID string, elements string) (createObjectResponse, error) {
	res, err := s.Send("CreateObject", "<ContainerID>"+containerID+"</ContainerID><Elements>"+elements+"</Elements>", "s:Body")
	if err != nil {
		return createObjectResponse{}, err
	}
	data := createObjectResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (s *ContentDirectory) DestroyObject(objectID string) error {
	_, err := s.Send("DestroyObject", "<ObjectID>"+objectID+"</ObjectID>", "")
	return err
}

// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (s *ContentDirectory) FindPrefix(objectID string, prefix string) (findPrefixResponse, error) {
	res, err := s.Send("FindPrefix", "<ObjectID>"+objectID+"</ObjectID><Prefix>"+prefix+"</Prefix>", "s:Body")
	if err != nil {
		return findPrefixResponse{}, err
	}
	data := findPrefixResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

// Returns one of `WMP`, `ITUNES` or `NONE`
func (s *ContentDirectory) GetAlbumArtistDisplayOption() (AlbumArtistDisplayOption string, err error) {
	return s.Send("GetAlbumArtistDisplayOption", "", "AlbumArtistDisplayOption")
}

// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (s *ContentDirectory) GetAllPrefixLocations(objectID string) (getAllPrefixLocationsResponse, error) {
	res, err := s.Send("GetAllPrefixLocations", "<ObjectID>"+objectID+"</ObjectID>", "s:Body")
	if err != nil {
		return getAllPrefixLocationsResponse{}, err
	}
	data := getAllPrefixLocationsResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *ContentDirectory) GetBrowseable() (IsBrowseable bool, err error) {
	res, err := s.Send("GetBrowseable", "", "IsBrowseable")
	return res == "1", err
}

func (s *ContentDirectory) GetLastIndexChange() (LastIndexChange time.Time, err error) {
	res, err := s.Send("GetLastIndexChange", "", "LastIndexChange")
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse("S"+time.DateTime, res)
}

func (s *ContentDirectory) GetSearchCapabilities() (SearchCaps string, err error) {
	return s.Send("GetSearchCapabilities", "", "SearchCaps")
}

func (s *ContentDirectory) GetShareIndexInProgress() (IsIndexing bool, err error) {
	res, err := s.Send("GetShareIndexInProgress", "", "IsIndexing")
	return res == "1", err
}

func (s *ContentDirectory) GetSortCapabilities() (SortCaps string, err error) {
	return s.Send("GetSortCapabilities", "", "SortCaps")
}

func (s *ContentDirectory) GetSystemUpdateID() (Id int, err error) {
	res, err := s.Send("GetSystemUpdateID", "", "Id")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

// `albumArtistDisplayOption` should be one of `s.AlbumArtistDisplayOptionModes.*`.
func (s *ContentDirectory) RefreshShareIndex(albumArtistDisplayOption string) error {
	_, err := s.Send("RefreshShareIndex", "<AlbumArtistDisplayOption>"+albumArtistDisplayOption+"</AlbumArtistDisplayOption>", "")
	return err
}

func (s *ContentDirectory) RequestResort(sortOrder string) error {
	_, err := s.Send("RequestResort", "<SortOrder>sortOrder</SortOrder>", "")
	return err
}

func (s *ContentDirectory) SetBrowseable(state bool) error {
	_, err := s.Send("SetBrowseable", "<Browseable>"+lib.BoolTo10(state)+"</Browseable>", "")
	return err
}

// `objectID` may be one of `Gonos.ContentTypes.*` or a custom id
func (s *ContentDirectory) UpdateObject(objectID string, currentTagValue string, newTagValue string) error {
	out, err := s.Send("UpdateObject", "<ObjectID>"+objectID+"</ObjectID><CurrentTagValue>"+currentTagValue+"</CurrentTagValue><NewTagValue>"+newTagValue+"</NewTagValue>", "")
	fmt.Println(out)
	return err
}

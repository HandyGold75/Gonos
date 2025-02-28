package Queue

import (
	"Gonos/lib"
	"encoding/xml"
	"strconv"
)

type (
	Queue struct {
		Send func(action, body, targetTag string) (string, error)
		// Should be one of `Gonos.ContentTypes.*` // OR maybe just 0?
		QueueID string
		// Should always be `0`
		UpdateID int
	}

	addMultipleURIsResponse struct {
		XMLName                  xml.Name `xml:"AddMultipleURIsResponse"`
		FirstTrackNumberEnqueued int
		NumTracksAdded           int
		NewQueueLength           int
		NewUpdateID              int
	}

	addURIResponse struct {
		XMLName                  xml.Name `xml:"AddURIResponse"`
		FirstTrackNumberEnqueued int
		NumTracksAdded           int
		NewQueueLength           int
		NewUpdateID              int
	}

	attachQueueResponse struct {
		XMLName           xml.Name `xml:"AttachQueueResponse"`
		QueueID           int
		QueueOwnerContext string
	}

	browseResponse struct {
		XMLName        xml.Name `xml:"BrowseResponse"`
		Result         string
		NumberReturned int
		TotalMatches   int
		UpdateID       int
	}

	replaceAllTracksResponse struct {
		XMLName        xml.Name `xml:"ReplaceAllTracksResponse"`
		NewQueueLength int
		NewUpdateID    int
	}
)

func New(send func(action, body, targetTag string) (string, error)) Queue {
	return Queue{Send: send, QueueID: lib.ContentTypes.QueueMain, UpdateID: 0}
}

func (s *Queue) AddMultipleURIs(containerURI string, containerMetaData string, desiredFirstTrackNumberEnqueued int, enqueueAsNext bool, numberOfURIs int, enqueuedURIsAndMetaData string) (addMultipleURIsResponse, error) {
	res, err := s.Send("AddMultipleURIs ", "<QueueID>"+s.QueueID+"</QueueID><UpdateID>"+strconv.Itoa(s.UpdateID)+"</UpdateID><ContainerURI>"+containerURI+"</ContainerURI><ContainerMetaData>"+containerMetaData+"</ContainerMetaData><DesiredFirstTrackNumberEnqueued>"+strconv.Itoa(desiredFirstTrackNumberEnqueued)+"</DesiredFirstTrackNumberEnqueued><EnqueueAsNext>"+lib.BoolTo10(enqueueAsNext)+"</EnqueueAsNext><NumberOfURIs>"+strconv.Itoa(numberOfURIs)+"</NumberOfURIs><EnqueuedURIsAndMetaData>"+enqueuedURIsAndMetaData+"</EnqueuedURIsAndMetaData>", "s:Body")
	if err != nil {
		return addMultipleURIsResponse{}, err
	}
	data := addMultipleURIsResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *Queue) AddURI(enqueuedURI string, enqueuedURIMetaData string, desiredFirstTrackNumberEnqueued int, enqueueAsNext bool) (addURIResponse, error) {
	res, err := s.Send("AddURI ", "<QueueID>"+s.QueueID+"</QueueID><UpdateID>"+strconv.Itoa(s.UpdateID)+"</UpdateID><EnqueuedURI>"+enqueuedURI+"</EnqueuedURI><EnqueuedURIMetaData>"+enqueuedURIMetaData+"</EnqueuedURIMetaData><DesiredFirstTrackNumberEnqueued>"+strconv.Itoa(desiredFirstTrackNumberEnqueued)+"</DesiredFirstTrackNumberEnqueued><EnqueueAsNext>"+lib.BoolTo10(enqueueAsNext)+"</EnqueueAsNext>", "s:Body")
	if err != nil {
		return addURIResponse{}, err
	}
	data := addURIResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *Queue) AttachQueue(queueOwnerID string) (attachQueueResponse, error) {
	res, err := s.Send("AttachQueue ", "<QueueOwnerID>"+queueOwnerID+"</QueueOwnerID>", "")
	if err != nil {
		return attachQueueResponse{}, err
	}
	data := attachQueueResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *Queue) Backup() (string, error) {
	return s.Send("Backup ", "", "")
}

func (s *Queue) Browse(startingIndex int, requestedCount int) (browseResponse, error) {
	res, err := s.Send("Browse ", "<QueueID>"+s.QueueID+"</QueueID><StartingIndex>"+strconv.Itoa(startingIndex)+"</StartingIndex><RequestedCount>"+strconv.Itoa(requestedCount)+"</RequestedCount>", "")
	if err != nil {
		return browseResponse{}, err
	}
	data := browseResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *Queue) CreateQueue(queueOwnerID string, queueOwnerContext string, queuePolicy string) (int, error) {
	res, err := s.Send("CreateQueue ", "<QueueOwnerID>"+queueOwnerID+"</QueueOwnerID><QueueOwnerContext>"+queueOwnerContext+"</QueueOwnerContext><QueuePolicy>"+queuePolicy+"</QueuePolicy>", "QueueID")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *Queue) RemoveAllTracks() (int, error) {
	res, err := s.Send("RemoveAllTracks ", "<QueueID>"+s.QueueID+"</QueueID><UpdateID>"+strconv.Itoa(s.UpdateID)+"</UpdateID>", "NewUpdateID")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *Queue) RemoveTrackRange(startingIndex int, numberOfTracks int) (int, error) {
	res, err := s.Send("RemoveTrackRange ", "<QueueID>"+s.QueueID+"</QueueID><UpdateID>"+strconv.Itoa(s.UpdateID)+"</UpdateID><StartingIndex>"+strconv.Itoa(startingIndex)+"</StartingIndex><NumberOfTracks>"+strconv.Itoa(numberOfTracks)+"</NumberOfTracks>", "NewUpdateID")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *Queue) ReorderTracks(startingIndex int, numberOfTracks int, insertBefore int) (int, error) {
	res, err := s.Send("ReorderTracks ", "<QueueID>"+s.QueueID+"</QueueID><StartingIndex>"+strconv.Itoa(startingIndex)+"</StartingIndex><NumberOfTracks>"+strconv.Itoa(numberOfTracks)+"</NumberOfTracks><InsertBefore>"+strconv.Itoa(insertBefore)+"</InsertBefore><UpdateID>"+strconv.Itoa(s.UpdateID)+"</UpdateID>", "NewUpdateID")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (s *Queue) ReplaceAllTracks(containerURI string, containerMetaData string, currentTrackIndex int, newCurrentTrackIndices string, numberOfURIs int, enqueuedURIsAndMetaData string) (replaceAllTracksResponse, error) {
	res, err := s.Send("ReplaceAllTracks ", "<QueueID>"+s.QueueID+"</QueueID><UpdateID>"+strconv.Itoa(s.UpdateID)+"</UpdateID><ContainerURI>"+containerURI+"</ContainerURI><ContainerMetaData>"+containerMetaData+"</ContainerMetaData><CurrentTrackIndex>"+strconv.Itoa(currentTrackIndex)+"</CurrentTrackIndex><NewCurrentTrackIndices>"+newCurrentTrackIndices+"</NewCurrentTrackIndices><NumberOfURIs>"+strconv.Itoa(numberOfURIs)+"</NumberOfURIs><EnqueuedURIsAndMetaData>"+enqueuedURIsAndMetaData+"</EnqueuedURIsAndMetaData>", "")
	if err != nil {
		return replaceAllTracksResponse{}, err
	}
	data := replaceAllTracksResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *Queue) SaveAsSonosPlaylist(title string, objectID string) (string, error) {
	return s.Send("SaveAsSonosPlaylist ", "<QueueID>"+s.QueueID+"</QueueID><Title>"+title+"</Title><ObjectID>"+objectID+"</ObjectID>", "AssignedObjectID")
}

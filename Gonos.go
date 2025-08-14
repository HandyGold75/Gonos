package Gonos

import (
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/HandyGold75/Gonos/lib"

	"github.com/HandyGold75/Gonos/AVTransport"
	"github.com/HandyGold75/Gonos/AlarmClock"
	"github.com/HandyGold75/Gonos/AudioIn"
	"github.com/HandyGold75/Gonos/ConnectionManager"
	"github.com/HandyGold75/Gonos/ContentDirectory"
	"github.com/HandyGold75/Gonos/DeviceProperties"
	"github.com/HandyGold75/Gonos/GroupManagement"
	"github.com/HandyGold75/Gonos/GroupRenderingControl"
	"github.com/HandyGold75/Gonos/HTControl"
	"github.com/HandyGold75/Gonos/MusicServices"
	"github.com/HandyGold75/Gonos/QPlay"
	"github.com/HandyGold75/Gonos/Queue"
	"github.com/HandyGold75/Gonos/RenderingControl"
	"github.com/HandyGold75/Gonos/SystemProperties"
	"github.com/HandyGold75/Gonos/VirtualLineIn"
	"github.com/HandyGold75/Gonos/ZoneGroupTopology"
)

type (
	ZonePlayer struct {
		// Full url address packets will be send to.
		URL string
		// GetZoneInfo call is made to confirm if the requested ZonePlayer exists opon creation, might as well store the returned data.
		ZoneInfo ZoneInfo

		// Sonos SOAP Service `AlarmClock`.
		//
		// Prefer methods in `zp` over methods in `zp.AlarmClock`.
		//
		// Control the sonos alarms and times.
		AlarmClock AlarmClock.AlarmClock

		// Sonos SOAP Service `AudioIn`.
		//
		// Prefer methods in `zp` over methods in `zp.AudioIn`.
		//
		// Control line in.
		AudioIn AudioIn.AudioIn

		// Sonos SOAP Service `AVTransport`.
		//
		// Prefer methods in `zp` over methods in `zp.AVTransport`.
		//
		// Service that controls stuff related to transport (play/pause/next/special URLs).
		AVTransport AVTransport.AVTransport

		// Sonos SOAP Service `ConnectionManager`.
		//
		// Prefer methods in `zp` over methods in `zp.ConnectionManager`.
		//
		// Services related to connections and protocols.
		ConnectionManager ConnectionManager.ConnectionManager

		// Sonos SOAP Service `ContentDirectory`.
		//
		// Prefer methods in `zp` over methods in `zp.ContentDirectory`.
		//
		// Browse for local content.
		ContentDirectory ContentDirectory.ContentDirectory

		// Sonos SOAP Service `DeviceProperties`.
		//
		// Prefer methods in `zp` over methods in `zp.DeviceProperties`.
		//
		// Modify device properties, like LED status and stereo pairs.
		DeviceProperties DeviceProperties.DeviceProperties

		// Sonos SOAP Service `GroupManagement`.
		//
		// Prefer methods in `zp` over methods in `zp.GroupManagement`.
		//
		// Services related to groups.
		GroupManagement GroupManagement.GroupManagement

		// Sonos SOAP Service `GroupRenderingControl`.
		//
		// Prefer methods in `zp` over methods in `zp.GroupRenderingControl`.
		//
		// Volume related controls for groups.
		GroupRenderingControl GroupRenderingControl.GroupRenderingControl

		// Sonos SOAP Service `HTControl`.
		//
		// Prefer methods in `zp` over methods in `zpTControl`.
		//
		// Service related to the TV remote control.
		HTControl HTControl.HTControl

		// Sonos SOAP Service `MusicServices`.
		//
		// Prefer methods in `zp` over methods in `zp.MusicServices`.
		//
		// Access to external music services, like Spotify or Youtube Music.
		MusicServices MusicServices.MusicServices

		// Sonos SOAP Service `QPlay`.
		//
		// Prefer methods in `zp` over methods in `zp.QPlay`.
		//
		// Services related to Chinese Tencent Qplay service.
		QPlay QPlay.QPlay

		// Sonos SOAP Service `Queue`.
		//
		// Prefer methods in `zp` over methods in `zp.Queue`.
		//
		// Modify and browse queues.
		Queue Queue.Queue

		// Sonos SOAP Service `RenderingControl`.
		//
		// Prefer methods in `zp` over methods in `zp.RenderingControl`.
		//
		// Volume related controls.
		RenderingControl RenderingControl.RenderingControl

		// Sonos SOAP Service `SystemProperties`.
		//
		// Prefer methods in `zp` over methods in `zp.SystemProperties`.
		//
		// Manage system-wide settings, mainly account stuff.
		SystemProperties SystemProperties.SystemProperties

		// Sonos SOAP Service `VirtualLineIn`.
		//
		// Prefer methods in `zp` over methods in `zp.//`.
		VirtualLineIn VirtualLineIn.VirtualLineIn

		// Sonos SOAP Service `ZoneGroupTopology`.
		//
		// Prefer methods in `zp` over methods in `zp.ZoneGroupTopology`.
		//
		// Zone config stuff, eg getting all the configured sonos zones.
		ZoneGroupTopology ZoneGroupTopology.ZoneGroupTopology
	}
)

var (
	ErrSonos                      = lib.ErrSonos
	ContentTypes                  = lib.ContentTypes
	SeekModes                     = lib.SeekModes
	PlayModes                     = lib.PlayModes
	PlayModeMap                   = lib.PlayModeMap
	PlayModeMapReversed           = lib.PlayModeMapReversed
	RecurrenceModes               = lib.RecurrenceModes
	AlbumArtistDisplayOptionModes = lib.AlbumArtistDisplayOptionModes
	UpdateTypes                   = lib.UpdateTypes
)

// Create new ZonePlayer for controling a Sonos speaker.
func NewZonePlayer(ipAddress string) (*ZonePlayer, error) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return &ZonePlayer{}, ErrSonos.ErrInvalidIPAdress
	}

	zp := &ZonePlayer{URL: "http://" + ip.String() + ":1400"}
	zp.AlarmClock = AlarmClock.New(zp.SendAlarmClock)
	zp.AudioIn = AudioIn.New(zp.SendAudioIn)
	zp.AVTransport = AVTransport.New(zp.SendAVTransport)
	zp.ConnectionManager = ConnectionManager.New(zp.SendConnectionManager)
	zp.ContentDirectory = ContentDirectory.New(zp.SendContentDirectory)
	zp.DeviceProperties = DeviceProperties.New(zp.SendDeviceProperties)
	zp.GroupManagement = GroupManagement.New(zp.SendGroupManagement)
	zp.GroupRenderingControl = GroupRenderingControl.New(zp.SendGroupRenderingControl)
	zp.HTControl = HTControl.New(zp.SendHTControl)
	zp.MusicServices = MusicServices.New(zp.SendMusicServices)
	zp.QPlay = QPlay.New(zp.SendQPlay)
	zp.Queue = Queue.New(zp.SendQueue)
	zp.RenderingControl = RenderingControl.New(zp.SendRenderingControl)
	zp.SystemProperties = SystemProperties.New(zp.SendSystemProperties)
	zp.VirtualLineIn = VirtualLineIn.New(zp.SendVirtualLineIn)
	zp.ZoneGroupTopology = ZoneGroupTopology.New(zp.SendZoneGroupTopology)

	info, err := zp.GetZoneInfo()
	if err != nil {
		return &ZonePlayer{}, ErrSonos.ErrNoZonePlayerFound
	}
	zp.ZoneInfo = info
	return zp, nil
}

// Create new ZonePlayer using discovery for controling a Sonos speaker.
//
// `timeout` of 1 second is recomended.
func DiscoverZonePlayer(timeout time.Duration) ([]*ZonePlayer, error) {
	conn, err := net.DialUDP("udp", &net.UDPAddr{Port: 1900}, &net.UDPAddr{IP: net.IPv4(239, 255, 255, 250), Port: 1900})
	if err != nil {
		return []*ZonePlayer{}, err
	}
	defer func() { _ = conn.Close() }()

	for range 3 {
		_, _ = conn.Write([]byte("M-SEARCH * HTTP/1.1\r\nHOST: 239.255.255.250:1900\r\nMAN: \"ssdp:discover\"\r\nMX: 1\r\nST: urn:schemas-upnp-org:device:ZonePlayer:1\r\n\r\n"))
	}

	zps := []*ZonePlayer{}
	for {
		buf := make([]byte, 1024)
		if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
			return zps, err
		}
		_, addr, err := conn.ReadFrom(buf)
		if err.(*net.OpError).Timeout() {
			break
		} else if err != nil {
			return zps, err
		}

		zp, err := NewZonePlayer(strings.Split(addr.String(), ":")[0])
		if err != nil {
			continue
		}
		zps = append(zps, zp)
	}

	if len(zps) <= 0 {
		return zps, ErrSonos.ErrNoZonePlayerFound
	}
	return zps, nil
}

// Create new ZonePlayer using network scanning for controling a Sonos speaker.
//
// `timeout` of 1 second is recomended.
func ScanZonePlayer(cidr string, timeout time.Duration) ([]*ZonePlayer, error) {
	incIP := func(ip net.IP) {
		for j := len(ip) - 1; j >= 0; j-- {
			ip[j]++
			if ip[j] > 0 {
				break
			}
		}
	}

	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return []*ZonePlayer{}, err
	}

	wg, zps := sync.WaitGroup{}, []*ZonePlayer{}
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()

			conn, err := net.DialTimeout("tcp", ip+":"+"1400", time.Second)
			if err != nil {
				return
			}
			defer func() { _ = conn.Close() }()

			zp, err := NewZonePlayer(ip)
			if err != nil {
				return
			}
			zps = append(zps, zp)
		}(ip.String())
	}
	wg.Wait()

	if len(zps) <= 0 {
		return zps, ErrSonos.ErrNoZonePlayerFound
	}
	return zps, nil
}

func (zp *ZonePlayer) SendAVTransport(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/MediaRenderer/AVTransport/Control", "AVTransport", action, "<InstanceID>0</InstanceID>"+body, targetTag)
}

func (zp *ZonePlayer) SendAlarmClock(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/AlarmClock/Control", "AlarmClock", action, body, targetTag)
}

func (zp *ZonePlayer) SendAudioIn(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/AudioIn/Control", "AudioIn", action, body, targetTag)
}

func (zp *ZonePlayer) SendConnectionManager(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/MediaRenderer/ConnectionManager/Control", "ConnectionManager", action, body, targetTag)
}

func (zp *ZonePlayer) SendContentDirectory(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/MediaServer/ContentDirectory/Control", "ContentDirectory", action, body, targetTag)
}

func (zp *ZonePlayer) SendDeviceProperties(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/DeviceProperties/Control", "DeviceProperties", action, body, targetTag)
}

func (zp *ZonePlayer) SendGroupManagement(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/GroupManagement/Control", "GroupManagement", action, body, targetTag)
}

func (zp *ZonePlayer) SendGroupRenderingControl(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/MediaRenderer/GroupRenderingControl/Control", "GroupRenderingControl", action, "<InstanceID>0</InstanceID>"+body, targetTag)
}

func (zp *ZonePlayer) SendHTControl(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/HTControl/Control", "HTControl", action, body, targetTag)
}

func (zp *ZonePlayer) SendMusicServices(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/MusicServices/Control", "MusicServices", action, body, targetTag)
}

func (zp *ZonePlayer) SendQPlay(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/QPlay/Control", "QPlay", action, body, targetTag)
}

func (zp *ZonePlayer) SendQueue(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/MediaRenderer/Queue/Control", "Queue", action, body, targetTag)
}

func (zp *ZonePlayer) SendRenderingControl(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/MediaRenderer/RenderingControl/Control", "RenderingControl", action, "<InstanceID>0</InstanceID>"+body, targetTag)
}

func (zp *ZonePlayer) SendSystemProperties(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/SystemProperties/Control", "SystemProperties", action, body, targetTag)
}

func (zp *ZonePlayer) SendVirtualLineIn(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/MediaRenderer/VirtualLineIn/Control", "VirtualLineIn", action, "<InstanceID>0</InstanceID>"+body, targetTag)
}

func (zp *ZonePlayer) SendZoneGroupTopology(action, body, targetTag string) (string, error) {
	return SendAndVerify(zp.URL+"/ZoneGroupTopology/Control", "ZoneGroupTopology", action, body, targetTag)
}

func SendAndVerify(url string, endpoint string, action string, body string, targetTag string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(`<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body><u:`+action+` xmlns:u="urn:schemas-upnp-org:service:`+endpoint+`:1">`+body+`</u:`+action+`></s:Body></s:Envelope>`))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPACTION", "urn:schemas-upnp-org:service:"+endpoint+":1#"+action)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	resb, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	res := string(resb[:])

	if targetTag != "" {
		res, err = lib.ExtractTag(res, targetTag)
		if err != nil {
			return res, ErrSonos.ErrUnexpectedResponse
		}
		return res, nil
	}
	if res != `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body><u:`+action+`Response xmlns:u="urn:schemas-upnp-org:service:`+endpoint+`:1"></u:`+action+`Response></s:Body></s:Envelope>` {
		return res, ErrSonos.ErrUnexpectedResponse
	}
	return res, nil
}

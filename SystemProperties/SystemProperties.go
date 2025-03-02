package SystemProperties

import (
	"encoding/xml"
	"strconv"

	"github.com/HandyGold75/Gonos/lib"
)

type (
	SystemProperties struct {
		Send func(action, body, targetTag string) (string, error)
	}

	addOAuthAccountXResponse struct {
		XMLName         xml.Name `xml:"AddOAuthAccountXResponse"`
		AccountUDN      string
		AccountNickname string
	}

	provisionCredentialedTrialAccountXResponse struct {
		XMLName    xml.Name `xml:"ProvisionCredentialedTrialAccountXResponse"`
		IsExpired  bool
		AccountUDN string
	}
)

func New(send func(action, body, targetTag string) (string, error)) SystemProperties {
	return SystemProperties{Send: send}
}

func (s *SystemProperties) AddAccountX(accountType int, accountID string, accountPassword string) (string, error) {
	return s.Send("AddAccountX", "<AccountType>"+strconv.Itoa(accountType)+"</AccountType><AccountID>"+accountID+"</AccountID><AccountPassword>"+accountPassword+"</AccountPassword>", "AccountUDN")
}

func (s *SystemProperties) AddOAuthAccountX(accountType int, accountToken string, accountKey string, oAuthDeviceID string, authorizationCode string, redirectURI string, userIdHashCode string, accountTier int) (addOAuthAccountXResponse, error) {
	res, err := s.Send("AddOAuthAccountX", "<AccountType>"+strconv.Itoa(accountType)+"</AccountType><AccountToken>"+accountToken+"</AccountToken><AccountKey>"+accountKey+"</AccountKey><OAuthDeviceID>"+oAuthDeviceID+"</OAuthDeviceID><AuthorizationCode>"+authorizationCode+"</AuthorizationCode><RedirectURI>"+redirectURI+"</RedirectURI><UserIdHashCode>"+userIdHashCode+"</UserIdHashCode><AccountTier>"+strconv.Itoa(accountTier)+"</AccountTier>", "s:Body")
	if err != nil {
		return addOAuthAccountXResponse{}, err
	}
	data := addOAuthAccountXResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *SystemProperties) DoPostUpdateTasks() error {
	_, err := s.Send("DoPostUpdateTasks", "", "")
	return err
}

func (s *SystemProperties) EditAccountMd(accountType int, accountID string, newAccountMd string) error {
	_, err := s.Send("EditAccountMd", "<AccountType>"+strconv.Itoa(accountType)+"</AccountType><AccountID>"+accountID+"</AccountID><NewAccountMd>"+newAccountMd+"</NewAccountMd>", "")
	return err
}

func (s *SystemProperties) EditAccountPasswordX(accountType int, accountID string, newAccountPassword string) error {
	_, err := s.Send("EditAccountPasswordX", "<AccountType>"+strconv.Itoa(accountType)+"</AccountType><AccountID>"+accountID+"</AccountID><NewAccountPassword>"+newAccountPassword+"</NewAccountPassword>", "")
	return err
}

func (s *SystemProperties) EnableRDM(state bool) error {
	_, err := s.Send("EnableRDM", "<RDMValue>"+lib.BoolTo10(state)+"</RDMValue>", "")
	return err
}

func (s *SystemProperties) GetRDM() (bool, error) {
	res, err := s.Send("GetRDM", "", "RDMValue")
	return res == "1", err
}

func (s *SystemProperties) GetString(variableName string) (string, error) {
	return s.Send("GetString", "<VariableName>"+variableName+"</VariableName>", "StringValue")
}

func (s *SystemProperties) GetWebCode(accountType string) (string, error) {
	return s.Send("GetWebCode", "<AccountType>"+accountType+"</AccountType>", "WebCode")
}

func (s *SystemProperties) ProvisionCredentialedTrialAccountX(accountType int, accountID string, accountPassword string) (provisionCredentialedTrialAccountXResponse, error) {
	res, err := s.Send("ProvisionCredentialedTrialAccountX", "<AccountType>"+strconv.Itoa(accountType)+"</AccountType><AccountID>"+accountID+"</AccountID><AccountPassword>"+accountPassword+"</AccountPassword>", "s:Body")
	if err != nil {
		return provisionCredentialedTrialAccountXResponse{}, err
	}
	data := provisionCredentialedTrialAccountXResponse{}
	err = xml.Unmarshal([]byte(res), &data)
	return data, err
}

func (s *SystemProperties) RefreshAccountCredentialsX(accountType int, accountUID int, accountToken string, accountKey string) error {
	_, err := s.Send("RefreshAccountCredentialsX", "<AccountType>"+strconv.Itoa(accountType)+"</AccountType><AccountUID>"+strconv.Itoa(accountUID)+"</AccountUID><AccountToken>"+accountToken+"</AccountToken><AccountKey>"+accountKey+"</AccountKey>", "")
	return err
}

func (s *SystemProperties) Remove(variableName string) error {
	_, err := s.Send("Remove", "<VariableName>"+variableName+"</VariableName>", "")
	return err
}

func (s *SystemProperties) RemoveAccount(accountType int, accountID string) error {
	_, err := s.Send("RemoveAccount", "<AccountType>accountType</AccountType><AccountID>accountID</AccountID>", "")
	return err
}

func (s *SystemProperties) ReplaceAccountX(accountUDN string, newAccountID string, newAccountPassword string, accountToken string, accountKey string, oAuthDeviceID string) (string, error) {
	return s.Send("ReplaceAccountX", "<AccountUDN>"+accountUDN+"</AccountUDN><NewAccountID>"+newAccountID+"</NewAccountID><NewAccountPassword>"+newAccountPassword+"</NewAccountPassword><AccountToken>"+accountToken+"</AccountToken><AccountKey>"+accountKey+"</AccountKey><OAuthDeviceID>"+oAuthDeviceID+"</OAuthDeviceID>", "NewAccountUDN")
}

func (s *SystemProperties) ResetThirdPartyCredentials() error {
	_, err := s.Send("ResetThirdPartyCredentials", "", "")
	return err
}

func (s *SystemProperties) SetAccountNicknameX(accountUDN string, accountNickname string) error {
	_, err := s.Send("SetAccountNicknameX", "<AccountUDN>"+accountUDN+"</AccountUDN><AccountNickname>"+accountNickname+"</AccountNickname>", "")
	return err
}

func (s *SystemProperties) SetString(variableName string, stringValue string) error {
	_, err := s.Send("SetString", "<VariableName>"+variableName+"</VariableName><StringValue>"+stringValue+"</StringValue>", "")
	return err
}

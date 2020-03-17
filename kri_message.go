package gov

import (
	"time"
)

type (
	KriMessage struct {
		MdmId          string `json:"mdm_id"`
		Version        int64
		Deleted        bool
		FirstName      string    `json:"first_name"`
		MiddleName     string    `json:"middle_name"`
		LastName       string    `json:"last_name"`
		BirthDate      time.Time `json:"birth_date"`
		RipDate        string    `json:"rip_date"`
		Ids            Ids
		Documents      Documents
		Contacts       Contacts
		Addresses      Addresses
		RegNum         string    `json:"reg_num"`
		RegDate        time.Time `json:"reg_date"`
		RegStatus      string    `json:"reg_status"`
		ValidationMask int       `json:"validation_mask"`
	}

	Documents struct {
		Passport []Passport `json:"doc_passport_rf"`
		Snils    []Document `json:"doc_snils"`
	}

	Document struct {
		Value      string `json:"value"`
		Deleted    bool   `json:"deleted"`
		Validation bool   `json:"validation"`
	}

	Passport struct {
		Document
		Series string `json:"series"`
	}

	Contacts struct {
		Phone []Contact `json:"contact_mobile_registration"`
	}

	Addresses struct {
		Address []Address `json:"addr_registration"`
	}

	Address struct {
		Unom               string `json:"unom"`
		Unad               string `json:"unad"`
		Deleted            bool   `json:"deleted"`
		Validation         bool   `json:"validation"`
		Description        string `json:"description"`
		CityName           string `json:"city_name"`
		Flat               string `json:"flat"`
		StreetNameEx       string `json:"street_name_ex"`
		AddrLineOne        string `json:"addr_line_one"`
		HouseNo            string `json:"house_no"`
		BuildingName       string `json:"building_name"`
		MoscowDistrictName string `json:"moscow_district_name"`
	}

	Ids struct {
		Sso []Sso `json:"SSO"`
	}

	Sso struct {
		Value string
	}

	Contact struct {
		Value   string `json:"value"`
		Deleted bool   `json:"deleted"`
	}
)

func (msg *KriMessage) FindSsoId() (ssoId string) {
	for _, sso := range msg.Ids.Sso {
		return sso.Value
	}
	return ""
}

func (msg *KriMessage) FindPassport() (number, series string) {
	for _, passport := range msg.Documents.Passport {
		if passport.Validation && !passport.Deleted {
			return passport.Value, passport.Series
		}
	}
	return "", ""
}

func (msg *KriMessage) FindSnils() (snils string) {
	for _, snils := range msg.Documents.Passport {
		if snils.Validation && !snils.Deleted {
			return snils.Value
		}
	}
	return ""
}

func (msg *KriMessage) FindPhone() (phone string) {
	for _, phone := range msg.Contacts.Phone {
		if !phone.Deleted {
			return phone.Value
		}
	}
	return ""
}

func (msg *KriMessage) FindRegAddr() Address {
	for i := 0; i < len(msg.Addresses.Address); i++ {
		if msg.Addresses.Address[i].Validation && !msg.Addresses.Address[i].Deleted {
			return msg.Addresses.Address[i]
		}
	}
	return Address{}
}

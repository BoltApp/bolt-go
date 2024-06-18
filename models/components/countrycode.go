// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// CountryCode - The country (in its ISO 3166 alpha-2 format) associated with this address.
type CountryCode string

const (
	CountryCodeAf CountryCode = "AF"
	CountryCodeAx CountryCode = "AX"
	CountryCodeAl CountryCode = "AL"
	CountryCodeDz CountryCode = "DZ"
	CountryCodeAs CountryCode = "AS"
	CountryCodeAd CountryCode = "AD"
	CountryCodeAo CountryCode = "AO"
	CountryCodeAi CountryCode = "AI"
	CountryCodeAq CountryCode = "AQ"
	CountryCodeAg CountryCode = "AG"
	CountryCodeAr CountryCode = "AR"
	CountryCodeAm CountryCode = "AM"
	CountryCodeAw CountryCode = "AW"
	CountryCodeAu CountryCode = "AU"
	CountryCodeAt CountryCode = "AT"
	CountryCodeAz CountryCode = "AZ"
	CountryCodeBh CountryCode = "BH"
	CountryCodeBs CountryCode = "BS"
	CountryCodeBd CountryCode = "BD"
	CountryCodeBb CountryCode = "BB"
	CountryCodeBy CountryCode = "BY"
	CountryCodeBe CountryCode = "BE"
	CountryCodeBz CountryCode = "BZ"
	CountryCodeBj CountryCode = "BJ"
	CountryCodeBm CountryCode = "BM"
	CountryCodeBt CountryCode = "BT"
	CountryCodeBo CountryCode = "BO"
	CountryCodeBq CountryCode = "BQ"
	CountryCodeBa CountryCode = "BA"
	CountryCodeBw CountryCode = "BW"
	CountryCodeBv CountryCode = "BV"
	CountryCodeBr CountryCode = "BR"
	CountryCodeIo CountryCode = "IO"
	CountryCodeBn CountryCode = "BN"
	CountryCodeBg CountryCode = "BG"
	CountryCodeBf CountryCode = "BF"
	CountryCodeBi CountryCode = "BI"
	CountryCodeKh CountryCode = "KH"
	CountryCodeCm CountryCode = "CM"
	CountryCodeCa CountryCode = "CA"
	CountryCodeCv CountryCode = "CV"
	CountryCodeKy CountryCode = "KY"
	CountryCodeCf CountryCode = "CF"
	CountryCodeTd CountryCode = "TD"
	CountryCodeCl CountryCode = "CL"
	CountryCodeCn CountryCode = "CN"
	CountryCodeCx CountryCode = "CX"
	CountryCodeCc CountryCode = "CC"
	CountryCodeCo CountryCode = "CO"
	CountryCodeKm CountryCode = "KM"
	CountryCodeCg CountryCode = "CG"
	CountryCodeCd CountryCode = "CD"
	CountryCodeCk CountryCode = "CK"
	CountryCodeCr CountryCode = "CR"
	CountryCodeCi CountryCode = "CI"
	CountryCodeHr CountryCode = "HR"
	CountryCodeCu CountryCode = "CU"
	CountryCodeCw CountryCode = "CW"
	CountryCodeCy CountryCode = "CY"
	CountryCodeCz CountryCode = "CZ"
	CountryCodeDk CountryCode = "DK"
	CountryCodeDj CountryCode = "DJ"
	CountryCodeDm CountryCode = "DM"
	CountryCodeDo CountryCode = "DO"
	CountryCodeEc CountryCode = "EC"
	CountryCodeEg CountryCode = "EG"
	CountryCodeSv CountryCode = "SV"
	CountryCodeGq CountryCode = "GQ"
	CountryCodeEr CountryCode = "ER"
	CountryCodeEe CountryCode = "EE"
	CountryCodeEt CountryCode = "ET"
	CountryCodeFk CountryCode = "FK"
	CountryCodeFo CountryCode = "FO"
	CountryCodeFj CountryCode = "FJ"
	CountryCodeFi CountryCode = "FI"
	CountryCodeFr CountryCode = "FR"
	CountryCodeGf CountryCode = "GF"
	CountryCodePf CountryCode = "PF"
	CountryCodeTf CountryCode = "TF"
	CountryCodeGa CountryCode = "GA"
	CountryCodeGm CountryCode = "GM"
	CountryCodeGe CountryCode = "GE"
	CountryCodeDe CountryCode = "DE"
	CountryCodeGh CountryCode = "GH"
	CountryCodeGi CountryCode = "GI"
	CountryCodeGr CountryCode = "GR"
	CountryCodeGl CountryCode = "GL"
	CountryCodeGd CountryCode = "GD"
	CountryCodeGp CountryCode = "GP"
	CountryCodeGu CountryCode = "GU"
	CountryCodeGt CountryCode = "GT"
	CountryCodeGg CountryCode = "GG"
	CountryCodeGn CountryCode = "GN"
	CountryCodeGw CountryCode = "GW"
	CountryCodeGy CountryCode = "GY"
	CountryCodeHt CountryCode = "HT"
	CountryCodeHm CountryCode = "HM"
	CountryCodeVa CountryCode = "VA"
	CountryCodeHn CountryCode = "HN"
	CountryCodeHk CountryCode = "HK"
	CountryCodeHu CountryCode = "HU"
	CountryCodeIs CountryCode = "IS"
	CountryCodeIn CountryCode = "IN"
	CountryCodeID CountryCode = "ID"
	CountryCodeIr CountryCode = "IR"
	CountryCodeIq CountryCode = "IQ"
	CountryCodeIe CountryCode = "IE"
	CountryCodeIm CountryCode = "IM"
	CountryCodeIl CountryCode = "IL"
	CountryCodeIt CountryCode = "IT"
	CountryCodeJm CountryCode = "JM"
	CountryCodeJp CountryCode = "JP"
	CountryCodeJe CountryCode = "JE"
	CountryCodeJo CountryCode = "JO"
	CountryCodeKz CountryCode = "KZ"
	CountryCodeKe CountryCode = "KE"
	CountryCodeKi CountryCode = "KI"
	CountryCodeKp CountryCode = "KP"
	CountryCodeKr CountryCode = "KR"
	CountryCodeKw CountryCode = "KW"
	CountryCodeKg CountryCode = "KG"
	CountryCodeLa CountryCode = "LA"
	CountryCodeLv CountryCode = "LV"
	CountryCodeLb CountryCode = "LB"
	CountryCodeLs CountryCode = "LS"
	CountryCodeLr CountryCode = "LR"
	CountryCodeLy CountryCode = "LY"
	CountryCodeLi CountryCode = "LI"
	CountryCodeLt CountryCode = "LT"
	CountryCodeLu CountryCode = "LU"
	CountryCodeMo CountryCode = "MO"
	CountryCodeMk CountryCode = "MK"
	CountryCodeMg CountryCode = "MG"
	CountryCodeMw CountryCode = "MW"
	CountryCodeMy CountryCode = "MY"
	CountryCodeMv CountryCode = "MV"
	CountryCodeMl CountryCode = "ML"
	CountryCodeMt CountryCode = "MT"
	CountryCodeMh CountryCode = "MH"
	CountryCodeMq CountryCode = "MQ"
	CountryCodeMr CountryCode = "MR"
	CountryCodeMu CountryCode = "MU"
	CountryCodeYt CountryCode = "YT"
	CountryCodeMx CountryCode = "MX"
	CountryCodeFm CountryCode = "FM"
	CountryCodeMd CountryCode = "MD"
	CountryCodeMc CountryCode = "MC"
	CountryCodeMn CountryCode = "MN"
	CountryCodeMe CountryCode = "ME"
	CountryCodeMs CountryCode = "MS"
	CountryCodeMa CountryCode = "MA"
	CountryCodeMz CountryCode = "MZ"
	CountryCodeMm CountryCode = "MM"
	CountryCodeNa CountryCode = "NA"
	CountryCodeNr CountryCode = "NR"
	CountryCodeNp CountryCode = "NP"
	CountryCodeNl CountryCode = "NL"
	CountryCodeNc CountryCode = "NC"
	CountryCodeNz CountryCode = "NZ"
	CountryCodeNi CountryCode = "NI"
	CountryCodeNe CountryCode = "NE"
	CountryCodeNg CountryCode = "NG"
	CountryCodeNu CountryCode = "NU"
	CountryCodeNf CountryCode = "NF"
	CountryCodeMp CountryCode = "MP"
	CountryCodeNo CountryCode = "NO"
	CountryCodeOm CountryCode = "OM"
	CountryCodePk CountryCode = "PK"
	CountryCodePw CountryCode = "PW"
	CountryCodePs CountryCode = "PS"
	CountryCodePa CountryCode = "PA"
	CountryCodePg CountryCode = "PG"
	CountryCodePy CountryCode = "PY"
	CountryCodePe CountryCode = "PE"
	CountryCodePh CountryCode = "PH"
	CountryCodePn CountryCode = "PN"
	CountryCodePl CountryCode = "PL"
	CountryCodePt CountryCode = "PT"
	CountryCodePr CountryCode = "PR"
	CountryCodeQa CountryCode = "QA"
	CountryCodeRe CountryCode = "RE"
	CountryCodeRo CountryCode = "RO"
	CountryCodeRu CountryCode = "RU"
	CountryCodeRw CountryCode = "RW"
	CountryCodeBl CountryCode = "BL"
	CountryCodeSh CountryCode = "SH"
	CountryCodeKn CountryCode = "KN"
	CountryCodeLc CountryCode = "LC"
	CountryCodeMf CountryCode = "MF"
	CountryCodePm CountryCode = "PM"
	CountryCodeVc CountryCode = "VC"
	CountryCodeWs CountryCode = "WS"
	CountryCodeSm CountryCode = "SM"
	CountryCodeSt CountryCode = "ST"
	CountryCodeSa CountryCode = "SA"
	CountryCodeSn CountryCode = "SN"
	CountryCodeRs CountryCode = "RS"
	CountryCodeSc CountryCode = "SC"
	CountryCodeSl CountryCode = "SL"
	CountryCodeSg CountryCode = "SG"
	CountryCodeSx CountryCode = "SX"
	CountryCodeSk CountryCode = "SK"
	CountryCodeSi CountryCode = "SI"
	CountryCodeSb CountryCode = "SB"
	CountryCodeSo CountryCode = "SO"
	CountryCodeZa CountryCode = "ZA"
	CountryCodeGs CountryCode = "GS"
	CountryCodeSs CountryCode = "SS"
	CountryCodeEs CountryCode = "ES"
	CountryCodeLk CountryCode = "LK"
	CountryCodeSd CountryCode = "SD"
	CountryCodeSr CountryCode = "SR"
	CountryCodeSj CountryCode = "SJ"
	CountryCodeSz CountryCode = "SZ"
	CountryCodeSe CountryCode = "SE"
	CountryCodeCh CountryCode = "CH"
	CountryCodeSy CountryCode = "SY"
	CountryCodeTw CountryCode = "TW"
	CountryCodeTj CountryCode = "TJ"
	CountryCodeTz CountryCode = "TZ"
	CountryCodeTh CountryCode = "TH"
	CountryCodeTl CountryCode = "TL"
	CountryCodeTg CountryCode = "TG"
	CountryCodeTk CountryCode = "TK"
	CountryCodeTo CountryCode = "TO"
	CountryCodeTt CountryCode = "TT"
	CountryCodeTn CountryCode = "TN"
	CountryCodeTr CountryCode = "TR"
	CountryCodeTm CountryCode = "TM"
	CountryCodeTc CountryCode = "TC"
	CountryCodeTv CountryCode = "TV"
	CountryCodeUg CountryCode = "UG"
	CountryCodeUa CountryCode = "UA"
	CountryCodeAe CountryCode = "AE"
	CountryCodeGb CountryCode = "GB"
	CountryCodeUs CountryCode = "US"
	CountryCodeUm CountryCode = "UM"
	CountryCodeUy CountryCode = "UY"
	CountryCodeUz CountryCode = "UZ"
	CountryCodeVu CountryCode = "VU"
	CountryCodeVe CountryCode = "VE"
	CountryCodeVn CountryCode = "VN"
	CountryCodeVg CountryCode = "VG"
	CountryCodeVi CountryCode = "VI"
	CountryCodeWf CountryCode = "WF"
	CountryCodeEh CountryCode = "EH"
	CountryCodeYe CountryCode = "YE"
	CountryCodeZm CountryCode = "ZM"
	CountryCodeZw CountryCode = "ZW"
)

func (e CountryCode) ToPointer() *CountryCode {
	return &e
}
func (e *CountryCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AF":
		fallthrough
	case "AX":
		fallthrough
	case "AL":
		fallthrough
	case "DZ":
		fallthrough
	case "AS":
		fallthrough
	case "AD":
		fallthrough
	case "AO":
		fallthrough
	case "AI":
		fallthrough
	case "AQ":
		fallthrough
	case "AG":
		fallthrough
	case "AR":
		fallthrough
	case "AM":
		fallthrough
	case "AW":
		fallthrough
	case "AU":
		fallthrough
	case "AT":
		fallthrough
	case "AZ":
		fallthrough
	case "BH":
		fallthrough
	case "BS":
		fallthrough
	case "BD":
		fallthrough
	case "BB":
		fallthrough
	case "BY":
		fallthrough
	case "BE":
		fallthrough
	case "BZ":
		fallthrough
	case "BJ":
		fallthrough
	case "BM":
		fallthrough
	case "BT":
		fallthrough
	case "BO":
		fallthrough
	case "BQ":
		fallthrough
	case "BA":
		fallthrough
	case "BW":
		fallthrough
	case "BV":
		fallthrough
	case "BR":
		fallthrough
	case "IO":
		fallthrough
	case "BN":
		fallthrough
	case "BG":
		fallthrough
	case "BF":
		fallthrough
	case "BI":
		fallthrough
	case "KH":
		fallthrough
	case "CM":
		fallthrough
	case "CA":
		fallthrough
	case "CV":
		fallthrough
	case "KY":
		fallthrough
	case "CF":
		fallthrough
	case "TD":
		fallthrough
	case "CL":
		fallthrough
	case "CN":
		fallthrough
	case "CX":
		fallthrough
	case "CC":
		fallthrough
	case "CO":
		fallthrough
	case "KM":
		fallthrough
	case "CG":
		fallthrough
	case "CD":
		fallthrough
	case "CK":
		fallthrough
	case "CR":
		fallthrough
	case "CI":
		fallthrough
	case "HR":
		fallthrough
	case "CU":
		fallthrough
	case "CW":
		fallthrough
	case "CY":
		fallthrough
	case "CZ":
		fallthrough
	case "DK":
		fallthrough
	case "DJ":
		fallthrough
	case "DM":
		fallthrough
	case "DO":
		fallthrough
	case "EC":
		fallthrough
	case "EG":
		fallthrough
	case "SV":
		fallthrough
	case "GQ":
		fallthrough
	case "ER":
		fallthrough
	case "EE":
		fallthrough
	case "ET":
		fallthrough
	case "FK":
		fallthrough
	case "FO":
		fallthrough
	case "FJ":
		fallthrough
	case "FI":
		fallthrough
	case "FR":
		fallthrough
	case "GF":
		fallthrough
	case "PF":
		fallthrough
	case "TF":
		fallthrough
	case "GA":
		fallthrough
	case "GM":
		fallthrough
	case "GE":
		fallthrough
	case "DE":
		fallthrough
	case "GH":
		fallthrough
	case "GI":
		fallthrough
	case "GR":
		fallthrough
	case "GL":
		fallthrough
	case "GD":
		fallthrough
	case "GP":
		fallthrough
	case "GU":
		fallthrough
	case "GT":
		fallthrough
	case "GG":
		fallthrough
	case "GN":
		fallthrough
	case "GW":
		fallthrough
	case "GY":
		fallthrough
	case "HT":
		fallthrough
	case "HM":
		fallthrough
	case "VA":
		fallthrough
	case "HN":
		fallthrough
	case "HK":
		fallthrough
	case "HU":
		fallthrough
	case "IS":
		fallthrough
	case "IN":
		fallthrough
	case "ID":
		fallthrough
	case "IR":
		fallthrough
	case "IQ":
		fallthrough
	case "IE":
		fallthrough
	case "IM":
		fallthrough
	case "IL":
		fallthrough
	case "IT":
		fallthrough
	case "JM":
		fallthrough
	case "JP":
		fallthrough
	case "JE":
		fallthrough
	case "JO":
		fallthrough
	case "KZ":
		fallthrough
	case "KE":
		fallthrough
	case "KI":
		fallthrough
	case "KP":
		fallthrough
	case "KR":
		fallthrough
	case "KW":
		fallthrough
	case "KG":
		fallthrough
	case "LA":
		fallthrough
	case "LV":
		fallthrough
	case "LB":
		fallthrough
	case "LS":
		fallthrough
	case "LR":
		fallthrough
	case "LY":
		fallthrough
	case "LI":
		fallthrough
	case "LT":
		fallthrough
	case "LU":
		fallthrough
	case "MO":
		fallthrough
	case "MK":
		fallthrough
	case "MG":
		fallthrough
	case "MW":
		fallthrough
	case "MY":
		fallthrough
	case "MV":
		fallthrough
	case "ML":
		fallthrough
	case "MT":
		fallthrough
	case "MH":
		fallthrough
	case "MQ":
		fallthrough
	case "MR":
		fallthrough
	case "MU":
		fallthrough
	case "YT":
		fallthrough
	case "MX":
		fallthrough
	case "FM":
		fallthrough
	case "MD":
		fallthrough
	case "MC":
		fallthrough
	case "MN":
		fallthrough
	case "ME":
		fallthrough
	case "MS":
		fallthrough
	case "MA":
		fallthrough
	case "MZ":
		fallthrough
	case "MM":
		fallthrough
	case "NA":
		fallthrough
	case "NR":
		fallthrough
	case "NP":
		fallthrough
	case "NL":
		fallthrough
	case "NC":
		fallthrough
	case "NZ":
		fallthrough
	case "NI":
		fallthrough
	case "NE":
		fallthrough
	case "NG":
		fallthrough
	case "NU":
		fallthrough
	case "NF":
		fallthrough
	case "MP":
		fallthrough
	case "NO":
		fallthrough
	case "OM":
		fallthrough
	case "PK":
		fallthrough
	case "PW":
		fallthrough
	case "PS":
		fallthrough
	case "PA":
		fallthrough
	case "PG":
		fallthrough
	case "PY":
		fallthrough
	case "PE":
		fallthrough
	case "PH":
		fallthrough
	case "PN":
		fallthrough
	case "PL":
		fallthrough
	case "PT":
		fallthrough
	case "PR":
		fallthrough
	case "QA":
		fallthrough
	case "RE":
		fallthrough
	case "RO":
		fallthrough
	case "RU":
		fallthrough
	case "RW":
		fallthrough
	case "BL":
		fallthrough
	case "SH":
		fallthrough
	case "KN":
		fallthrough
	case "LC":
		fallthrough
	case "MF":
		fallthrough
	case "PM":
		fallthrough
	case "VC":
		fallthrough
	case "WS":
		fallthrough
	case "SM":
		fallthrough
	case "ST":
		fallthrough
	case "SA":
		fallthrough
	case "SN":
		fallthrough
	case "RS":
		fallthrough
	case "SC":
		fallthrough
	case "SL":
		fallthrough
	case "SG":
		fallthrough
	case "SX":
		fallthrough
	case "SK":
		fallthrough
	case "SI":
		fallthrough
	case "SB":
		fallthrough
	case "SO":
		fallthrough
	case "ZA":
		fallthrough
	case "GS":
		fallthrough
	case "SS":
		fallthrough
	case "ES":
		fallthrough
	case "LK":
		fallthrough
	case "SD":
		fallthrough
	case "SR":
		fallthrough
	case "SJ":
		fallthrough
	case "SZ":
		fallthrough
	case "SE":
		fallthrough
	case "CH":
		fallthrough
	case "SY":
		fallthrough
	case "TW":
		fallthrough
	case "TJ":
		fallthrough
	case "TZ":
		fallthrough
	case "TH":
		fallthrough
	case "TL":
		fallthrough
	case "TG":
		fallthrough
	case "TK":
		fallthrough
	case "TO":
		fallthrough
	case "TT":
		fallthrough
	case "TN":
		fallthrough
	case "TR":
		fallthrough
	case "TM":
		fallthrough
	case "TC":
		fallthrough
	case "TV":
		fallthrough
	case "UG":
		fallthrough
	case "UA":
		fallthrough
	case "AE":
		fallthrough
	case "GB":
		fallthrough
	case "US":
		fallthrough
	case "UM":
		fallthrough
	case "UY":
		fallthrough
	case "UZ":
		fallthrough
	case "VU":
		fallthrough
	case "VE":
		fallthrough
	case "VN":
		fallthrough
	case "VG":
		fallthrough
	case "VI":
		fallthrough
	case "WF":
		fallthrough
	case "EH":
		fallthrough
	case "YE":
		fallthrough
	case "ZM":
		fallthrough
	case "ZW":
		*e = CountryCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CountryCode: %v", v)
	}
}

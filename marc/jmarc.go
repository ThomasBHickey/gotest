package jmarc

import(
    "encoding/xml"
    "fmt"
)

type MARC21 struct {
    XMLName xml.Name `xml:"http://www.loc.gov/MARC21/slim record"`
    Leader LeaderT `xml:"leader"`
    //Leader string `xml:"leader,chardata"`
    ControlFields [] ControlFieldT `xml:"controlfield"`
    DataFields [] DataFieldT `xml:"datafield"`
}

type LeaderT struct {
    Leader string `xml:",chardata"`
}

type ControlFieldT struct{
    Tag string `xml:"tag,attr"`
    Data string `xml:",chardata"` 
}

type DataFieldT struct {
    Tag string `xml:"tag,attr"`
    Ind1 string `xml:"ind1,attr"`
    Ind2 string `xml:"ind2,attr"`
    Subfield [] SubfieldT `xml:"subfield"`
}

type SubfieldT struct{
    Code string `xml:"code,attr"`
    Data string `xml:",chardata"`
}

type outXML struct {
    XMLName xml.Name `xml:"record"`
}

func CreateFromXML(xmldata []byte)(MARC21, error){
    //v := MARC21 {}
    var v MARC21
    if err := xml.Unmarshal([]byte(xmldata), &v); err != nil {
        fmt.Printf("error: %v", err)
        return v, err
    }
    return v, nil
 }

func GenerateXML(v MARC21)([]byte, error){
    output, err := xml.MarshalIndent(v, " ", "  ")
    if err != nil {
        fmt.Printf("Error creating XML", err)
    }
    return output, err
}

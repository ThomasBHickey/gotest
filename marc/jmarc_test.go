package jmarc

import (
    "testing"
    "fmt"
    "io/ioutil"
)

func TestXMLImport(t *testing.T){
    fmt.Println("In TestXMLImport")
    dat, err := ioutil.ReadFile("test.marc.xml")
    if err != nil {
        t.Errorf("Unable to read test file 'test.marc.xml")
    }
    marc, err := CreateFromXML(dat)
    if err != nil {
        t.Errorf("CreateFromXML failed")
    }
    fmt.Printf("marc: %#v\n", marc)
    fmt.Printf("leader %#v\n", marc.Leader)
    //fmt.Printf("marc.Leader.Leader %#v\n", marc.Leader.Leader)
    for _, cf := range marc.ControlFields {
        fmt.Printf("cf tag %#v\n", cf.Tag)
    }
    fmt.Println("control fields", marc.ControlFields)
    if newxml, err := GenerateXML(marc); err!=nil{
        t.Errorf("GenerateXML failed")
    } else{
        fmt.Println("newxml", string(newxml))

    }
}
package buzzwords

import (
  "testing"
  "fmt"
  "encoding/json"
  "bytes"
)

func TestReadFile(t *testing.T) {
  tt := []struct {
    name string
    path string
    expectError bool
  } {
    {"open existing json-file", "./testdata/good.json", false},
    {"open non-existent file", "./testdata/nonexistentfile", true},
    {"open bad-json-file", "./testdata/bad.json", true},
  }

  for _, tc := range tt {
    t.Run(tc.name, func(t *testing.T) {
      err := readFile(tc.path)
      if err != nil {
        if tc.expectError == false {
          t.Errorf("Did not expect error: %v\n", err.Error())
        }
      } else if tc.expectError == true {
        t.Error("Expected error but didn't get one")
      }
    })
  }
}

func TestInit(t *testing.T) {
  tt := []struct {
    name string
    path string
    expectError bool
  } {
    {"initialize with existing file", "./testdata/good.json", false},
    {"initialize with non-existent file", "./testdata/nonexistentfile", true},
  }

  for _, tc := range tt {
    t.Run(tc.name, func(t *testing.T) {
      err := Init(tc.path)
      if err != nil {
        if tc.expectError == false {
          t.Errorf("Did not expect error: %v\n", err.Error())
        }
      } else if tc.expectError == true {
        t.Error("Expected error but didn't get one")
      }
    })
  }
}

func TestGetRandomWords(t *testing.T) {
  s := GetRandomWords()
  if s == "" {
    t.Errorf("Expected output but got empty string")
  }
}

func TestFormatHTML(t *testing.T) {
  b := "IoT"
  e := fmt.Sprintf("<html><head><title>Buzzwords As a Service</title></head><body><center style=\"font-family: monospace\"><h3>%s</h3></center></body></html>", b)

  bs := FormatHTML(b)
  if bs != e {
    t.Errorf("\nExpected:   %s\nBut got:    %s", e, bs)
  }
}

func TestFormatJSON(t *testing.T) {
  b := "IoT"
  o := Buzzword{b}

  om, err := json.Marshal(o)
  if err != nil {
    t.Errorf("Error marshalling reference json: %v", err)
  }

  bs, err := FormatJSON(b)
  if err != nil {
    t.Errorf("Error formatting JSON: %v", err)
  }

  if !bytes.Equal(om, bs){
    t.Errorf("Mashalled JSON objects are not the same:\nGot:   %s\nExpected: %s", bs, om)
  }


}

package buzzwords

import (
  "testing"
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

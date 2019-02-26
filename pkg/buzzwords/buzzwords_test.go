package buzzwords

import (
  "testing"
)

func TestBuzzwords(t *testing.T) {
  tt := []struct {
    name string
    path string
    expectError bool
  } {
    {"open existing json-file", "../../buzzwords.json", false},
    {"open non-existent file", "./nonexistentfile", true},
    {"open non-json-file", "./buzzwords.go", true},
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

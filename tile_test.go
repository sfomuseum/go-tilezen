package tilezen

import (
	"testing"
)

func TestParseURI(t *testing.T) {

	tests := map[string][3]int{
		"tiles/11/327/792.mvt": [3]int{11, 327, 792},
	}

	for uri, zxy := range tests {

		new_t, err := ParseURI(uri)

		if err != nil {
			t.Fatal(err)
		}

		expected_z := zxy[0]
		expected_x := zxy[1]
		expected_y := zxy[2]

		if new_t.Z != expected_z {
			t.Fatal("Failed to parse Z")
		}

		if new_t.X != expected_x {
			t.Fatal("Failed to parse X")
		}

		if new_t.Y != expected_y {
			t.Fatal("Failed to parse Y")
		}

	}

}

package elastic_test

import (
	"encoding/json"
	"math"
	"testing"

	testtarget "github.com/olivere/elastic/v7"
)

func TestMarshalScore(t *testing.T) {
	type testset struct {
		Title string
		Value interface{}
		Want  string
	}

	type hasEntity struct {
		Field testtarget.Score
	}
	type hasPointer struct {
		Field *testtarget.Score
	}

	null := (*testtarget.Score)(nil)
	zero := testtarget.Score(0)
	valid := testtarget.Score(42.195)
	posiInfinity := testtarget.Score(math.Inf(0))
	negaInfinity := testtarget.Score(math.Inf(-1))

	cases := []testset{{
		Title: "zero",
		Value: zero,
		Want:  "0",
	}, {
		Title: "nil",
		Value: null,
		Want:  "null",
	}, {
		Title: "zero in a structure",
		Value: hasEntity{zero},
		Want:  `{"Field":0}`,
	}, {
		Title: "nil in a structure",
		Value: hasPointer{null},
		Want:  `{"Field":null}`,
	}, {
		Title: "omit default value in a structure",
		Value: struct {
			Field testtarget.Score `json:",omitempty"`
		}{zero},
		Want: `{}`,
	}, {
		Title: "omit nil in a structure",
		Value: struct {
			Field *testtarget.Score `json:",omitempty"`
		}{null},
		Want: `{}`,
	}, {
		Title: "valid value",
		Value: valid,
		Want:  "42.195",
	}, {
		Title: "pointer having valid value",
		Value: &valid,
		Want:  "42.195",
	}, {
		Title: "valid value in a structure",
		Value: hasEntity{valid},
		Want:  `{"Field":42.195}`,
	}, {
		Title: "pointer having valid value in a structure",
		Value: hasPointer{&valid},
		Want:  `{"Field":42.195}`,
	}, {
		Title: "infinity",
		Value: posiInfinity,
		Want:  `"Infinity"`,
	}, {
		Title: "negative infinity",
		Value: negaInfinity,
		Want:  `"-Infinity"`,
	}, {
		Title: "pointer having infinity",
		Value: &posiInfinity,
		Want:  `"Infinity"`,
	}, {
		Title: "pointer having negative inifinity",
		Value: &negaInfinity,
		Want:  `"-Infinity"`,
	}, {
		Title: "infinity in a structure",
		Value: hasEntity{posiInfinity},
		Want:  `{"Field":"Infinity"}`,
	}, {
		Title: "negative infinity in a structure",
		Value: hasEntity{negaInfinity},
		Want:  `{"Field":"-Infinity"}`,
	}, {
		Title: "pointer having infinity in a structure",
		Value: hasPointer{&posiInfinity},
		Want:  `{"Field":"Infinity"}`,
	}, {
		Title: "pointer having negative inifinity in a structure",
		Value: hasPointer{&negaInfinity},
		Want:  `{"Field":"-Infinity"}`,
	}}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			got, err := json.Marshal(c.Value)
			if err != nil {
				t.Fatalf("failed to marshal: %#v", err)
			}
			if string(got) != c.Want {
				t.Errorf("marshalled value %q does not equal to wanted value %q", got, c.Want)
			}
		})
	}
}

func TestUnmarshalScore(t *testing.T) {
	type testset struct {
		Title string
		Value string
		Want  interface{}
	}

	type hasEntity struct {
		Field testtarget.Score
	}
	type hasPointer struct {
		Field *testtarget.Score
	}

	null := (*testtarget.Score)(nil)
	zero := testtarget.Score(0)
	valid := testtarget.Score(42.195)
	posiInfinity := testtarget.Score(math.Inf(0))
	negaInfinity := testtarget.Score(math.Inf(-1))

	t.Run("zero", func(t *testing.T) {
		var value = "0"
		var want = zero
		var got testtarget.Score
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("nil", func(t *testing.T) {
		var value = "null"
		var want = null
		var got *testtarget.Score
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("zero in a structure", func(t *testing.T) {
		var value = `{"Field":0}`
		var want = hasEntity{zero}
		var got hasEntity
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("nil in a structure", func(t *testing.T) {
		var value = `{"Field":null}`
		var want = hasPointer{null}
		var got hasPointer
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("valid value", func(t *testing.T) {
		var value = "42.195"
		var want = valid
		var got testtarget.Score
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("pointer may have valid value", func(t *testing.T) {
		var value = "42.195"
		var want = valid
		var got *testtarget.Score
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got == nil {
			t.Errorf("unmarshalled value %#v is nil", got)
		}
		if *got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", *got, want)
		}
	})

	t.Run("valid value in a structure", func(t *testing.T) {
		var want = hasEntity{valid}
		var value = `{"Field":42.195}`
		var got hasEntity
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("pointer may have valid value in a structure", func(t *testing.T) {
		var want = valid
		var value = `{"Field":42.195}`
		var got hasPointer
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got.Field == nil {
			t.Errorf("unmarshalled field %#v is nil", got)
		}
		if *got.Field != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", *got.Field, want)
		}
	})

	t.Run("infinity", func(t *testing.T) {
		var want = posiInfinity
		var value = `"Infinity"`
		var got testtarget.Score
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("negative infinity", func(t *testing.T) {
		var want = negaInfinity
		var value = `"-Infinity"`
		var got testtarget.Score
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("pointer may have infinity", func(t *testing.T) {
		var want = posiInfinity
		var value = `"Infinity"`
		var got *testtarget.Score
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got == nil {
			t.Errorf("unmarshalled value %#v is nil", got)
		}
		if *got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", *got, want)
		}
	})

	t.Run("pointer may have negative inifinity", func(t *testing.T) {
		var want = negaInfinity
		var value = `"-Infinity"`
		var got *testtarget.Score
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got == nil {
			t.Errorf("unmarshalled value %#v is nil", got)
		}
		if *got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", *got, want)
		}
	})

	t.Run("infinity in a structure", func(t *testing.T) {
		var want = hasEntity{posiInfinity}
		var value = `{"Field":"Infinity"}`
		var got hasEntity
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("pointer may have infinity in a structure", func(t *testing.T) {
		var want = posiInfinity
		var value = `{"Field":"Infinity"}`
		var got hasPointer
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got.Field == nil {
			t.Errorf("unmarshalled field %#v is nil", got)
		}
		if *got.Field != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", *got.Field, want)
		}
	})

	t.Run("negative infinity in a structure", func(t *testing.T) {
		var want = hasEntity{negaInfinity}
		var value = `{"Field":"-Infinity"}`
		var got hasEntity
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", got, want)
		}
	})

	t.Run("pointer may have negative infinity in a structure", func(t *testing.T) {
		var want = negaInfinity
		var value = `{"Field":"-Infinity"}`
		var got hasPointer
		if err := json.Unmarshal([]byte(value), &got); err != nil {
			t.Fatalf("failed to unmarshal %q", value)
		}
		if got.Field == nil {
			t.Errorf("unmarshalled field %#v is nil", got)
		}
		if *got.Field != want {
			t.Errorf("unmarshalled value %#v does not equal to wanted value %#v", *got.Field, want)
		}
	})

}

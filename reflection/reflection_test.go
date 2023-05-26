package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	D    Data
}

type Data struct {
	Location string
	Thing    uint
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			Name:     "Single string field",
			Input:    Person{Name: "Eric", D: Data{Location: "b"}},
			Expected: []string{"Eric", "b"},
		},
		{
			Name:     "Multiple fields",
			Input:    Person{Name: "Eric", D: Data{Location: "k"}},
			Expected: []string{"Eric", "k"},
		},
		{
			Name:     "Nested structs",
			Input:    Person{"Eric", Data{"xd", 32}},
			Expected: []string{"Eric", "xd"},
		},
		{
			Name: "But pointers",
			Input: &Person{
				Name: "Eric", D: Data{"lel", 32},
			},
			Expected: []string{"Eric", "lel"},
		},
		{
			Name: "slices",
			Input: []Person{
				{
					Name: "Eric", D: Data{"lel", 32},
				},
				{
					Name: "Gogo", D: Data{"lel", 32},
				},
			},
			Expected: []string{"Eric", "lel", "Gogo", "lel"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(str string) {
				got = append(got, str)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("got %v, want %v", got, test.Expected)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		input := map[string]Person{
			"eric": {
				Name: "Eric", D: Data{"lel", 32},
			},
			"gogo": {
				Name: "Gogo", D: Data{"lel", 32},
			},
		}

		var got []string

		Walk(input, func(str string) {
			got = append(got, str)
		})

		assertContains(t, got, "Eric")
		assertContains(t, got, "lel")
		assertContains(t, got, "Gogo")
	})

	t.Run("chans", func(t *testing.T) {
		c := make(chan Person)

		go func() {
			c <- Person{Name: "Eric", D: Data{"lel", 32}}
			c <- Person{Name: "bos", D: Data{"lel", 32}}
			close(c)
		}()

		var got []string
		want := []string{"Eric", "lel", "bos", "lel"}

		Walk(c, func(inp string) {
			got = append(got, inp)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("funcs", func(t *testing.T) {
		aFunc := func() (Person, Person) {
			return Person{Name: "eric", D: Data{}}, Person{"bro", Data{}}
		}

		var got []string
		want := []string{"eric", "", "bro", ""}

		Walk(aFunc, func(inp string) {
			got = append(got, inp)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, v := range haystack {
		if v == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expecting %v to contain %q", haystack, needle)
	}
}

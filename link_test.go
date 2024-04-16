package link

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetLinks(t *testing.T) {

	var tests = []struct {
		filename string
		want     []Link
	}{
		{
			filename: "testdata/ex1.html",
			want: []Link{
				{
					Href: "/other-page",
					Text: "A link to another page",
				},
			},
		},
		{
			filename: "testdata/ex2.html",
			want: []Link{
				{
					Href: "https://www.twitter.com/joncalhoun",
					Text: "Check me out on twitter",
				},
				{
					Href: "https://github.com/gophercises",
					Text: "Gophercises is on Github!",
				},
			},
		},
		{
			filename: "testdata/ex3.html",
			want: []Link{
				{
					Href: "#",
					Text: "Login",
				},
				{
					Href: "/lost",
					Text: "Lost? Need help?",
				},
				{
					Href: "https://twitter.com/marcusolsson",
					Text: "@marcusolsson",
				},
			},
		},
		{
			filename: "testdata/ex4.html",
			want: []Link{
				{
					Href: "/dog-cat",
					Text: "dog cat",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.filename, func(t *testing.T) {
			f, err := os.Open(tc.filename)
			if err != nil {
				t.Error(err)
			}
			got, err := Parse(f)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
			}
		})
	}

}

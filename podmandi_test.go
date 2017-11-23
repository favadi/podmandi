package podmandi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mmcdole/gofeed"
)

const validFeedURL = "https://changelog.com/gotime/feed"

// testParser is the test parser to use in test.
type testParser struct {
	data map[string][]byte
}

func newTestParser() *testParser {
	return &testParser{data: map[string][]byte{
		validFeedURL: []byte(`
{
  "title": "Go Time",
  "description": "A panel of Go experts and special guests discuss the Go programming language, the community, and everything in between.\n This show records LIVE on Thursdays at 3pm US/Eastern time. Hosts Erik St. Martin, Carlisia Pinto, and Brian Ketelsen welcome guests from around the Go community. Topics include Golang, DevOps, web development, infrastructure, Kubernetes, and more. If you develop in Go – or aspire to – this is the show for you.",
  "link": "http://gotime.fm",
  "author": {
    "name": "Changelog Media"
  },
  "language": "en-us",
  "copyright": "All rights reserved",
  "categories": [
    "go",
    " golang",
    " open source",
    " software",
    " development",
    "Technology",
    "Software How-To"
  ],
  "extensions": {
    "atom": {
      "link": [
        {
          "name": "link",
          "value": "",
          "attrs": {
            "href": "https://changelog.com/gotime",
            "rel": "alternate",
            "type": "application/rss+xml"
          },
          "children": {}
        }
      ]
    },
    "itunes": {
      "author": [
        {
          "name": "author",
          "value": "Changelog Media",
          "attrs": {},
          "children": {}
        }
      ],
      "category": [
        {
          "name": "category",
          "value": "",
          "attrs": {
            "text": "Technology"
          },
          "children": {
            "category": [
              {
                "name": "category",
                "value": "",
                "attrs": {
                  "text": "Software How-To"
                },
                "children": {}
              },
              {
                "name": "category",
                "value": "",
                "attrs": {
                  "text": "Tech News"
                },
                "children": {}
              }
            ]
          }
        }
      ],
      "explicit": [
        {
          "name": "explicit",
          "value": "no",
          "attrs": {},
          "children": {}
        }
      ],
      "image": [
        {
          "name": "image",
          "value": "",
          "attrs": {
            "href": "https://cdn.changelog.com/images/podcasts/gotime-cover-art-2f51f845f9ffc5a946394231c4b41c21.png?vsn=d"
          },
          "children": {}
        }
      ],
      "keywords": [
        {
          "name": "keywords",
          "value": "go, golang, open source, software, development",
          "attrs": {},
          "children": {}
        }
      ],
      "owner": [
        {
          "name": "owner",
          "value": "",
          "attrs": {},
          "children": {
            "email": [
              {
                "name": "email",
                "value": "editors@changelog.com",
                "attrs": {},
                "children": {}
              }
            ],
            "name": [
              {
                "name": "name",
                "value": "Changelog Media",
                "attrs": {},
                "children": {}
              }
            ]
          }
        }
      ],
      "subtitle": [
        {
          "name": "subtitle",
          "value": "Go Time",
          "attrs": {},
          "children": {}
        }
      ],
      "summary": [
        {
          "name": "summary",
          "value": "A panel of Go experts and special guests discuss the Go programming language, the community, and everything in between.\n This show records LIVE on Thursdays at 3pm US/Eastern time. Hosts Erik St. Martin, Carlisia Pinto, and Brian Ketelsen welcome guests from around the Go community. Topics include Golang, DevOps, web development, infrastructure, Kubernetes, and more. If you develop in Go – or aspire to – this is the show for you.",
          "attrs": {},
          "children": {}
        }
      ]
    }
  },
  "items": [
    {
      "title": "60: Why WADL When You Can Swagger? with Ivan Porto Carrero",
      "description": "<p>Ivan Porto Carrero joined the show to talk about generating documentation (with Swagger), pks, kubo, and other interesting Go projects and news.</p>\n\n  <p>Sponsors</p>\n  <ul>\n    <li>\n      <a href=\"https://linode.com/changelog\">Linode</a> – \n<strong>Our cloud server of choice.</strong> Get one of the fastest, most efficient SSD cloud servers for only $5/mo. Use the code <code>changelog2017</code> to get 4 months free!\n    </li>\n    <li>\n      <a href=\"https://www.fastly.com/?utm_source=changelog&amp;utm_medium=podcast&amp;utm_campaign=changelog-sponsorship\">Fastly</a> – \n<strong>Our bandwidth partner.</strong> Fastly powers fast, secure, and scalable digital experiences. Move beyond your content delivery network to their powerful edge cloud platform.\n    </li>\n  </ul>\n\n  <p>Featuring</p>\n  <ul>\n    <li>Ivan Porto Carrero &ndash; <a href=\"https://twitter.com/casualjim\">Twitter</a>, <a href=\"https://github.com/casualjim\">GitHub</a>, <a href=\"http://flanders.co.nz\">Website</a></li>\n    <li>Erik St. Martin &ndash; <a href=\"https://twitter.com/erikstmartin\">Twitter</a>, <a href=\"https://github.com/erikstmartin\">GitHub</a></li>\n    <li>Carlisia Pinto &ndash; <a href=\"https://twitter.com/carlisia\">Twitter</a>, <a href=\"https://github.com/carlisia\">GitHub</a></li>\n    <li>Brian Ketelsen &ndash; <a href=\"https://twitter.com/bketelsen\">Twitter</a>, <a href=\"https://github.com/bketelsen\">GitHub</a></li>\n  </ul>\n<p>Notes and Links</p>\n<p><a href=\"https://github.com/go-swagger/go-swagger\">go-swagger</a></p>\n<p><a href=\"https://github.com/scalatra/scalatra\">scalatra</a></p>\n<p><a href=\"https://github.com/go-openapi/spec\">openapi specification object model</a></p>\n<hr />\n<h3>Interesting Go Projects and News</h3>\n<p><a href=\"https://groups.google.com/forum/#!topic/golang-announce/s_hLxKF9ApA\">Go 1.9.2 released</a>\n<em>&quot;These releases include fixes to the compiler, linker, runtime, documentation, go command, and the crypto/x509, database/sql, log, and net/smtp packages. They include a fix to a bug introduced in Go 1.9.1 and Go 1.8.4 that broke &quot;go get&quot; of non-Git repositories under certain conditions.&quot;</em></p>\n<p><a href=\"https://github.com/contribsys/faktory\">Factory</a> (by <a href=\"http://www.mikeperham.com/\">Mike Perham</a>, <a href=\"https://github.com/mperham/sidekiq\">Sidekiq</a> author)</p>\n<p><a href=\"https://github.com/IMQS/authaus\">Authentication Haus</a></p>\n<p><a href=\"https://github.com/rgburke/grv\">GRV - Git view on command line</a></p>\n<p><a href=\"https://www.youtube.com/watch?v=ySy3sR1LFCQ\">Video: Using the Go Tracer</a></p>\n<p><a href=\"https://github.com/golang/dep/releases/tag/v0.3.2\">Dep 0.3.2</a> - support for importing from additional dependency management tools (gvt, gb), various other fixes and improvements</p>\n<p><a href=\"https://github.com/hybridgroup/gobot/releases/tag/v1.7.0\">GoBot 1.7.0</a></p>\n<p><a href=\"https://www.goinggo.net/2017/10/the-behavior-of-channels.html\">The Behavior Of Channels</a></p>\n<p><a href=\"https://github.com/meqaio/swagger_meqa\">Automated testing for swagger api</a></p>\n<hr />\n<h3>Free Software Friday!</h3>\n<p>Each week on the show we give a shout out to an open source project or community (or maintainer) that's made an impact in our day to day developer lives.</p>\n<p>Erik - <a href=\"https://www.gonum.org/\">Gonum Numerical Packages</a></p>\n<p>Carlisia - <a href=\"https://twitter.com/goinggodotnet\">Bill Kennedy</a></p>\n<p>Brian  - <a href=\"https://twitter.com/francesc\">Francesc Campoy</a></p>\n",
      "link": "https://changelog.com/gotime/60",
      "published": "Fri, 17 Nov 2017 22:50:38 +0000",
      "publishedParsed": "2017-11-17T22:50:38Z",
      "author": {
        "name": "Erik St. Martin, Carlisia Pinto, and Brian Ketelsen"
      },
      "guid": "changelog.com/2/448",
      "enclosures": [
        {
          "url": "https://cdn.changelog.com/uploads/gotime/60/go-time-60.mp3",
          "length": "77430347",
          "type": "audio/mpeg"
        }
      ],
      "extensions": {
        "dc": {
          "creator": [
            {
              "name": "creator",
              "value": "Erik St. Martin, Carlisia Pinto, and Brian Ketelsen",
              "attrs": {},
              "children": {}
            }
          ]
        },
        "itunes": {
          "author": [
            {
              "name": "author",
              "value": "Erik St. Martin, Carlisia Pinto, and Brian Ketelsen",
              "attrs": {},
              "children": {}
            }
          ],
          "duration": [
            {
              "name": "duration",
              "value": "53:39",
              "attrs": {},
              "children": {}
            }
          ],
          "explicit": [
            {
              "name": "explicit",
              "value": "no",
              "attrs": {},
              "children": {}
            }
          ],
          "image": [
            {
              "name": "image",
              "value": "",
              "attrs": {
                "href": "https://cdn.changelog.com/images/podcasts/gotime-cover-art-2f51f845f9ffc5a946394231c4b41c21.png?vsn=d"
              },
              "children": {}
            }
          ],
          "keywords": [
            {
              "name": "keywords",
              "value": "go, golang, open source, software, development",
              "attrs": {},
              "children": {}
            }
          ],
          "subtitle": [
            {
              "name": "subtitle",
              "value": "Why WADL When You Can Swagger?",
              "attrs": {},
              "children": {}
            }
          ],
          "summary": [
            {
              "name": "summary",
              "value": "Ivan Porto Carrero joined the show to talk about generating documentation (with Swagger), pks, kubo, and other interesting Go projects and news.",
              "attrs": {},
              "children": {}
            }
          ]
        }
      }
    }
  ],
  "feedType": "rss",
  "feedVersion": "2.0"
}`),
	}}
}

func (ts *testParser) ParseURL(feedURL string) (feed *gofeed.Feed, err error) {
	for url, encoded := range ts.data {
		if feedURL == url {
			feed = &gofeed.Feed{}
			err = json.Unmarshal(encoded, feed)
			return
		}
	}
	err = fmt.Errorf("URL not found: %q", feedURL)
	return
}

func TestManagerAdd(t *testing.T) {
	m, err := NewManager(newTestParser())
	if err != nil {
		t.Fatal(err)
	}
	invalidFeedURL := "123"
	if err = m.Add(invalidFeedURL); err == nil {
		t.Fatalf("invalid URL is allowed: %q", invalidFeedURL)
	}
	if err = m.Add(validFeedURL); err != nil {
		t.Fatal(err)
	}
	if err = m.Add(validFeedURL); err == nil {
		t.Fatal("duplicated URL is allow")
	}
}

func TestManagerWithDataFile(t *testing.T) {
	tmp, err := ioutil.TempFile("", "podmandi_test")
	if err != nil {
		t.Fatal(err)
	}
	if err = tmp.Close(); err != nil {
		t.Fatal(err)
	}

	if err = os.Remove(tmp.Name()); err != nil {
		t.Fatal(err)
	}

	m, err := NewManager(newTestParser(), WithDataFile(tmp.Name()))
	if err != nil {
		t.Fatal(err)
	}
	testURL := validFeedURL
	if err = m.Add(testURL); err != nil {
		t.Fatal(err)
	}

	fd, err := os.Open(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	decoder := json.NewDecoder(fd)
	var podcasts []Podcast
	if err = decoder.Decode(&podcasts); err != nil {
		t.Fatal(err)
	}
	if len(podcasts) != 1 {
		t.Fatalf("wrong number of podcasts, expected: %d, got: %d", 1, len(podcasts))
	}
	pod := podcasts[0]
	if pod.URL != testURL {
		t.Fatalf("wrong URL stored in persisten file, expected: %q, got: %q", testURL, pod.URL)
	}
}

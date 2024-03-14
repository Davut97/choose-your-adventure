package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

var defualtHandletTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Choose Your Own Adventure</title>
</head>
<body>
	<section class="page">
		<h1>{{.Title}}</h1>
		{{range .Story}}
			<p>{{.}}</p>
		{{end}}
		<ul>
			{{range .Options}}
				<li><a href="/{{.Arc}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</section>
	<style>
		body {
			font-family: helvetica, arial;
		}
		h1 {
			text-align: center;
			position: relative;
		}
		.page {
			width: 80%;
			max-width: 500px;
			margin: 40px auto;
			padding: 80px;
			background: #FFFCF6;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #777;
		}
		ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		}
		li {
			padding-top: 10px;
		}
		a,
		a:visited {
			text-decoration: none;
			color: #6295b5;
		}
		a:active,
		a:hover {
			color: #7792a2;
		}
		p {
			text-indent: 1em;
		}
	</style>
</body>
</html>`

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Story map[string]StoryArc
type handler struct {
	s Story
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defualtHandletTemplate))
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	if arc, ok := h.s[path]; ok {
		err := tpl.Execute(w, arc)
		if err != nil {
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
}

func JsonStory(r io.Reader) (Story, error) {
	dec := json.NewDecoder(r)
	var story Story
	err := dec.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func main() {
	// jsonData, err := os.ReadFile("gopher.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// story, err := parseStory(jsonData)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Println(story["intro"])
}

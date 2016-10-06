package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cmars/tfclient"
)

var addr = flag.String("addr", "", "remote address")
var samples = flag.String("samples", "", "Directory containing samples")

var addrs []string
var samplesHandler http.Handler
var sampleFiles []string

func main() {
	flag.Parse()
	if *addr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	addrs = strings.Split(*addr, ",")

	if *samples != "" {
		files, err := ioutil.ReadDir(*samples)
		if err != nil {
			log.Printf("failed to read samples: %v", err)
		} else {
			for _, fi := range files {
				if strings.HasSuffix(strings.ToLower(fi.Name()), ".jpg") {
					sampleFiles = append(sampleFiles, fi.Name())
				}
			}
			samplesHandler = http.StripPrefix("/samples/", http.FileServer(http.Dir(*samples)))
		}
	}

	http.HandleFunc("/", page)
	http.ListenAndServe(":8080", nil)
}

func chooseAddr() string {
	return addrs[rand.Intn(len(addrs))]
}

type response struct {
	Predictions   []tfclient.Prediction
	ImageLocation string
	Samples       []string
	Error         error
}

var pageTemplate = template.Must(template.New("page").Parse(`<html>
<head><title>Image Classifier</title>
<style>
h1, p, form, div {
	text-align: center;
}
form, img, div {
	display: block;
	margin: auto;
}
img {
  max-width:300px;
  max-height:400px;
  width: auto;
  height: auto;
}
table {
	margin: auto;
}
form input[type=text] {
	width: 50%;
}
</style>
</head>
<body>
<h1>Image Classifier</h1>
<form action="/" method="GET">
	<input name="image" type="text" />
	<input type="submit" value="What is it?" />
</form>
{{ if .ImageLocation -}}
<p>Samples:{{ range .Samples -}}&nbsp;|&nbsp;<a href="/?sample={{.}}">{{.}}</a>{{ end -}}</p>
<div class="result">
<p>This?</p>
</div>
<img src="{{ .ImageLocation }}" />
{{ else -}}
<div class="result">
<p>Paste a URL to a jpg image, and click the button.</p>
<p>Samples:{{ range .Samples -}}&nbsp;|&nbsp;<a href="/?sample={{.}}">{{.}}</a>{{ end -}}</p>
</div>
{{ end -}}
{{ if .Predictions }}<div class="result">
<p>I think it's a...</p></div>
<table><tr><th>Class</th><th>Score</th></tr>
{{ range .Predictions -}}
<tr><td>{{ .Class }}</td><td>{{ .Score }}</td></tr>
{{ end -}}
</table>{{ end -}}
{{ if .Error -}}
<div class="error">
An error occurred when trying to classify this image:<br/>
<i>{{ . }}</i>
</div>
{{ end -}}
</body>
</html>
`))

func page(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	hasSamples := samplesHandler != nil
	if hasSamples && strings.HasPrefix(r.URL.Path, "/samples") {
		samplesHandler.ServeHTTP(w, r)
		return
	}

	fail := func(err error) {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = pageTemplate.Execute(w, response{Error: err})
		if err != nil {
			log.Println(err)
		}
	}
	err := r.ParseForm()
	if err != nil {
		fail(err)
		return
	}
	imgdata, imageLocation, err := getImage(r, hasSamples)
	if err != nil {
		fail(err)
		return
	}
	if imgdata == nil {
		err := pageTemplate.Execute(w, response{Samples: sampleFiles})
		if err != nil {
			log.Println(err)
		}
		return
	}

	client, err := tfclient.NewClient(chooseAddr())
	if err != nil {
		fail(err)
		return
	}
	predictions, err := client.Predict(imgdata)
	if err != nil {
		fail(err)
		return
	}
	err = pageTemplate.Execute(w, response{
		Predictions:   predictions,
		ImageLocation: imageLocation,
		Samples:       sampleFiles,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func getImage(r *http.Request, hasSamples bool) ([]byte, string, error) {
	fail := func(err error) ([]byte, string, error) {
		return nil, "", err
	}

	sampleLocation := r.Form.Get("sample")
	if sampleLocation != "" && hasSamples {
		imageLocation := "/samples/" + sampleLocation
		log.Println("open", imageLocation, sampleLocation)
		sampleFile, err := os.Open(filepath.Join(*samples, sampleLocation))
		if err != nil {
			return fail(err)
		}
		defer sampleFile.Close()
		imgdata, err := ioutil.ReadAll(sampleFile)
		if err != nil {
			return fail(err)
		}
		return imgdata, imageLocation, nil
	}

	imageLocation := r.Form.Get("image")
	if imageLocation == "" {
		return nil, "", nil
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Get(imageLocation)
	if err != nil {
		return fail(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fail(fmt.Errorf("request failed"))
	}
	imgdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fail(err)
	}
	return imgdata, imageLocation, nil
}

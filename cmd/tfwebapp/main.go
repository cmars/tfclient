package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/cmars/tfclient"
)

var addr = flag.String("addr", "", "remote address")

var addrs []string

func main() {
	flag.Parse()
	if *addr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	addrs = strings.Split(*addr, ",")

	http.HandleFunc("/", page)
	http.ListenAndServe(":8080", nil)
}

func chooseAddr() string {
	return addrs[rand.Intn(len(addrs))]
}

type response struct {
	Predictions   []tfclient.Prediction
	ImageLocation string
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
<div class="result">
<p>This?</p>
</div>
<img src="{{ .ImageLocation }}" />
{{ else -}}
<div class="result">
<p>Paste a URL to a jpg image, and click the button.</p>
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
	imageLocation := r.Form.Get("image")
	if imageLocation == "" {
		err := pageTemplate.Execute(w, response{})
		if err != nil {
			log.Println(err)
		}
		return
	}
	resp, err := http.Get(imageLocation)
	if err != nil {
		fail(err)
		return
	}
	defer resp.Body.Close()
	imgdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fail(err)
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
	})
	if err != nil {
		log.Println(err)
		return
	}
}

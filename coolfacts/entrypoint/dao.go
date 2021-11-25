package main


var factRequest struct {
	Image string `json:"image"`
	Description string `json:"description"`
}

var newsTemplate = `<!DOCTYPE html>
<html>
  <head><style>/* copy coolfacts/styles.css for some color 🎨*/</style></head>
  <body>
  <h1>Facts List</h1>
  <div>
    {{ range . }}
       <article>
            <h3>{{.Description}}</h3>
            <img src="{{.Image}}" width="100%" />
       </article>
    {{ end }}
  <div>
  </body>
</html>`

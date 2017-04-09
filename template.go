package remark

type tmplParamsIndex struct {
	SrcPath string
}

var tmplIndex = `
<!DOCTYPE html>
<html>
  <head>
    <title>Title</title>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
    <style type="text/css">
      @import url(http://fonts.googleapis.com/css?family=Yanone+Kaffeesatz);
      @import url(http://fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic);
      @import url(http://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700,400italic);

      body { font-family: 'Droid Serif'; }
      h1, h2, h3 {
        font-family: 'Yanone Kaffeesatz';
        font-weight: normal;
      }
      .remark-code, .remark-inline-code { font-family: 'Ubuntu Mono'; }
    </style>
    <link rel="stylesheet" href="/static/custom.css" />
  </head>
  <body>
    <script src="http://gnab.github.io/remark/downloads/remark-latest.min.js" type="text/javascript">
    </script>
    <script type="text/javascript">
      function load() {
        remark.create({ sourceUrl: "/index.md" });
      }

      function reset() {
        document.getElementsByTagName("body").forEach(function(elem) {
          elem.classList.remove("remark-container");
        });

        document.getElementsByTagName("html").forEach(function(elem) {
          elem.classList.remove("remark-container");
        });

        [
          "remark-slides-area",
          "remark-notes-area",
          "remark-preview-area",
          "remark-backdrop",
          "remark-pause",
          "remark-help"
        ].forEach(function(className) {
          document.getElementsByClassName(className).forEach(function(elem) {
            elem.remove();
          });
        });
      }

      document.addEventListener("keyup", function (e) {
        if (e.key === "r") {
          reset();
          load();
        }
      });

      load();
    </script>
  </body>
</html>
`

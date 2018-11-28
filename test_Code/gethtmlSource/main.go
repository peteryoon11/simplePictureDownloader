// http_get101.go
//
// get the contents of a web page with given URL
//
// for imported package info see ...
// http://golang.org/pkg/fmt/
// http://golang.org/pkg/io/ioutil/
// http://golang.org/pkg/net/http/
//
// tested with Go version 1.4.2   by vegaseat  28apr2015
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://tour.golang.org/welcome/1"
	//url := "https://kissme2145.tistory.com/1287"
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
}

/* result ...
HTML code of http://tour.golang.org/welcome/1 ...
<!doctype html>
<html lang="en" ng-app="tour">
<head>
    <meta charset="utf-8">
    <title>A Tour of Go</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="mobile-web-app-capable" content="yes">
    <link rel="shortcut icon" sizes="196x196" href="/static/img/favicon.ico">
    <link rel="stylesheet" href="/static/css/app.css" />
    <link rel="stylesheet" href="/static/lib/codemirror/lib/codemirror.css">
    <link href='//fonts.googleapis.com/css?family=Inconsolata' rel='stylesheet' type='text/css'>
<script type="text/javascript">
var _gaq = _gaq || [];
_gaq.push(["_setAccount", "UA-11222381-5"]);
_gaq.push(["b._setAccount", "UA-49880327-6"]);
window.trackPageview = function() {
  _gaq.push(["_trackPageview", location.pathname+location.hash]);
  _gaq.push(["b._trackPageview", location.pathname+location.hash]);
};
window.trackPageview();
window.trackEvent = function(category, action, opt_label, opt_value, opt_noninteraction) {
  _gaq.push(["_trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
  _gaq.push(["b._trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
};
</script>
</head>
<body>
    <div class="bar top-bar">
        <a class="left logo" href="/list">A Tour of Go</a>
        <div table-of-contents-button=".toc"></div>
        <div feedback-button></div>
    </div>
    <div table-of-contents></div>
    <div ng-view ng-cloak class="ng-cloak"></div>
    <script src="/script.js"></script>
    <script>
    window.transport = HTTPTransport();
    window.socketAddr = "";
    function highlight(selector) {
        var speed = 50;
        var obj = $(selector).stop(true, true)
        for (var i = 0; i < 5; i++) {
            obj.addClass("highlight", speed)
            obj.delay(speed)
            obj.removeClass("highlight", speed)
        }
    }
    function highlightAndClick(selector) {
        highlight(selector);
        setTimeout(function() {
            $(selector)[0].click()
        }, 750);
    }
    function click(selector) {
        $(selector)[0].click();
    }
    </script>
<script type="text/javascript">
(function() {
  var ga = document.createElement("script"); ga.type = "text/javascript"; ga.async = true;
  ga.src = ("https:" == document.location.protocol ? "https://ssl" : "http://www") + ".google-analytics.com/ga.js";
  var s = document.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ga, s);
})();
</script>
</body>
</html>
*/

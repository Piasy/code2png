/**
 * Created by piasy on 03/11/2016.
 */

function code2png(lang, code) {
  post("http://code2png.babits.top/v1/code_images/",
    "lang=java&code=" + encodeURIComponent(code.replace(/\n/g, "\\n").replace(/\r/g, "\\r")),
    function(resp) {
      var image = JSON.parse(resp);
      downloadImage(image.url);
    });
}

function post(url, body, callback) {
  var xhr = new XMLHttpRequest();
  xhr.open('POST', url, true);
  xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status == 200) {
      callback(xhr.responseText);
    }
  };
  xhr.send(body);
}

function downloadImage(url) {
  var xhr = new XMLHttpRequest();
  xhr.open("GET", url, true);
  xhr.responseType = 'blob';
  xhr.onload = function(e) {
    download(xhr.response, url.substring(url.lastIndexOf("/") + 1), "image/png");
  };
  xhr.send();
}

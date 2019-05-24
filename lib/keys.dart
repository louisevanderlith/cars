import 'dart:html';

String getObjKey() {
  var path = window.location.pathname;
  return path.substring(path.lastIndexOf('/') + 1).replaceFirst('%60', '`');
}
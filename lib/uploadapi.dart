import 'dart:async';
import 'dart:convert';
import 'dart:html';
import 'keys.dart';
import 'pathlookup.dart';

String _imageURL;

Future<HttpRequest> createUpload(FormData data, Function callback) async {
  var url = await buildPath('Artifact.API', 'upload', new List<String>());
  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.setRequestHeader(
      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd.listen((e) => compltr.complete(request));
  request.send(data);

  return compltr.future;
}

void uploadFile(Event e) {
  if (e.target is FileUploadInputElement) {
    FileUploadInputElement fileElem = e.target;
    var files = fileElem.files;

    var forAttr = fileElem.dataset['for'];
    var nameAttr = fileElem.dataset['name'];
    var ctrlID = fileElem.id;
    var infoObj = {
      "For": forAttr,
      "ItemName": nameAttr,
      "ItemKey": getObjKey()
    };

    if (files.length > 0) {
      File firstFile = files[0];

      doUpload(firstFile, infoObj, ctrlID);
    }
  }
}

void doUpload(File file, Map<String, String> infoObj, String ctrlID) {
  var formData = new FormData();
  formData.appendBlob("file", file);
  formData.append("info", jsonEncode(infoObj));

  createUpload(formData, (obj) {
    var resp = jsonDecode(obj);
    finishUpload(resp, infoObj, ctrlID);
  });
}

void finishUpload(
    dynamic obj, Map<String, String> infoObj, String ctrlID) async {
  if (obj['Error'].length > 0) {
    print(obj['Error']);
    return;
  }

  if (_imageURL?.isEmpty ?? true) {
    print('getting path first time.');
    _imageURL = await buildPath('Artifact.API', "upload", ["file"]);
  }

  var data = obj['Data'];
  var fullURL = "${_imageURL}/${data}";

  var imageHolder = querySelector("#${ctrlID.replaceFirst('Img', 'View')}");
  var uploader = querySelector("#${ctrlID}");

  imageHolder.classes.remove('is-hidden');
  imageHolder.setAttribute('src', fullURL);

  uploader.dataset['id'] = data;
  uploader.attributes.remove('required');
}

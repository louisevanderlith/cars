import 'dart:async';
import 'dart:html';
import 'dart:convert';

Map<String, String> _pastNames = {
  "Router.API": routerDefault()
};

String routerDefault() {
  HiddenInputElement hostPath = querySelector("#HostID");
  return "https://router${hostPath.value}v1/discovery/";
}

Future<String> getRouterPath(String apiName) async {
  HiddenInputElement instanceElem = querySelector("#InstanceID");
  var routerURL = await getServiceURL("Router.API");

  return "${routerURL}${instanceElem.value}/${apiName}/true";
}

Future<String> doLookup(String apiName) async {
  var routerPath = await getRouterPath(apiName);
  var resp = await HttpRequest.request(routerPath, method: "GET", onProgress: lookupProgress);
  final json = jsonDecode(resp.response);

  print(json);
  return json["Data"];
}

void lookupProgress(ProgressEvent info) {
  print(info.eventPhase);
}

Future<String> getServiceURL(String apiName) async {
  var serviceURL = _pastNames[apiName];

  if (serviceURL == null) {
    serviceURL = await doLookup(apiName);
    _pastNames[apiName] = serviceURL;
  }

  return serviceURL;
}

Future<String> buildPath(
    String apiName, String controller, List<String> params) async {
  var url = await getServiceURL(apiName);
  
  var result = url + 'v1/' + controller;

  for (var i = 0; i < params.length; i++) {
    result += "/" + params[i];
  }

  return result;
}

void requestComplete(HttpRequest request, Function callback) {
  switch (request.status) {
    case 200:
      callback(request.responseText);
      return;
    default:
      print(request.status);
  }
}
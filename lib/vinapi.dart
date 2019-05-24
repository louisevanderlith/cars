import 'dart:async';
import 'dart:html';

import 'package:Cars.APP/pathlookup.dart';

Future<HttpRequest> lookupVIN(String vin) async {
  var url = await buildPath("VIN.API", "lookup", [vin]);
  return HttpRequest.request(url,method: "GET", onProgress: onProgress);
  /*final request = HttpRequest();
  request.open("GET", url);
  request.setRequestHeader(      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd.listen((e) => requestComplete(request, callback));
  request.onProgress.listen(onProgress);
  request.send();*/
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}
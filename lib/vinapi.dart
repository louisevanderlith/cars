import 'dart:async';
import 'dart:html';

import 'package:Cars.APP/pathlookup.dart';

Future<HttpRequest> validateVIN(String vin) async {
  final url = await buildPath("VIN.API", "validate", [vin]);
  
  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("GET", url);
  request.setRequestHeader("Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd.listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send();
  
  return compltr.future;
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}


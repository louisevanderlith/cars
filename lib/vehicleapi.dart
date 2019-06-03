import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:Cars.APP/pathlookup.dart';

Future<HttpRequest> submitVehicle(String vinKey, String vin) async {
  final url = await buildPath("Vehicle.API", "vehicle", []);

  var data = jsonEncode({
    'VINKey': vinKey,
    'FullVIN': vin,
    'Series': {
      'Manufacturer': '',
    },
    'Colour': '',
    'PaintNo': '',
    'Month': '',
    'Year': '',
    'Engine': '',
    'Gearbox': '',
    'BodyStyle': '',
    'Doors': '',
    'Trim': '',
    'Extra': [],
  });

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.setRequestHeader(
      "Authorization", "Bearer " + window.localStorage['avosession']);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}

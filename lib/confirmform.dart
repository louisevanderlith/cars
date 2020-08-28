import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_vehicle/bodies/engine.dart';
import 'package:mango_vehicle/bodies/gearbox.dart';
import 'package:mango_vehicle/bodies/series.dart';
import 'package:mango_vehicle/bodies/vehicle.dart';
import 'package:mango_vehicle/vehicleapi.dart';

class ConfirmForm extends FormState {
  HiddenInputElement _vinKey;
  HiddenInputElement _vin;
  CheckboxInputElement _accept;

  ParagraphElement _error;

  ConfirmForm(String idElem, String vinKeyElem, String vinElem,
      String acceptElem, String submitBtn)
      : super(idElem, submitBtn) {
    _vinKey = querySelector(vinKeyElem);
    _vin = querySelector(vinElem);
    _accept = querySelector(acceptElem);

    _error = querySelector("${idElem}Err");

    querySelector(submitBtn).onClick.listen(onSend);
  }

  Key get vinKey {
    return new Key(_vinKey.value);
  }

  String get vin {
    return _vin.value;
  }

  bool get accepted {
    return _accept.checked;
  }

  Series get series {
    return new Series(0, 0, "None", "Unknown", "None", "Somewhere");
  }

  String get makeCountry {
    return "Unknown";
  }

  String get drive {
    return "RHD";
  }

  String get transmission {
    return "Unknown";
  }

  String get body {
    return "Unknown";
  }

  String get enginePos {
    return "Unknown";
  }

  num get mileage {
    return 0;
  }

  num get price {
    return 0;
  }

  String get condition {
    return "Unknown";
  }

  String get issues {
    return "";
  }

  bool get spare {
    return false;
  }

  bool get service {
    return false;
  }

  String get bodytype {
    return "Unknown";
  }

  num get doors {
    return 0;
  }

  String get colour {
    return "Unknown";
  }

  String get paintNo {
    return "unknown";
  }

  Engine get engine {
    return new Engine("XXX", "None", 0);
  }

  Gearbox get gearbox {
    return new Gearbox("XXX", "None", 0, "Manual");
  }

  List<String> get extra {
    return new List<String>();
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var vh = new Vehicle(
          vinKey,
          vin,
          series,
          colour,
          paintNo,
          engine,
          gearbox,
          bodytype,
          doors,
          extra,
          spare,
          service,
          condition,
          issues,
          mileage);
      var result = await submitVehicle(vh);
      var obj = jsonDecode(result.response);

      if (result.status == 200) {
        var data = obj['Data'];
        print(data);
        if (data) {
          window.location.replace('/create/step3/${vin}');
        }
      } else {
        _error.text = obj['Error'];
      }
    }
  }
}

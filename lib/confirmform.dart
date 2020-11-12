import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
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
  SelectElement cboBodyStyle;
  CheckboxInputElement _accept;

  ConfirmForm(String idElem, String vinKeyElem, String vinElem,
      String styleElem, String acceptElem, String submitBtn)
      : super(idElem, submitBtn) {
    _vinKey = querySelector(vinKeyElem);
    _vin = querySelector(vinElem);
    cboBodyStyle = querySelector(styleElem);
    _accept = querySelector(acceptElem);

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

  num get bodystyle {
    return cboBodyStyle.selectedIndex;
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
          bodystyle,
          doors,
          extra,
          spare,
          service,
          condition,
          issues,
          mileage);

      var result = await submitVehicle(vh);

      if (result.status == 200) {
        var obj = jsonDecode(result.response);
        var data = obj['Data'];
        if (data) {
          window.location.replace('/create/step3/${vin}');
        }
      } else {
        Toast.error(
            title: "Failed!",
            message: result.response,
            position: ToastPos.bottomLeft);
      }
    }
  }
}

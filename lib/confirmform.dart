import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/services/vehicleapi.dart';

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

  String get vinKey {
    return _vinKey.value;
  }

  String get vin {
    return _vin.value;
  }

  bool get accepted {
    return _accept.checked;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var result = await submitVehicle(vinKey, vin);
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

/*
VINKey    husk.Key
	FullVIN   string
	Series    SeriesInfo
	Colour    string
	PaintNo   string
	Month     int
	Year      int
	Engine    Engine
	Gearbox   Gearbox
	BodyStyle string
	Doors     int
	Trim      string
	Extra     []string
 */

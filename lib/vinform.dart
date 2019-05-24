import 'dart:convert';
import 'dart:html';

import 'package:Cars.APP/formstate.dart';
import 'package:Cars.APP/pathlookup.dart';
import 'package:Cars.APP/vinapi.dart';

class VINForm extends FormState {
  TextInputElement _vin;

  VINForm(String idElem, String vinElem, String submitBtn)
      : super(idElem, submitBtn) {
    _vin = querySelector(vinElem);

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get vin {
    return _vin.value;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);
      print('Validating ${vin}');
      var result = await lookupVIN(vin);
      onComplete(result);
    }
  }

  void onComplete(HttpRequest req) {
    var obj = jsonDecode(req.response);
    print(obj);

    var data = obj['Data'];
    print(data);
    window.localStorage['Step1'] = htmlEscape.convert(data.toString());
    window.location.replace('/create/step2');
  }
}

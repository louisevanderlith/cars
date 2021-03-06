import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/formstate.dart';
import 'package:mango_ui/services/vinapi.dart';

class VINForm extends FormState {
  TextInputElement _vin;
  ParagraphElement _error;

  VINForm(String idElem, String vinElem, String submitBtn)
      : super(idElem, submitBtn) {
    _vin = querySelector(vinElem);
    _error = querySelector("${idElem}Err");

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get vin {
    return _vin.value;
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var result = await validateVIN(vin);
      var obj = jsonDecode(result.response);

      if (result.status == 200) {
        var data = obj['Data'];
        print(data);
        if (data) {
          window.location.replace('/create/step2/${vin}');
        }
      } else {
        _error.text = obj['Error'];
      }
    }
  }
}

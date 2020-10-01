import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_ui/formstate.dart';
import 'package:mango_vin/vinapi.dart';

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

      var result = await validateVIN(vin);
      var data = jsonDecode(result.response);

      if (result.status == 200) {
        print(data);
        if (data) {
          window.location.replace('/create/step2/${vin}');
        }
      } else {
        new Toast.error(
            title: "Error!",
            message: "Validation Failed\n${data}",
            position: ToastPos.bottomLeft);
      }
    }
  }
}

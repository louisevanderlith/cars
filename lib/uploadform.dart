import 'dart:convert';
import 'dart:html';

import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_ui/formstate.dart';

class UploadForm extends FormState {
  FileUploadInputElement _front;
  FileUploadInputElement _left;
  FileUploadInputElement _right;
  FileUploadInputElement _back;
  FileUploadInputElement _intera;
  FileUploadInputElement _interb;
  FileUploadInputElement _interc;

  UploadForm(
      String idElem,
      String frontElem,
      String leftElem,
      String rightElem,
      String backElem,
      String interAElem,
      String interBElem,
      String interCElem,
      String submitBtn)
      : super(idElem, submitBtn) {
    _front = querySelector(frontElem);
    _left = querySelector(leftElem);
    _right = querySelector(rightElem);
    _back = querySelector(backElem);
    _intera = querySelector(interAElem);
    _interb = querySelector(interBElem);
    _interc = querySelector(interCElem);

    querySelector(submitBtn).onClick.listen(onSend);

    _front.onChange.listen(uploadFile);
    _left.onChange.listen(uploadFile);
    _right.onChange.listen(uploadFile);
    _back.onChange.listen(uploadFile);
    _intera.onChange.listen(uploadFile);
    _interb.onChange.listen(uploadFile);
    _interc.onChange.listen(uploadFile);
  }

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);
      //print('Validating ${vin}');
      //var result = await lookupVIN(vin);
      //onComplete(result);
    }
  }

  void onComplete(HttpRequest req) {
    var obj = jsonDecode(req.response);

    var data = obj['Data'];
    window.localStorage['Step1'] =
        data.toString(); //htmlEscape.convert(data.toString());
    window.location.replace('/create/step2');
  }
}

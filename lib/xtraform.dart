import 'dart:html';

import 'package:mango_ui/formstate.dart';

class XtraForm extends FormState {
  ElementList<CheckboxInputElement> _extras;

  XtraForm(String idElem, String chkExtra, String submitBtn)
      : super(idElem, submitBtn) {
    _extras = querySelectorAll(chkExtra);

    querySelector(submitBtn).onClick.listen(onSend);
  }

  List<String> get extras {
    var result = new List<String>();

    for (var xtra in _extras) {
      result.add(xtra.value);
    }

    return result;
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      window.localStorage['Step3'] = extras.toString();
      window.location.replace('/create/step4');
    }
  }
}

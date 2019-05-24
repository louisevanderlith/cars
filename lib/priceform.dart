import 'dart:html';
import 'dart:svg';

import 'package:Cars.APP/formstate.dart';

class PriceForm extends FormState {
  NumberInputElement _price;
  NumberInputElement _mileage;
  CheckboxInputElement _natis;
  CheckboxInputElement _finance;

  PriceForm(String idElem, String priceElem, String mileageElem,
      String natisElem, String financeElem, String submitBtn)
      : super(idElem, submitBtn) {
    _price = querySelector(priceElem);
    _mileage = querySelector(mileageElem);
    _natis = querySelector(natisElem);
    _finance = querySelector(financeElem);

    querySelector(submitBtn).onClick.listen(onSend);
  }

  num get price {
    return _price.valueAsNumber;
  }

  num get mileage {
    return _mileage.valueAsNumber;
  }

  bool get hasNatis {
    return _natis.checked;
  }

  bool get canFinance {
    return _finance.checked;
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      print('do something, like verify the Price, in the Ad creation');
    }
  }
}

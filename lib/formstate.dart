import 'dart:html';

class FormState {
  FormElement _form;
  ButtonElement _sendBtn;

  FormState(String formID, String submitID) {
    _form = querySelector(formID);
    _sendBtn = querySelector(submitID);

    disableSubmit(true);

    _form.onInput.listen(validateElement);
  }

  bool isFormValid() {
    return _form.checkValidity();
  }

  void disableSubmit(bool disable) {
    _sendBtn.disabled = disable;
  }

  void validateElement(Event e) {
    var isInput = e.target is InputElement;
    
    if (!isInput) {
      return;
    }

    var elem = e.target as InputElement;
    if (elem != null) {
      var elemValid = elem.checkValidity();

      if (!elemValid) {
        elem.setAttribute("invalid", "");
      } else {
        elem.removeAttribute("invalid");
      }

      elem.nextElementSibling.text = elem.validationMessage;

      disableSubmit(!isFormValid());
    }
  }

  void pressEnter(Event e) {
    if (e.runtimeType.toString() != "KeyboardEvent") {
      return;
    }

    var keyEvent = e as KeyboardEvent;

    if (keyEvent.key != 'Enter') {
      return;
    }

    _sendBtn.click();
    e.preventDefault();
  }
}

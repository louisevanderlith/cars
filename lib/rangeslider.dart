import 'dart:convert';
import 'dart:html';

class RangeSlider {
  RangeInputElement _min;
  RangeInputElement _max;
  SpanElement _minValue;
  SpanElement _maxValue;

  RangeSlider(String minElem, String maxElem) {
    _min = querySelector(minElem);
    _minValue = querySelector("${minElem}Val");
    _max = querySelector(maxElem);
    _maxValue = querySelector("${maxElem}Val");

    _min.onInput.listen(onMinChange);
    _max.onInput.listen(onMaxChange);
  }

  num get minimum {
    return num.parse(_min.value);
  }

  num get maximum {
    return num.parse(_max.value);
  }

  void onMinChange(Event e) {
    final almost = maximum - 500;
    if (minimum > almost) {
      _min.value = almost.toString();
    }

    _minValue.text = minimum.toString();
  }

  void onMaxChange(Event e) {
    if (maximum - 500 < minimum) {
      var almost = minimum + 500;
      _max.value = almost.toString();
    }

    _maxValue.text = maximum.toString();
  }

  String toJson() {
    return jsonEncode({
      "PriceMin": minimum,
      "PriceMax": maximum,
    });
  }
}

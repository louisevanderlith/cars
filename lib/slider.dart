import 'dart:html';

class Slider {
  ButtonElement _nxt;
  ButtonElement _bck;
  ImageElement _active;

  List<String> _images;
  num _activeIdx;

  Slider(String nextElem, String backElem, String activeElem) {
    _nxt = querySelector(nextElem);
    _bck = querySelector(backElem);
    _active = querySelector(activeElem);
    _activeIdx = 0;

    _nxt.onClick.listen(moveNext);
    _bck.onClick.listen(moveBack);
  }

  void moveNext(MouseEvent e) {
    _activeIdx++;
    _active.src = images[_activeIdx];
  }

  void moveBack(MouseEvent e) {
    _activeIdx--;
    _active.src = images[_activeIdx];

    if (_activeIdx == 0) {
      _active.setAttribute("disabled", "");
      _active.classes.add("is-disabled");
    } else {
      _active.removeAttribute("disabled");
      _active.classes.remove("is-disabled");
    }
  }

  List<String> get images {
    return _images;
  }

  void set images(List<String> imgs) {
    _images = imgs;
  }
}

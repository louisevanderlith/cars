class Slider {
    DivElement _nxt;
    DevElement _bck;
    ImageElement _active;
    //List<
    List<String> _images;
    number _activeIdx;

    Slider(String nextElem, String backElem, String activeElem){
        _nxt = querySelector(nextElem);
        _bck = querySelector(backElem);
        _active = querySelector(activeElem);
        _activeIdx = 0;
    }

    void moveNext(){
        _activeIdx++;
        _active.src = images[_activeIdx];
    }

    void moveBack(){
        _activeIdx--;
        _active.src = images[_activeIdx];
    }


    List<String> get images {
        return _images;
    }

    void set images(List<String> imgs){
        _images = imgs;
    }
}
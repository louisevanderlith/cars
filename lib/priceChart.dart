import 'dart:html';

import 'package:modern_charts/modern_charts.dart';

class PriceChart {
  DivElement _price;

  PriceChart(String priceElm, DataTable table) {
    _price = querySelector(priceElm);

    var chart = GaugeChart(_price);
    chart.draw(table, {
      'animation': {
        'easing': (double t) {
          t = 4 * t - 2;
          return (t * t * t - t) / 12 + .5;
        },
        'onEnd': () {}
      },
      'gaugeLabels': {'enabled': false},
      'title': {'text': 'Vehicles for Current Search'},
    });
  }
}

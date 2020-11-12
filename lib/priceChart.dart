import 'dart:html';

import 'package:modern_charts/modern_charts.dart';

class PriceChart {
  DivElement _price;

  PriceChart(String priceElm, DataTable table) {
    _price = querySelector(priceElm);

    var chart = PieChart(_price);
    chart.draw(table, {
      'pieHole': .5,
      'series': {
        'counterclockwise': true,
        'labels': {'enabled': true},
        'startAngle': 90 + 10 * 360,
      },
      'title': {'text': 'Pie Chart Demo'},
    });
  }
}

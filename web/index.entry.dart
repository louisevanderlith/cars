import 'package:Cars.APP/priceChart.dart';
import 'package:Cars.APP/rangeslider.dart';
import 'package:modern_charts/modern_charts.dart';

void main() {
  new RangeSlider("#minPrice", "#maxPrice");

  var table = DataTable([
    ['Cars', 'Total'],
    ['Toyota', 25],
    ['Honda', 75],
    ['BMW', 40]
  ]);
  new PriceChart("#grphPrice", table);
}

import 'package:Cars.APP/priceChart.dart';
import 'package:Cars.APP/rangeslider.dart';
import 'package:modern_charts/modern_charts.dart';

void main() {
  new RangeSlider("#minPrice", "#maxPrice");

   var table = DataTable([
    ['Make', 'Percentage'],
    ['Toyota', 40],
    ['Ford', 20],
    ['Honda', 10],
    ['Isuzu', 10],
    ['Subaru', 10],
    ['Fiat', 10]
  ]);
  new PriceChart("#grphPrice", table);
}

import 'dart:html';
import 'package:charts_common/src/chart/line/line_renderer.dart';
import 'package:charts_common/src/chart/common/processed_series.dart'
    show MutableSeries, ImmutableSeries;
import 'package:charts_common/src/common/color.dart';
import 'package:charts_common/src/common/material_palette.dart'
    show MaterialPalette;

class PriceChart {
  CanvasElement _price;

  PriceChart(String priceElm) {
    _price = querySelector(priceElm);
  }
}

//https://github.com/google/charts/blob/master/charts_common/test/chart/line/line_renderer_test.dart
class MyRow {
  final String campaignString;
  final int campaign;
  final int clickCount;
  final Color color;
  final List<int> dashPattern;
  final double strokeWidthPx;
  MyRow(this.campaignString, this.campaign, this.clickCount, this.color,
      this.dashPattern, this.strokeWidthPx);
}

void renderChart() {
  LineRenderer renderer;
//_price.value = "jou chsrt";
  List<MutableSeries<int>> numericSeriesList;
  List<MutableSeries<String>> ordinalSeriesList;

  List<MyRow> myFakeDesktopData;
  List<MyRow> myFakeTabletData;
  List<MyRow> myFakeMobileData;

  setUp(() {
    myFakeDesktopData = [
      new MyRow(
          'MyCampaign1', 1, 5, MaterialPalette.blue.shadeDefault, null, 2.0),
      new MyRow(
          'MyCampaign2', 2, 25, MaterialPalette.green.shadeDefault, null, 2.0),
      new MyRow(
          'MyCampaign3', 3, 100, MaterialPalette.red.shadeDefault, null, 2.0),
      new MyRow('MyOtherCampaign', 4, 75, MaterialPalette.red.shadeDefault,
          null, 2.0),
    ];

    myFakeTabletData = [
      new MyRow(
          'MyCampaign1', 1, 5, MaterialPalette.blue.shadeDefault, [2, 2], 2.0),
      new MyRow(
          'MyCampaign2', 2, 25, MaterialPalette.blue.shadeDefault, [3, 3], 2.0),
      new MyRow('MyCampaign3', 3, 100, MaterialPalette.blue.shadeDefault,
          [4, 4], 2.0),
      new MyRow('MyOtherCampaign', 4, 75, MaterialPalette.blue.shadeDefault,
          [4, 4], 2.0),
    ];

    myFakeMobileData = [
      new MyRow(
          'MyCampaign1', 1, 5, MaterialPalette.blue.shadeDefault, null, 2.0),
      new MyRow(
          'MyCampaign2', 2, 25, MaterialPalette.blue.shadeDefault, null, 3.0),
      new MyRow(
          'MyCampaign3', 3, 100, MaterialPalette.blue.shadeDefault, null, 4.0),
      new MyRow('MyOtherCampaign', 4, 75, MaterialPalette.blue.shadeDefault,
          null, 4.0),
    ];

    numericSeriesList = [
      new MutableSeries<int>(new Series<MyRow, int>(
          id: 'Desktop',
          colorFn: (_, __) => MaterialPalette.blue.shadeDefault,
          domainFn: (dynamic row, _) => row.campaign,
          measureFn: (dynamic row, _) => row.clickCount,
          measureOffsetFn: (_, __) => 0,
          data: myFakeDesktopData)),
      new MutableSeries<int>(new Series<MyRow, int>(
          id: 'Tablet',
          colorFn: (_, __) => MaterialPalette.red.shadeDefault,
          domainFn: (dynamic row, _) => row.campaign,
          measureFn: (dynamic row, _) => row.clickCount,
          measureOffsetFn: (_, __) => 0,
          strokeWidthPxFn: (_, __) => 1.25,
          data: myFakeTabletData)),
      new MutableSeries<int>(new Series<MyRow, int>(
          id: 'Mobile',
          colorFn: (_, __) => MaterialPalette.green.shadeDefault,
          domainFn: (dynamic row, _) => row.campaign,
          measureFn: (dynamic row, _) => row.clickCount,
          measureOffsetFn: (_, __) => 0,
          strokeWidthPxFn: (_, __) => 3.0,
          data: myFakeMobileData))
    ];

    ordinalSeriesList = [
      new MutableSeries<String>(new Series<MyRow, String>(
          id: 'Desktop',
          colorFn: (_, __) => MaterialPalette.blue.shadeDefault,
          domainFn: (dynamic row, _) => row.campaignString,
          measureFn: (dynamic row, _) => row.clickCount,
          measureOffsetFn: (_, __) => 0,
          data: myFakeDesktopData)),
      new MutableSeries<String>(new Series<MyRow, String>(
          id: 'Tablet',
          colorFn: (_, __) => MaterialPalette.red.shadeDefault,
          domainFn: (dynamic row, _) => row.campaignString,
          measureFn: (dynamic row, _) => row.clickCount,
          measureOffsetFn: (_, __) => 0,
          strokeWidthPxFn: (_, __) => 1.25,
          data: myFakeTabletData)),
      new MutableSeries<String>(new Series<MyRow, String>(
          id: 'Mobile',
          colorFn: (_, __) => MaterialPalette.green.shadeDefault,
          domainFn: (dynamic row, _) => row.campaignString,
          measureFn: (dynamic row, _) => row.clickCount,
          measureOffsetFn: (_, __) => 0,
          strokeWidthPxFn: (_, __) => 3.0,
          data: myFakeMobileData))
    ];
  });
}

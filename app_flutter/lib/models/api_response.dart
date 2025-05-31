

class ApiResponse<T> {
  final int? code;
  final String? statusMessage;
  final T? data;

  ApiResponse({this.code, this.statusMessage, this.data});

}
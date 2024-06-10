void main() {
  try {
    int result = 10 ~/ 0;
    print(result);
  } catch (e) {
    print('Error: $e');
  } finally {
    print('This will always execute.');
  }
}

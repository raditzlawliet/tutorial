Future<void> fetchUserOrder() {
  return Future.delayed(Duration(seconds: 2), () => print('User order'));
}

void main() async {
  print('Fetching user order...');
  await fetchUserOrder();
  print('Order fetched.');
}

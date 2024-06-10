void say(String message) {
  print(message);
}

int add(int x, int y) {
  return x + y;
}

String greet(String name) {
  return 'Hello, $name!';
}

void main() {
  say("you know who!"); // you know who!
  print(add(3, 4)); // 7
  print(greet('Dart')); // Hello, Dart!
}

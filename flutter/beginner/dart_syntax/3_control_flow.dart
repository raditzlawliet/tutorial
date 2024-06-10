void main() {
  // If-Else Statement
  int number = 10;
  if (number > 0) {
    print('$number is positive.');
  } else if (number < 0) {
    print('$number is negative.');
  } else {
    print('$number is zero.');
  }

  // Switch Statement
  String grade = 'A';
  switch (grade) {
    case 'A':
      print('Excellent!');
      break;
    case 'B':
      print('Good!');
      break;
    case 'C':
      print('Fair.');
      break;
    case 'D':
      print('Poor.');
      break;
    default:
      print('Invalid grade.');
  }

  // For Loop
  for (int i = 0; i < 5; i++) {
    print('i = $i');
  }

  // While Loop
  int j = 0;
  while (j < 5) {
    print('j = $j');
    j++;
  }

  // Do-While Loop
  int k = 0;
  do {
    print('k = $k');
    k++;
  } while (k < 5);
}

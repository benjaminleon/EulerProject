#include<iostream>

double recursive(int, int);
int maxmoves = 20;
  
int main() {
  std::cout << recursive(0,0) << std::endl;
}

double recursive(int x, int y) {
  if ((x == maxmoves) && (y == maxmoves)){
      return 1;
    }
  else if (x == y) {
    return 2*recursive(x+1,y);
  }
  else if ((x < maxmoves) && (y < maxmoves)) {
    return recursive(x+1,y) + recursive(x,y+1);
  }
  else {
    return 1;
  }
}

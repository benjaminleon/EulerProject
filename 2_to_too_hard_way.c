#include<stdio.h>
#include "math.h"

int mult(int* firstptr, int* secondptr, int* totalProdptr, int totalProdlen);

int main(){
  int numberContainer1 [17] = {1,1,2,5,8,9,9,9,0,6,8,4,2,6,2,4,-1};
  int numberContainer2 [17] = {1,1,2,5,8,9,9,9,0,6,8,4,2,6,2,4,-1};
  int totalProdLen = 30;
  int numberContainer3 [totalProdLen] = {0};
  
  int ans = 0;
  ans = mult(&numberContainer1[0], &numberContainer2[0], &numberContainer3[0], totalProdLen);
  printf("Done, the sum of the digits is %d\n", ans);

  long long unsigned num = 2;
}

int mult(int* firstptr, int* secondptr, int* totalProdptr, int totalProdLen){
  int digsum = 0;
  int decimal = 0;
  int current = 0;
  int firstLen = 0;
  int secondLen = 0;
  int i, j, num1, num2, prod;
  
  // Find the number of digits in the numbers
  for (i = 0; *(firstptr + i) != -1; i++) {
    firstLen = i;
  }
  for (j = 0; *(secondptr + j) != -1; j++) {
    secondLen = j;
  }

  // Here is where the work is done
  for (i = 0; i <= firstLen; i++) {
    decimal = 0;
    num1 = *(firstptr + firstLen - i);
    
    for (j = 0; j <= secondLen; j++) {
      num2 = *(secondptr + secondLen - j);
      prod = num1*num2 + decimal;
      
      if (prod + *(totalProdptr + totalProdLen - i - j) >= 10) {
	decimal = (prod + *(totalProdptr + totalProdLen - i - j)) / 10;
	current = prod - 10*decimal;
      }

      else {
	decimal = 0;
	current = prod;
      }

      totalProdptr[totalProdLen - i - j] += current;
      
    }
    totalProdptr[totalProdLen  - i - j] += decimal;
  }

  digsum = 0;

  printf("%d", *(totalProdptr));
  digsum += *(totalProdptr);
  for (int x = 1; x <= totalProdLen; x++) {
    digsum += *(totalProdptr + x);
    printf(",%d", *(totalProdptr + x));
  }
  printf("\n");
  
  return digsum;
}


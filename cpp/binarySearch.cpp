
/*

Binary Search algorithm (C++)
By CatProgrammerYT

*/

#include<iostream>
using namespace std;
int binarySearch(int arr[], int p, int r, int num) {
   int mid = (p + r)/2;
   if (p <= r) {
      if (arr[mid] == num)
      return mid ;
      else if (arr[mid] > num)
      return binarySearch(arr, p, mid-1, num);
      else if (arr[mid] > num)
      return binarySearch(arr, mid+1, r, num);
   }
   return -1;
}
int main(void) {
   int arr[] = {1, 3, 7, 15, 18, 20, 25, 33, 36, 40};
   int n = sizeof(arr)/ sizeof(arr[0]);
   int num = 33;
   int index = binarySearch (arr, 0, n-1, num);
   if(index == -1)
   cout<< num <<" is not present in the array";
   else if (index != -1)
   cout<< num <<" is present at index "<< index <<" in the array";
   while (true) {
	   int x;
	   cin >> x;
   } 
   return 0;
}
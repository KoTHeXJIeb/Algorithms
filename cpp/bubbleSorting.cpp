
/*

Binary Search algorithm (C++)
By CatProgrammerYT

*/
#include<iostream>

using namespace std;

int numarray[15] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15};
int n;

int main() {

	cout << "Enter the number, you are searching for: " << endl;
	cin >> n;
	
	binarySearch(&numarray[15], n);
	return 0;
}

void binarySearch(int* array, int number) {
	// All stuff going here
	
	int len = sizeOf(array);
	cout << len;


}


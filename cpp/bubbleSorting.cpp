
/*

Bubble Sorting algorithm (C++)
By CatProgrammerYT

Early Access! This is not final version of the file!

*/

#include<iostream>
#include <cstdint>
#include <stdint.h>

using namespace std;

int* binarySearch(int* array) {
	// All stuff going here
	for(int i = 0; i < sizeof(array); i++) {
		for(int j = 0; j < sizeof(array)-1; i++) {
			if (array[j+1] > array[j]) {
				int temp;
				temp = array[j];
				array[j] = array[j+1];
				array[j+1] = temp;
				return array;
			}
		}
	}
}

int main() {

	int numarray[15] = {12,3125,12,52,16,12,67,327,17,13};

	int* sortedArray[15] = {binarySearch(&numarray[15])};
	for(int j = 0; j < sizeof(sortedArray); i++) {
		sortedarray = sortedarray[j];
	int sortedarray[15] = { &sortedArray[15] };
	for(int i = 0; i < sizeof(sortedarray); i++) {
		cout << i;
	}
	return 1;
}
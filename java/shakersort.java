
/*

Shaker Sorting Algorithm (Java)
By CatProgrammerYT

*/

public class shakersort {

        public static void shakerSorting(int[] array) {
                for (int i = 0; i < array.length/2; i++) {
                    boolean isSwapped = false;
                    for (int j = i; j < array.length - i - 1; j++) {
                        if (array[j] < array[j+1]) {
                            int temp = array[j];
                            array[j] = array[j+1];
                            array[j+1] = temp;
                            isSwapped = true;
                        }
                    }
                    for (int j = array.length - 2 - i; j > i; j--) {
                        if (array[j] > array[j-1]) {
                            int temp = array[j];
                            array[j] = array[j-1];
                            array[j-1] = temp;
                            isSwapped = true;
                        }
                    }
                    if(!isSwapped) break;
                }
            }
}
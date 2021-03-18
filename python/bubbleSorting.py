
"""

Bubble sorting algorithm
By CatProgrammerYT

"""

# Well, I can make a generator of random array, but I`m too lazy to actually
# do that, so... YES
array = [1,23,45,152,612,6,7,132,-213,-12050125,-1]

"""print(array) - modern ways to debug, ya know"""

def bubbleSort(array):
        for i in range(len(array)): # Looping array (will be used as next number then)
                for j in range(len(array)-1): # Looping array (will be used as previous number then)
                        if array[j+1] > array[j]: # If next number is greater than previous
                                array[j], array[j+1] = array[j+1], array[j] # Then change theirs places in array
        return array

sortedArray = bubbleSort(array)

# Debugging tools, ya know
print(sortedArray)

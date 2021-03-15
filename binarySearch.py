
"""

Binary search algorithm (Python)
By CatProgrammerYT

"""

# Well, I can make a random generator of an array, but I`m too lazy
array = [
1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23,24,25,26,27,28,29,30,
31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50
]
# Getting the number, which user are looking for
n = int(input('Search for the number (1-50): '))

def binarySearch(array, n):
    array.sort()                                            
    l = 0
    r = len(array)  - 1 
    central = 0
    
    while r >= l:
        central = (l + r) // 2
        if array[central] < n:
            l = central + 1
        elif array[central] > n:
            r = central - 1
        else:
            return central
    return -1 


res = binarySearch(array=array, n=n)

if res != -1:
    print('Element ' + str(n) + ' is in the list with index ' + str(res))
else:
    print('There is no element  ' + str(n) + ' in this list!')
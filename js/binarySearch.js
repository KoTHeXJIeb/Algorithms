
/*

Binary Search algorithm (JavaScript)
By CatProgrammerYT

*/


// Code below is not supossed to be editable, but if you need - you can change it as you want
let binarySearch = function (arr, number) { 

        // Start is the start position of the array (who could guess?)
	let start=0;
        // End is, well, end position
        let end=arr.length-1; 
		
	while (start<=end){ 
		let mid=Math.floor((start + end)/2); 
		if (arr[mid]===number) return true; 
		else if (arr[mid] < number) 
			start = mid + 1; 
		else
			end = mid - 1; 
	} 

	return false; 
} 

// This code is editable (just for checking, if all working purfect :3)
let arr = [1, 3, 5, 7, 8, 9]; 
let number = 5;

if (binarySearch(arr, number)) document.write("Element found!<br>"); 
else { document.write("Element not found!<br>") } 
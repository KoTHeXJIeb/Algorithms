
/*

Binary Search algorithm (JavaScript)
By CatProgrammerYT

*/

// Code below is not supossed to be editable, but if you need - you can change it as you want
let binarySearch = function (arr, number, position) { 

        // Start is the start position of the array (who could guess?)
	let start=0;
        // The end position (ye, cool)
        let end=arr.length-1; 

	while (start<=end){ 
		let mid=Math.floor((start + end)/2); 
		if (arr[mid]===number) return true; 
		else if (arr[mid] < number) start = mid + 1; 
		else end = mid - 1; 
	} 

	return false;

} 

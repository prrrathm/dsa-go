function BinarySearch(arr: number[], x: number, left: number, right: number) {
	// // let left = 0;
	// let right = arr.length;
	if (arr.length == 0) {
		console.log("Array is Empty");
		return -1;
	}
	if (left > right) return -1;
	let mid = Math.floor((right + left) / 2);
	if (arr[mid] === x) return mid;
	if (x > arr[mid]) {
		return BinarySearch(arr, x, mid + 1, right);
	}
	if (x < arr[mid]) {
		return BinarySearch(arr, x, left, mid - 1);
	}
}

function BinarySearchMain(arr: number[], x: number) {
	return BinarySearch(arr, x, 0, arr.length);
}
console.log(BinarySearchMain([1, 2, 3, 4, 5, 6, 7, 8, 10], 4));

function getTHArrayFromElement(coll) {
	let arr = [].slice.call(coll.children);
	let thArr = arr.filter((el) => el.tagName == "TH");
	return thArr
}

function getTRArrayFromTable() {
	let table = document.getElementById("table");

	let arr = [].slice.call(table.children);
	let trArr = arr.filter((el) => el.tagName == "TR");
	return trArr
}

function toggleUI() {
	let navigation = document.getElementById("navigation");
	let insertForm = document.getElementById("insertForm");
	let downloadPDF = document.getElementById("downloadPDF");
	let week = document.getElementById("week");
	let tableHead = document.getElementById("tableHead");
	let headers = getTHArrayFromElement(tableHead);
	let trArr = getTRArrayFromTable();

	let current = navigation.style.display;
	navigation.style.display = (current == "") ? "none" : "";
	insertForm.style.display = (current == "") ? "none" : "";
	downloadPDF.style.display = (current == "") ? "" : "none";
	week.style.display = (current == "") ? "" : "none";
	headers[7].style.display = (current == "") ? "none" : "";
	for (tr of trArr) {
		let trTHArr = getTHArrayFromElement(tr)
		trTHArr[7].style.display = (current == "") ? "none" : "";
	}
}

function printTable() {
	let filterClassContainer = document.getElementById("filterClassContainer");

	filterClassContainer.style.display = "none";
	window.print()
	filterClassContainer.style.display = "";


}

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

function filterTable() {
	let columns = new Object();
	columns["group"] = 4;
	columns["cabinet"] = 2;
	columns["prof"] = 1;
	columns["date"] = 5;

	let columnInput = document.getElementById("column");
	let dateStartInput = document.getElementById("dateStart");
	let dateEndInput = document.getElementById("dateEnd");
	let searchTermInput = document.getElementById("searchTerm");

	let dateInd = columns["date"];
	let thInd = columns[columnInput.value];

	let dateStart = new Date(dateStartInput.value);
	let dateEnd = new Date(dateEndInput.value);

	if (isNaN(dateStart) || isNaN(dateEnd)) {
		flatNotify().alert("Дата должна быть формата ГГГГ-ММ-ДД", 3000);
	}

	let searchTerm = searchTermInput.value;

	let tableHead = document.getElementById("tableHead");
	let headers = getTHArrayFromElement(tableHead);
	headers[columns["group"]].style.display = "";
	headers[columns["cabinet"]].style.display = "";
	headers[columns["prof"]].style.display = "";
	headers[thInd].style.display = "none";

	let trArr = getTRArrayFromTable();
	for (tr of trArr) {
		let trTHArr = getTHArrayFromElement(tr);

		trTHArr[columns["group"]].style.display = "";
		trTHArr[columns["cabinet"]].style.display = "";
		trTHArr[columns["prof"]].style.display = "";
		trTHArr[thInd].style.display = "none";

		let trDate = new Date(trTHArr[dateInd].getElementsByTagName("input")[0].value);
		if (trDate < dateStart || trDate > dateEnd || !trTHArr[thInd].getElementsByTagName("input")[0].value.includes(searchTerm)) {
			tr.style.display = "none";
		} else {
			tr.style.display = "";
		}
	}
}

function returnTableRows() {
	let columns = new Object();
	columns["group"] = 4;
	columns["cabinet"] = 2;
	columns["prof"] = 1;

	let tableHead = document.getElementById("tableHead");
	let headers = getTHArrayFromElement(tableHead);
	headers[columns["group"]].style.display = "";
	headers[columns["cabinet"]].style.display = "";
	headers[columns["prof"]].style.display = "";

	let trArr = getTRArrayFromTable();
	for (tr of trArr) {
		let trTHArr = getTHArrayFromElement(tr)
		trTHArr[columns["group"]].style.display = "";
		trTHArr[columns["cabinet"]].style.display = "";
		trTHArr[columns["prof"]].style.display = "";
		tr.style.display = "";
	}
}

function toggleUI() {
	let navigation = document.getElementById("navigation");
	let insertForm = document.getElementById("insertForm");
	let downloadPDF = document.getElementById("downloadPDF");
	let filterContainer = document.getElementById("filterContainer");
	let tableHead = document.getElementById("tableHead");
	let headers = getTHArrayFromElement(tableHead);
	let trArr = getTRArrayFromTable();

	let current = navigation.style.display;
	navigation.style.display = (current == "") ? "none" : "";
	insertForm.style.display = (current == "") ? "none" : "";
	downloadPDF.style.display = (current == "") ? "" : "none";
	filterContainer.style.display = (current == "") ? "" : "none";
	headers[7].style.display = (current == "") ? "none" : "";
	for (tr of trArr) {
		let trTHArr = getTHArrayFromElement(tr)
		trTHArr[7].style.display = (current == "") ? "none" : "";
	}

	if (current = "none") {
		returnTableRows()
	}
}

function printTable() {
	let filterClassContainer = document.getElementById("filterClassContainer");
	let classFiltersTitle = document.getElementById("classFiltersTitle");
	let tableTitle = document.getElementById("tableTitle");
	let dateStartInput = document.getElementById("dateStart");
	let dateEndInput = document.getElementById("dateEnd");
	let searchTermInput = document.getElementById("searchTerm");

	let dateStart = dateStartInput.value;
	let dateEnd = dateEndInput.value;

	let searchTerm = searchTermInput.value;

	classFiltersTitle.innerHTML = "Расписание для " + searchTerm + " на " + dateStart + " по " + dateEnd + ":"

	filterClassContainer.style.display = "none";
	tableTitle.style.display = "none";
	classFiltersTitle.style.display = "";
	window.print()
	classFiltersTitle.style.display = "none";
	tableTitle.style.display = "";
	filterClassContainer.style.display = "";
}

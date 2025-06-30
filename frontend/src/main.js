import {Greet} from '../wailsjs/go/main/App';

let nameElement = document.getElementById("name");
nameElement.focus();
let resultElement = document.getElementById("result");

// Setup the greet function
window.greet = async function () {
	// Get name
	let name = nameElement.value;

	// Check if the input is empty
	if (name === "") return;

	let result;
	try {
		result = await Greet(name);
		resultElement.innerText = result;
	} catch (ex) {
		console.error(ex);
	}
};
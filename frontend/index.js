const $ = (query) => document.querySelector(query);

const structTypeInput = $('#struct-type');
const jsonDataInput = $('#json');

const statusElems = {
	error: $('#result #error'),
	valid: $('#result #valid'),
	invalid: $('#result #invalid'),
	default: $('#result #default'),
};

const errMessageElem = $('#error-message');
const invalidTagsElem = $('#invalid #tags');
const invalidFieldsElem = $('#invalid #fields');

const toggleResultElem = (status) => {
	for (const key in statusElems) {
		if (status == key) {
			statusElems[status].style.visibility = 'visible';
			statusElems[status].style.display = 'inline-block';
			continue;
		}

		statusElems[key].style.visibility = 'hidden';
		statusElems[key].style.display = 'none';
	}
};

const reset = () => {
	toggleResultElem('default');
	errMessageElem.innerText = '';
	invalidTagsElem.innerText = '';
	invalidFieldsElem.innerText = '';
};

const render = () => {
	reset();

	const structType = structTypeInput.value;
	const jsonData = jsonDataInput.value;

	if (structType == '' || jsonData == '') {
		return;
	}

	const result = validateStruct(structType, jsonData);

	if (result == null) {
		errMessageElem.innerText = 'Panic!';
		toggleResultElem('error');
		return;
	}

	switch (result.status) {
		case 'valid':
			toggleResultElem('valid');
			break;

		case 'invalid':
			toggleResultElem('invalid');
			invalidTagsElem.innerText = result.invalid_result.tags.join(', ');
			invalidFieldsElem.innerText = result.invalid_result.fields.join(', ');
			break;

		case 'error':
			errMessageElem.innerText = result.error;
			toggleResultElem('error');
			break;

		default:
			break;
	}
};

// Instantiate WASM and first render
const go = new Go();
WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then((result) => {
	go.run(result.instance);
	jsonDataInput.innerText = `{ "foo": 42, "bar": 1 }`;

	structTypeInput.addEventListener('input', render);
	jsonDataInput.addEventListener('input', render);

	render();
});

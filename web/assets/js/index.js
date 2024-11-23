const select = document.getElementById('categories-select');

select.addEventListener('change', (e) => {
    const categoriesContainer = document.querySelector('.create-post-categories');
    // Parse the value as JSON to extract id and label
    const selectedValue = JSON.parse(e.target.value);
    const { id, label } = selectedValue;
    // console.log('ID:', id, 'Label:', label);

    // create the elemenet for the category
    const span = document.createElement('span');
    span.textContent = label;
    span.classList.add('selected-category');


    // create hidden input to hold the id of selected category
    const input = document.createElement('input')
    input.type = 'hidden';
    input.value = id
    input.name = 'categories'

    // add the elements (span and hidden input) 
    // at the first  position of the categories container
    categoriesContainer.prepend(input, span);


    // disable the option selected in the select
    e.target.options[e.target.selectedIndex].disabled = true;
});

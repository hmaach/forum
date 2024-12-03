document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('createPostForm');
    const titleInput = document.getElementById('postTitle');
    const contentTextarea = document.getElementById('postContent');
    const categorySelect = document.getElementById('categories-select');

    const titleError = document.getElementById('titleError');
    const contentError = document.getElementById('contentError');
    const categoryError = document.getElementById('categoryError');

    form.addEventListener('submit', (event) => {
        let isValid = true;

        titleError.textContent = '';
        contentError.textContent = '';
        categoryError.textContent = '';

        const maxTitleLength = 100;
        const maxContentLength = 1000;

        const title = titleInput.value.trim();
        const content = contentTextarea.value.trim();
        const category = categorySelect.value;

        if (title.length === 0 || title.length > maxTitleLength) {
            titleError.textContent = `Title must be between 1 and ${maxTitleLength} characters.`;
            isValid = false;
        }

        if (content.length === 0 || content.length > maxContentLength) {
            contentError.textContent = `Content must be between 1 and ${maxContentLength} characters.`;
            isValid = false;
        }

        if (!category || category === "") {
            categoryError.textContent = 'Please select a category.';
            isValid = false;
        }

        if (!isValid) {
            event.preventDefault();
        }
    });
});

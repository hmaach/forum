document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('createPostForm');
    const titleInput = document.getElementById('postTitle');
    const contentTextarea = document.getElementById('postContent');
    const categorySelect = document.getElementById('categories-select');

    form.addEventListener('submit', (event) => {
        const title = titleInput.value.trim();
        const content = contentTextarea.value.trim();
        const category = categorySelect.value;


        const maxTitleLength = 100;
        const maxContentLength = 1000;


        if (title.length === 0 || title.length > maxTitleLength) {
            alert(`The title must be between 1 and ${maxTitleLength} characters long.`);
            event.preventDefault();
            return;
        }

        if (content.length === 0 || content.length > maxContentLength) {
            alert(`The content must be between 1 and ${maxContentLength} characters long.`);
            event.preventDefault();
            return;
        }

        if (!category || category === "") {
            alert("Please select a category for your post.");
            event.preventDefault();
            return;
        }

        alert("Post validation successful. Submitting...");
    });
});

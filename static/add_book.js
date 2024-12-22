const API_ENDPOINT = '/add-book';

document.getElementById('add-author').addEventListener('click', () => {
    const authorsDiv = document.getElementById('authors');
    const index = authorsDiv.children.length;
    const newAuthor = document.createElement('div');
    newAuthor.innerHTML = `
        <input type="text" name="authors[${index}][first_name]" placeholder="Author First Name" required>
        <input type="text" name="authors[${index}][last_name]" placeholder="Author Last Name" required>
    `;
    authorsDiv.appendChild(newAuthor);
});

document.getElementById('add-book-form').addEventListener('submit', async (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    const authors = Array.from(document.querySelectorAll('#authors .author-group')).map(authorGroup => ({
        first_name: authorGroup.querySelector('[name*="first_name"]').value,
        last_name: authorGroup.querySelector('[name*="last_name"]').value,
    }));

    const payload = {
        title: formData.get('title'),
        publication_year: parseInt(formData.get('publication_year'), 10),
        genre: formData.get('genre'),
        authors,
    };

    try {
        const response = await fetch(API_ENDPOINT, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload),
        });

        const result = await response.json();
        if (response.ok) {
            showAlert(result.message || 'Book added successfully!');
            window.location.href = '/';
        } else {
            showAlert(result.error || 'Failed to add book.');
        }
    } catch (error) {
        console.error('Error adding book:', error);
        showAlert('An unexpected error occurred. Please try again.');
    }
});



function showAlert(message) {
    const alertContainer = document.createElement('div');
    alertContainer.className = `custom-alert`;
    alertContainer.setAttribute('aria-live', 'assertive');
    alertContainer.innerHTML = `
        <div class="alert-body">
            <p>${message}</p>
        </div>
    `;

    document.body.appendChild(alertContainer);

    setTimeout(() => {
        alertContainer.style.opacity = '0';
        alertContainer.style.transition = 'opacity 0.5s';
        setTimeout(() => alertContainer.remove(), 500);
    }, 2000);
}

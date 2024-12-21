const API_ENDPOINTS = {
    books: '/books-with-authors',
};

async function fetchBooksWithAuthors() {
    try {
        const books = await fetch(API_ENDPOINTS.books).then(res => res.json());
        const booksList = document.getElementById('books-list');
        booksList.innerHTML = books.map(book => `
            <p>
                <a href="/book/${book.book_id}" class="book-link">${book.title}</a>
            </p>
        `).join('');
    } catch (error) {
        console.error(`Error fetching books: ${error}`);
        document.getElementById('books-list').innerHTML = `
            <p class="error">Error loading books: ${error.message}</p>`;
    }
}

document.addEventListener('DOMContentLoaded', fetchBooksWithAuthors);
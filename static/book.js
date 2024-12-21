const API_ENDPOINTS = {
    book: id => `/book-data/${id}`,
};

async function fetchBookDetails(bookId) {
    try {
        const book = await fetch(API_ENDPOINTS.book(bookId)).then(res => res.json());
        document.getElementById('book-details').innerHTML = `
            <h2>${book.title}</h2>
            <p><strong>Publication Year:</strong> ${book.publication_year}</p>
            <p><strong>Genre:</strong> ${book.genre}</p>
            <p><strong>Authors:</strong> ${book.authors || 'Author unknown'}</p>
        `;
    } catch (error) {
        console.error(`Error fetching book details: ${error}`);
        document.getElementById('book-details').innerHTML = `
            <p class="error">Error loading book details: ${error.message}</p>`;
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const bookId = window.location.pathname.split('/').pop();
    fetchBookDetails(bookId);
});

async function fetchBooksWithAuthors() {
    try {
        const response = await fetch('/books-with-authors');
        if (!response.ok) {
            throw new Error('Failed to fetch books with authors');
        }
        const books = await response.json();
        renderBooksWithAuthors(books);
    } catch (error) {
        console.error('Error fetching books with authors:', error);
        document.getElementById('books-list').innerHTML = `<p class="error">Error loading books with authors.</p>`;
    }
}

function renderBooksWithAuthors(books) {
    const tableBody = books.map(book => `
        <tr>
            <td>${book.title}</td>
            <td>${book.publication_year}</td>
            <td>${book.genre}</td>
            <td>${book.authors || 'Author unknown'}</td>
        </tr>
    `).join('');
    document.getElementById('books-list').innerHTML = `
        <table>
            <thead>
                <tr>
                    <th>Title</th>
                    <th>Publication Year</th>
                    <th>Genre</th>
                    <th>Authors</th>
                </tr>
            </thead>
            <tbody>${tableBody}</tbody>
        </table>
    `;
}

document.addEventListener('DOMContentLoaded', fetchBooksWithAuthors);

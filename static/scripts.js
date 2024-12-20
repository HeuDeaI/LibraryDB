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

document.getElementById('signup-form').addEventListener('submit', async (event) => {
    event.preventDefault();
    const reader = {
        first_name: document.getElementById('first-name').value,
        last_name: document.getElementById('last-name').value,
        phone_number: document.getElementById('phone-number').value,
        email: document.getElementById('email').value,
    };

    try {
        const response = await fetch('/signup-reader', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(reader),
        });
        if (!response.ok) throw new Error('Failed to sign up');
        const data = await response.json();
        alert(`Signed up successfully! Reader ID: ${data.reader_id}`);
    } catch (error) {
        console.error('Error signing up:', error);
    }
});

async function loanBook(bookId, readerId) {
    try {
        const response = await fetch('/loan-book', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ book_id: bookId, reader_id: readerId }),
        });
        if (!response.ok) throw new Error('Failed to loan book');
        const data = await response.json();
        alert(`Book loaned successfully! Loan ID: ${data.loan_id}`);
    } catch (error) {
        console.error('Error loaning book:', error);
    }
}

document.addEventListener('DOMContentLoaded', fetchBooksWithAuthors);

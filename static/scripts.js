const API_ENDPOINTS = {
    books: '/books-with-authors',
    signup: '/signup-reader',
    loan: '/loan-book',
};

async function fetchWithErrorHandling(url, options = {}) {
    try {
        const response = await fetch(url, options);
        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || `HTTP Error: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error(`Error in fetch request: ${error.message}`);
        throw error;
    }
}

async function fetchBooksWithAuthors() {
    try {
        const books = await fetchWithErrorHandling(API_ENDPOINTS.books);
        renderBooksWithAuthors(books);
    } catch (error) {
        document.getElementById('books-list').innerHTML = `
            <p class="error">Error loading books: ${error.message}</p>`;
    }
}

function renderBooksWithAuthors(books) {
    const tableBody = books.map(book => `
        <tr>
            <td>${book.title}</td>
            <td>${book.publication_year}</td>
            <td>${book.genre}</td>
            <td>${book.authors || 'Author unknown'}</td>
            <td>
                <button onclick="promptAndLoanBook(${book.book_id})">Loan</button>
            </td>
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
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>${tableBody}</tbody>
        </table>`;
}

async function promptAndLoanBook(bookId) {
    const readerId = parseInt(prompt('Enter your Reader ID:'), 10);
    if (isNaN(readerId)) {
        alert("Invalid Reader ID. Please enter a valid number.");
        return;
    }

    try {
        const loanData = { book_id: bookId, reader_id: readerId };
        const response = await fetchWithErrorHandling(API_ENDPOINTS.loan, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(loanData),
        });
        alert(`Book loaned successfully! Loan ID: ${response.loan_id}`);
    } catch (error) {
        alert(`Error loaning book: ${error.message}`);
    }
}

async function handleSignup(event) {
    event.preventDefault();

    const reader = {
        first_name: document.getElementById('first-name').value.trim(),
        last_name: document.getElementById('last-name').value.trim(),
        phone_number: document.getElementById('phone-number').value.trim(),
        email: document.getElementById('email').value.trim(),
    };

    if (!reader.first_name || !reader.last_name || !reader.phone_number || !reader.email) {
        alert("All fields are required.");
        return;
    }

    try {
        const response = await fetchWithErrorHandling(API_ENDPOINTS.signup, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(reader),
        });
        alert(`Signed up successfully! Reader ID: ${response.reader_id}`);
        document.getElementById('signup-form').reset();
    } catch (error) {
        alert(`Error signing up: ${error.message}`);
    }
}

document.addEventListener('DOMContentLoaded', () => {
    fetchBooksWithAuthors();
    document.getElementById('signup-form').addEventListener('submit', handleSignup);
});

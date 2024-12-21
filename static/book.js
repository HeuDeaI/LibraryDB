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

document.addEventListener('DOMContentLoaded', () => {
    const bookId = window.location.pathname.split('/').pop();
    fetchBookDetails(bookId);

    const loanFormContainer = document.getElementById('loan-form-container');
    const openLoanFormButton = document.getElementById('open-loan-form');
    const closeLoanFormButton = document.getElementById('close-loan-form');

    openLoanFormButton.addEventListener('click', () => {
        loanFormContainer.classList.remove('hidden');
    });

    closeLoanFormButton.addEventListener('click', () => {
        loanFormContainer.classList.add('hidden');
    });

    document.getElementById('loan-form').addEventListener('submit', async (e) => {
        e.preventDefault();

        const formData = new FormData(e.target);
        const payload = Object.fromEntries(formData.entries());
        payload.book_id = bookId;

        try {
            const response = await fetch('/loan-book', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload),
            });

            const result = await response.json();
            if (response.ok) {
                alert(result.message || 'Book loaned successfully!');
                loanFormContainer.classList.add('hidden');
            } else {
                alert(result.error || 'Failed to loan book.');
            }
        } catch (error) {
            console.error('Error loaning book:', error);
            alert('An unexpected error occurred.');
        }
    });
});

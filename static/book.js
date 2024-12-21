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

    openLoanFormButton.addEventListener('click', () => {
        loanFormContainer.classList.remove('hidden');
    });

    document.addEventListener('click', (event) => {
        if (!loanFormContainer.classList.contains('hidden') && !loanFormContainer.contains(event.target) && event.target !== openLoanFormButton) {
            loanFormContainer.classList.add('hidden');
        }
    });

    document.getElementById('loan-form').addEventListener('submit', async (e) => {
        e.preventDefault();

        const formData = new FormData(e.target);
        const payload = Object.fromEntries(formData.entries());
        payload.book_id = bookId;

        const firstName = payload.first_name.trim();
        const lastName = payload.last_name.trim();
        const phoneNumber = payload.phone_number.trim();
        const email = payload.email.trim();

        const phonePattern = /^\+375[0-9]{9}$/;
        const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

        if (!firstName || firstName.length > 31) {
            showAlert('First Name is required and must be 31 characters or fewer.');
            return;
        }
        if (!lastName || lastName.length > 31) {
            showAlert('Last Name is required and must be 31 characters or fewer.');
            return;
        }
        if (phoneNumber && !phonePattern.test(phoneNumber)) {
            showAlert('Phone Number must match the format +375XXXXXXXXX.');
            return;
        }
        if (!email || !emailPattern.test(email)) {
            showAlert('Email is required and must be a valid email address.');
            return;
        }

        try {
            const response = await fetch('/loan-book', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload),
            });

            const result = await response.json();
            if (response.ok) {
                showAlert(result.message || 'Book loaned successfully!');
                loanFormContainer.classList.add('hidden');
            } else {
                showAlert(result.error || 'Failed to loan book.');
            }
        } catch (error) {
            console.error('Error loaning book:', error);
            showAlert('An unexpected error occurred.');
        }
    });
});

function showAlert(message, type = 'info') {
    const alertContainer = document.createElement('div');
    alertContainer.className = `custom-alert ${type}`;
    alertContainer.innerHTML = `
        <div class="alert-header">
            <h4>${type.toUpperCase()}</h4>
            <button class="close-alert-btn">&times;</button>
        </div>
        <div class="alert-body">
            <p>${message}</p>
        </div>
    `;

    document.body.appendChild(alertContainer);

    const closeButton = alertContainer.querySelector('.close-alert-btn');
    closeButton.addEventListener('click', () => {
        alertContainer.remove();
    });

    setTimeout(() => {
        alertContainer.remove();
    }, 5000);
}
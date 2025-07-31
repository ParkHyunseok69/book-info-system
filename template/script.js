let currentBookId = null;

        // Sample book data - replace with your Go backend calls
        const sampleBooks = {
            1: {
                id: 1,
                title: '1984',
                author: 'George Orwell',
                genre: 'Science Fiction',
                isbn: '9780451524935',
                publicationYear: 1949,
                pages: 328,
                status: 'available',
                purchasePrice: 12.99,
                purchaseDate: '2023-01-15',
                notes: 'Classic dystopian novel'
            },
            2: {
                id: 2,
                title: 'Pride and Prejudice',
                author: 'Jane Austen',
                genre: 'Romance',
                isbn: '9780141439518',
                publicationYear: 1813,
                pages: 432,
                status: 'borrowed',
                purchasePrice: 9.99,
                purchaseDate: '2023-02-10',
                notes: 'Borrowed to Sarah on March 1st'
            }
        };

        function viewBook(bookId) {
            const book = sampleBooks[bookId]; // Replace with API call
            if (!book) return;

            const detailsHtml = `
                <div class="detail-item">
                    <span class="detail-label">Title:</span>
                    <span class="detail-value">${book.title}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">Author:</span>
                    <span class="detail-value">${book.author}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">Genre:</span>
                    <span class="detail-value">${book.genre}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">ISBN:</span>
                    <span class="detail-value">${book.isbn}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">Published:</span>
                    <span class="detail-value">${book.publicationYear}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">Pages:</span>
                    <span class="detail-value">${book.pages}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">Status:</span>
                    <span class="detail-value">${book.status}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">Purchase Price:</span>
                    <span class="detail-value">$${book.purchasePrice}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">Purchase Date:</span>
                    <span class="detail-value">${book.purchaseDate}</span>
                </div>
                <div class="detail-item">
                    <span class="detail-label">Notes:</span>
                    <span class="detail-value">${book.notes || 'No notes'}</span>
                </div>
            `;

            document.getElementById('bookDetails').innerHTML = detailsHtml;
            openModal('viewModal');
        }

        function editBook(bookId) {
            const book = sampleBooks[bookId]; // Replace with API call
            if (!book) return;

            currentBookId = bookId;

            // Populate form fields
            document.getElementById('editTitle').value = book.title;
            document.getElementById('editAuthor').value = book.author;
            document.getElementById('editGenre').value = book.genre.toLowerCase().replace(' ', '-');
            document.getElementById('editIsbn').value = book.isbn;
            document.getElementById('editYear').value = book.publicationYear;
            document.getElementById('editPages').value = book.pages;
            document.getElementById('editStatus').value = book.status;
            document.getElementById('editPrice').value = book.purchasePrice;
            document.getElementById('editNotes').value = book.notes || '';

            openModal('editModal');
        }

        function deleteBook(bookId) {
            const book = sampleBooks[bookId]; // Replace with API call
            if (!book) return;

            currentBookId = bookId;
            document.getElementById('deleteBookTitle').textContent = book.title;
            openModal('deleteModal');
        }

        function confirmDelete() {
            if (!currentBookId) return;

            const deleteBtn = event.target;
            const loading = document.getElementById('deleteLoading');
            
            deleteBtn.disabled = true;
            loading.style.display = 'inline-block';

            // Simulate API call
            setTimeout(() => {
                // Replace with actual API call to delete book
                console.log('Deleting book:', currentBookId);
                
                // Remove book card from UI
                const bookCard = document.querySelector(`[data-book-id="${currentBookId}"]`);
                if (bookCard) {
                    bookCard.remove();
                }

                closeModal('deleteModal');
                showSuccess('Book deleted successfully!');
                
                deleteBtn.disabled = false;
                loading.style.display = 'none';
                currentBookId = null;
            }, 1000);
        }

        // Form submission handler
        document.getElementById('editBookForm').addEventListener('submit', function(e) {
            e.preventDefault();
            
            const saveBtn = e.target.querySelector('button[type="submit"]');
            const loading = document.getElementById('saveLoading');
            
            saveBtn.disabled = true;
            loading.style.display = 'inline-block';

            // Collect form data
            const formData = {
                id: currentBookId,
                title: document.getElementById('editTitle').value,
                author: document.getElementById('editAuthor').value,
                genre: document.getElementById('editGenre').value,
                isbn: document.getElementById('editIsbn').value,
                publicationYear: parseInt(document.getElementById('editYear').value),
                pages: parseInt(document.getElementById('editPages').value),
                status: document.getElementById('editStatus').value,
                purchasePrice: parseFloat(document.getElementById('editPrice').value),
                notes: document.getElementById('editNotes').value
            };

            // Simulate API call
            setTimeout(() => {
                console.log('Updating book:', formData);
                // Replace with actual API call
                
                closeModal('editModal');
                showSuccess('Book updated successfully!');
                
                saveBtn.disabled = false;
                loading.style.display = 'none';
                currentBookId = null;
                
                // Refresh the page or update the book card
                // location.reload(); // Simple approach
            }, 1000);
        });

        function openModal(modalId) {
            document.getElementById(modalId).style.display = 'block';
            document.body.style.overflow = 'hidden';
        }

        function closeModal(modalId) {
            document.getElementById(modalId).style.display = 'none';
            document.body.style.overflow = 'auto';
        }

        function showSuccess(message) {
            document.getElementById('successMessage').textContent = message;
            openModal('successModal');
        }

        function openAddBookModal() {
            // Reset form and open edit modal for adding
            document.getElementById('editBookForm').reset();
            currentBookId = null;
            document.querySelector('#editModal .modal-title').textContent = 'Add New Book';
            openModal('editModal');
        }

        // Close modal when clicking outside
        window.onclick = function(event) {
            if (event.target.classList.contains('modal')) {
                event.target.style.display = 'none';
                document.body.style.overflow = 'auto';
            }
        }

        // Close modal with Escape key
        document.addEventListener('keydown', function(event) {
            if (event.key === 'Escape') {
                const openModal = document.querySelector('.modal[style*="block"]');
                if (openModal) {
                    openModal.style.display = 'none';
                    document.body.style.overflow = 'auto';
                }
            }
        });
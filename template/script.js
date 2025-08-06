let currentBookID = null;

        function listBooks() {
            const booksGrid = document.getElementById('booksGrid');
            booksGrid.innerHTML = '';
            fetch('http://localhost:8080/books')
            .then(res => res.json())
            .then(books => {
                books.forEach(book => {
                    const card = document.createElement('div');
                    card.classList.add('book-card');
                    card.setAttribute('data-book-id', book.id);
                    card.innerHTML =  `<h3>${book.title}</h3>
                    <p>by ${book.author}</p>
                    <p>${book.publication_year} · ${book.pages} pages </p>
                    <span class="status ${book.status === 'COMPLETED' ? 'Completed' : 'Ongoing'}">
                        ${book.status}
                    </span>
                    <div class="book-actions">
                        <button class="action-btn btn-view" onclick="viewBook('${book.id}')">View</button>
                        <button class="action-btn btn-edit" onclick="editBook('${book.id}')">Edit</button>
                        <button class="action-btn btn-delete" onclick="deleteBook('${book.id}')">Delete</button>
                    </div>
                `;
                document.getElementById('booksGrid').appendChild(card);
                });
            })
        }
        function viewBook(id) {
            fetch(`http://localhost:8080/book/${encodeURIComponent(id)}`)
            .then(res => res.json())
            .then(book => {
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
                        <span class="detail-label">Published:</span>
                        <span class="detail-value">${book.publication_year}</span>
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
                        <span class="detail-label">Purchase Date:</span>
                        <span class="detail-value">${book.date_acquired}</span>
                    </div>
                    <div class="detail-item">
                        <span class="detail-label">Notes:</span>
                        <span class="detail-value">${book.summary || 'No Summary'}</span>
                    </div>
                `;

                document.getElementById('bookDetails').innerHTML = detailsHtml;
                openModal('viewModal');
            })
        }

        function editBook(id) {
            fetch(`http://localhost:8080/book/${encodeURIComponent(id)}`)
            .then(res => res.json())
            .then(book => {
                // Populate form fields
                currentBookID = book.id;
                document.getElementById('editTitle').value = book.title;
                document.getElementById('editAuthor').value = book.author;
                document.getElementById('editSummary').value = book.summary || '';
                document.getElementById('editGenre').value = book.genre.replace(' ', '-');
                document.getElementById('editYear').value = book.publication_year;
                document.getElementById('editPages').value = book.pages;
                document.getElementById('editAcquired').value = book.date_acquired;
                document.getElementById('editStatus').value = book.status;
                openModal('editModal');
            })
        }

        function deleteBook(id) {
            currentBookID = id; 

            const bookCard = document.querySelector(`[data-book-id="${id}"]`);
            const title = bookCard?.querySelector('h3')?.textContent || 'this book';

            document.getElementById('deleteBookTitle').textContent = title;
            openModal('deleteModal');
        }

        function confirmDelete() {
            if (!currentBookID) return;

            const deleteBtn = event.target;
            const loading = document.getElementById('deleteLoading');
            
            deleteBtn.disabled = true;
            loading.style.display = 'inline-block';

            fetch(`http://localhost:8080/book/${encodeURIComponent(currentBookID)}`, {
            method: 'DELETE',
            })
            .then(res => {
                if (!res.ok) throw new Error('Failed to delete');
                const bookCard = document.querySelector(`[data-book-id="${currentBookID}"]`);
                if (bookCard) {
                    bookCard.remove();
                }

                closeModal('deleteModal');
                showSuccess('Book deleted successfully!');
            })
            .catch(err => {
                console.error(err);
                alert('Failed to delete book.')
            })
            .finally(()=> {
                deleteBtn.disabled = false;
                loading.style.display = 'none';
                currentBookID= null;
            })
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
                title: document.getElementById('editTitle').value,
                author: document.getElementById('editAuthor').value,
                summary: document.getElementById('editSummary').value,
                genre: document.getElementById('editGenre').value,
                publication_year: document.getElementById('editYear').value,
                pages: parseInt(document.getElementById('editPages').value),
                date_acquired: document.getElementById('editAcquired').value,
                status: document.getElementById('editStatus').value,
                
            };
            const method = currentBookID ? 'PUT' : 'POST';
            const url = currentBookID 
                ? `http://localhost:8080/book/${encodeURIComponent(currentBookID)}`
                : 'http://localhost:8080/book';

            fetch(url, {
                method: method,
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(formData)
            })
            .then(res => {
                if (!res.ok) throw new Error('Failed to save book');
                return res.json();
            })
            .then(() => {
                closeModal('editModal');
                showSuccess(currentBookID ? 'Book updated successfully!' : 'Book added successfully!');
                listBooks();
            })
            .catch(err => {
                console.error('Error saving book:', err);
                alert('Save failed!');
            })
            .finally(() => {
                saveBtn.disabled = false;
                loading.style.display = 'none';
                currentBookID = null;
            });
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
            currentBookID= null;
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
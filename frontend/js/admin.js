const API_BASE = '/api';

// Handle form submission
document.getElementById('articleForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const formData = new FormData();
    const imageFile = document.getElementById('image').files[0];

    let imageUrl = null;

    // Upload image if selected
    if (imageFile) {
        const imageFormData = new FormData();
        imageFormData.append('image', imageFile);

        try {
            const uploadResponse = await fetch(`${API_BASE}/upload-image`, {
                method: 'POST',
                body: imageFormData
            });

            if (!uploadResponse.ok) {
                alert('ছবি আপলোড করতে ব্যর্থ হয়েছে');
                return;
            }

            const uploadResult = await uploadResponse.json();
            imageUrl = uploadResult.url;
        } catch (error) {
            alert('ছবি আপলোড করতে ত্রুটি: ' + error.message);
            return;
        }
    }

    // Prepare article data
    const articleData = {
        title: document.getElementById('title').value,
        content: document.getElementById('content').value,
        category: document.getElementById('category').value,
        author: document.getElementById('author').value,
        featured: document.getElementById('featured').checked,
        image: imageUrl
    };

    // Add article
    try {
        const response = await fetch(`${API_BASE}/add-article`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(articleData)
        });

        if (!response.ok) {
            const errorData = await response.json();
            alert('খবর যোগ করতে ব্যর্থ হয়েছে: ' + (errorData.message || 'Unknown error'));
            return;
        }

        const result = await response.json();
        alert('খবর সফলভাবে যোগ করা হয়েছে!');
        
        // Reset form
        document.getElementById('articleForm').reset();
        document.getElementById('imagePreview').innerHTML = '';
    } catch (error) {
        alert('খবর যোগ করতে ত্রুটি: ' + error.message);
    }
});

// Image preview
document.getElementById('image').addEventListener('change', function(e) {
    const file = e.target.files[0];
    const preview = document.getElementById('imagePreview');

    if (file) {
        const reader = new FileReader();
        reader.onload = function(e) {
            preview.innerHTML = `<img src="${e.target.result}" alt="Preview" style="max-width: 200px; max-height: 200px;">`;
        };
        reader.readAsDataURL(file);
    } else {
        preview.innerHTML = '';
    }
});
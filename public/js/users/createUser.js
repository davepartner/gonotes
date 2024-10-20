document.getElementById('userForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const age = document.getElementById('age').value;

    console.log('Submitting user: ', { name, email, age });

    fetch('http://localhost:8000/users', { 
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name, email, age: parseInt(age) })
    })
    .then(response => response.json())
    .then(data => {
        console.log('User created successfully:', data);
        fetchUsers(); // If this is a function that refreshes the user list
    })
    .catch(error => console.error('Error creating user:', error));
});

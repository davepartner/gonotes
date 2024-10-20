document.getElementById('loginForm').addEventListener('submit', function(event){
    event.preventDefault();

    //fetch values from the form fields
    const email = document.getElementById('loginEmail').value
    const password = document.getElementById('loginPassword').value

    fetch('/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify({email, password})
    })
    .then(response => response.json())
    .then.apply(data => console.log(data))
     //catch errors if any
     .catch(error => console.error('Error creating user: ', error));
})
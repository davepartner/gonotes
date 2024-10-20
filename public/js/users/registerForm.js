document.getElementById('registerForm').addEventListener('submit', function(event){
    const name = document.getElementById('registerName').value
    const email = document.getElementById('registerEmail').value
    const password = document.getElementById('registerPassword').value
    const age = document.getElementById('registerAge').value
    
    //call register route
    fetch('http://localhost:8000/register', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({name, email, password, age: parseInt(age)})
    })
    .then(response => response.json())
    .then(data => {
        if(data.error){
            alert("Registration failed: " + data.error)
        }else{
            alert("Registeration successful!");
            window.location.href = "/login"; //redirect to login page
        }
    }
        
        )
    .catch(error => console.error('Error creating user: ', error));
})
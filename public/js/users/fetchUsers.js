
//fetch users from the server
function fetchUsers(){
    fetch('/users')
    .then(response => response.json())
    .then(data => {
        //find the tbody tag in the table
        const tbody = document.querySelector('#userTable tbody');
        //now insert
        tbody.innerHTML = ''; //empty the tbody
        data.forEach(user => {
            const row = document.createElement('tr');
            row.innerHTML = `
            <td>${user.id}</td>
            <td>${user.name}</td>
            <td>${user.email}</td>
            <td>${user.age}</td>
            <td>
                <button onclick="deleteUser(${user.id})">Delete</button>
            </td>
            `;
            tbody.appendChild(row);
        })
    })
    //catch errors if any
    .catch(error => console.error('Error fetching users: ', error));
}

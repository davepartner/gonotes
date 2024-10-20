//delete a user
function deleteUser(id){
    fetch(`/users/${id}`, {method: 'DELETE'})
    .then(() => fetchUsers())
    //catch errors if any
    .catch(error => console.error('Error deleting user: ', error));
}
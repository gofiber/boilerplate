$(document).ready(function () {
  listUsers()
});

function listUsers() {
  $.getJSON("/api/v1/users", (data) => {
    var users = ''
    for (var i = 0; i < data.user.length; i++) {
      users += '<li class="list-group-item">' + data.user[i].name + '</li>'
    }
    $('#users').html('')
    $('#users').append(users)
  })
}

$('#add_user').on('click', (e) => {
  var user = $('#user').val()
  $.post("/api/v1/users", "user=" + user, (data) => {
    $('#users').prepend('<li class="list-group-item">' + data.user.name + '</li>')
  })
})
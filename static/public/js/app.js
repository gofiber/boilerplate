$(document).ready(function () {
  listUsers()
});

function listUsers() {
  $.getJSON("/api/v1/users", (data) => {
    var users = ''
    $.each(data.user, function(index, val) {
      users += '<li class="list-group-item">' + val.name + '</li>'
    });
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
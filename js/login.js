document.addEventListener('DOMContentLoaded', function() {

	var form = document.getElementById('form');
	var usernameInput = document.getElementById('username');
	var passwordInput = document.getElementById('password');

	usernameInput.focus();

	usernameInput.addEventListener('keydown', function(e) {
		if (e.keyCode === 13) {
			e.preventDefault();
			if (usernameInput.value.trim() !== '') {
				passwordInput.focus();
			}
		}
	});

	passwordInput.addEventListener('keydown', function(e) {
		if (e.keyCode === 13) {
			e.preventDefault();
			if (passwordInput.value.trim() !== '') {
				form.submit();
			}
		}
	});

});

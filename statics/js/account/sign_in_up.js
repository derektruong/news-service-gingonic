
const signUpButton = document.getElementById('signUp');
const emailSignUpInput = document.getElementById('emailSignUp');
const signInButton = document.getElementById('signIn');
const container = document.getElementById('container');

signUpButton.addEventListener('click', () => {
	container.classList.add("right-panel-active");
});

signInButton.addEventListener('click', () => {
	container.classList.remove("right-panel-active");
});



function help() {
	fetch('../statics/data.json')
		.then(function (response) {
			return response.json();
		})
		.then(function (data) {
			(function() {
				let flag = false
				for (var i = 0; i < data.length; i++) {
					if(emailSignUpInput.value === data[i].email) {
						document.getElementById("emailError").innerHTML 
							= "This email is already registered!";
						flag = true;
					}
					if(flag) break;
					else {
						document.getElementById("emailError").innerHTML 
							= "";
					}
				}
			})();
		})
		.catch(function (err) {
			console.log('error: ' + err);
		});

	
}


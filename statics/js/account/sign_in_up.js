
const signUpButton = document.getElementById('signUp');
const emailSignUpInput = document.getElementById('emailSignUp');
const signInButton = document.getElementById('signIn');
const btnSignInSubmit = document.getElementById('btnSignInSubmit');
const container = document.getElementById('container');

signUpButton.addEventListener('click', () => {
	container.classList.add("right-panel-active");
});

signInButton.addEventListener('click', () => {
	container.classList.remove("right-panel-active");
});



function helpEmailSignUp() {
	fetch('../statics/data.json')
		.then(function (response) {
			return response.json();
		})
		.then(function (data) {
			(function() {
				let flag = false
				for (var i = 0; i < data.length; i++) {
					if(emailSignUpInput.value === data[i].email) {
						document.getElementById("emailExist").innerHTML 
							= "<p class='text-warning m-0'>This email is already registered!</p>";

						document.getElementById("btnSignUpSubmit").disabled = true;
						flag = true;
					}
					if(flag) break;
					else {
						document.getElementById("emailExist").innerHTML 
							= "";
						document.getElementById("btnSignUpSubmit").disabled = false;
					}
				}
			})();
		})
		.catch(function (err) {
			console.log('error: ' + err);
		});

	
}

function formSignInSubmit() {
	const formData = new FormData();

	formData.append('email', document.getElementById("emailSignIn").value);
	formData.append('password', document.getElementById("passwordSignIn").value);
	
	fetch('/signin', {
		method: 'POST', 
		body: formData,
	  }).then(function (response) {
			return response.json();
		})
		.then(function (data) {
			(function() {
				if(data.message === "Email wrong") {
					document.getElementById("notExistEmail").innerHTML 
							= "<p class='text-warning m-0'>"+ data.text +"</p>";
					document.getElementById("wrongPass").innerHTML 
							= "";
					document.getElementById("btnSignInSubmit").disabled = true;
				} else if(data.message === "Pass wrong") {
					document.getElementById("wrongPass").innerHTML 
							= "<p class='text-warning m-0'>"+ data.text +"</p>";
					document.getElementById("notExistEmail").innerHTML 
							= "";
					document.getElementById("btnSignInSubmit").disabled = true;

				} else {
					document.getElementById("notExistEmail").innerHTML 
							= "";
					document.getElementById("wrongPass").innerHTML 
							= "";
					document.getElementById("btnSignInSubmit").disabled = false;

				}
			})();
		})
		.catch(function (err) {
			console.log('error: ' + err);
		});
}



btnSignInSubmit.addEventListener("click", () => {
	const formData = new FormData();

	formData.append('email', document.getElementById("emailSignIn").value);
	formData.append('password', document.getElementById("passwordSignIn").value);

	fetch('/signin', {
		method: 'POST', 
		body: formData,
	  }).then(function (response) {
			return response.json();
		})
		.then(function (data) {
			console.log(data);
			(function() {
				if(data.message === "Set cookie successfully") {
					window.location.href = "/";
				}
				
			})();
		})
		.catch(function (err) {
			console.log('error: ' + err);
		});
})

// Handle cookies and session

// function setCookie(cname,cvalue,exdays) {
// 	const d = new Date();
// 	d.setTime(d.getTime() + (exdays*24*60*60*1000));
// 	let expires = "expires=" + d.toGMTString();
// 	document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
//   }
  
//   function getCookie(cname) {
// 	let name = cname + "=";
// 	let decodedCookie = decodeURIComponent(document.cookie);
// 	let ca = decodedCookie.split(';');
// 	for(let i = 0; i < ca.length; i++) {
// 	  let c = ca[i];
// 	  while (c.charAt(0) == ' ') {
// 		c = c.substring(1);
// 	  }
// 	  if (c.indexOf(name) == 0) {
// 		return c.substring(name.length, c.length);
// 	  }
// 	}
// 	return "";
//   }
  
//   function checkCookie() {
// 	let user = getCookie("username");
// 	if (user != "") {
// 	  alert("Welcome again " + user);
// 	} else {
// 	   user = prompt("Please enter your name:","");
// 	   if (user != "" && user != null) {
// 		 setCookie("username", user, 25);
// 	   }
// 	}
//   }

  //
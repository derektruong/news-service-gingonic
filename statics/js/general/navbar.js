document.getElementById("navigation").innerHTML = 
`<!-- Navbar -->
<nav class="shadow-lg navbar navbar-expand-lg navbar-dark fixed-top">
	<!-- Container wrapper -->
	<div class="container">
	  <!-- Navbar brand -->
	  <a
		class="navbar-brand me-2"
		href=""
		target="_self"
	  >
		<img
		  src="../statics/image/logo.png"
		  height="50"
		  alt=""
		  loading="lazy"
		  style="margin-top: -1px"
		/>
	  </a>

	  <!-- Toggle button -->
	  <button
		class="navbar-toggler"
		type="button"
		data-mdb-toggle="collapse"
		data-mdb-target="#navbarButtonsExample"
		aria-controls="navbarButtonsExample"
		aria-expanded="false"
		aria-label="Toggle navigation"
	  >
		<i class="fas fa-bars"></i>
	  </button>

	  <!-- Collapsible wrapper -->
	  <div class="collapse navbar-collapse" id="navbarButtonsExample">
		<!-- Left links -->
		<ul class="navbar-nav me-auto mb-2 mb-sm-1">
		  <li class="nav-item">
			<a
			  class="nav-link text-light fs-5 fw-bold me-2"
			  href="/"
			  >Cyber News</a
			>
		  </li>
		</ul>

		<div class="navbar-collapse collapse show" id="navbarNavAltMarkup">
		  <div class="navbar-nav">
			<a
			  class="nav-link active text-light fs-5 me-2"
			  aria-current="page"
			  href="/headlines"
			  >Headlines</a
			>
			<a
			  class="nav-link text-light fs-5"
			  href="/stocks"
			  >Stocks</a
			>
			<a
			  class="nav-link text-light fs-5"
			  href="/technology"
			  >Technologies</a
			>
			<a
			  class="nav-link text-light fs-5"
			  href="/science"
			  >Science</a
			>
			<a
			  class="nav-link text-light fs-5"
			  href="/sport"
			  >Sport</a
			>
		  </div>
		</div>

		<!-- Left links -->
		<form class="d-flex input-group w-auto" action="/search" method="GET">
		  <input
		  	id="search-box"
			required
			type="search"
			class="form-control rounded"
			placeholder="Search"
			value="{{ .Query }}"
			name="q"
			aria-label="Search"
			aria-describedby="search-addon"
		  />
		  <span class="input-group-text border-0" id="search-addon">
			<i class="fas fa-search"></i>
		  </span>
		</form>

		<div class="d-flex align-items-center">
		  <a
			class="btn btn-dark px-3 me-2"
			href="https://github.com/derektruong/news-app"
			role="button"
			><i class="fab fa-github"></i
		  ></a>

		  <a href="/signin">
			  <button class="btn btn-primary me-3">
				  Sign in
			  </button>
		  </a>
			
			
		  </form>
		</div>
	  </div>
	  <!-- Collapsible wrapper -->
	</div>
	<!-- Container wrapper -->
  </nav>
  <!-- Navbar -->`;

  if(document.getElementById("search-box").value === "{{ .Query }}") {
	document.getElementById("search-box").value = "";
  }
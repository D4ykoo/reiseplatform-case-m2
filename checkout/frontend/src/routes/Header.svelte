<script>
	import { page } from '$app/stores';
	import logo from '$lib/images/svelte-logo.svg';
	import github from '$lib/images/github.svg';
</script>

<header>
	<div class="navbar bg-base-100 top-0 mb-6" id="top">
    <div class="navbar-start">
      <div class="dropdown">
        <div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h8m-8 6h16"
            />
          </svg>
        </div>
        <ul
          tabindex="-1"
          class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
        >
          <li><a href="#top">Login/ Register</a></li>
          <li><a href="#top">Usermanagement</a></li>
        </ul>
      </div>
      <a class="btn btn-ghost text-xl" href="/">travMa</a>
    </div>
    <div class="navbar-center hidden lg:flex">
      <ul class="menu menu-horizontal px-1">
        <li><a href="#top">Login/ Register</a></li>
        <li><a href="#top">Usermanagement</a></li>
      </ul>
    </div>
    <div class="navbar-end">
      <div class="mr-4">
        <a class="btn" href="#top">Overview</a>
      </div>
      <div class="mr-4">
        <a class="btn" href="#top">Checkout</a>
      </div>
      <div>
        <a class="btn" href="#top">Logout</a>
      </div>
    </div>
  </div>
</header>
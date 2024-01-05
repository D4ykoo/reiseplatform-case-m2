<template>
  <div class="navbar bg-base-100 top-0 mb-6">
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
          tabindex="0"
          class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
        >
          <li><a href="{{  }}">Travelmangement</a></li>
          <li><a>Login/ Register</a></li>
          <li><a>Usermanagement</a></li>
        </ul>
      </div>
      <a class="btn btn-ghost text-xl" href="/">travMa</a>
    </div>
    <div class="navbar-center hidden lg:flex">
      <ul class="menu menu-horizontal px-1">
        <li><a :href="travelmanagementUrl">Travelmangement</a></li>
        <li><a href="/">Login/ Register</a></li>
        <li>
          <router-link :to="{ name: 'usermanagement' }"
            ><a>Usermanagement</a></router-link
          >
        </li>
        <li><a :href="monitoringUrl">Logging</a></li>
      </ul>
    </div>
    <div class="navbar-end">
      <div class="mr-4">
        <a :href="checkoutUrl">
          <button><i class="fi fi-sr-shopping-cart"></i></button>
        </a>
      </div>
      <div>
        <a class="btn" @click="logout">Logout</a>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { LoginRegisterService } from "@/services/LoginRegisterService";

let loginRegisterService = new LoginRegisterService();

export default {
  name: "NavBar",
  data() {
    let checkoutUrl = import.meta.env.VITE_CHECKOUT_URL;
    let travelmanagementUrl = import.meta.env.VITE_TRAVELMANAGEMENT_URL;
    let monitoringUrl = import.meta.env.VITE_MONITORING_URL;

    return {
      checkoutUrl,
      travelmanagementUrl,
      monitoringUrl,
    };
  },
  methods: {
    logout() {
      loginRegisterService.LogoutRequest().subscribe((res: any) => {
        if (res.status === 200) {
          this.$router.push("/");
        }
      });
    },
  },
};
</script>

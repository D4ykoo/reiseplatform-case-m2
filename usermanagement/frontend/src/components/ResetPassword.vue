<template>
  <div class="ml-auto mr-auto w-96">
    <label class="form-control w-full">
      <div class="label">
        <span class="label-text">Username</span>
      </div>
      <input
        v-model="username"
        type="text"
        placeholder="Username"
        class="input input-bordered w-full"
      />
    </label>

    <label class="form-control w-full">
      <div class="label">
        <span class="label-text">E-Mail</span>
      </div>
      <input
        v-model="email"
        type="text"
        placeholder="foo@bar.de"
        class="input input-bordered w-full"
      />
    </label>
    <label class="form-control w-full">
      <div class="label">
        <span class="label-text">New Password</span>
      </div>
      <input
        v-model="newPassword"
        type="password"
        placeholder="New Password"
        class="input input-bordered w-full"
      />
    </label>

    <div class="flex flex-row items-center justify-evenly align-middle">
      <button @click="back" class="btn btn-error btn-outline mt-6 w-2/5 flex">
        Cancel
      </button>
      <button @click="reset" class="btn btn-primary w-2/5 mt-6 hover:scale-105">
        Reset
      </button>
    </div>
    <div class="ml-auto mr-auto mt-4 flex justify-center">
      <div role="alert" class="alert">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          class="stroke-info shrink-0 w-6 h-6"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          ></path>
        </svg>
        <span>
          <b>Info:</b> Since the email reset is just mocked simply put the password in
          the input above!</span
        >
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import type { ResetUser } from "@/models/UserModel";
import { LoginRegisterService } from "@/services/LoginRegisterService";
import { ref, type Ref } from "vue";

let loginRegisterService = new LoginRegisterService();

export default {
  name: "ResetPassword",
  setup() {
    const username = ref("");
    const email = ref("");
    const newPassword = ref("");

    return {
      username,
      email,
      newPassword,
    };
  },
  methods: {
    reset() {
      let resetUser: ResetUser = {
        username: this.username,
        email: this.email,
        newPassword: this.newPassword,
      };

      loginRegisterService
        .ResetPasswordRequest(resetUser)
        .subscribe((res: any) => {
          if (res.status === 200){
            this.$router.push("/users");
          }
        });
    },
    back() {
      this.$router.push("/");
    },
  },
};
</script>
